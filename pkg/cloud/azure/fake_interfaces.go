package azure

import (
	"context"
	"fmt"
	"reflect"
	"github.com/Azure/go-autorest/autorest"
)

type FakeSuccessService struct{}
type FakeFailureService struct{}
type FakeNotFoundService struct{}
type FakeCachedService struct{ Cache *map[string]int }

func (s *FakeSuccessService) Get(ctx context.Context, spec Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func (s *FakeSuccessService) CreateOrUpdate(ctx context.Context, spec Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
func (s *FakeSuccessService) Delete(ctx context.Context, spec Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

type FakeStruct struct{}

func (s *FakeFailureService) Get(ctx context.Context, spec Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return FakeStruct{}, fmt.Errorf("Failed to Get service")
}
func (s *FakeFailureService) CreateOrUpdate(ctx context.Context, spec Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Errorf("Failed to Create")
}
func (s *FakeFailureService) Delete(ctx context.Context, spec Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Errorf("Failed to Delete")
}
func (s *FakeNotFoundService) Get(ctx context.Context, spec Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, autorest.DetailedError{StatusCode: 404}
}
func (s *FakeNotFoundService) CreateOrUpdate(ctx context.Context, spec Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autorest.DetailedError{StatusCode: 404}
}
func (s *FakeNotFoundService) Delete(ctx context.Context, spec Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return autorest.DetailedError{StatusCode: 404}
}
func (s *FakeCachedService) Get(ctx context.Context, spec Spec) (interface{}, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil, nil
}
func (s *FakeCachedService) CreateOrUpdate(ctx context.Context, spec Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if spec == nil {
		return nil
	}
	v := reflect.ValueOf(spec).Elem()
	(*s.Cache)[v.FieldByName("Name").String()]++
	return nil
}
func (s *FakeCachedService) Delete(ctx context.Context, spec Spec) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}
