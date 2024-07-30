package handler_test

import (
	"context"
	"errors"
	"onion/domain/model"
	"onion/interface/handler"
	entrypb "onion/proto/entry"
	"onion/usecase"
	"reflect"
	"testing"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type MockEntryUsecase struct {
	usecase.EntryUsecase
	FakeSearch func(context context.Context, id string) (*model.Entry, error)
	FakeCreate func(context context.Context, user string) (*model.Entry, error)
}

func (m MockEntryUsecase) Search(context context.Context, id string) (*model.Entry, error) {
	return m.FakeSearch(context, id)
}
func (m MockEntryUsecase) Create(context context.Context, user string) (*model.Entry, error) {
	return m.FakeCreate(context, user)
}

func TestEntryHandlerSearch(t *testing.T) {
	td := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	zerotd := time.Date(01, 01, 01, 0, 0, 0, 0, time.UTC)

	type args struct {
		ctx context.Context
		req *entrypb.SearchRequest
	}
	type test struct {
		name    string
		fields  model.Entry
		args    args
		err     error
		want    *entrypb.SearchResponse
		wantErr bool
	}
	tests := []test{
		{
			"success",
			model.Entry{ID: "1", User: "test", Status: "active", CreatedAt: td, DeletedAt: zerotd, UpdatedAt: td},
			args{context.TODO(), &entrypb.SearchRequest{Id: "1"}},
			nil,
			&entrypb.SearchResponse{Id: "1", User: "test", Status: "active", CreatedAt: timestamppb.New(td), DeletedAt: nil, UpdatedAt: timestamppb.New(td)},
			false,
		},
		{
			"success with deletedAt",
			model.Entry{ID: "2", User: "test", Status: "active", CreatedAt: td, DeletedAt: td, UpdatedAt: td},
			args{context.TODO(), &entrypb.SearchRequest{Id: "2"}},
			nil,
			&entrypb.SearchResponse{Id: "2", User: "test", Status: "active", CreatedAt: timestamppb.New(td), DeletedAt: timestamppb.New(td), UpdatedAt: timestamppb.New(td)},
			false,
		},
		{
			"error",
			model.Entry{},
			args{context.TODO(), &entrypb.SearchRequest{Id: "3"}},
			errors.New("usecase error"),
			nil,
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			eh := handler.NewEntryHandler(MockEntryUsecase{
				FakeSearch: func(context context.Context, id string) (*model.Entry, error) {
					return &tt.fields, tt.err
				},
			})

			got, err := eh.Search(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("entryHandler.Search() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.DeepEqual(got, tt.want) {
				t.Logf("ok")
			} else {
				t.Errorf("entryHandler.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntryHandlerCreate(t *testing.T) {
	td := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)
	zerotd := time.Date(01, 01, 01, 0, 0, 0, 0, time.UTC)

	type args struct {
		ctx context.Context
		req *entrypb.CreateRequest
	}
	type test struct {
		name    string
		fields  model.Entry
		args    args
		err     error
		want    *entrypb.CreateResponse
		wantErr bool
	}
	tests := []test{
		{
			"success",
			model.Entry{ID: "1", User: "test", Status: "active", CreatedAt: td, DeletedAt: zerotd, UpdatedAt: td},
			args{context.TODO(), &entrypb.CreateRequest{User: "test"}},
			nil,
			&entrypb.CreateResponse{Id: "1", User: "test", Status: "active", CreatedAt: timestamppb.New(td), DeletedAt: nil, UpdatedAt: timestamppb.New(td)},
			false,
		},
		{
			"error",
			model.Entry{},
			args{context.TODO(), &entrypb.CreateRequest{User: "test"}},
			errors.New("usecase error"),
			nil,
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			eh := handler.NewEntryHandler(MockEntryUsecase{
				FakeCreate: func(context context.Context, user string) (*model.Entry, error) {
					return &tt.fields, tt.err
				},
			})
			got, err := eh.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("entryHandler.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Logf("ok")
			} else {
				t.Errorf("entryHandler.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
