package responder

type Matcher interface {
	Match(string) bool
}
