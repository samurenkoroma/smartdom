package http

import (
	"smartdom/services/library/internal/useCase"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Delivery struct {
	ucBook useCase.BookReader
	router *gin.Engine
}

func New(ucBook useCase.BookReader) *Delivery {
	var d = &Delivery{
		ucBook: ucBook,
	}

	d.router = d.initRouter()
	return d
}

func (d *Delivery) Run() error {
	return d.router.Run(":8079")
}

func (d *Delivery) initRouter() *gin.Engine {

	if viper.GetBool("IS_PRODUCTION") {
		switch strings.ToUpper(strings.TrimSpace(viper.GetString("LOG_LEVEL"))) {
		case "DEBUG":
			gin.SetMode(gin.DebugMode)
		default:
			gin.SetMode(gin.ReleaseMode)
		}
	} else {
		gin.SetMode(gin.DebugMode)
	}

	var router = gin.New()

	d.routerBooks(router.Group("/books"))

	return router
}
