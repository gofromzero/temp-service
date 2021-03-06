// Code generated by goctl. DO NOT EDIT!
// Source: admin.proto

package server

import (
	"context"

	"temp-service/apps/user/admin/admin"
	"temp-service/apps/user/admin/internal/logic"
	"temp-service/apps/user/admin/internal/svc"
)

type AdminServer struct {
	svcCtx *svc.ServiceContext
	admin.UnimplementedAdminServer
}

func NewAdminServer(svcCtx *svc.ServiceContext) *AdminServer {
	return &AdminServer{
		svcCtx: svcCtx,
	}
}

func (s *AdminServer) Ping(ctx context.Context, in *admin.Request) (*admin.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
