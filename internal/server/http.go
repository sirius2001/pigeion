package server

import (
	"context"
	"net/http"
	"pigeon2/internal/conf"
	"pigeon2/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type HTTPServer struct {
	srv  *gin.Engine
	addr string
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, wokerFlow *service.WorkFlowService, logger log.Logger) *HTTPServer {
	srv := gin.Default()
	g := srv.Group("/v1")
	gr := g.Group("/greeter")
	{
		gr.GET("/sayHello", webHandler(greeter.SayHello))
		gr.GET("/sayBye", webHandler(greeter.SayBye))
	}
	return &HTTPServer{srv: srv, addr: c.Http.Addr}
}

// Start start the http server.
func (s *HTTPServer) Start(ctx context.Context) error {
	return s.srv.Run(s.addr)
}

// Stop stop the http server.
func (s *HTTPServer) Stop(ctx context.Context) error {
	return nil
}

func getJsonData(ctx *gin.Context, data any) error {
	if ctx.Request.Method == "GET" {
		if err := ctx.BindQuery(data); err != nil {
			return err
		}
		return nil
	}

	if ctx.Request.Method == "POST" || ctx.Request.Method == "PUT" {
		if err := ctx.ShouldBindJSON(data); err != nil {
			return err
		}
	}
	return nil
}

func webHandler[T, R any](handler func(ctx context.Context, arg *T) (*R, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := new(T)
		if err := getJsonData(ctx, data); err != nil {
			failed(ctx, http.StatusBadRequest, err)
			return
		}

		// 调用处理函数
		resp, err := handler(ctx.Request.Context(), data)
		if err != nil {
			failed(ctx, http.StatusInternalServerError, err)
			return
		}

		// 返回 JSON 响应
		success(ctx, resp)
	}
}

func success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}

func failed(ctx *gin.Context, code int, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": code,
		"msg":  err.Error(),
		"data": nil,
	})
}
