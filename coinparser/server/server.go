package server

import (
	"context"
	"log/slog"
	"net"
	"strconv"

	"github.com/gintokos/coinder/backend/models"
	"github.com/gintokos/coinder/backend/pkg/sl"
	"github.com/gintokos/coinder/coinparser/config"
	"github.com/gintokos/coinder/coinparser/parser"
	"github.com/gintokos/coinder/coinparser/storage"
	pb "github.com/gintokos/coinder/protos/coinupdateprotos"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCoinServiceServer
	db      *storage.Storage
	parser  parser.Parser
	gserver *grpc.Server
}

func (s *Server) UpdateCoins(ctx context.Context, pbcoins *pb.Coins) (*pb.Coins, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		idList := make([]string, 0, len(pbcoins.Coins))
		for _, coin := range pbcoins.Coins {
			idList = append(idList, strconv.Itoa(int(coin.Id)))
		}
		pscoins, err := s.parser.GetListWithoutMeta(idList)

		// TO DO REFACTOR THIS PART TO USE FIFO QUEEE OR SMT LIKE THIS
		go func() {
			dbcoins := make([]models.DBCoin, 0, len(pscoins))

			for _, coin := range pscoins {
				dbcoins = append(dbcoins, parser.ToDBcoin(&coin))
			}

			err := s.db.UpdateCoins(dbcoins)
			if err != nil {
				slog.Error("error on updating coins: ", "error", sl.Err(err))
			}
		}()

		for k := range pbcoins.Coins {
			pbcoins.Coins[k] = parser.ToPBcoin(pscoins[k])
		}

		return pbcoins, err
	}
}

func NewServer(storage *storage.Storage) (*Server, error) {
	prs := parser.New(storage)
	return &Server{
		parser: prs,
		db:     storage,
	}, nil
}

func (s *Server) MustRun() {
	err := s.Run()
	if err != nil {
		slog.Error("error on running server", sl.Err(err))
		panic(err)
	}
}

func (s *Server) Run() error {
	cfg := config.Server()

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		return err
	}

	gs := grpc.NewServer()
	s.gserver = gs
	pb.RegisterCoinServiceServer(gs, s)

	slog.Info("server running")
	if err := gs.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *Server) GraceFullShutDown(ctx context.Context) error {
	s.gserver.GracefulStop()
	return nil
}
