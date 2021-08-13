package youtube

import (
	"fmt"
	// "time"
	// try "github.com/manucorporat/try"
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

// Does Not Work
// youtube.OpenURI( "https://youtu.be/dltRr8w8oCo?t=245" )
// youtube.OpenURI( "https://www.youtube.com/playlist?list=PLcW8xNfZoh7fCLYJi0m3JXLs0LdcAsc0R" )
func ( c *Controller ) OpenURI( uri_string string ) ( result string ) {
	adb := adb.Connect( c.Settings.BinaryPath , c.Settings.Host , c.Settings.Port )
	fmt.Printf( "Opening === %s\n" , uri_string )
	result = adb.OpenURI( uri_string )
	return
}