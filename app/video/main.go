package main

import (
	_ "VideoRoom/app/video/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "VideoRoom/app/video/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"VideoRoom/app/video/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
