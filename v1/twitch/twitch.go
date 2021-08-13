package twitch

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

func ( c *Controller ) OpenURI( uri_string string ) ( result string ) {
	adb := adb.Connect( c.Settings.BinaryPath , c.Settings.Host , c.Settings.Port )
	fmt.Printf( "Opening twitch://%s\n" , uri_string )
	result = adb.OpenURI( fmt.Sprintf( "twitch://%s" , uri_string ) )
	return
}

func ( c *Controller ) OpenStream( username string ) ( result string ) {
	adb := adb.Connect( c.Settings.BinaryPath , c.Settings.Host , c.Settings.Port )
	fmt.Printf( "Opening === https://twitch.tv/%s\n" , username )
	result = adb.OpenURI( fmt.Sprintf( "twitch://stream/%s" , username ) )
	return
}