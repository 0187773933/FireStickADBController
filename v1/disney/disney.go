package disney

import (
	"fmt"
	"time"
	"strconv"
	try "github.com/manucorporat/try"
	adb "github.com/0187773933/ADBWrapper/v1/wrapper"
)


func OpenURI( uri_string string ) ( result string ) {
	adb := adb.Connect(
		"/Users/morpheous/Library/Android/sdk/platform-tools/adb" ,
		"192.168.1.120" ,
		"5555" ,
	)
	try.This(func() {
		distance_from_profile_selection_screen := adb.CurrentScreenSimilarityToReferenceImage( "/Users/morpheous/WORKSPACE/GO/FireStickADBController/v1/disney/images/profile_selection.png" )
		distance_from_profile_selection_screen_float , _ := strconv.ParseFloat( distance_from_profile_selection_screen , 64 )
		if distance_from_profile_selection_screen_float < 3.0 {
			fmt.Println( "on profile selection screen" )
			adb.PressButtonSequence( 21 , 21 , 21 , 21 , 21 , 21 , 21 , 21 , 21 )
			time.Sleep( 300 * time.Millisecond )
			adb.PressButtonSequence( 22 , 22 , 22 )
			time.Sleep( 300 * time.Millisecond )
			adb.PressButtonSequence( 23 )
			time.Sleep( 1500 * time.Millisecond )
		}
		fmt.Println( distance_from_profile_selection_screen )
	}).Catch(func(e try.E) {
	})
	result = adb.OpenURI( uri_string )
	// adb.Screenshot() // monkaS
	return
}