package main

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/getlantern/systray"
)

func main() {
	if runOnBackground {
		cmd := exec.Command(os.Args[0], "--src", srcLanguage, "--out", outLanguage, "--sec", strconv.Itoa(waitSec))
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Start()
		os.Exit(1)
	}
	systray.Run(setup, exit)
}

func setup() {
	systray.SetIcon(resourceGoPng)
	systray.SetTitle("Clipboard Translater :D")
	systray.SetTooltip("Translate your clipboard realtime?")

	enableDisable := systray.AddMenuItem("Disable", "")
	systray.AddSeparator()
	quit := systray.AddMenuItem("Quit", "Close program")
	go requestEnableDisableClipboardHandle(enableDisable)
	go requestCloseApp(quit)
	//quit.Show()

	go loop()

}

func exit() {
	systray.Quit()
	os.Exit(0)
}

func requestCloseApp(menu *systray.MenuItem) {
	<-menu.ClickedCh
	exit()
}

func requestEnableDisableClipboardHandle(menu *systray.MenuItem) {
	for range menu.ClickedCh {
		enable = !enable
		if enable {
			menu.SetTitle("Disable")
		} else {
			menu.SetTitle("Enable")
		}
	}
}
