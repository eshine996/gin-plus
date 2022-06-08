package ginp

import "github.com/gin-gonic/gin"

type router struct {
	engine *gin.Engine
}

func (r *router) Group(path string, callback func(Group)) {
	callback(Group{
		engine: r.engine,
		path:   path,
	})
}

func NewRouter(engine *gin.Engine) *router {
	return &router{engine: engine}
}
