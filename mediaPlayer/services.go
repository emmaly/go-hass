package mediaPlayer

import "github.com/emmaly/go-hass/common"

var (
	// defined in Home Assistant source code
	// at https://github.com/home-assistant/core/blob/dev/homeassistant/components/media_player/strings.json
	ServiceTurnOn = common.Service{
		ServiceID:   "turn_on",
		Name:        "Turn on",
		Description: "Turns on the power of the media player.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureTurnOn.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceTurnOff = common.Service{
		ServiceID:   "turn_off",
		Name:        "Turn off",
		Description: "Turns off the power of the media player.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureTurnOff.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceToggle = common.Service{
		ServiceID:   "toggle",
		Name:        "Toggle",
		Description: "Toggles a media player on/off.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureTurnOn.AsFeatureString(),
							FeatureTurnOff.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceVolumeUp = common.Service{
		ServiceID:   "volume_up",
		Name:        "Turn up volume",
		Description: "Turns up the volume.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureVolumeSet.AsFeatureString(),
							FeatureVolumeStep.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceVolumeDown = common.Service{
		ServiceID:   "volume_down",
		Name:        "Turn down volume",
		Description: "Turns down the volume.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureVolumeSet.AsFeatureString(),
							FeatureVolumeStep.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceVolumeMute = common.Service{
		ServiceID:   "volume_mute",
		Name:        "Mute/unmute volume",
		Description: "Mutes or unmutes the media player.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureVolumeMute.AsFeatureString(),
						},
					},
				},
			},
		},
		Fields: []common.ServiceField{
			{
				FieldID:     "is_volume_muted",
				Name:        "Muted",
				Description: "Defines whether or not it is muted.",
				Required:    true,
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeBool,
					},
				},
			},
		},
	}
	ServiceVolumeSet = common.Service{
		ServiceID:   "volume_set",
		Name:        "Set volume",
		Description: "Sets the volume level.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureVolumeSet.AsFeatureString(),
						},
					},
				},
			},
		},
		Fields: []common.ServiceField{
			{
				FieldID:     "volume_level",
				Name:        "Level",
				Description: "The volume. 0 is inaudible, 1 is the maximum volume.",
				Required:    true,
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeNumber,
						Number: &common.NumberSelector{
							Min:  0,
							Max:  1,
							Step: 0.01,
						},
					},
				},
			},
		},
	}
	ServiceMediaPlayPause = common.Service{
		ServiceID:   "media_play_pause",
		Name:        "Play/Pause",
		Description: "Toggles play/pause.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeaturePlay.AsFeatureString(),
							FeaturePause.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceMediaPlay = common.Service{
		ServiceID:   "media_play",
		Name:        "Play",
		Description: "Starts playing.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeaturePlay.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceMediaPause = common.Service{
		ServiceID:   "media_pause",
		Name:        "Pause",
		Description: "Pauses.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeaturePause.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceMediaStop = common.Service{
		ServiceID:   "media_stop",
		Name:        "Stop",
		Description: "Stops playing.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureStop.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceMediaNextTrack = common.Service{
		ServiceID:   "media_next_track",
		Name:        "Next",
		Description: "Selects the next track.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureNextTrack.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceMediaPreviousTrack = common.Service{
		ServiceID:   "media_previous_track",
		Name:        "Previous",
		Description: "Selects the previous track.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeaturePreviousTrack.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceMediaSeek = common.Service{
		ServiceID:   "media_seek",
		Name:        "Seek",
		Description: "Allows you to go to a different part of the media that is currently playing.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureSeek.AsFeatureString(),
						},
					},
				},
			},
		},
		Fields: []common.ServiceField{
			{
				FieldID:     "seek_position",
				Name:        "Position",
				Description: "Target position in the currently playing media. The format is platform dependent.",
				Required:    true,
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeNumber,
						Number: &common.NumberSelector{
							Min:  0,
							Max:  9223372036854775807,
							Step: 0.01,
							Mode: common.NumberSelectorModeBox,
						},
					},
				},
			},
		},
	}
	ServicePlayMedia = common.Service{
		ServiceID:   "play_media",
		Name:        "Play media",
		Description: "Starts playing specified media.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeaturePlayMedia.AsFeatureString(),
						},
					},
				},
			},
		},
		Fields: []common.ServiceField{
			{
				FieldID:     "media_content_id",
				Name:        "Content ID",
				Description: "The ID of the content to play. Platform dependent.",
				Required:    true,
				Example:     "https://home-assistant.io/images/cast/splash.png",
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeText,
					},
				},
			},
			{
				FieldID:     "media_content_type",
				Name:        "Content type",
				Description: "The type of the content to play. Such as image, music, tv show, video, episode, channel, or playlist.",
				Required:    true,
				Example:     MediaTypeMusic,
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeText,
					},
				},
			},
			{
				FieldID:     "enqueue",
				Name:        "Enqueue",
				Description: "If the content should be played now or be added to the queue.",
				Required:    false,
				Filter: &common.ServiceFieldFilter{
					SupportedFeatures: []common.SupportedFeature{
						FeatureMediaEnqueue.AsFeatureString(),
					},
				},
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeSelect,
						Select: &common.SelectSelector{
							Options: []string{
								"play",
								"next",
								"add",
								"replace",
							},
							TranslationKey: "enqueue",
						},
					},
				},
			},
			{
				FieldID:     "announce",
				Name:        "Announce",
				Description: "If the media should be played as an announcement.",
				Required:    false,
				Example:     true,
				Filter: &common.ServiceFieldFilter{
					SupportedFeatures: []common.SupportedFeature{
						FeatureMediaAnnounce.AsFeatureString(),
					},
				},
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeBool,
					},
				},
			},
		},
	}
	ServiceSelectSource = common.Service{
		ServiceID:   "select_source",
		Name:        "Select source",
		Description: "Sends the media player the command to change input source.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureSelectSource.AsFeatureString(),
						},
					},
				},
			},
		},
		Fields: []common.ServiceField{
			{
				FieldID:     "source",
				Name:        "Source",
				Description: "Name of the source to switch to. Platform dependent.",
				Required:    true,
				Example:     "video1",
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeText,
					},
				},
			},
		},
	}
	ServiceSelectSoundMode = common.Service{
		ServiceID:   "select_sound_mode",
		Name:        "Select sound mode",
		Description: "Selects a specific sound mode.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureSelectSoundMode.AsFeatureString(),
						},
					},
				},
			},
		},
		Fields: []common.ServiceField{
			{
				FieldID:     "sound_mode",
				Name:        "Sound mode",
				Description: "Name of the sound mode to switch to.",
				Example:     "Music",
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeText,
					},
				},
			},
		},
	}
	ServiceClearPlaylist = common.Service{
		ServiceID:   "clear_playlist",
		Name:        "Clear playlist",
		Description: "Clears the playlist.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureClearPlaylist.AsFeatureString(),
						},
					},
				},
			},
		},
	}
	ServiceShuffleSet = common.Service{
		ServiceID:   "shuffle_set",
		Name:        "Shuffle",
		Description: "Playback mode that selects the media in randomized order.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureShuffleSet.AsFeatureString(),
						},
					},
				},
			},
		},
		Fields: []common.ServiceField{
			{
				FieldID:     "shuffle",
				Name:        "Shuffle",
				Description: "Whether or not shuffle mode is enabled.",
				Required:    true,
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeBool,
					},
				},
			},
		},
	}
	ServiceRepeatSet = common.Service{
		ServiceID:   "repeat_set",
		Name:        "Repeat",
		Description: "Playback mode that plays the media in a loop.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureRepeatSet.AsFeatureString(),
						},
					},
				},
			},
		},
		Fields: []common.ServiceField{
			{
				FieldID:     "repeat",
				Name:        "Repeat mode",
				Description: "Repeat mode to set.",
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeSelect,
						Select: &common.SelectSelector{
							Options: []string{
								"off",
								"all",
								"one",
							},
							TranslationKey: "repeat",
						},
					},
				},
			},
		},
	}
	ServiceJoin = common.Service{
		ServiceID:   "join",
		Name:        "Join",
		Description: "Groups media players together for synchronous playback. Only works on supported multiroom audio systems.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureGrouping.AsFeatureString(),
						},
					},
				},
			},
		},
		Fields: []common.ServiceField{
			{
				FieldID:     "group_members",
				Name:        "Group members",
				Description: "The players which will be synced with the playback specified in `target`.",
				Required:    true,
				Example: []string{
					"media_player.multiroom_player2",
					"media_player.multiroom_player3",
				},
				Selector: []common.Selector{
					{
						Type: common.SelectorTypeEntity,
						Entity: &common.EntitySelector{
							Multiple: true,
							Filter: []common.SelectorEntityFilter{
								{
									Domain: "media_player",
								},
							},
						},
					},
				},
			},
		},
	}
	ServiceUnjoin = common.Service{
		ServiceID:   "unjoin",
		Name:        "Unjoin",
		Description: "Removes the player from a group. Only works on platforms which support player groups.",
		Target: []common.TargetSelector{
			{
				Entity: []common.SelectorEntityFilter{
					{
						Domain: "media_player",
						SupportedFeatures: []common.SupportedFeature{
							FeatureGrouping.AsFeatureString(),
						},
					},
				},
			},
		},
	}
)
