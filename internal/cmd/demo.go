package cmd

import (
	"gin-plus/internal/controller"
	"gin-plus/internal/pkg/ginp"
	"github.com/gin-gonic/gin"
)

func TestStart() {
	engine := gin.Default()
	_ = engine.SetTrustedProxies(nil)

	r := ginp.NewRouter(engine)
	r.Group("/app", func(group ginp.Group) {
		group.Bind2(
			controller.User,
		)
	})

	//server := &http.Server{
	//	//Addr:    fmt.Sprintf("%s:%d", g.Config.App.IP, g.Config.App.Port),
	//	Handler: r,
	//}
	engine.Run()
}
