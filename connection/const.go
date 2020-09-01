package connection

const (
	OrderedRecvQueueSize    = 24 * 8        // OrderedRecvQueue channel cap
	RecvQueueSize           = 24 * 8        // RecvQueue channel cap
	OutboundRecvBuffer      = 16 * 1024 * 8 // 16K receive buffer for Outbound Connection
	OutboundBlockTimeoutSec = 1             // Wait the period and check exit signal
	PacketWaitTimeoutSec    = 3             // If block processor is waiting for a "hole", and no packet comes within this limit, the Connection will be closed
)
