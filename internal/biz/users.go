package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)
type Users struct{
	username string;
	password string;
	id_card int64;
	phone string;
	sex int64;
	age int64;
	true_name string;
	carrer int64;
}
type Users_Change struct{
	username string;
	password string;
	id_card int64;
	phone string;
	sex int64;
	age int64;
	true_name string;
	carrer int64;
}
type Users_Response struct{
	id int64;
	username string;
	password string;
	id_card int64;
	phone string;
	sex int64;
	age int64;
	true_name string;
	carrer int64;
}
type UserRepo interface{
	CreateUser(context.Context,*Users) error//注册账户
	UpdateUser(context.Context,*Users_Change) error
	FindUserByID(context.Context,int64) (error,[]Users_Response)//传入id查询用户信息
	ListUser(context.Context) (error,[]Users_Response)//不需要传递参数
	DelUser(context.Context,int64)error//按照id删除
}

type UserUsecase struct{
	repo UserRepo//依赖倒置
	log *log.Helper
}

func NewUserUsecase(repo UserRepo,logger log.Logger) *UserUsecase{
	return &UserUsecase{repo: repo,log: log.NewHelper(logger)}
}

func(uc *UserUsecase) Create(ctx context.Context,u *Users) error{
	return uc.repo.CreateUser(ctx,u)
}
func(uc *UserUsecase) Del(ctx context.Context,ID int64) error{
	return uc.repo.DelUser(ctx,ID)
}
func(uc *UserUsecase) UpdateUser(ctx context.Context,u *Users_Change) error{
	return uc.repo.UpdateUser(ctx,u)
}
func(uc *UserUsecase) ListUser(ctx context.Context)(error,[]Users_Response){
	return uc.repo.ListUser(ctx)
}
func(uc *UserUsecase) FindUser(ctx context.Context,ID int64)(error,[]Users_Response){
	return uc.repo.FindUserByID(ctx,ID)
}