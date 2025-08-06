package server

import (
	"github.com/nebulabox/rabbit-tcp/logger"
	"github.com/nebulabox/rabbit-tcp/peer"
	"github.com/nebulabox/rabbit-tcp/tunnel"
	"net"
)

type Server struct {
	peerGroup peer.PeerGroup
	logger    *logger.Logger
}

func NewServer(cipher tunnel.Cipher) Server {
	return Server{
		peerGroup: peer.NewPeerGroup(cipher),
		logger:    logger.NewLogger("[Server]"),
	}
}

func (s *Server) Serve(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			s.logger.Errorf("Error when accept connection: %v.\n", err)
			continue
		}
		err = s.peerGroup.AddTunnelFromConn(conn)
		if err != nil {
			s.logger.Errorf("Error when add tunnel to tunnel pool: %v.\n", err)
		}
	}
}
