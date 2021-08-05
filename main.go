package main

import (
	// "fmt"
	"time"
	spotify "github.com/0187773933/FireStickADBController/v1/spotify"
	disney "github.com/0187773933/FireStickADBController/v1/disney"
	twitch "github.com/0187773933/FireStickADBController/v1/twitch"
	youtube "github.com/0187773933/FireStickADBController/v1/youtube"
)

func main() {
	spotify.OpenPlaylist( "46CkdOm6pd6tsREVoIgZWw" )
	time.Sleep( 30 * time.Second )
	disney.OpenURI( "https://www.disneyplus.com/video/74351ae5-f6cd-4464-aeb6-8f4b10ca2649" )
	time.Sleep( 30 * time.Second )
	twitch.OpenStream( "exbc" )
	time.Sleep( 30 * time.Second )
	youtube.OpenURI( "https://www.youtube.com/watch?v=naOsvWxeYgo&list=PLcW8xNfZoh7fCLYJi0m3JXLs0LdcAsc0R&index=1" )
}