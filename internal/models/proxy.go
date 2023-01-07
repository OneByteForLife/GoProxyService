package models

import (
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

	/*
		Проверка прокси
			Сохранение в базу
	*/

	return "Success saving"
}
