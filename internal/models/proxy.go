package models

import (
	"GoProxyService/internal/database"
	"time"

	gojson "github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

type ProxyData struct {
	Types []string `json:"protocols"`
	Data  struct {
		IP      string `json:"ip"`
		Port    string `json:"port"`
		Speed   int    `json:"speed"`
		AnonLvL string `json:"anon_lvl"`
		Geo     struct {
			City    string `json:"city"`
			Country string `json:"country"`
		} `json:"geo"`
	} `json:"data"`
}

func SaveData(body []byte) string {
	var pd []ProxyData

	if err := gojson.Unmarshal(body, &pd); err != nil {
		logrus.Errorf("Err unmarshal proxy data - %s", err)
		return "Incorrect data"
	}

	logrus.Infof("%d objects were accepted", len(pd))

	db, err := database.ConnectDataBase()
	if err != nil {
		logrus.Error(err)
		return "Err saving data"
	}

	query := "INSERT INTO proxy_list (types, ip, port, speed, anonlvl, city, country, last_check) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	for _, val := range pd {
		if _, err := db.Exec(query, val.Types[0], val.Data.IP, val.Data.Port, val.Data.Speed, val.Data.AnonLvL, val.Data.Geo.City, val.Data.Geo.Country, time.Now().Format("2006-01-02 15:04:05")); err != nil {
			logrus.Errorf("Err write data to database - %s", err)
			return "Err saving data"
		}
	}
	return "Success saving"
}
