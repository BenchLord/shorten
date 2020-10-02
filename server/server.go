package server

import (
	"context"

	pb "shorten/protos"
	"shorten/server/ports"
)

type Service struct {
	storage ports.Storage
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) WithStorage(storage ports.Storage) *Service {
	s.storage = storage
	return s
}

func (s *Service) CreateUrl(ctx context.Context, req *pb.CreateUrlReq) (*pb.CreateUrlRes, error) {
	return s.storage.CreateUrl(ctx, req)
}

func (s *Service) GetUrl(ctx context.Context, req *pb.GetUrlReq) (*pb.Url,
	error) {
	return s.storage.GetUrl(ctx, req)
}
