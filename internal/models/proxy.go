package models

import (
	gojson "github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

/*
	Алгоритм прост
	1 - Собираем прокси (Единоразовая операция отдельный сервис)
	2 - Проверяем все прокси перед сохранением
	3 - Валидные добавляем в бд
	4 - Добавляем время последней проверки
	5 - Каждые n минут проверяем прокси в базе и удаляем невалидные
	6 - Каждые 3 часов база с прокси обновляется
*/

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

	return "Success savind"
}
