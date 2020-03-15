package livematches

type errRedisOp struct {
	Key string
	Err error
}

func (*errRedisOp) Error() string {
	return "redis error"
}

func (err *errRedisOp) Unwrap() error {
	return err.Err
}

type errRedisPubsub struct {
	Topic string
	Err   error
}

func (*errRedisPubsub) Error() string {
	return "redis error"
}

func (err *errRedisPubsub) Unwrap() error {
	return err.Err
}
