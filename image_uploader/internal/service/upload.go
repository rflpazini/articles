package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SaveUploadedFile(file multipart.File, fileName string) error {
	if err := os.MkdirAll("uploads", 0755); err != nil {
		return fmt.Errorf("erro ao criar pasta de uploads: %v", err)
	}

	uploadPath := filepath.Join("uploads", fileName)
	out, err := os.Create(uploadPath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo no servidor: %v", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return fmt.Errorf("erro ao gravar arquivo: %v", err)
	}

	return nil
}
