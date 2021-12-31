package service

import (
	"context"

	pb "seckill/api/admin/v1"
	"seckill/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	uc *biz.AdminUsecase//抽象接口
	log *log.Helper
}

func NewAdminService(uc *biz.AdminUsecase,logger log.Logger) *AdminService {
	return &AdminService{uc:uc,log:log.NewHelper(logger)}
}

func (s *AdminService) CreateAdmin(ctx context.Context, req *pb.CreateAdminRequest) (*pb.CreateAdminReply, error) {
	Admin := &biz.Admin{
		Username: req.User,
		Password: req.Pwd,
		Permission: req.Permission,
	}
	err := s.uc.CreateAdmin(ctx,Admin)
	reply := &pb.CommonReply{}
	if err != nil{
		reply.Code = 500
		reply.Msg = err.Error()
		return &pb.CreateAdminReply{Reply:reply},err
	}
	reply.Code = 200
	reply.Msg = "success"
	return &pb.CreateAdminReply{Reply: reply}, err
}
func (s *AdminService) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminRequest) (*pb.UpdateAdminReply, error) {
	Admin := &biz.Admin_Change{
		ID:req.Id,
		Password: req.Pwd,
		Permission: req.Permission,
	}
	err := s.uc.UpdateAdmin(ctx,Admin)
	if err != nil{
		
	}
	reply := &pb.CommonReply{}
	if err != nil{
		reply.Code = 500
		reply.Msg = err.Error()
		return &pb.UpdateAdminReply{Reply:reply},err
	}
	reply.Code = 200
	reply.Msg = "success"
	return &pb.UpdateAdminReply{Reply: reply}, err
}
func (s *AdminService) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminRequest) (*pb.DeleteAdminReply, error) {
	return &pb.DeleteAdminReply{}, nil
}
func (s *AdminService) GetAdmin(ctx context.Context, req *pb.GetAdminRequest) (*pb.GetAdminReply, error) {
	return &pb.GetAdminReply{}, nil
}
func (s *AdminService) ListAdmin(ctx context.Context, req *pb.ListAdminRequest) (*pb.ListAdminReply, error) {
	return &pb.ListAdminReply{}, nil
}
