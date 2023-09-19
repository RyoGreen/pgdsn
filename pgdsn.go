package pgdsn

import (
	"fmt"
)

type SslModeType string

const (
	Disable    SslModeType = "disable"
	Require    SslModeType = "require"
	VerifyCa   SslModeType = "verify-ca"
	VerifyFull SslModeType = "verify-full"
)

type Config struct {
	User     string
	DbName   string
	Password string
	Host     string
	Port     int
	SslMode  SslModeType
}

func (c *Config) FormatDSN() string {
	var dsn = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%v", c.User, c.DbName, c.Password, c.Host, c.Port)
	if c.SslMode != "" {
		dsn += fmt.Sprintf(" sslmode=%v", c.SslMode)
	}
	return dsn
}
