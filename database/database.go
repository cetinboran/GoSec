package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// mysql giriş bilgileri.
	dbUser = "root"
	dbPass = "Boran123."
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "themanager"
	DBConn *sql.DB
)

func InitDB() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	DBConn, err = sql.Open("mysql", dsn) // Mysql bağlandık.

	if err != nil {
		return err
	}

	// Bağlantı sağlandı mı kontrol ediyoruz.
	err = DBConn.Ping()
	if err != nil {
		return fmt.Errorf("veritabanına bağlanırken hata oluştu: %v", err)
	}
	return nil
}
