package main

import (
	"github.com/b3liv3r/balance-for-gym/config"
	"github.com/b3liv3r/balance-for-gym/modules/balance/brpc/server"
	"github.com/b3liv3r/balance-for-gym/modules/balance/repository"
	"github.com/b3liv3r/balance-for-gym/modules/balance/service"
	"github.com/b3liv3r/balance-for-gym/modules/db"
	walletv1 "github.com/b3liv3r/protos-for-gym/gen/go/wallet"
	"gitlab.com/golight/loggerx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	appConf := config.MustLoadConfig(".env")

	logger := loggerx.InitLogger(appConf.Name, appConf.Production)

	sqlDB, err := db.NewSqlDB(logger, appConf.Db)
	if err != nil {
		logger.Fatal("failed to connect to db", zap.Error(err))
	}

	repo := repository.NewWalletRepositoryDB(sqlDB)
	service := service.NewWalletService(repo, logger)
	s := InitRPC(service)
	lis, err := net.Listen("tcp", appConf.GrpcServerPort)
	if err != nil {
		logger.Error("failed to listen:", zap.Error(err))
	}
	logger.Info("grpc server listening at", zap.Stringer("address", lis.Addr()))
	if err = s.Serve(lis); err != nil {
		logger.Fatal("failed to serve:", zap.Error(err))
	}
}

func InitRPC(wservice service.Walleter) *grpc.Server {
	s := grpc.NewServer()
	walletv1.RegisterWalletServer(s, server.NewWalletRPCServer(wservice))

	return s
}
