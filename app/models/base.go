package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/hidenari-yuda/todo_app/config"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser      = "users"
	tableNameSession   = "sessions"
	tableNameItem      = "items"
	tableNameChatGroup = "chatgroups"
	tableNameMessage   = "messages"
)

func init() {
	/*url := os.Getenv("DATABASE_URL")
	conncetion, _ := pq.ParseURL(url)
	conncetion += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, conncetion)
	if err != nil {
		log.Fatalln(err)
	}*/

	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		created_at DATETIME,
		name STRING,
		nick_name STRING NULL,
		email STRING,
		password STRING,
		icon_url STRING NULL,
		phone STRING NULL,
		address STRING NULL,
		birthday STRING NULL,
		)`, tableNameUser)

	Db.Exec(cmdU)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id STRING,
		created_at DATETIME)`, tableNameSession)

	Db.Exec(cmdS)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		created_at DATETIME,
		title STRING,
		content TEXT NULL,
		category STRING NULL,
		price INTEGER,
		)`, tableNameItem)
	Db.Exec(cmdT)

	cmdCG := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		created_at DATETIME,
		chat_member STRING NULL,
		chat_name STRING NULL)`, tableNameChatGroup)

	Db.Exec(cmdCG)

	cmdM := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id STRING,
		created_at DATETIME,
		group_id STRING,
		user_name STRING NULL)`, tableNameMessage)

	Db.Exec(cmdM)
}

func createUUID() (uuidObj uuid.UUID) {
	uuidObj, _ = uuid.NewUUID()
	return uuidObj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
