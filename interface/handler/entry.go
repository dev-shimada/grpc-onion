package handler

import (
	"context"

	entrypb "onion/pkg/grpc"
	"onion/usecase"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type EntryHandler interface {
	Search(ctx context.Context, req *entrypb.SearchRequest) (*entrypb.SearchResponse, error)
	Create(ctx context.Context, req *entrypb.CreateRequest) (*entrypb.CreateResponse, error)
}

type entryHandler struct {
	eu usecase.EntryUsecase
}

func NewEntryHandler(eu usecase.EntryUsecase) EntryHandler {
	return &entryHandler{eu: eu}
}

func (eh *entryHandler) Search(ctx context.Context, req *entrypb.SearchRequest) (*entrypb.SearchResponse, error) {
	res, err := eh.eu.Search(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	var deletedAt *timestamppb.Timestamp
	if res.DeletedAt.IsZero() {
		deletedAt = nil
	} else {
		timestamppb.New(res.DeletedAt)
	}

	return &entrypb.SearchResponse{
		Id:        res.ID,
		User:      res.User,
		Status:    res.Status,
		CreatedAt: timestamppb.New(res.CreatedAt),
		DeletedAt: deletedAt,
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}

func (eh *entryHandler) Create(ctx context.Context, req *entrypb.CreateRequest) (*entrypb.CreateResponse, error) {
	res, err := eh.eu.Create(ctx, req.User)
	if err != nil {
		return nil, err
	}

	var deletedAt *timestamppb.Timestamp
	if res.DeletedAt.IsZero() {
		deletedAt = nil
	} else {
		timestamppb.New(res.DeletedAt)
	}

	return &entrypb.CreateResponse{
		Id:        res.ID,
		User:      res.User,
		Status:    res.Status,
		CreatedAt: timestamppb.New(res.CreatedAt),
		DeletedAt: deletedAt,
		UpdatedAt: timestamppb.New(res.UpdatedAt),
	}, nil
}
