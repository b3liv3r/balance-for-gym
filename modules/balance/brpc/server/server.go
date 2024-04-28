package server

import (
	"context"
	"fmt"
	"github.com/b3liv3r/balance-for-gym/modules/balance/service"
	walletv1 "github.com/b3liv3r/protos-for-gym/gen/go/wallet"
	"github.com/golang/protobuf/ptypes"
)

type WalletRPCServer struct {
	walletv1.UnimplementedWalletServer
	srv service.Walleter
}

func NewWalletRPCServer(srv service.Walleter) walletv1.WalletServer {
	return &WalletRPCServer{srv: srv}
}

func (s *WalletRPCServer) Create(ctx context.Context, req *walletv1.CreateRequest) (*walletv1.CreateResponse, error) {
	message, err := s.srv.Create(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &walletv1.CreateResponse{
		Message: message,
	}, nil
}

func (s *WalletRPCServer) GetBalance(ctx context.Context, req *walletv1.GetBalanceRequest) (*walletv1.GetBalanceResponse, error) {
	wallet, err := s.srv.GetByID(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &walletv1.GetBalanceResponse{
		Balance: wallet.Balance,
	}, nil
}

func (s *WalletRPCServer) Transaction(ctx context.Context, req *walletv1.TransactionRequest) (*walletv1.TransactionResponse, error) {
	err := s.srv.Update(ctx, int(req.UserId), req.Amount, req.Type, req.Description)
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Transaction of %.2f completed successfully", req.Amount)

	wallet, err := s.srv.GetByID(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &walletv1.TransactionResponse{
		Message: message,
		Balance: wallet.Balance,
	}, nil
}

func (s *WalletRPCServer) History(ctx context.Context, req *walletv1.HistoryRequest) (*walletv1.HistoryResponse, error) {
	transactions, err := s.srv.ListTransactions(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}

	protoTransactions := []*walletv1.Transaction{}
	for _, t := range transactions {
		timestamp, _ := ptypes.TimestampProto(t.Date)
		protoTransaction := &walletv1.Transaction{
			Id:          fmt.Sprintf("%d", t.Id),
			Type:        t.Type,
			Amount:      t.Amount,
			Description: t.Description,
			Timestamp:   timestamp,
		}
		protoTransactions = append(protoTransactions, protoTransaction)
	}

	return &walletv1.HistoryResponse{
		Transactions: protoTransactions,
	}, nil
}
