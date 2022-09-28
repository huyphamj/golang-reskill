package handler

import (
	"context"
	"flight-management/common/proto"
	"flight-management/flight/dao"
)

type Server struct {
	proto.FlightServiceServer
	FlightService dao.IFlightService
}

func (s *Server) SearchFlight(ctx context.Context, req *proto.QueryFlightDetail) (*proto.FlightDetailResponse, error) {
	found, err := s.FlightService.GetById(int(req.FlightId))
	if err != nil {
		return nil, err
	}
	return &proto.FlightDetailResponse{
		Id:            int32(found.ID),
		Code:          found.Code,
		TotalSlot:     int32(found.TotalSlot),
		DepartureTime: found.DepartureTime.String(),
		ArrivalTime:   found.ArrivalTime.String(),
	}, nil
}
