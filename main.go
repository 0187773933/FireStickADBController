package main

import (
	"time"
	adb_controller "github.com/0187773933/FireStickADBController/v1/controller"
)

func main() {
	controller := adb_controller.NewController(
		"/Users/morpheous/Library/Android/sdk/platform-tools/adb" ,
		"192.168.1.120" ,
		"5555" ,
	)
	controller.Spotify.OpenPlaylist( "46CkdOm6pd6tsREVoIgZWw" )
	time.Sleep( 30 * time.Second )
	controller.Disney.OpenURI( "https://www.disneyplus.com/video/74351ae5-f6cd-4464-aeb6-8f4b10ca2649" )
	time.Sleep( 30 * time.Second )
	controller.Twitch.OpenStream( "exbc" )
	time.Sleep( 30 * time.Second )
	controller.Youtube.OpenURI( "https://www.youtube.com/watch?v=naOsvWxeYgo&list=PLcW8xNfZoh7fCLYJi0m3JXLs0LdcAsc0R&index=1" )
}