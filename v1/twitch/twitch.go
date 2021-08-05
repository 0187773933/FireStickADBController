package twitch

import (
	"fmt"
	// "time"
	adb "github.com/0187773933/ADBWrapper/v1/wrapper"
)

func OpenURI( uri_string string ) ( result string ) {
	adb := adb.Connect(
		"/Users/morpheous/Library/Android/sdk/platform-tools/adb" ,
		"192.168.1.120" ,
		"5555" ,
	)
	fmt.Printf( "Opening twitch://%s\n" , uri_string )
	result = adb.OpenURI( fmt.Sprintf( "twitch://%s" , uri_string ) )
	return
}

func OpenStream( username string ) ( result string ) {
	adb := adb.Connect(
		"/Users/morpheous/Library/Android/sdk/platform-tools/adb" ,
		"192.168.1.120" ,
		"5555" ,
	)
	fmt.Printf( "Opening === https://twitch.tv/%s\n" , username )
	result = adb.OpenURI( fmt.Sprintf( "twitch://stream/%s" , username ) )
	return
}