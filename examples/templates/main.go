package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/emmaly/go-hass"
	"github.com/emmaly/go-hass/mediaPlayer"
	"github.com/joho/godotenv"
)

func getHA() (ha *hass.Access, err error) {
	godotenv.Load(".env")
	host := os.Getenv("HASS_HOST")
	if host == "" {
		panic("HASS_HOST environment variable not set")
	}

	token := os.Getenv("HASS_TOKEN")
	if token == "" {
		panic("HASS_TOKEN environment variable not set")
	}

	ha = hass.NewAccess(host, "")
	ha.SetBearerToken(token)
	if err := ha.CheckAPI(); err != nil {
		return nil, err
	}
	return
}

func main() {
	ha, err := getHA()
	if err != nil {
		panic(err)
	}
	_ = ha

	domains := []string{
		"media_player",
	}
	domainsJSON, err := json.Marshal(domains)
	if err != nil {
		panic(err)
	}

	entityAreaNames := []string{
		// "", // "" == none aka associated with no area
		"Living Room",
		// "Purple Room",
	}
	entityAreaNamesJSON, err := json.Marshal(entityAreaNames)
	if err != nil {
		panic(err)
	}

	deviceAreaNames := []string{
		// "", // "" == none aka associated with no area
		"Living Room",
		// "Purple Room",
	}
	deviceAreaNamesJSON, err := json.Marshal(deviceAreaNames)
	if err != nil {
		panic(err)
	}

	out, err := ha.RenderTemplate(fmt.Sprintf(`
		{%% set domains = %s %%}
		{%% set entity_area_names = %s %%}
		{%% set device_area_names = %s %%}

		{%% set area_names_length = (entity_area_names + device_area_names) | length %%}
		{%% set area_names = namespace(entity=entity_area_names, device=device_area_names) %%}
		{%% set ns = namespace(entity_area_ids=[], device_area_ids=[], entityObjects=[]) %%}

		{%% for key in ('entity', 'device') %%}
			{%% set a = namespace(ok=false,names=[],id=none) %%}
			{%% if key == 'entity' %%}
				{%% set a.ok=true %%}
				{%% set a.names=area_names.entity %%}
			{%% elif key == 'device' %%}
				{%% set a.ok=true %%}
				{%% set a.names=area_names.device %%}
			{%% endif %%}
			{%% for area_name in a.names %%}
				{%% set a.id = area_name %%}
				{%% if a.id == '' %%}
					{%% set a.id = none %%}
				{%% endif %%}
				{%% if a.id is not none %%}
					{%% set a.id = area_id(a.id) %%}
					{%% if a.id is none %%}
						{%% set a.ok = false %%}
					{%% endif %%}
				{%% endif %%}
				{%% if a.ok %%}
					{%% if key == 'entity' %%}
						{%% set ns.entity_area_ids = ns.entity_area_ids + [a.id] %%}
					{%% elif key == 'device' %%}
						{%% set ns.device_area_ids = ns.device_area_ids + [a.id] %%}
					{%% endif %%}
				{%% endif %%}
			{%% endfor %%}
		{%% endfor %%}

		{%% for domain in domains %%}
			{%% for entity in states[domain] %%}
				{%% set entityId = entity.entity_id %%}
				{%% set entityName = entity.name %%}
				{%% set entityState = entity.state %%}
				{%% set entityArea = area_id(entityId) %%}
				{%% set deviceId = device_id(entityId) %%}
				{%% set deviceArea = area_id(deviceId) %%}
				{%% if entityArea in ns.entity_area_ids or deviceArea in ns.device_area_ids or area_names_length == 0 %%}
					{%% set e = namespace(object_str="{") %%}

					{%% set attributes = [
						{'key': 'entity_id', 'value': entityId},
						{'key': 'domain', 'value': domain},
						{'key': 'name', 'value': entityName},
						{'key': 'state', 'value': entityState},
						{'key': 'entity_area', 'value': entityArea},
						{'key': 'device_id', 'value': deviceId},
						{'key': 'device_area', 'value': deviceArea},
					] %%}

					{%% for attr in attributes %%}
						{%% if attr.value is not none %%}
							{%% set e.object_str = e.object_str ~ (attr.key | tojson) ~ ':' ~ (attr.value | tojson) ~ ', ' %%}
						{%% endif %%}
					{%% endfor %%}

					{%% set e.object_str = e.object_str ~ ('attributes' | tojson) ~ ': {' %%}

					{%% for attr in entity.attributes %%}
						{%% set key = (attr ~ '') | tojson %%}
						{%% set value = entity.attributes[attr] %%}
						{%% if value is not none %%}
							{%% if value is number or value is boolean or value is sequence  %%}
								{%% set value = value | tojson %%}
								{%% set e.object_str = e.object_str ~ key ~ ':' ~ value ~ ', ' %%}
							{%% elif value | string | regex_match('^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}') %%}
								{%% set value = (value.utcnow().strftime('%%Y-%%m-%%dT%%H:%%M:%%S') ~ 'Z') | tojson %%}
								{%% set e.object_str = e.object_str ~ key ~ ':' ~ value ~ ', ' %%}
							{%% else %%}
								{%% set value = (value ~ '') | tojson %%}
								{%% set e.object_str = e.object_str ~ key ~ ':' ~ value ~ ', ' %%}
							{%% endif %%}
						{%% endif %%}
					{%% endfor %%}

					{%% set e.object_str = e.object_str[:-2] ~ '}}' %%}

					{%% set object = e.object_str | from_json %%}
					{%% set ns.entityObjects = ns.entityObjects + [object] %%}
				{%% endif %%}
			{%% endfor %%}
		{%% endfor %%}

		{{ ns.entityObjects  | tojson(indent=2) }}
	`, domainsJSON, entityAreaNamesJSON, deviceAreaNamesJSON))
	if err != nil {
		panic(err)
	}

	// println(out)
	// return

	var players []mediaPlayer.MediaPlayer
	err = json.Unmarshal([]byte(out), &players)
	if err != nil {
		panic(err)
	}

	for _, player := range players {
		fmt.Printf("%s\n", player.EntityID)

		// if player.Name != "" {
		// 	fmt.Printf("\tName: %s\n", player.Name)
		// }
		// if player.FriendlyName != "" && player.FriendlyName != player.Name {
		// 	fmt.Printf("\tFriendlyName: %s\n", player.FriendlyName)
		// }
		// if player.EntityArea != "" {
		// 	fmt.Printf("\tArea: %s\n", player.EntityArea)
		// }
		// if player.DeviceID != "" {
		// 	fmt.Printf("\tDevice ID: %s\n", player.DeviceID)
		// 	if player.DeviceClass != "" {
		// 		fmt.Printf("\t\t- Class: %s\n", player.DeviceClass)
		// 	}
		// 	if player.DeviceArea != "" {
		// 		fmt.Printf("\t\t- Area : %s\n", player.DeviceArea)
		// 	}
		// }
		// if player.State != "" {
		// 	fmt.Printf("\tState: %s\n", player.State)
		// }
		// if player.Attributes.EntityPicture.Valid {
		// 	fmt.Printf("\tEntityPicture: %s\n", player.Attributes.EntityPicture.URL(ha))
		// }
		// if player.Attributes.EntityPictureLocal.Valid {
		// 	fmt.Printf("\tEntityPictureLocal: %s\n", player.Attributes.EntityPictureLocal.URL(ha))
		// }
		// if len(player.Attributes.GroupMembers) > 0 {
		// 	fmt.Printf("\tGroupMembers: %s\n", player.Attributes.GroupMembers)
		// }
		// if player.Attributes.InputSource != "" {
		// 	fmt.Printf("\tInputSource: %s\n", player.Attributes.InputSource)
		// }
		// if player.Attributes.InputSourceList.Valid {
		// 	fmt.Printf("\tInputSourceList:\n\t\t- %s\n", player.Attributes.InputSourceList.String())
		// }
		// if player.Attributes.MediaAnnounce != "" {
		// 	fmt.Printf("\tMediaAnnounce: %s\n", player.Attributes.MediaAnnounce)
		// }
		// if player.Attributes.MediaAlbumArtist != "" {
		// 	fmt.Printf("\tMediaAlbumArtist: %s\n", player.Attributes.MediaAlbumArtist)
		// }
		// if player.Attributes.MediaAlbumName != "" {
		// 	fmt.Printf("\tMediaAlbumName: %s\n", player.Attributes.MediaAlbumName)
		// }
		// if player.Attributes.MediaArtist != "" {
		// 	fmt.Printf("\tMediaArtist: %s\n", player.Attributes.MediaArtist)
		// }
		// if player.Attributes.MediaChannel != "" {
		// 	fmt.Printf("\tMediaChannel: %s\n", player.Attributes.MediaChannel)
		// }
		// if player.Attributes.MediaContentID != "" {
		// 	fmt.Printf("\tMediaContentID: %s\n", player.Attributes.MediaContentID)
		// }
		// if player.Attributes.MediaContentType != "" {
		// 	fmt.Printf("\tMediaContentType: %s\n", player.Attributes.MediaContentType)
		// }
		// if player.Attributes.MediaDuration > 0 {
		// 	fmt.Printf("\tMediaDuration: %0.2f\n", player.Attributes.MediaDuration)
		// }
		// if player.Attributes.MediaEnqueue != "" {
		// 	fmt.Printf("\tMediaEnqueue: %s\n", player.Attributes.MediaEnqueue)
		// }
		// if player.Attributes.MediaExtra != "" {
		// 	fmt.Printf("\tMediaExtra: %s\n", player.Attributes.MediaExtra)
		// }
		// if player.Attributes.MediaEpisode != "" {
		// 	fmt.Printf("\tMediaEpisode: %s\n", player.Attributes.MediaEpisode)
		// }
		// if player.Attributes.MediaPlaylist != "" {
		// 	fmt.Printf("\tMediaPlaylist: %s\n", player.Attributes.MediaPlaylist)
		// }
		// if player.Attributes.MediaPosition > 0 {
		// 	fmt.Printf("\tMediaPosition: %0.2f\n", player.Attributes.MediaPosition)
		// }
		// if player.Attributes.MediaPositionUpdatedAt != nil {
		// 	fmt.Printf("\tMediaPositionUpdatedAt: %s\n", player.Attributes.MediaPositionUpdatedAt)
		// }
		// if player.Attributes.MediaRepeat != "" {
		// 	fmt.Printf("\tMediaRepeat: %s\n", player.Attributes.MediaRepeat)
		// }
		// if player.Attributes.MediaSeason != "" {
		// 	fmt.Printf("\tMediaSeason: %s\n", player.Attributes.MediaSeason)
		// }
		// if player.Attributes.MediaSeekPosition > 0 {
		// 	fmt.Printf("\tMediaSeekPosition: %0.2f\n", player.Attributes.MediaSeekPosition)
		// }
		// if player.Attributes.MediaSeriesTitle != "" {
		// 	fmt.Printf("\tMediaSeriesTitle: %s\n", player.Attributes.MediaSeriesTitle)
		// }
		// if player.Attributes.MediaShuffle.Valid {
		// 	fmt.Printf("\tMediaShuffle: %s\n", player.Attributes.MediaShuffle.String())
		// }
		// if player.Attributes.MediaTitle != "" {
		// 	fmt.Printf("\tMediaTitle: %s\n", player.Attributes.MediaTitle)
		// }
		// if player.Attributes.MediaTrack > 0 {
		// 	fmt.Printf("\tMediaTrack: %d\n", player.Attributes.MediaTrack)
		// }
		// if player.Attributes.MediaVolumeLevel.Valid {
		// 	fmt.Printf("\tMediaVolumeLevel: %s\n", player.Attributes.MediaVolumeLevel.String())
		// }
		// if player.Attributes.MediaVolumeMuted.Valid {
		// 	fmt.Printf("\tMediaChannel: %s\n", player.Attributes.MediaVolumeMuted.String())
		// }
		// if player.Attributes.SoundMode != "" {
		// 	fmt.Printf("\tSoundMode: %s\n", player.Attributes.SoundMode)
		// }
		// if player.Attributes.SoundModeList != "" {
		// 	fmt.Printf("\tSoundModeList: %s\n", player.Attributes.SoundModeList)
		// }
		// if player.Attributes.SupportedFeatures.Valid {
		// 	fmt.Printf("\tMediaPlayerEntityFeatures:\n\t\t- %s\n", player.Attributes.SupportedFeatures.String())
		// }
		// fmt.Println()

		// player.Attributes.SupportedFeatures.Value = 0
		player.Attributes.SupportedFeatures.Valid = false
		// hi := mediaPlayer.InputSource("Hello")
		// player.Attributes.InputSource = &hi
		// player.Attributes.InputSourceList = &mediaPlayer.InputSourceList{
		// 	Valid: true,
		// 	Value: []mediaPlayer.InputSource{
		// 		mediaPlayer.InputSource("Hello"),
		// 		mediaPlayer.InputSource("World"),
		// 	},
		// }
		// player.Attributes.MediaContentType = mediaPlayer.MediaClass("tvshow")

		jb, err := json.MarshalIndent(player, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jb))
		fmt.Println()
	}
}
