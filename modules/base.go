package modules

import (
	"database/sql"
	"di_sample/domain"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "user"
)

func init() {
	dbConf := "root:root@tcp(127.0.0.1:3306)/test_database?charset=utf8mb4"
	Db, err = sql.Open("mysql", dbConf)

	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		user_id int PRIMARY KEY AUTO_INCREMENT,
		user_name varchar(64),
		status int)`, tableNameUser)

	_, err = Db.Exec(cmdU)

	if err != nil {
		log.Fatalln(err)
	}

	cmdInsert := `INSERT INTO user (user_id, user_name, status) VALUES(?,?,?)`
	_, _ = Db.Exec(cmdInsert, 1, "normal_user", domain.Normal)

	_, _ = Db.Exec(cmdInsert, 2, "prime_user", domain.Prime)

	_, _ = Db.Exec(cmdInsert, 99, "error_user", 999)
}
