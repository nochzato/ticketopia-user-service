package grpcapi

import (
	"context"

	db "github.com/nochzato/ticketopia-user-service/internal/db/sqlc"
	"github.com/nochzato/ticketopia-user-service/pkg/hashpass"
	pb "github.com/nochzato/ticketopia-user-service/pkg/pb/user/v1"
	"github.com/nochzato/ticketopia-user-service/pkg/validator"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := hashpass.Hash(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash the password: %s", err)
	}

	arg := db.CreateUserParams{
		Username: req.GetUsername(),
		FullName: req.GetFullName(),
		Password: hashedPassword,
		Email:    req.GetEmail(),
	}

	user, err := s.queries.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "username or email already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create the user: %s", err)
	}

	res := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return res, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validator.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if err := validator.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	if err := validator.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}
	if err := validator.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	return violations
}
