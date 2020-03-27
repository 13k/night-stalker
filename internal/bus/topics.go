package bus

const (
	TopicSteamEvents                                      = "steam.events"
	TopicSession                                          = "session"
	TopicSessionSteam                                     = "session.steam"
	TopicSessionDota                                      = "session.dota"
	TopicLiveMatchesReplace                               = "live_matches.replace"
	TopicLiveMatchesAdd                                   = "live_matches.add"
	TopicLiveMatchesUpdate                                = "live_matches.update"
	TopicLiveMatchesRemove                                = "live_matches.remove"
	TopicPatternLiveMatchesAll                            = "live_matches.*"
	TopicLiveMatchStatsAdd                                = "live_match_stats.add"
	TopicPatternLiveMatchStatsAll                         = "live_match_stats.*"
	TopicGCDispatcherSend                                 = "gc.dispatcher.send"
	TopicGCDispatcherReceivedMatchesMinimalResponse       = "gc.dispatcher.received.matches_minimal_response"
	TopicGCDispatcherReceivedFindTopSourceTVGamesResponse = "gc.dispatcher.received.find_top_source_tv_games_response"
	TopicWebLiveMatchesReplace                            = "web.live_matches.replace"
	TopicWebLiveMatchesAdd                                = "web.live_matches.add"
	TopicWebLiveMatchesUpdate                             = "web.live_matches.update"
	TopicWebLiveMatchesRemove                             = "web.live_matches.remove"
	TopicWebPatternLiveMatchesAll                         = "web.live_matches.*"
)
