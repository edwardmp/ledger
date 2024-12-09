// Code generated by MockGen. DO NOT EDIT.
//
// Generated by this command:
//
//	mockgen -write_source_comment=false -write_package_comment=false -source ../bucket/bucket.go -destination buckets_generated_test.go -package driver --mock_names Factory=BucketFactory . Factory
package driver

import (
	context "context"
	reflect "reflect"

	migrations "github.com/formancehq/go-libs/v2/migrations"
	ledger "github.com/formancehq/ledger/internal"
	bucket "github.com/formancehq/ledger/internal/storage/bucket"
	gomock "go.uber.org/mock/gomock"
)

// MockBucket is a mock of Bucket interface.
type MockBucket struct {
	ctrl     *gomock.Controller
	recorder *MockBucketMockRecorder
}

// MockBucketMockRecorder is the mock recorder for MockBucket.
type MockBucketMockRecorder struct {
	mock *MockBucket
}

// NewMockBucket creates a new mock instance.
func NewMockBucket(ctrl *gomock.Controller) *MockBucket {
	mock := &MockBucket{ctrl: ctrl}
	mock.recorder = &MockBucketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBucket) EXPECT() *MockBucketMockRecorder {
	return m.recorder
}

// AddLedger mocks base method.
func (m *MockBucket) AddLedger(ctx context.Context, ledger ledger.Ledger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLedger", ctx, ledger)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLedger indicates an expected call of AddLedger.
func (mr *MockBucketMockRecorder) AddLedger(ctx, ledger any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLedger", reflect.TypeOf((*MockBucket)(nil).AddLedger), ctx, ledger)
}

// GetMigrationsInfo mocks base method.
func (m *MockBucket) GetMigrationsInfo(ctx context.Context) ([]migrations.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMigrationsInfo", ctx)
	ret0, _ := ret[0].([]migrations.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMigrationsInfo indicates an expected call of GetMigrationsInfo.
func (mr *MockBucketMockRecorder) GetMigrationsInfo(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMigrationsInfo", reflect.TypeOf((*MockBucket)(nil).GetMigrationsInfo), ctx)
}

// HasMinimalVersion mocks base method.
func (m *MockBucket) HasMinimalVersion(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasMinimalVersion", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasMinimalVersion indicates an expected call of HasMinimalVersion.
func (mr *MockBucketMockRecorder) HasMinimalVersion(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasMinimalVersion", reflect.TypeOf((*MockBucket)(nil).HasMinimalVersion), ctx)
}

// IsInitialized mocks base method.
func (m *MockBucket) IsInitialized(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInitialized", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsInitialized indicates an expected call of IsInitialized.
func (mr *MockBucketMockRecorder) IsInitialized(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInitialized", reflect.TypeOf((*MockBucket)(nil).IsInitialized), arg0)
}

// IsUpToDate mocks base method.
func (m *MockBucket) IsUpToDate(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUpToDate", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUpToDate indicates an expected call of IsUpToDate.
func (mr *MockBucketMockRecorder) IsUpToDate(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUpToDate", reflect.TypeOf((*MockBucket)(nil).IsUpToDate), ctx)
}

// Migrate mocks base method.
func (m *MockBucket) Migrate(ctx context.Context, opts ...migrations.Option) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Migrate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockBucketMockRecorder) Migrate(ctx any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockBucket)(nil).Migrate), varargs...)
}

// BucketFactory is a mock of Factory interface.
type BucketFactory struct {
	ctrl     *gomock.Controller
	recorder *BucketFactoryMockRecorder
}

// BucketFactoryMockRecorder is the mock recorder for BucketFactory.
type BucketFactoryMockRecorder struct {
	mock *BucketFactory
}

// NewBucketFactory creates a new mock instance.
func NewBucketFactory(ctrl *gomock.Controller) *BucketFactory {
	mock := &BucketFactory{ctrl: ctrl}
	mock.recorder = &BucketFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *BucketFactory) EXPECT() *BucketFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *BucketFactory) Create(name string) bucket.Bucket {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name)
	ret0, _ := ret[0].(bucket.Bucket)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *BucketFactoryMockRecorder) Create(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*BucketFactory)(nil).Create), name)
}

// GetMigrator mocks base method.
func (m *BucketFactory) GetMigrator(b string) *migrations.Migrator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMigrator", b)
	ret0, _ := ret[0].(*migrations.Migrator)
	return ret0
}

// GetMigrator indicates an expected call of GetMigrator.
func (mr *BucketFactoryMockRecorder) GetMigrator(b any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMigrator", reflect.TypeOf((*BucketFactory)(nil).GetMigrator), b)
}
