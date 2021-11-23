package connection

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewConnectionToDB() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4", "kimpv", "G6BdxfTSusn23Huu", "10.3.80.60:3306", "lotus_layer")
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("Failed to connect database - err: %v\n", err)
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	fmt.Println("Database connection success!")

	return db
}

func CloseConnection(db *sqlx.DB) error {
	err := db.Close()
	if err != nil {
		fmt.Printf("Failed to clase connection db - err: %v\n", err)
	}

	return err
}
