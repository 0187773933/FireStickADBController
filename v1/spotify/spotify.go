package spotify

import (
	"fmt"
	"time"
	adb "github.com/0187773933/ADBWrapper/v1/wrapper"
)


func OpenURI( uri_string string ) ( result string ) {
	adb := adb.Connect(
		"/Users/morpheous/Library/Android/sdk/platform-tools/adb" ,
		"192.168.1.120" ,
		"5555" ,
	)
	result = adb.OpenURI( fmt.Sprintf( "spotify:%s" , uri_string ) )
	return
}

func OpenPlaylist( playlist_id string ) ( result string ) {
	adb := adb.Connect(
		"/Users/morpheous/Library/Android/sdk/platform-tools/adb" ,
		"192.168.1.120" ,
		"5555" ,
	)
	result = adb.OpenURI( fmt.Sprintf( "spotify:playlist:%s:play" , playlist_id ) )
	time.Sleep( 1000 * time.Millisecond )
	adb.PressButtonSequence( 21 , 21 , 23 )
	time.Sleep( 1000 * time.Millisecond )
	adb.PressButtonSequence( 22 , 22 , 22 )
	time.Sleep( 400 * time.Millisecond )
	adb.PressButtonSequence( 23 )
	time.Sleep( 2000 * time.Millisecond )
	adb.Tap( 500 , 50 )
	return
}