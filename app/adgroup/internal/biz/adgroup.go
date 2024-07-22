package biz

import (
	"context"

	v1 "ad-marketing/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type Adgroup struct {
	Hello string
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

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *AdgroupUseCase) CreateGreeter(ctx context.Context, g *Adgroup) (*Adgroup, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
