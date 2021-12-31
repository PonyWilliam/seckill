package data

import (
	"context"
	"errors"
	"seckill/common"
	"seckill/internal/biz"
	"strconv"

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
type Admin struct{
	Id int `db:"id"`
	User string `db:"user"`
	Password string `db:"password"`
	Permission int `db:"perm"`
}
var results []Admin
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
	if admin.Password != "" && admin.Permission != -1{
		//一起插入
		_,err := r.data.Db.Exec("update admin set pwd=?,perm=?",common.EncodePassword(admin.Password),admin.Permission)
		if err != nil{
			r.log.Error(err)
			return err
		}
		return nil
	}
	conn,err := r.data.Db.Db.Begin()
	if err != nil{
		r.log.Error(err)
		return err
	}
	if admin.Password != ""{
		_,err := r.data.Db.Exec("update admin set pwd=? where id=?",common.EncodePassword(admin.Password),admin.ID)
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
	conn.Commit()//没有错误，原子性保持完整，直接写入数据库
	return nil
}
func(r *adminRepo)DelAdmin(ctx context.Context,id int64) error{
	_,err := r.data.Db.Del("admin","id",strconv.FormatInt(id,10))
	if err != nil{
		r.log.Error(err)
		return err
	}
	return nil
}
func(r *adminRepo)ListAdmin(ctx context.Context) (error,[]biz.Admin_Response){
	err := r.data.Db.Select(&results,"select * from admin")
	if err != nil{
		r.log.Info(err)
		return err,nil
	}
	temp := make([]biz.Admin_Response,len(results))
	i := 0
	for _,v := range results{
		reverse_Admin(&v,&temp[i])
		i+=1
	}
	return nil,temp
}
func(r *adminRepo)FindAdminByID(ctx context.Context,id int64)(error,[]biz.Admin_Response){
	res := &Admin{}
	err := r.data.Db.Select(&res,"select * from admin where id = ? limit 1",id)
	if err != nil{
		r.log.Error(err)
		return err,nil
	}
	temp := make([]biz.Admin_Response,1)
	reverse_Admin(res,&temp[0])
	return nil,temp
}

func reverse_Admin(src *Admin,dest *biz.Admin_Response){
	dest.ID = int64(src.Id)
	dest.Permission = int64(src.Permission)
	dest.Username = src.User
}