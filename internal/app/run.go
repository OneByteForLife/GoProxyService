package app

import (
	"GoProxyService/internal/middleware"
	"GoProxyService/internal/routes"
	"time"

	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("api/v1/get", routes.FindProxyToTotal)
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

	if err := app.Listen(":80"); err != nil {
		logrus.Fatalf("Err up server - %s", err)
	}

	logrus.Info("Service is up!")
}
