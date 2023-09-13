package mediaPlayer

import (
	"encoding/json"

	"github.com/emmaly/go-hass/common"
)

type Features struct {
	Valid    bool
	Value    int
	Features []Feature
}

func (f *Features) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &f.Value); err != nil {
		return err
	}
	for _, v := range []Feature{
		FeaturePause,
		FeatureSeek,
		FeatureVolumeSet,
		FeatureVolumeMute,
		FeaturePreviousTrack,
		FeatureNextTrack,
		FeatureTurnOn,
		FeatureTurnOff,
		FeaturePlayMedia,
		FeatureVolumeStep,
		FeatureSelectSource,
		FeatureStop,
		FeatureClearPlaylist,
		FeaturePlay,
		FeatureShuffleSet,
		FeatureSelectSoundMode,
		FeatureBrowseMedia,
		FeatureRepeatSet,
		FeatureGrouping,
		FeatureMediaAnnounce,
		FeatureMediaEnqueue,
	} {
		if f.Value&int(v) != 0 {
			f.Features = append(f.Features, v)
		}
	}
	f.Valid = true
	return nil
}

func (f Features) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(f.Value)
}

func (f *Features) String() string {
	if !f.Valid {
		return ""
	}
	if len(f.Features) == 0 {
		return "(none)"
	}
	var s string
	for _, v := range f.Features {
		s += v.String() + " "
	}
	return s
}

type Feature int

const (
	// defined in Home Assistant source code
	// https://github.com/home-assistant/core/blob/dev/homeassistant/components/media_player/const.py#L177
	FeaturePause         Feature = 1 << iota // 1
	FeatureSeek                              // 2
	FeatureVolumeSet                         // 4
	FeatureVolumeMute                        // 8
	FeaturePreviousTrack                     // 16
	FeatureNextTrack                         // 32
	_
	_
	FeatureTurnOn          // 128
	FeatureTurnOff         // 256
	FeaturePlayMedia       // 512
	FeatureVolumeStep      // 1024
	FeatureSelectSource    // 2048
	FeatureStop            // 4096
	FeatureClearPlaylist   // 8192
	FeaturePlay            // 16384
	FeatureShuffleSet      // 32768
	FeatureSelectSoundMode // 65536
	FeatureBrowseMedia     // 131072
	FeatureRepeatSet       // 262144
	FeatureGrouping        // 524288
	FeatureMediaAnnounce   // 1048576
	FeatureMediaEnqueue    // 2097152
)

func (f Feature) String() string {
	switch f {
	case FeaturePause:
		return "pause"
	case FeatureSeek:
		return "seek"
	case FeatureVolumeSet:
		return "volume_set"
	case FeatureVolumeMute:
		return "volume_mute"
	case FeaturePreviousTrack:
		return "previous_track"
	case FeatureNextTrack:
		return "next_track"
	case FeatureTurnOn:
		return "turn_on"
	case FeatureTurnOff:
		return "turn_off"
	case FeaturePlayMedia:
		return "play_media"
	case FeatureVolumeStep:
		return "volume_step"
	case FeatureSelectSource:
		return "select_source"
	case FeatureStop:
		return "stop"
	case FeatureClearPlaylist:
		return "clear_playlist"
	case FeaturePlay:
		return "play"
	case FeatureShuffleSet:
		return "shuffle_set"
	case FeatureSelectSoundMode:
		return "select_sound_mode"
	case FeatureBrowseMedia:
		return "browse_media"
	case FeatureRepeatSet:
		return "repeat_set"
	case FeatureGrouping:
		return "grouping"
	case FeatureMediaAnnounce:
		return "media_announce"
	case FeatureMediaEnqueue:
		return "media_enqueue"
	default:
		return "unknown"
	}
}

func (f Feature) AsFeatureString() common.SupportedFeature {
	switch f {
	case FeaturePause:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.PAUSE")
	case FeatureSeek:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.SEEK")
	case FeatureVolumeSet:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.VOLUME_SET")
	case FeatureVolumeMute:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.VOLUME_MUTE")
	case FeaturePreviousTrack:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.PREVIOUS_TRACK")
	case FeatureNextTrack:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.NEXT_TRACK")
	case FeatureTurnOn:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.TURN_ON")
	case FeatureTurnOff:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.TURN_OFF")
	case FeaturePlayMedia:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.PLAY_MEDIA")
	case FeatureVolumeStep:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.VOLUME_STEP")
	case FeatureSelectSource:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.SELECT_SOURCE")
	case FeatureStop:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.STOP")
	case FeatureClearPlaylist:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.CLEAR_PLAYLIST")
	case FeaturePlay:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.PLAY")
	case FeatureShuffleSet:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.SHUFFLE_SET")
	case FeatureSelectSoundMode:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.SELECT_SOUND_MODE")
	case FeatureBrowseMedia:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.BROWSE_MEDIA")
	case FeatureRepeatSet:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.REPEAT_SET")
	case FeatureGrouping:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.GROUPING")
	case FeatureMediaAnnounce:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.MEDIA_ANNOUNCE")
	case FeatureMediaEnqueue:
		return common.SupportedFeature("media_player.MediaPlayerEntityFeature.MEDIA_ENQUEUE")
	default:
		return common.SupportedFeature("unknown")
	}
}
