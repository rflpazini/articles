package config

import "time"

type Config struct {
	Server         *ServerConfig   `json:"server"`
	DatabaseConfig *DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	AppName           string        `json:"appName"`
	AppVersion        string        `json:"appVersion"`
	Port              string        `json:"port"`
	PprofPort         string        `json:"pprofPort"`
	Mode              string        `json:"mode"`
	JwtSecretKey      string        `json:"jwtSecretKey"`
	CookieName        string        `json:"cookieName"`
	ReadTimeout       time.Duration `json:"readTimeout"`
	WriteTimeout      time.Duration `json:"writeTimeout"`
	SSL               bool          `json:"ssl"`
	CtxDefaultTimeout time.Duration `json:"ctxDefaultTimeout"`
	CSRF              bool          `json:"csrf"`
	Debug             bool          `json:"debug"`
	Repository        string        `json:"repository"`
}

type DatabaseConfig struct {
	URI     string `json:"uri"`
	Name    string `json:"name"`
	Timeout int    `json:"timeout"`
}
