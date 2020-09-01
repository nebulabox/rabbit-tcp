package connection

const (
	OrderedRecvQueueSize    = 32        // OrderedRecvQueue channel cap
	RecvQueueSize           = 32        // RecvQueue channel cap
	OutboundRecvBuffer      = 16 * 1024 // 16K receive buffer for Outbound Connection
	OutboundBlockTimeoutSec = 2         // Wait the period and check exit signal
	PacketWaitTimeoutSec    = 5         // If block processor is waiting for a "hole", and no packet comes within this limit, the Connection will be closed
)
