package data

import (
	"context"
	"errors"
	"seckill/common"
	"seckill/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// type AdminRepo interface{
// 	CreateAdmin(context.Context,*Admin) error
// 	UpdateAdmin(context.Context,*Admin_Change) error
// 	DelAdmin(context.Context,int64) error
// 	ListAdmin(context.Context) (error,[]Admin_Response)//根据ID查找并列举
// 	FindAdminByID(context.Context,int64)(error,[]Admin_Response)//方便返空
// }


type adminRepo struct{
	data *Data
	log *log.Helper
}

func NewAdminRepo(data *Data,logger log.Logger)biz.AdminRepo{
	return &adminRepo{
		data:data,
		log:log.NewHelper(logger),
	}
}

//与sql交互，service调用这一层实现
func(r *adminRepo) CreateAdmin(ctx context.Context,admin *biz.Admin) error{
	key := []string{"user","pwd"}
	val := []string{admin.Username,common.EncodePassword(admin.Password)}
	err := r.data.Db.Insert("admin",key,val)
	return err
}
func(r *adminRepo)UpdateAdmin(ctx context.Context,admin *biz.Admin_Change) error{
	err,res := r.FindAdminByID(ctx,admin.ID)
	if err != nil{
		r.log.Error(err)
		return err
	}
	if len(res) != 1{
		r.log.Error("查询账户出错")
		return errors.New("查询账户出错")
	}
	//Username不允许修改,使用事务进行修改
	conn,err := r.data.Db.Db.Begin()
	if err != nil{
		r.log.Error(err)
		return err
	}
	if admin.Password != ""{
		_,err := r.data.Db.Exec("update admin set pwd=? where id=?",admin.Password,admin.ID)
		if err != nil{
			conn.Rollback()
			r.log.Error(err)
			return err
		}
	}
	if admin.Permission != -1{
		_,err := r.data.Db.Exec("update admin set perm=? where id=?",admin.Password,admin.ID)
		if err != nil{
			conn.Rollback()
			r.log.Error(err)
			return err
		}
	}
	conn.Commit()
	return nil
}
func(r *adminRepo)DelAdmin(ctx context.Context,id int64) error{
	return nil
}
func(r *adminRepo)ListAdmin(ctx context.Context) (error,[]biz.Admin_Response){
	return nil,nil
}
func(r *adminRepo)FindAdminByID(ctx context.Context,id int64)(error,[]biz.Admin_Response){
	return nil,nil
}