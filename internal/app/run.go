package app

import (
	"GoProxyService/internal/middleware"
	"GoProxyService/internal/models"
	"GoProxyService/internal/routes"
	"time"

	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("api/v1/add", middleware.CheckContentType(), routes.SavingData)
}

func Run() {
	// Конифгурация сервера
	app := fiber.New(fiber.Config{
		JSONDecoder:  gojson.Unmarshal,
		JSONEncoder:  gojson.Marshal,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	})

	app.Use(logger.New(), middleware.CheckJwtToken())

	SetUpRoutes(app)

	go Cron()

	if err := app.Listen(":8080"); err != nil {
		logrus.Fatalf("Err up server - %s", err)
	}

	logrus.Info("Service is up!")
}

func Cron() {
	for {
		time.Sleep(time.Minute * 2)
		count := models.CheckList()
		logrus.Infof("Proxy count from check - %d", count)
	}
}
