package main

import (
	"database/sql"
	"fmt"
	"log"
	"ssproxy/back/cmd/api"
	"ssproxy/back/configs"
	logger "ssproxy/back/internal/pkg"
	"time"

	"ssproxy/back/db"

	"github.com/go-sql-driver/mysql"
)

func init() {
	current_time := time.Now().Local()
	fileName := "main-" + current_time.Format("2006-01-02") + ".log"
	logger.InitLog("../logs/" + fileName)
}

func main() {
	logger.Info.Println("Server is starting")

	cfg := mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPwd,
		Addr:                 configs.Envs.DBAds,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, _ := db.NewMySQLStorage(cfg)
	pingDB(db)

	logger.Info.Println("ping DataBase")

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func pingDB(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		logger.Error.Println(err)
	}
}
