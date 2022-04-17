package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type DBSetting struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

const driverName = "mysql"

func MustGetDB(s DBSetting) *sql.DB {
	db, err := sql.Open(driverName, fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true", s.User, s.Password, s.Host, s.Port, s.Name),
	)
	if err != nil {
		log.Fatalln("fail db open", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		log.Fatalln("fail db ping", err)
	}

	return db
}
