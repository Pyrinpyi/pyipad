package appmessage

// MsgUnexpectedPruningPoint represents a pyrin UnexpectedPruningPoint message
type MsgUnexpectedPruningPoint struct {
	baseMessage
}

// Command returns the protocol command string for the message
func (msg *MsgUnexpectedPruningPoint) Command() MessageCommand {
	return CmdUnexpectedPruningPoint
}

// NewMsgUnexpectedPruningPoint returns a new pyrin UnexpectedPruningPoint message
func NewMsgUnexpectedPruningPoint() *MsgUnexpectedPruningPoint {
	return &MsgUnexpectedPruningPoint{}
}
