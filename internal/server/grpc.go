package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/programzheng/rent-house-crawler/config"
	"github.com/programzheng/rent-house-crawler/ent"
	"github.com/programzheng/rent-house-crawler/internal"
	pb "github.com/programzheng/rent-house-crawler/pkg/proto/rent-house"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProxyServer
}

type RentHome interface {
	[]*ent.Rent591Home
}

func (s *server) GetRentHousesByConditions(ctx context.Context, req *pb.GetRentHousesByConditionsRequest) (*pb.GetRentHousesResponse, error) {
	r5hs, err := internal.GetNewRent591HomesByCity(ctx, &req.City)
	if err != nil {
		return &pb.GetRentHousesResponse{
			Status:        "fail",
			StatusMessage: err.Error(),
		}, nil
	}
	pbrhs, err := covertEntRentHomesToResponseRentHomes[[]*ent.Rent591Home](r5hs)
	if err != nil {
		return &pb.GetRentHousesResponse{
			Status:        "fail",
			StatusMessage: err.Error(),
		}, nil
	}
	return &pb.GetRentHousesResponse{
		Status:        "success",
		StatusMessage: "",
		RentHouses:    pbrhs,
	}, nil
}

func (s *server) GetRentHousesByTimestamp(ctx context.Context, req *pb.GetRentHousesByTimestampRequest) (*pb.GetRentHousesResponse, error) {
	startTimestamp := req.StartTimestamp.AsTime().Local()
	endTimestamp := req.EndTimestamp.AsTime().Local()

	r5hs, err := internal.GetRent591HomesFilterCreatedAtByStartTimestampAndEndTimestamp(ctx, &startTimestamp, &endTimestamp)
	if err != nil {
		return &pb.GetRentHousesResponse{
			Status:        "fail",
			StatusMessage: err.Error(),
		}, nil
	}
	pbrhs, err := covertEntRentHomesToResponseRentHomes[[]*ent.Rent591Home](r5hs)
	if err != nil {
		return &pb.GetRentHousesResponse{
			Status:        "fail",
			StatusMessage: err.Error(),
		}, nil
	}
	return &pb.GetRentHousesResponse{
		Status:        "success",
		StatusMessage: "",
		RentHouses:    pbrhs,
	}, nil
}

func covertEntRentHomesToResponseRentHomes[T RentHome](rentHomes T) ([]*pb.RentHouse, error) {
	r5hs := rentHomes
	pbrhs := make([]*pb.RentHouse, 0, len(r5hs))
	for _, r5h := range r5hs {
		pbrhs = append(pbrhs, &pb.RentHouse{
			Title:            r5h.Title,
			Type:             int64(r5h.Type),
			PostId:           int64(r5h.PostID),
			KindName:         r5h.KindName,
			RoomStr:          r5h.RoomStr,
			FloorStr:         r5h.FloorStr,
			Community:        r5h.Community,
			Price:            int64(r5h.Price),
			PriceUnit:        r5h.PriceUnit,
			PhotoList:        r5h.PhotoList,
			RegionName:       r5h.RegionName,
			SectionName:      r5h.SectionName,
			StreetName:       r5h.StreetName,
			Location:         r5h.Location,
			Area:             r5h.Area,
			RoleName:         r5h.RoleName,
			Contact:          r5h.Contact,
			RefreshTime:      r5h.RefreshTime,
			YesterdayHit:     int64(r5h.YesterdayHit),
			IsVip:            int32(r5h.IsVip),
			IsCombine:        int32(r5h.IsCombine),
			Hurry:            int32(r5h.Hurry),
			IsSocial:         int32(r5h.IsSocial),
			DiscountPriceStr: r5h.DiscountPriceStr,
			CasesId:          int64(r5h.CasesID),
			IsVideo:          int32(r5h.IsVideo),
			Preferred:        int32(r5h.Preferred),
			Cid:              int64(r5h.Cid),
			DetailUrl:        fmt.Sprintf("%s/%d", config.Cfg.GetString("crawlers.591.detail-url"), r5h.PostID),
		})
	}
	return pbrhs, nil
}

func RunGrpcServer(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProxyServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
