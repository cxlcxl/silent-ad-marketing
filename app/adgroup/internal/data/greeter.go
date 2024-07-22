package data

import (
	"ad-marketing/app/adgroup/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type adgroupRepo struct {
	data *Data
	log  *log.Helper
}

// NewAdgroupRepo .
func NewAdgroupRepo(data *Data, logger log.Logger) biz.AdgroupRepo {
	return &adgroupRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *adgroupRepo) Save(ctx context.Context, g *biz.Adgroup) (*biz.Adgroup, error) {
	return g, nil
}

func (r *adgroupRepo) Update(ctx context.Context, g *biz.Adgroup) (*biz.Adgroup, error) {
	return g, nil
}

func (r *adgroupRepo) FindByID(context.Context, int64) (*biz.Adgroup, error) {
	return nil, nil
}

func (r *adgroupRepo) ListByHello(context.Context, string) ([]*biz.Adgroup, error) {
	return nil, nil
}

func (r *adgroupRepo) ListAll(context.Context) ([]*biz.Adgroup, error) {
	return nil, nil
}
