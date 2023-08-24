package main

import (
	"go-web-example/conf"
	"go-web-example/router"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	// load config
	conf.Load()

	// set run mode
	gin.SetMode(conf.ServerConf.RunMode)

	r := router.InitRouter()

	// address
	addr := strings.Builder{}
	addr.Grow(32)

	addr.WriteString(conf.ServerConf.Host)
	addr.WriteString(":")
	addr.WriteString(strconv.Itoa(conf.ServerConf.Port))

	if err := r.Run(addr.String()); err != nil {
		panic(err)
	}
}
