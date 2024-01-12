package mysql

import (
	"database/sql"
	"fmt"
	"taskema/pkg/richerror"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Port     uint   `koanf:"port"`
	Host     string `koanf:"host"`
	DBName   string `koanf:"db_name"`
}

type MYSQL struct {
	cfg   Config
	mysql *sql.DB
}

func New(cfg Config) (*MYSQL, error) {
	op := "mysql.New"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		return nil,
			richerror.New(op).
				WithMessage(err.Error()).
				WithCode(richerror.CodeUnexpected)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	mysql := MYSQL{
		cfg:   cfg,
		mysql: db,
	}

	return &mysql, nil
}

func (mysql *MYSQL) Conn() *sql.DB {
	return mysql.mysql
}
