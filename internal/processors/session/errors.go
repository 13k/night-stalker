package session

type ErrInvalidServerAddress struct {
	Address string
}

func (*ErrInvalidServerAddress) Error() string {
	return "invalid server address"
}
