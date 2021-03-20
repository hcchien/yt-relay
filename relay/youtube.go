package relay

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	ytrelay "github.com/mirror-media/yt-relay"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// YouTubeServiceV3 implements the VideoRelay interface and provides api for searching videos with youtube sdk v3
type YouTubeServiceV3 struct {
	youtubeService *youtube.Service
}

func New(apiKey string) (*YouTubeServiceV3, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("apikey is empty for youtube service")
	}
	s, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	return &YouTubeServiceV3{
		youtubeService: s,
	}, err
}

// Search supports the following parameters: part, channelId, eventType, q, maxResults, pageToken, order, safeSearch, type
func (s *YouTubeServiceV3) Search(options ytrelay.Options) (resp interface{}, err error) {
	yt := s.youtubeService
	call := yt.Search.List(strings.Split(options.Part, ","))
	if !isZero(options.ChannelID) {
		call.ChannelId(options.ChannelID)
	}
	if !isZero(options.EventType) {
		call.EventType(options.EventType)
	}
	if !isZero(options.Query) {
		call.Q(options.Query)
	}
	if !isZero(options.MaxResults) {
		call.MaxResults(options.MaxResults)
	}
	if !isZero(options.PageToken) {
		call.PageToken(options.PageToken)
	}
	if !isZero(options.Order) {
		call.Order(options.Order)
	}
	if !isZero(options.SafeSearch) {
		call.SafeSearch(options.SafeSearch)
	}
	if !isZero(options.Type) {
		call.Type(options.Type)
	}

	return call.Do()
}

// ListByVideoIDs supports the following parameters: part, id, maxResults, pageToken
func (s *YouTubeServiceV3) ListByVideoIDs(options ytrelay.Options) (resp interface{}, err error) {
	yt := s.youtubeService
	call := yt.Videos.List(strings.Split(options.Part, ","))
	if !isZero(options.IDs) {
		call.Id(strings.Split(options.IDs, ",")...)
	} else {
		return nil, fmt.Errorf("parameter \"id\" is mandantory")
	}
	if !isZero(options.PageToken) {
		call.PageToken(options.PageToken)
	}
	if !isZero(options.MaxResults) {
		call.MaxResults(options.MaxResults)
	}
	return call.Do()
}

// ListPlaylistVideos supports the following parameters: part, playlistId, maxResults, pageToken
func (s *YouTubeServiceV3) ListPlaylistVideos(options ytrelay.Options) (resp interface{}, err error) {
	yt := s.youtubeService
	call := yt.PlaylistItems.List(strings.Split(options.Part, ","))
	if !isZero(options.Fields) {
		call.PlaylistId(options.Fields)
	}
	if !isZero(options.PlaylistID) {
		call.PlaylistId(options.PlaylistID)
	}
	if !isZero(options.PageToken) {
		call.PageToken(options.PageToken)
	}
	if !isZero(options.MaxResults) {
		call.MaxResults(options.MaxResults)
	}
	return call.Do()
}

func isZero(i interface{}) bool {
	v := reflect.ValueOf(i)
	return !v.IsValid() || reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
