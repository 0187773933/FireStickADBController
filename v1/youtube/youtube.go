package youtube

import (
	"fmt"
	// "time"
	adb "github.com/0187773933/ADBWrapper/v1/wrapper"
)

// Does Not Work
// youtube.OpenURI( "https://youtu.be/dltRr8w8oCo?t=245" )
// youtube.OpenURI( "https://www.youtube.com/playlist?list=PLcW8xNfZoh7fCLYJi0m3JXLs0LdcAsc0R" )
func OpenURI( uri_string string ) ( result string ) {
	adb := adb.Connect(
		"/Users/morpheous/Library/Android/sdk/platform-tools/adb" ,
		"192.168.1.120" ,
		"5555" ,
	)
	fmt.Printf( "Opening === %s\n" , uri_string )
	result = adb.OpenURI( uri_string )
	return
}