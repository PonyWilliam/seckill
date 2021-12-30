package data

import (
	"context"
	"seckill/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)
type usersRepo struct{
	data *Data
	log *log.Helper
}

// type UserRepo interface{
// 	CreateUser(context.Context,*Users) error//注册账户
// 	UpdateUser(context.Context,*Users) error
// 	FindUserByID(context.Context,int64) (error,Users_Response)//传入id查询用户信息
// 	ListUser(context.Context) (error,[]Users_Response)//不需要传递参数
// 	DelUser(context.Context,int64)error//按照id删除
// }

func NewUserRepo(data *Data,logger log.Logger) biz.UserRepo{
	return &usersRepo{
		data:data,
		log:log.NewHelper(logger),
	}
}

func(r *usersRepo) CreateUser(ctx context.Context,u *biz.Users) error{
	//与sql进行交互
	return nil
}
func(r *usersRepo) UpdateUser(ctx context.Context,u *biz.Users_Change) error{
	return nil
}
func(r *usersRepo) FindUserByID(ctx context.Context,ID int64) (error,[]biz.Users_Response){
	return nil,nil
}
func(r *usersRepo) ListUser(ctx context.Context) (error,[]biz.Users_Response){
	return nil,nil
}
func(r *usersRepo) DelUser(ctx context.Context,ID int64) error{
	return nil
}