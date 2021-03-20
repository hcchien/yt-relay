package ytrelay

// Options are used to store the supported parsed queries and passed to VideoRelay service
type Options struct {
	ChannelID  string `form:"channelId"`  // For YouTube
	EventType  string `form:"eventType"`  // For YouTube
	Fields     string `form:"fields"`     // For YouTube
	IDs        string `form:"id"`         // For YouTube
	MaxResults int64  `form:"maxResults"` // For YouTube
	Order      string `form:"order"`      // For YouTube
	PageToken  string `form:"pageToken"`  // For YouTube
	Part       string `form:"part"`       // For YouTube
	PlaylistID string `form:"playlistId"` // For YouTube
	Query      string `form:"q"`          // For YouTube
	SafeSearch string `form:"safeSearch"` // For YouTube
	Type       string `form:"type"`       // For YouTube
}

// VideoRelay is responsible to bypass the api request to the video service
type VideoRelay interface {
	Search(options Options) (resp interface{}, err error)
	ListByVideoIDs(options Options) (resp interface{}, err error)
	ListPlaylistVideos(options Options) (resp interface{}, err error)
}

// APIWhitelist is responsible to validate some options to prevent abusive requests
type APIWhitelist interface {
	// ValidateParameters(options Options) bool
	ValidateChannelID(channelID string) bool
	ValidatePlaylistIDs(playlistID string) bool
}
