package ginp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type Group struct {
	engine *gin.Engine
	path   string
}

func (g Group) Group(path string, callback func(Group)) {
	g.path += path
	callback(g)
}

func (g *Group) Handle(method string, relativePath string, handler ...gin.HandlerFunc) {
	if relativePath == "/" {
		relativePath = g.path
	} else {
		relativePath = g.path + relativePath
	}

	g.engine.Handle(method, relativePath, handler...)
}

func (g *Group) GET(relativePath string, handler IHandlerFunc, middleware ...gin.HandlerFunc) {
	handlers := g.mergeHandlers(handler, middleware...)
	g.Handle(http.MethodGet, relativePath, handlers...)
}

func (g *Group) POST(relativePath string, handler IHandlerFunc, middleware ...gin.HandlerFunc) {
	handlers := g.mergeHandlers(handler, middleware...)
	g.Handle(http.MethodPost, relativePath, handlers...)
}

func (g *Group) PUT(relativePath string, handler IHandlerFunc, middleware ...gin.HandlerFunc) {
	handlers := g.mergeHandlers(handler, middleware...)
	g.Handle(http.MethodPut, relativePath, handlers...)
}

func (g *Group) DELETE(relativePath string, handler IHandlerFunc, middleware ...gin.HandlerFunc) {
	handlers := g.mergeHandlers(handler, middleware...)
	g.Handle(http.MethodDelete, relativePath, handlers...)
}
func (g *Group) Bind(relativePath string, handler interface{}, middleware ...gin.HandlerFunc) {
	handlers := g.mergeHandlers(handler, middleware...)
	g.Handle(http.MethodPost, relativePath, handlers...)
}

func (g *Group) Bind2(handlerOrObject ...interface{}) {
	for _, v := range handlerOrObject {
		reflectType := reflect.TypeOf(v)
		reflectValue := reflect.ValueOf(v)

		switch reflectType.Kind() {
		case reflect.Func:
			err := checkHandlerFunc(reflectValue.Interface())
			if err != nil {

				fmt.Println(err.Error())
				panic(err)
			}

			objectReq := reflect.New(reflectType.In(1).Elem())
			path, method := parseApiTag(objectReq.Type())
			g.Handle(method, path, g.mergeHandlers(v)...)
		case reflect.Struct:
			newValue := reflect.New(reflectType)
			newValue.Elem().Set(reflectValue)

			for i := 0; i < newValue.NumMethod(); i++ {
				err := checkHandlerFunc(newValue.Method(i).Interface())
				if err != nil {
					fmt.Println(err.Error())
					panic(err)
				}

				objectReq := reflect.New(newValue.Method(i).Type().In(1).Elem())
				path, method := parseApiTag(objectReq.Type())
				g.Handle(method, path, g.mergeHandlers(newValue.Method(i).Interface())...)
			}
		}
	}
}

func parseApiTag(rType reflect.Type) (path, method string) {
	field, _ := rType.Elem().FieldByName("Api")
	path = field.Tag.Get("path")
	method = field.Tag.Get("method")
	return
}

func (g *Group) Use(middleware ...gin.HandlerFunc) {
	g.engine.RouterGroup.Use(middleware...)
}

func (g *Group) mergeHandlers(handler interface{}, middleware ...gin.HandlerFunc) gin.HandlersChain {
	mergedHandlers := make(gin.HandlersChain, len(middleware)+1)
	mergedHandlers[0] = handlerWithUserInfoWrapper(handler)
	copy(mergedHandlers, middleware)

	return mergedHandlers
}
