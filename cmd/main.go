package main

import (
	"fmt"
	"log/slog"
	"net"
	"onion/di"
	"onion/infrastructure/database"
	"onion/router"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

func main() {
	db := database.NewDB()
	defer database.CloseDB(db)

	// listenerの作成
	port := 3000
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		slog.Error(fmt.Sprintf("failed to listen: %v", err))
		panic(err)
	}

	s := grpc.NewServer()
	es := router.NewEntryServer(di.Entry(db))
	router.Register(s, es)

	// grpc serverを起動
	go func() {
		slog.Info(fmt.Sprintf("start gRPC server on port %v", port))
		s.Serve(listener)
	}()

	// サーバーをシグナルで停止
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("stopping gRPC server...")
	s.GracefulStop()
}
