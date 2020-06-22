package middleware

import (
	"context"
	"reflect"
	"testing"

	"github.com/awslabs/smithy-go/middleware"
)

type mockInitalizeHandler func(context.Context, middleware.InitializeInput) (middleware.InitializeOutput, middleware.Metadata, error)

func (m mockInitalizeHandler) HandleInitialize(
	ctx context.Context, in middleware.InitializeInput,
) (middleware.InitializeOutput, middleware.Metadata, error) {
	return m(ctx, in)
}

func TestServiceMetadataProvider(t *testing.T) {
	m := RegisterServiceMetadata{
		ServiceName:    "Foo",
		ServiceID:      "Bar",
		EndpointPrefix: "Baz",
		SigningName:    "Jaz",
		Region:         "Fuz",
		OperationName:  "FooOp",
	}

	_, _, err := m.HandleInitialize(context.Background(), middleware.InitializeInput{}, mockInitalizeHandler(func(
		ctx context.Context, input middleware.InitializeInput,
	) (o middleware.InitializeOutput, m middleware.Metadata, err error) {
		t.Helper()
		if e, a := "Foo", GetServiceName(ctx); e != a {
			t.Errorf("expected %v, got %v", e, a)
		}
		if e, a := "Bar", GetServiceID(ctx); e != a {
			t.Errorf("expected %v, got %v", e, a)
		}
		if e, a := "Baz", GetEndpointPrefix(ctx); e != a {
			t.Errorf("expected %v, got %v", e, a)
		}
		if e, a := "Jaz", GetSigningName(ctx); e != a {
			t.Errorf("expected %v, got %v", e, a)
		}
		if e, a := "Fuz", GetRegion(ctx); e != a {
			t.Errorf("expected %v, got %v", e, a)
		}
		if e, a := "FooOp", GetOperationName(ctx); !reflect.DeepEqual(e, a) {
			t.Errorf("expected %v, got %v", e, a)
		}
		return o, m, err
	}))
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
