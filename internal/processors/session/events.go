package session

type SteamWebSessionIDEvent struct {
	SessionID string
}

type SteamWebLoggedOnEvent struct {
	AuthToken  string
	AuthSecret string
}
