package azure

import (
	"context"
)

const (
	UserAgent = "cluster-api-azure-services"
)

type Spec interface{}
type Service interface {
	Get(ctx context.Context, spec Spec) (interface{}, error)
	CreateOrUpdate(ctx context.Context, spec Spec) error
	Delete(ctx context.Context, spec Spec) error
}
