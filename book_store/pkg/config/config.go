package config

import "time"

type Config struct {
	Server   *ServerConfig   `json:"server"`
	Database *DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	CookieName        string        `json:"cookieName"`
	AppName           string        `json:"appName"`
	Port              string        `json:"port"`
	PprofPort         string        `json:"pprofPort"`
	Mode              string        `json:"mode"`
	JwtSecretKey      string        `json:"jwtSecretKey"`
	AppVersion        string        `json:"appVersion"`
	Repository        string        `json:"repository"`
	CtxDefaultTimeout time.Duration `json:"ctxDefaultTimeout"`
	WriteTimeout      time.Duration `json:"writeTimeout"`
	ReadTimeout       time.Duration `json:"readTimeout"`
	SSL               bool          `json:"ssl"`
	CSRF              bool          `json:"csrf"`
	Debug             bool          `json:"debug"`
}

type DatabaseConfig struct {
	URI     string `json:"uri"`
	Name    string `json:"name"`
	Timeout int    `json:"timeout"`
}
