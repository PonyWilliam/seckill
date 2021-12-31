package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)
type Admin struct{
	Username string;
	Password string;
	Permission int64
}
type Admin_Change struct{
	ID int64
	Username string;
	Password string;
	Permission int64
}
type Admin_Response struct{
	ID int64;
	Username string;
	Permission int64
}
type AdminRepo interface{
	CreateAdmin(context.Context,*Admin) error
	UpdateAdmin(context.Context,*Admin_Change) error
	DelAdmin(context.Context,int64) error
	ListAdmin(context.Context) (error,[]Admin_Response)//根据ID查找并列举
	FindAdminByID(context.Context,int64)(error,[]Admin_Response)//方便返空
}

type AdminUsecase struct{
	repo AdminRepo//依赖倒置
	log *log.Helper
}

func NewAdminUsecase(repo AdminRepo,logger log.Logger) *AdminUsecase{
	return &AdminUsecase{repo: repo,log: log.NewHelper(logger)}
}

func(uc *AdminUsecase)CreateAdmin(ctx context.Context,admin *Admin) error{
	return uc.repo.CreateAdmin(ctx,admin)
}
func(uc *AdminUsecase)UpdateAdmin(ctx context.Context,admin *Admin_Change) error{
	return uc.repo.UpdateAdmin(ctx,admin)
}
func(uc *AdminUsecase)DelAdmin(ctx context.Context,id int64) error{
	return uc.repo.DelAdmin(ctx,id)
}
func(uc *AdminUsecase)ListAdmin(ctx context.Context) (error,[]Admin_Response){
	return uc.repo.ListAdmin(ctx)
}
func(uc *AdminUsecase)FindAdminByID(ctx context.Context,id int64)(error,[]Admin_Response){
	return uc.repo.FindAdminByID(ctx,id)
}
