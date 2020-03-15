package matchdetails

type errMatchSave struct {
	MatchID uint64
	Err     error
}

func (*errMatchSave) Error() string {
	return "match save error"
}

func (err *errMatchSave) Unwrap() error {
	return err.Err
}
