package controller

import (
	utils "github.com/0187773933/FireStickADBController/v1/utils"
	spotify "github.com/0187773933/FireStickADBController/v1/spotify"
	disney "github.com/0187773933/FireStickADBController/v1/disney"
	twitch "github.com/0187773933/FireStickADBController/v1/twitch"
	youtube "github.com/0187773933/FireStickADBController/v1/youtube"
)

type Controller struct {
	Settings utils.ADBSettings
	Spotify spotify.Controller
	Disney disney.Controller
	Twitch twitch.Controller
	Youtube youtube.Controller
}

func NewController( adb_binary_path string , host string , port string ) ( result Controller ) {
	result = Controller{}
	result.Settings = utils.ADBSettings{}
	result.Settings.BinaryPath = adb_binary_path
	result.Settings.Host = host
	result.Settings.Port = port
	result.Spotify = spotify.NewController( &result.Settings )
	result.Disney = disney.NewController( &result.Settings )
	result.Twitch = twitch.NewController( &result.Settings )
	result.Youtube = youtube.NewController( &result.Settings )
	return
}