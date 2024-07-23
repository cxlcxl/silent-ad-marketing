package service

import (
	pb "ad-marketing/api/adgroup/v1"
	"ad-marketing/app/adgroup/internal/biz"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdgroupService)

type AdgroupService struct {
	pb.UnimplementedAdgroupServer

	uc *biz.AdgroupUseCase
}

func NewAdgroupService(adgroupUseCase *biz.AdgroupUseCase) *AdgroupService {
	return &AdgroupService{
		uc: adgroupUseCase,
	}
}
