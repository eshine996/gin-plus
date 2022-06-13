package ginp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

//IContext 自己的context ，携带了token的登录信息
type IContext struct {
	*gin.Context
	UserInfo
}

type ResponseData interface{} //响应的数据
type RequestBody interface{}  //请求的数据

//IHandlerFunc 函数包装器-->统一响应
type IHandlerFunc func(c *IContext) (ResponseData, error)
type IMiddlewareFunc func(c *gin.Context)

func checkHandlerFunc(handler interface{}) error {
	reflectType := reflect.TypeOf(handler)

	if reflectType.NumIn() != 2 || reflectType.NumOut() != 2 {
		return errors.New("invalid handler: 'func(context.Context, RequestBody) (ResponseData, error)' is required")
	}

	if reflectType.In(0).String() != "context.Context" {
		return errors.New("invalid handler: the first input parameter should be type of 'context.Context'")
	}

	if reflectType.Out(1).String() != "error" {
		return errors.New("the last output parameter should be type of 'error'")
	}

	if !strings.HasSuffix(reflectType.In(1).String(), `Req`) {
		return errors.New("invalid struct naming for request,it should be named with 'Req' suffix like 'xxxReq'")
	}

	if !strings.HasSuffix(reflectType.Out(0).String(), `Res`) {
		return errors.New("invalid struct naming for response,it should be named with 'Res' suffix like 'xxxRes'")
	}

	return nil
}

func handlerResponse(ctx *gin.Context, data ResponseData, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	} else {
		h := gin.H{
			"code":    http.StatusOK,
			"message": "操作成功",
			"data":    data,
		}

		ctx.JSON(http.StatusOK, h)
		return
	}
}

//包装器需要定响应协议
func handlerWithUserInfoWrapper(handler interface{}) gin.HandlerFunc {
	f, ok := handler.(IHandlerFunc)
	if ok {
		return func(ctx *gin.Context) {
			ic := IContext{Context: ctx}
			resp, err := f(&ic)
			handlerResponse(ctx, resp, err)
		}
	} else {
		return func(ctx *gin.Context) {
			defer func() {
				if err := recover(); err != nil {
					info := err.(string)
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"code":    http.StatusBadRequest,
						"message": info,
					})
					return
				}
			}()

			handlerType := reflect.TypeOf(handler)
			requestBodyType := handlerType.In(1).Elem()
			requestBodyValue := reflect.New(requestBodyType)

			//验证参数
			if ctx.Request.Method == http.MethodGet {
				if err := ctx.ShouldBindQuery(requestBodyValue.Interface()); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"code":    http.StatusBadRequest,
						"message": err.Error(),
						"error":   err.Error(),
					})
					return
				}
			}

			if err := ctx.ShouldBindJSON(requestBodyValue.Interface()); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"code":    http.StatusBadRequest,
					"message": err.Error(),
					"error":   err.Error(),
				})
				return
			}

			//执行handler函数
			results := reflect.ValueOf(handler).Call([]reflect.Value{
				reflect.ValueOf(ctx.Request.Context()),
				requestBodyValue,
			})

			var err error
			if !results[1].IsNil() {
				err = results[1].Interface().(error)
			}

			resp := results[0].Interface()
			handlerResponse(ctx, resp, err)
		}
	}
}
