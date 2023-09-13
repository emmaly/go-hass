package mediaPlayer

import (
	"time"

	"github.com/emmaly/go-hass/common"
)

type MediaPlayer struct {
	*common.Common
	State      *State      `json:"state,omitempty"`
	Attributes *Attributes `json:"attributes,omitempty"`
}

type Attributes struct {
	*common.CommonAttributes
	DeviceClass        DeviceClass    `json:"device_class,omitempty"`
	EntityPicture      *EntityPicture `json:"entity_picture,omitempty"`
	AppID              string         `json:"app_id,omitempty"`
	AppName            string         `json:"app_name,omitempty"`
	EntityPictureLocal *EntityPicture `json:"entity_picture_local,omitempty"`
	GroupMembers       []struct {
		EntityID string `json:"entity_id,omitempty"`
	} `json:"group_members,omitempty"`
	InputSource            *InputSource      `json:"source,omitempty"`
	InputSourceList        *InputSourceList  `json:"source_list,omitempty"`
	MediaAnnounce          string            `json:"announce,omitempty"`
	MediaAlbumArtist       string            `json:"media_album_artist,omitempty"`
	MediaAlbumName         string            `json:"media_album_name,omitempty"`
	MediaArtist            string            `json:"media_artist,omitempty"`
	MediaChannel           string            `json:"media_channel,omitempty"`
	MediaContentID         string            `json:"media_content_id,omitempty"`
	MediaContentType       MediaClass        `json:"media_content_type,omitempty"`
	MediaDuration          float64           `json:"media_duration,omitempty"`
	MediaEnqueue           string            `json:"enqueue,omitempty"`
	MediaExtra             string            `json:"extra,omitempty"`
	MediaEpisode           string            `json:"media_episode,omitempty"`
	MediaPlaylist          string            `json:"media_playlist,omitempty"`
	MediaPosition          float64           `json:"media_position,omitempty"`
	MediaPositionUpdatedAt *time.Time        `json:"media_position_updated_at,omitempty"`
	MediaRepeat            RepeatMode        `json:"repeat,omitempty"`
	MediaSeason            string            `json:"season,omitempty"`
	MediaSeekPosition      float64           `json:"media_seek_position,omitempty"`
	MediaSeriesTitle       string            `json:"media_series_title,omitempty"`
	MediaShuffle           *MediaShuffle     `json:"shuffle,omitempty"`
	MediaTitle             string            `json:"media_title,omitempty"`
	MediaTrack             int               `json:"media_track,omitempty"`
	MediaVolumeLevel       *MediaVolumeLevel `json:"volume_level,omitempty"`
	MediaVolumeMuted       *MediaVolumeMuted `json:"is_volume_muted,omitempty"`
	SoundMode              string            `json:"sound_mode,omitempty"`
	SoundModeList          []string          `json:"sound_mode_list,omitempty"`
	SupportedFeatures      *Features         `json:"supported_features,omitempty"`
}
