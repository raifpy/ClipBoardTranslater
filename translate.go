package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/raifpy/Go/errHandler"
)

const uri = "https://www.translate.com/translator/ajax_translate"

func translate(lang1, lang2, text string) (string, error) {

	fmt.Printf("Lang = %s to %s\nText = %s\n", lang1, lang2, text)
	client := &http.Client{}

	values := url.Values{
		"text_to_translate": {text},
		"source_lang":       {lang1},
		"translated_lang":   {lang2},
	}

	nreq, _ := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))
	nreq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ham, err := client.Do(nreq)
	if err != nil {
		return "", err
	}
	//ham, err := http.PostForm("https://www.translate.com/translator/ajax_translate", param) //
	//ham, err := http.PostForm("https://www.translate.com/translator/ajax_lang_auto_detect", param) // Bunlar da iş görür. Hediye

	if ham.StatusCode != 200 {
		//fmt.Printf("\033[31m200 dönmeyen kod {%d}\033[0m\n", ham.StatusCode)
		return "", errors.New("!200 Status Code ~ translate.com")

	}
	jsonYab := map[string]interface{}{}

	http, _ := ioutil.ReadAll(ham.Body)
	errHandler.HandlerExit(json.Unmarshal(http, &jsonYab)) // err dönerse program exit atacak; Olmaması gerek. El atarsınız oraya

	if jsonYab["result"] != "success" {
		return "", errors.New(jsonYab["result"].(string))

	}

	fmt.Println("\033[32mText \033[0;33m= \033[0m", jsonYab["translated_text"])

	return jsonYab["translated_text"].(string), nil

}
