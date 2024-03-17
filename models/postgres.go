package models

type PostgresInfo struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSL      string
}
