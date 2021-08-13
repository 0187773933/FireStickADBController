package spotify

import (
	"fmt"
	"time"
	adb "github.com/0187773933/ADBWrapper/v1/wrapper"
	utils "github.com/0187773933/FireStickADBController/v1/utils"
)

type Controller struct {
	Settings *utils.ADBSettings
}

func NewController( settings *utils.ADBSettings ) ( result Controller ) {
	result = Controller{}
	result.Settings = settings
	return
}

func ( c *Controller ) OpenURI( uri_string string ) ( result string ) {
	adb := adb.Connect( c.Settings.BinaryPath , c.Settings.Host , c.Settings.Port )
	result = adb.OpenURI( fmt.Sprintf( "spotify:%s" , uri_string ) )
	return
}

func ( c *Controller ) OpenPlaylist( playlist_id string ) ( result string ) {
	adb := adb.Connect( c.Settings.BinaryPath , c.Settings.Host , c.Settings.Port )
	result = adb.OpenURI( fmt.Sprintf( "spotify:playlist:%s:play" , playlist_id ) )
	time.Sleep( 3000 * time.Millisecond )
	adb.PressButtonSequence( 21 )
	adb.PressButtonSequence( 21 )
	adb.PressButtonSequence( 21 )
	adb.PressButtonSequence( 21 )
	adb.PressButtonSequence( 21 )
	adb.PressButtonSequence( 21 )
	time.Sleep( 1 * time.Second )
	adb.PressButtonSequence( 22 )
	time.Sleep( 1 * time.Second )
	adb.PressButtonSequence( 23 )
	time.Sleep( 1 * time.Second )
	adb.PressButtonSequence( 87 )
	time.Sleep( 2000 * time.Millisecond )
	adb.Tap( 500 , 500 )
	return
}


// func OpenURI( uri_string string ) ( result string ) {
// 	adb := adb.Connect(
// 		"/Users/morpheous/Library/Android/sdk/platform-tools/adb" ,
// 		"192.168.1.120" ,
// 		"5555" ,
// 	)
// 	result = adb.OpenURI( fmt.Sprintf( "spotify:%s" , uri_string ) )
// 	return
// }

// func OpenPlaylist( playlist_id string ) ( result string ) {
// 	adb := adb.Connect(
// 		"/Users/morpheous/Library/Android/sdk/platform-tools/adb" ,
// 		"192.168.1.120" ,
// 		"5555" ,
// 	)
// 	result = adb.OpenURI( fmt.Sprintf( "spotify:playlist:%s:play" , playlist_id ) )
// 	time.Sleep( 3000 * time.Millisecond )
// 	adb.PressButtonSequence( 21 )
// 	adb.PressButtonSequence( 21 )
// 	adb.PressButtonSequence( 21 )
// 	adb.PressButtonSequence( 21 )
// 	adb.PressButtonSequence( 21 )
// 	adb.PressButtonSequence( 21 )
// 	time.Sleep( 1 * time.Second )
// 	adb.PressButtonSequence( 22 )
// 	time.Sleep( 1 * time.Second )
// 	adb.PressButtonSequence( 23 )
// 	time.Sleep( 1 * time.Second )
// 	adb.PressButtonSequence( 87 )
// 	time.Sleep( 2000 * time.Millisecond )
// 	adb.Tap( 500 , 500 )
// 	return
// }