package ui

import (
	"github.com/rivo/tview"
)

type grid struct {
	*tview.Flex
	currentOption int
}

func (app *App) SplashPage() {

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

	splashGrid := createFlex()
	startButton := tview.NewForm().AddButton("Initiate", func() {
		
		app.loadDevicePage()
	}).AddButton("forgetit", func() {

		app.Stop()
	})
	main := newTextPrimitive(logo_art)

	splashGrid.
		AddItem(main, 0, 6, false).SetDirection(tview.FlexRow).
		AddItem(startButton, 0, 1, true)
	if err := app.SetRoot(splashGrid, true).Run(); err != nil {
		panic(err)
	}

}

func createFlex() grid {
	return grid{Flex: tview.NewFlex()}
}
