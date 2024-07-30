package router

import (
	"context"
	"onion/interface/handler"
	entrypb "onion/proto/entry"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type EntryServer struct {
	entrypb.UnimplementedEntryServiceServer
	eh handler.EntryHandler
}

func NewEntryServer(eh handler.EntryHandler) *EntryServer {
	return &EntryServer{eh: eh}
}

func Register(s *grpc.Server, e *EntryServer) {
	// サービスの登録
	entrypb.RegisterEntryServiceServer(s, e)
	// サービスのリフレクション登録
	reflection.Register(s)
}

func (es *EntryServer) Search(ctx context.Context, req *entrypb.SearchRequest) (*entrypb.SearchResponse, error) {
	return es.eh.Search(ctx, req)
}

func (es *EntryServer) Create(ctx context.Context, req *entrypb.CreateRequest) (*entrypb.CreateResponse, error) {
	return es.eh.Create(ctx, req)
}
