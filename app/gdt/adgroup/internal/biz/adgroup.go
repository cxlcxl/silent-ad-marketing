package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound("", "user not found")
)

type Adgroup struct {
	AdgroupId        int64
	AdgroupName      string
	AdvertiserId     int64
	CompanyId        int64
	ConfiguredStatus string
	BeginDate        string
	EndDate          string
	TimeSeries       string
	CreatedAt        int64
	UpdatedAt        int64
}

// AdgroupRepo is a Greater repo.
type AdgroupRepo interface {
	Save(context.Context, *Adgroup) (*Adgroup, error)
	Update(context.Context, *Adgroup) (*Adgroup, error)
	FindByID(context.Context, int64) (*Adgroup, error)
	ListByHello(context.Context, string) ([]*Adgroup, error)
	ListAll(context.Context) ([]*Adgroup, error)
}

// AdgroupUseCase .
type AdgroupUseCase struct {
	repo AdgroupRepo
	log  *log.Helper
}

// NewAdgroupUseCase new a Greeter usecase.
func NewAdgroupUseCase(repo AdgroupRepo, logger log.Logger) *AdgroupUseCase {
	return &AdgroupUseCase{repo: repo, log: log.NewHelper(logger)}
}

// CreateAdgroup creates a Greeter, and returns the new Greeter.
func (uc *AdgroupUseCase) CreateAdgroup(ctx context.Context, g *Adgroup) (*Adgroup, error) {
	uc.log.WithContext(ctx).Infof("CreateAdgroup: %v", g.AdgroupName)
	return uc.repo.Save(ctx, g)
}