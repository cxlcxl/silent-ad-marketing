package service

import (
	"ad-marketing/app/adgroup/internal/biz"
	"context"
	"github.com/jinzhu/copier"

	pb "ad-marketing/api/adgroup/v1"
)

func (s *AdgroupService) CreateAdgroup(ctx context.Context, req *pb.CreateAdgroupRequest) (*pb.CreateAdgroupReply, error) {
	var adgroupData biz.Adgroup
	err := copier.Copy(&adgroupData, req)
	if err != nil {
		return nil, err
	}
	adgroup, err := s.uc.CreateAdgroup(ctx, &adgroupData)
	return &pb.CreateAdgroupReply{
		AdgroupId: adgroup.AdgroupId,
	}, err
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
