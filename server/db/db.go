package db

import (
	"database/sql"
	"log"
)

// TODO : [#4] docker 세팅 후 mysql로 변경하며 커넥션 관련 개선
type DBSetting struct {
	FilePath string
}

const driverName = "sqlite3"

func MustGetDB(s DBSetting) *sql.DB {
	db, err := sql.Open(driverName, s.FilePath)
	if err != nil {
		log.Fatalln("fail db open", err)
	}
	return db
}
