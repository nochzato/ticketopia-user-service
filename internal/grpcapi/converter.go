package grpcapi

import (
	db "github.com/nochzato/ticketopia-user-service/internal/db/sqlc"
	pb "github.com/nochzato/ticketopia-user-service/internal/pb/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Id:        user.ID.String(),
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
