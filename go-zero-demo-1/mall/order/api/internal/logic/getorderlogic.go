package logic

import (
	"context"
	"errors"

	"go-zero-demo/mall/order/api/internal/svc"
	"go-zero-demo/mall/order/api/internal/types"
	"go-zero-demo/mall/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: "1",
	})
	if err != nil {
		return nil, err
	}

	if user.Name != "zero" {
		return nil, errors.New("user name is not zero")
	}
	return &types.OrderReply{
		Id:   req.Id,
		Name: "test order",
	}, nil
}
