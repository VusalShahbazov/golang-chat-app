package migrations

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"os"
)

func Run(baseDir string) error {
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?multiStatements=true&parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return  err
	}
	defer db.Close()
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	fmt.Sprintf("file://%vdatabase/migrations/"  , baseDir)
	m, err  := migrate.NewWithDatabaseInstance(
		"file://" + baseDir,
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		log.Println(err)
	}
	return  nil
}