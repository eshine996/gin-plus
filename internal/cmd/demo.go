package cmd

import (
	"fmt"
	"gin-plus/internal/controller"
	"gin-plus/internal/pkg/ginp"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()
	_ = engine.SetTrustedProxies(nil)

	r := ginp.NewRouter(engine)
	r.Group("/app", func(app ginp.Group) {
		app.Bind(
			controller.User,
		)
	})

	port := ginp.Config.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	err := engine.Run(fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
