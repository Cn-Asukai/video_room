package main

import (
	_ "video_room/app/video/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"video_room/app/video/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
