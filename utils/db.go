package utils

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/user"
)

var (
	Db       *sql.DB
	BasePath string
)

func init() {
	currentUser, err := user.Current()
	CheckErr(err)
	BasePath = currentUser.HomeDir
	dbFilePath := BasePath + "/" + "tsm.db"

	Db, err = sql.Open("sqlite3", dbFilePath)
	CheckErr(err)

	if exists, _ := PathExists(dbFilePath); exists {
		return
	}

	tableSql := `
CREATE TABLE IF NOT EXISTS "main"."servers" (
  "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
  "name" text NOT NULL,
  "host" text NOT NULL,
  "port" text NOT NULL,
  "user" text NOT NULL,
  "pass" text NOT NULL,
  "created_at" integer(10) NOT NULL,
  "updated_at" integer(10) NOT NULL,
  CONSTRAINT "uni_name" UNIQUE ("name")
);
`

	_, err = Db.Exec(tableSql)
	CheckErr(err)
}

// CheckErr 检测并panic error
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// PathExists 检测文件/目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
