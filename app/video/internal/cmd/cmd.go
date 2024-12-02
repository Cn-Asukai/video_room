package cmd

import (
	"VideoRoom/app/video/internal/controller/video"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/video", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					video.NewV1(),
				)
			})

			s.SetIndexFolder(true)
			//s.SetServerRoot("C:\\Users\\暮居池\\Desktop\\code\\go\\VideoRoom\\app\\video\\resource\\public\\resource")
			s.SetServerRoot("./resource/public/resource")
			s.Run()
			return nil
		},
	}
)
