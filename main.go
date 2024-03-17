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
	//err = server.Run(":8080")
	err = server.RunTLS(":8080", "./assets/cert/server.pem", "./assets/cert/server.key")
	if err != nil {
		slog.Error("failed to start server:%v", err)
	}
}
