package mediaPlayer

// MediaType is the Content Type of the media, also known as media_content_type
type MediaType string

const (
	MediaTypeAlbum              MediaType = "album"
	MediaTypeApp                MediaType = "app"
	MediaTypeApps               MediaType = "apps"
	MediaTypeArtist             MediaType = "artist"
	MediaTypeChannel            MediaType = "channel"
	MediaTypeChannels           MediaType = "channels"
	MediaTypeComposer           MediaType = "composer"
	MediaTypeContributingArtist MediaType = "contributing_artist"
	MediaTypeEpisode            MediaType = "episode"
	MediaTypeGame               MediaType = "game"
	MediaTypeGenre              MediaType = "genre"
	MediaTypeImage              MediaType = "image"
	MediaTypeMovie              MediaType = "movie"
	MediaTypeMusic              MediaType = "music"
	MediaTypePlaylist           MediaType = "playlist"
	MediaTypePodcast            MediaType = "podcast"
	MediaTypeSeason             MediaType = "season"
	MediaTypeTrack              MediaType = "track"
	MediaTypeTvShow             MediaType = "tvshow"
	MediaTypeUrl                MediaType = "url"
	MediaTypeVideo              MediaType = "video"
)

func (t MediaType) String() string {
	switch t {
	case MediaTypeAlbum:
		return "Album"
	case MediaTypeApp:
		return "App"
	case MediaTypeApps:
		return "Apps"
	case MediaTypeArtist:
		return "Artist"
	case MediaTypeChannel:
		return "Channel"
	case MediaTypeChannels:
		return "Channels"
	case MediaTypeComposer:
		return "Composer"
	case MediaTypeContributingArtist:
		return "Contributing Artist"
	case MediaTypeEpisode:
		return "Episode"
	case MediaTypeGame:
		return "Game"
	case MediaTypeGenre:
		return "Genre"
	case MediaTypeImage:
		return "Image"
	case MediaTypeMovie:
		return "Movie"
	case MediaTypeMusic:
		return "Music"
	case MediaTypePlaylist:
		return "Playlist"
	case MediaTypePodcast:
		return "Podcast"
	case MediaTypeSeason:
		return "Season"
	case MediaTypeTrack:
		return "Track"
	case MediaTypeTvShow:
		return "TV Show"
	case MediaTypeUrl:
		return "URL"
	case MediaTypeVideo:
		return "Video"
	default:
		return "Unknown"
	}
}
