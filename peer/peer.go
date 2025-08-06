package peer

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"github.com/nebulabox/rabbit-tcp/connection_pool"
	"github.com/nebulabox/rabbit-tcp/tunnel_pool"
	"io"
	"math/rand"
)

type Peer struct {
	peerID         uint32
	connectionPool connection_pool.ConnectionPool
	tunnelPool     tunnel_pool.TunnelPool
	ctx            context.Context
	cancel         context.CancelFunc
}

func (p *Peer) Stop() {
	p.cancel()
}

func initRand() error {
	seedSize := 8
	seedBytes := make([]byte, seedSize)
	_, err := io.ReadFull(crand.Reader, seedBytes)
	if err != nil {
		return err
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(seedBytes)))
	return nil
}
