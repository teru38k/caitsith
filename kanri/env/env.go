package env

import (
	"github.com/jackc/pgx"
	"github.com/rs/zerolog"
)

func PostgresDSN() string {
	dsn := "host=localhost user=root password=root dbname=db1 port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	return dsn
}

func LogLevel() LoggerLevel {
	lvl := "info"
	return getLogLevel(lvl)
}

func ZerologLevel() zerolog.Level {
	return LogLevel().ZerlogLevel()
}

func PgxlogLevel() pgx.LogLevel {
	return LogLevel().PgxLogLevel()
}

func EnableLogFile() bool {
	return true
}
