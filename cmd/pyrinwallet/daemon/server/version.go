package server

import (
	"context"
	"github.com/Pyrinpyi/pyipad/cmd/pyrinwallet/daemon/pb"
	"github.com/Pyrinpyi/pyipad/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
