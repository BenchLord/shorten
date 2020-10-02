package ports

import (
	"context"
	pb "shorten/protos"
)

type Storage interface {
	CreateUrl(context.Context, *pb.CreateUrlReq) (*pb.CreateUrlRes, error)
	GetUrl(context.Context, *pb.GetUrlReq) (*pb.Url, error)
}
