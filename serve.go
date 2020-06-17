package craftboard

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	logMiddleware "github.com/neko-neko/echo-logrus/v2"
	nekoLog "github.com/neko-neko/echo-logrus/v2/log"
	log "github.com/sirupsen/logrus"
)

type httpError struct {
	Error string
}

func constructHTTPError(err error) httpError {
	return httpError{err.Error()}
}

// Serve craftboard server
func Serve(config Config) {
	e := echo.New()

	nekoLog.Logger().SetOutput(os.Stdout)
	nekoLog.Logger().SetLevel(echoLog.INFO)
	// nekoLog.Logger().SetFormatter(&logrus.JSONFormatter{
	// 	TimestampFormat: time.RFC3339,
	// })
	e.Logger = nekoLog.Logger()
	e.Use(logMiddleware.Logger())

	// Cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.Static("/", "app/dist/")

	g := e.Group("/api/v1")
	g.GET("/board", func(c echo.Context) error {
		board, err := GetBoardFromFile(config.BoardFile)
		if err != nil {
			log.Warnf("couldn't get JSON board: %v", err)
			return c.JSON(http.StatusInternalServerError, constructHTTPError(err))
		}
		return c.JSON(http.StatusOK, board)
	})
	g.POST("/board", func(c echo.Context) error {

		board := Board{}

		decoder := json.NewDecoder(c.Request().Body)
		err := decoder.Decode(&board)
		if err != nil {
			log.Warnf("couldn't decode JSON board: %v", err)
			return c.JSON(http.StatusInternalServerError, constructHTTPError(err))
		}

		err = SaveBoardToFile(config.BoardFile, board)
		if err != nil {
			log.Warnf("couldn't save JSON board: %v", err)
			return c.JSON(http.StatusInternalServerError, constructHTTPError(err))
		}
		return c.JSON(http.StatusOK, nil)
	})

	e.Logger.Fatal(e.Start(":1234"))
}
