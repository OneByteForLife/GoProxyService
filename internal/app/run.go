package app

import (
	"GoProxyService/internal/controller/rest"
	"GoProxyService/internal/usecase/proxy"
	"GoProxyService/pkg/database"

	"github.com/sirupsen/logrus"
)

func Run() {
	db, err := database.ConnectDataBase()
	if err != nil {
		logrus.Error(err)
	}

	api := rest.NewHandler(proxy.NewService(proxy.NewStorage(db)))

	rest.Init(api)
}
