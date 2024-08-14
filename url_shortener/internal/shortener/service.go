package shortener

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"url-shortener/pkg/utils/base62"
)

type Service struct {
	Repository RepositoryInterface
	Json       jsoniter.API
}

func (s Service) Upsert(c echo.Context) error {
	u := new(URLInfo)
	if err := c.Bind(u); err != nil {
		fmt.Println(err)
		log.Errorf("error during upsert: %v", err.Error())
		return err
	}

	u.Short = base62.EncodeURL(u.Url)

	result, err := s.Repository.Get(context.Background(), u.Short)
	if result != "" {
		u = s.unmarshal(result)
		u.UpdatedAt = time.Now().UTC().Format(time.RFC3339Nano)
	} else {
		u.CreatedAt = time.Now().UTC().Format(time.RFC3339Nano)
	}

	err = s.Repository.Set(context.Background(), u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

func (s Service) Get(c echo.Context) error {
	key := c.QueryParam("url")

	if key == "" {
		result, err := s.Repository.GetAll(context.Background())
		if err != nil {
			log.Errorf("error during get all: %v", err.Error())
			return err
		}

		resultList, err := s.unmarshalList(result)
		if err != nil {
			fmt.Println(err)
			log.Errorf("error during get %s: %v", key, err.Error())
			return err
		}
		return c.JSON(http.StatusOK, resultList)
	}

	result, err := s.Repository.Get(context.Background(), key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return c.NoContent(http.StatusNotFound)
		}
		return c.NoContent(http.StatusInternalServerError)
	}

	u := s.unmarshal(result)
	return c.Redirect(http.StatusFound, u.Url)
}

func (s Service) unmarshal(result string) *URLInfo {
	var u URLInfo
	if err := s.Json.Unmarshal([]byte(result), &u); err != nil {
		log.Errorf("error while unmarshalling json: %v", err)
	}

	return &u
}

func (s Service) unmarshalList(outerMap map[string]string) (*[]URLInfo, error) {
	var resultList []URLInfo
	// Loop through the map and unmarshal each inner JSON string
	for key, value := range outerMap {
		urlInfo := new(URLInfo)
		err := s.Json.Unmarshal([]byte(value), &urlInfo)
		if err != nil {
			fmt.Printf("Error unmarshalling inner JSON for key %s: %v\n", key, err)
			continue
		}
		resultList = append(resultList, *urlInfo)
	}

	return &resultList, nil
}
