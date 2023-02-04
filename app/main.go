package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Env  string `mapstructure:"env"`
		Port string `mapstructure:"port"`
	} `mapstructure:"app"`
}

func readConfig(conf *AppConfig) error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "read in config error")
	}

	if err := viper.Unmarshal(conf); err != nil {
		return errors.Wrap(err, "convert config to struct error")
	}

	fmt.Printf("config: %+v", *conf)

	return nil
}

func main() {
	// read config
	conf := AppConfig{}
	if err := readConfig(&conf); err != nil {
		log.Fatalf("read config error: %+v", err)
	}

	// echo http server
	e := echo.New()
	e.Use(
		middleware.CORS(),
		middleware.Logger(),
		middleware.RequestID(),
	)

	// router
	e.GET("/", func(c echo.Context) error {
		res := Response{
			Message: "Hello, World",
			Env:     conf.App.Env,
		}
		return c.JSON(http.StatusOK, res)
	})

	// start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", conf.App.Port)))
}

type Response struct {
	Message string `json:"message"`
	Env     string `json:"env"`
}
