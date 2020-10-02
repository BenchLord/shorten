package adapters

import (
	"context"
	"time"

	pb "shorten/protos"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"github.com/renstrom/shortuuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RedisStorage struct {
	client *redis.Client
}

func NewRedisStorage(address, port string) (*RedisStorage, error) {
	c := redis.NewClient(&redis.Options{
		Addr: address + ":" + port,
	})

	_, err := c.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisStorage{c}, nil
}

func (rs *RedisStorage) CreateUrl(ctx context.Context, req *pb.CreateUrlReq) (*pb.CreateUrlRes, error) {
	if req.GetUrl() == "" {
		return nil, status.Error(codes.InvalidArgument, "url cannot be empty")
	}

	url := &pb.Url{
		Id:  shortuuid.New(),
		Url: req.GetUrl(),
	}

	bytes, err := proto.Marshal(url)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	duration := time.Second * time.Duration(req.GetKeepAlive())
	if err := rs.client.Set(ctx, url.Id, bytes, duration).Err(); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateUrlRes{Id: url.Id}, nil
}

func (rs *RedisStorage) GetUrl(ctx context.Context, req *pb.GetUrlReq) (*pb.Url, error) {
	if req.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "id cannot be empty")
	}

	entry, err := rs.client.Get(ctx, req.GetId()).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, status.Error(codes.NotFound, "url not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	url := &pb.Url{}
	if err := proto.Unmarshal(entry, url); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	duration := rs.client.TTL(ctx, req.GetId()).Val()
	url.KeepAlive = int64(duration.Seconds())

	return url, nil
}
