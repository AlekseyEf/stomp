package client

import (
	"errors"
	"fmt"
)

var (
	ErrNotConnected             = errors.New("expected CONNECT or STOMP frame")
	ErrUnexpectedCommand        = errors.New("unexpected frame command")
	ErrUnknownCommand           = errors.New("unknown command")
	ErrReceiptInConnect         = errors.New("receipt header prohibited in CONNECT or STOMP frame")
	ErrAuthenticationFailed     = errors.New("authentication failed")
	ErrTxAlreadyInProgress      = errors.New("transaction already in progress")
	ErrTxUnknown                = errors.New("unknown transaction")
	ErrUnsupportedVersion       = errors.New("unsupported version")
	ErrSubscriptionExists       = errors.New("subscription already exists")
	ErrSubscriptionNotFound     = errors.New("subscription not found")
	ErrInvalidFrameFormat       = errors.New("invalid frame format")
	ErrInvalidCommand           = errors.New("invalid command")
	ErrUnknownVersion           = errors.New("incompatible version")
	ErrNotConnectFrame          = errors.New("operation valid for STOMP and CONNECT frames only")
	ErrInvalidHeartBeat         = errors.New("invalid format for heart-beat")
	ErrInvalidOperationForFrame = errors.New("invalid operation for frame")
	ErrExceededMaxFrameSize     = errors.New("exceeded max frame size")
	ErrInvalidHeaderValue       = errors.New("invalid header value")
)

func MissingHeaderError(name string) error {
	return fmt.Errorf("missing header: " + name)
}
