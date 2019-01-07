package ui

import (
	"github.com/rivo/tview"
)

var app_home *tview.Pages

func SplashPage() *tview.Pages {
	app_home = tview.NewPages()
	const logo_art = `   
	
	
	
	
	
	
		
	
	
	
	
	
	
	
	___           ___                         ___           ___           ___           ___           ___                   
	/  /\         /  /\          ___          /  /\         /  /\         /  /\         /  /\         /  /\          ___     
   /  /::\       /  /::\        /  /\        /  /::\       /  /::\       /  /::\       /  /::\       /  /::\        /  /\    
  /  /:/\:\     /  /:/\:\      /  /:/       /  /:/\:\     /__/:/\:\     /  /:/\:\     /  /:/\:\     /  /:/\:\      /  /::\   
 /  /::\ \:\   /  /::\ \:\    /  /:/       /  /::\ \:\   _\_ \:\ \:\   /  /:/  \:\   /  /::\ \:\   /  /:/  \:\    /  /:/\:\  
/__/:/\:\ \:\ /__/:/\:\_\:\  /__/:/  ___  /__/:/\:\ \:\ /__/\ \:\ \:\ /__/:/ \__\:| /__/:/\:\_\:\ /__/:/ \__\:\  /  /::\ \:\ 
\  \:\ \:\_\/ \__\/  \:\/:/  |  |:| /  /\ \  \:\ \:\_\/ \  \:\ \:\_\/ \  \:\ /  /:/ \__\/~|::\/:/ \  \:\ /  /:/ /__/:/\:\_\:\
 \  \:\ \:\        \__\::/   |  |:|/  /:/  \  \:\ \:\    \  \:\_\:\    \  \:\  /:/     |  |:|::/   \  \:\  /:/  \__\/  \:\/:/
  \  \:\_\/        /  /:/    |__|:|__/:/    \  \:\_\/     \  \:\/:/     \  \:\/:/      |  |:|\/     \  \:\/:/        \  \::/ 
   \  \:\         /__/:/      \__\::::/      \  \:\        \  \::/       \__\::/       |__|:|~       \  \::/          \__\/  
	\__\/         \__\/           ~~~~        \__\/         \__\/            ~~         \__\|         \__\/                  

	 

`
	app_home.AddPage("home", tview.NewTextView().SetText(logo_art).SetTextAlign(tview.AlignCenter), false, false)
	return app_home
	// app_home.Draw(tview.NewButton("START"))
	// return frame
}
