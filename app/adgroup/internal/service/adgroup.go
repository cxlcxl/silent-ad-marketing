package service

import (
	"ad-marketing/app/adgroup/internal/biz"
	"context"

	pb "ad-marketing/api/adgroup/v1"
)

type AdgroupService struct {
	pb.UnimplementedAdgroupServer

	uc *biz.AdgroupUseCase
}

func NewAdgroupService(adgroupUseCase *biz.AdgroupUseCase) *AdgroupService {
	return &AdgroupService{
		uc: adgroupUseCase,
	}
}

func (s *AdgroupService) CreateAdgroup(ctx context.Context, req *pb.CreateAdgroupRequest) (*pb.CreateAdgroupReply, error) {
	return &pb.CreateAdgroupReply{}, nil
}
func (s *AdgroupService) UpdateAdgroup(ctx context.Context, req *pb.UpdateAdgroupRequest) (*pb.UpdateAdgroupReply, error) {
	return &pb.UpdateAdgroupReply{}, nil
}
func (s *AdgroupService) DeleteAdgroup(ctx context.Context, req *pb.DeleteAdgroupRequest) (*pb.DeleteAdgroupReply, error) {
	return &pb.DeleteAdgroupReply{}, nil
}
func (s *AdgroupService) GetAdgroup(ctx context.Context, req *pb.GetAdgroupRequest) (*pb.GetAdgroupReply, error) {
	return &pb.GetAdgroupReply{}, nil
}
func (s *AdgroupService) ListAdgroup(ctx context.Context, req *pb.ListAdgroupRequest) (*pb.ListAdgroupReply, error) {
	return &pb.ListAdgroupReply{}, nil
}
