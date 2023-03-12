package database

import (
	"GoProxyService/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func ConnectDataBase() (*sql.DB, error) {
	conf := config.ReadConfig()

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Dbname,
		conf.Database.Sslmode))
	if err != nil {
		logrus.Errorf("error open connect to database: %v", err)
		return nil, err
	}
	return db, nil
}
