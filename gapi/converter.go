package gapi

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	db "github.com/LuigiAzevedo/simplebank/db/sqlc"
	"github.com/LuigiAzevedo/simplebank/pb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
