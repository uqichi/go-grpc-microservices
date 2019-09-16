package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	pbActivity "github.com/uqichi/go-grpc-microservices/proto/activity"
	"github.com/uqichi/go-grpc-microservices/shared/md"
)

type ActivityService struct {
	db    *sql.DB
	store Store
}

func (s *ActivityService) CreateActivity(ctx context.Context,
	req *pbActivity.CreateActivityRequest) (*empty.Empty, error) {
	log.Println(">>>", os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("ACTIVITY_SERVICE_SERVICE_HOST"), os.Getenv("HOSTNAME"))
	userID := md.GetUserIDFromContext(ctx)
	if _, err := s.store.CreateActivity(&pbActivity.Activity{
		Content:   req.Content,
		UserId:    userID,
		CreatedAt: ptypes.TimestampNow(),
	}); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *ActivityService) FindActivities(ctx context.Context,
	_ *empty.Empty) (*pbActivity.FindActivitiesResponse, error) {
	log.Println(">>>", os.Getenv("KUBERNETES_SERVICE_HOST"), os.Getenv("ACTIVITY_SERVICE_SERVICE_HOST"), os.Getenv("HOSTNAME"))
	userID := md.GetUserIDFromContext(ctx)
	activities, err := s.store.FindActivities(userID)
	if err != nil {
		return nil, err
	}
	return &pbActivity.FindActivitiesResponse{Activities: activities}, nil
}
