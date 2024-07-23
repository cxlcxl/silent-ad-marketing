package service

import (
	pb "ad-marketing/api/gdt/adgroup/v1"
	"ad-marketing/app/gdt/adgroup/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdgroupService)

type AdgroupService struct {
	pb.UnimplementedAdgroupServer

	uc *biz.AdgroupUseCase
	l  log.Logger
}

func NewAdgroupService(adgroupUseCase *biz.AdgroupUseCase, logger log.Logger) *AdgroupService {
	return &AdgroupService{
		uc: adgroupUseCase,
		l:  logger,
	}
}
