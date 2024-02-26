package frame

// The AckMode type is an enumeration of the acknowledgement modes for a
// STOMP subscription.
type AckMode int

const (
	// AckAuto No acknowledgement is required, the server assumes that the client
	// received the message.
	AckAuto AckMode = iota

	// AckClient Client acknowledges messages. When a client acknowledges a message,
	// any previously received messages are also acknowledged.
	// Client sends ACK/NACK
	AckClient

	// AckClientIndividual Client acknowledges message. Each message is acknowledged individually.
	// Client sends ACK/NACK for individual messages
	AckClientIndividual
)

const (
	ackAuto             = "auto"
	ackClient           = "client"
	ackClientIndividual = "client-individual"
)

func ParseAckMode(value string) AckMode {
	switch value {
	case ackClient:
		return AckClient
	case ackClientIndividual:
		return AckClientIndividual
	default:
		return AckAuto
	}
}

// String returns the string representation of the AckMode value.
func (a AckMode) String() string {
	switch a {
	case AckClient:
		return ackClient
	case AckClientIndividual:
		return ackClientIndividual
	default:
		return ackAuto
	}
}

// ShouldAck returns true if this AckMode is an acknowledgement
// mode which requires acknowledgement. Returns true for all values
// except AckAuto, which returns false.
func (a AckMode) ShouldAck() bool {
	switch a {
	case AckAuto:
		return false
	case AckClient, AckClientIndividual:
		return true
	}
	panic("invalid AckMode value")
}
