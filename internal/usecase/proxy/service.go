package proxy

import (
	"GoProxyService/internal/entity"
	"errors"

	gojson "github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

type ProxyService struct {
	storage Storage
}

func NewService(s Storage) Service {
	return &ProxyService{s}
}

func (s *ProxyService) SaveData(body []byte) error {
	var pd []entity.ProxyData

	if err := gojson.Unmarshal(body, &pd); err != nil {
		logrus.Errorf("error unmarshal data: %v", err)
		return errors.New("invalid json obj, check your obj")
	}

	if err := s.storage.SaveData(pd); err != nil {
		logrus.Errorf("error save data: %v", err)
		return errors.New("error save data")
	}

	return nil
}
