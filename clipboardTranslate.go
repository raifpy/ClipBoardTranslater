package main

import (
	"flag"
	"time"

	"github.com/atotto/clipboard"
	"github.com/gen2brain/dlgs"
	"github.com/raifpy/Go/errHandler"
)

var runOnBackground bool

var srcLanguage string
var outLanguage string

var lastClipBoard string
var enable bool = true
var waitSec int = 2

//var command = [3]string{"xsel", "--output", "--clipboard"}

/*
func init() {
	_cmd := exec.Command(command[0], command[1:]...)
	_cmd.Stderr = os.Stderr
	//_cmd.Stdout = os.Stdout
	_cmd.Stdin = os.Stdin
	cmd = _cmd
	out, err := _cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	lastClipBoard = string(out)
	fmt.Println(lastClipBoard)

}*/
func init() {
	_wait := flag.Int("sec", 2, "Wait second on loop")
	_src := flag.String("src", "en", "-src <langCode> -src en")
	_out := flag.String("out", "tr", "-out <langCode> -out tr")

	_onBackground := flag.Bool("background", false, "For run on background")

	flag.Parse()

	waitSec = *_wait
	srcLanguage = *_src
	outLanguage = *_out
	runOnBackground = *_onBackground

}
func getClipBoard() (string, error) {
	/*out, err := exec.Command("xsel", "--output", "--clipboard").Output()
	if errHandler.HandlerBool(err) {
		return "", err
	}
	return string(out), nil*/
	return clipboard.ReadAll()
}

func loop() {

	for range time.NewTicker(time.Second * time.Duration(waitSec)).C {
		if !enable {
			continue
		}
		clip, err := getClipBoard()
		if err != nil {
			dlgs.Error("Error on getClipBoard", err.Error())
			continue
		}
		if clip != lastClipBoard {
			lastClipBoard = clip
			text, err := translate(srcLanguage, outLanguage, clip)
			if errHandler.HandlerBool(err) {
				dlgs.Error("Error on getClipBoard", err.Error())
				continue
			}
			dlgs.Info(clip, text)
			//go dlgs.Info(clip, text)//

		}
	}
}
