package main

import (
	"log/slog"
	"server/cmd"
)

func main() {
	server, err := cmd.InitApp()
	if err != nil {
		slog.Error("error:%v", err)
	}
	err = server.Run(":8080")
	if err != nil {
		slog.Error("failed to start server:%v", err)
	}
}
