package service

import (
	"context"

	pb "github.com/tonytheleg/resource-api/api/relationships/v1"
	"github.com/tonytheleg/resource-api/internal/biz"
)

type KesselRelationshipServiceService struct {
	pb.UnimplementedObjectSubjectRelationshipServiceServer
	uc *biz.RelationshipUsecase
}

func NewKesselRelationshipServiceService(uc *biz.RelationshipUsecase) *KesselRelationshipServiceService {
	return &KesselRelationshipServiceService{uc: uc}
}

func (s *KesselRelationshipServiceService) CreateRelationship(ctx context.Context, req *pb.CreateObjectSubjectRelationshipRequest) (*pb.CreateObjectSubjectRelationshipResponse, error) {
	return &pb.CreateObjectSubjectRelationshipResponse{}, nil
}
func (s *KesselRelationshipServiceService) UpdateRelationship(ctx context.Context, req *pb.UpdateObjectSubjectRelationshipRequest) (*pb.UpdateObjectSubjectRelationshipResponse, error) {
	return &pb.UpdateObjectSubjectRelationshipResponse{}, nil
}
func (s *KesselRelationshipServiceService) DeleteRelationship(ctx context.Context, req *pb.DeleteObjectSubjectRelationshipRequest) (*pb.DeleteObjectSubjectRelationshipResponse, error) {
	return &pb.DeleteObjectSubjectRelationshipResponse{}, nil
}
