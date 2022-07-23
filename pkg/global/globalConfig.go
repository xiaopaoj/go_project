package global

import (
	"time"
)

var (
	// Conf global global var
	Conf *Config
)

// Config global propertie
// nolint
type Config struct {
	Name              string
	Version           string
	Mode              string
	PprofPort         string
	URL               string
	JwtSecret         string
	JwtTimeout        int
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
	EnableTrace       bool
	EnablePprof       bool
	HTTP              ServerConfig
	GRPC              ServerConfig
	Log		  		  LogConfig
}

// ServerConfig server propertie.
type ServerConfig struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type LogConfig struct {
	FilePath string
	FileName string
}
