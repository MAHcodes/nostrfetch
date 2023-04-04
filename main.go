package main

import (
	"context"
	"encoding/json"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/casimir/xdg-go"
	"github.com/dolmen-go/kittyimg"
	"github.com/i582/cfmt/cmd/cfmt"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
	"github.com/spf13/viper"
)

type Config struct {
	relayUrl string
	npub     string
}

func readConfig() (relayUrl, npub string) {
	app := xdg.App{Name: "nostrfetch"}
	configFile := app.ConfigPath("config.yaml")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(cfmt.Errorf("Fatal error reading configuration file: %s \n", err))
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(cfmt.Errorf("Fatal error unmarshaling configuration: %s \n", err))
	}

	relayUrl = viper.GetString("relayUrl")
	npub = viper.GetString("npub")

	return relayUrl, npub
}

func getProfile(relayUrl, npub string) (event *nostr.Event) {
	relay, err := nostr.RelayConnect(context.Background(), relayUrl)
	if err != nil {
		panic(err)
	}

	var filters nostr.Filters
	if _, v, err := nip19.Decode(npub); err == nil {
		pub := v.(string)
		filters = []nostr.Filter{{
			Kinds:   []int{0},
			Authors: []string{pub},
			Limit:   1,
		}}
	} else {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	sub := relay.Subscribe(ctx, filters)

	go func() {
		<-sub.EndOfStoredEvents
		cancel()
	}()

	for ev := range sub.Events {
		return ev
	}
	return nil
}

type Data struct {
	Name    string `json:"name"`
	About   string `json:"about"`
	Picture string `json:"picture"`
	Banner  string `json:"banner"`
	Nip05   string `json:"nip05"`
	Lud06   string `json:"lud06"`
	Lud16   string `json:"lud16"`
}

func parseContent(content string) (data Data) {
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		cfmt.Println("Error parsing JSON:", err)
		return
	}
	return data
}

func fetchImg(url string) string {
	resp, err := http.Get(url)
	filename := path.Base(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create temporary file
	tmpfile, err := ioutil.TempFile("", "*-"+filename)
	if err != nil {
		panic(err)
	}
	defer tmpfile.Close()

	// Write file contents to temporary file
	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		panic(err)
	}

	return tmpfile.Name()
}

func displayImg(imgPath string) {
	f, err := os.Open(imgPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		panic(err)
	}

	cfmt.Print("\n")
	kittyimg.Fprintln(os.Stdout, img)
	cfmt.Print("\n")
}

func main() {
	relayUrl, npub := readConfig()
	event := getProfile(relayUrl, npub)
	data := parseContent(event.Content)

	profilePath := fetchImg(data.Picture)
	defer os.Remove(profilePath)

	displayImg(profilePath)

	relayStyle := "yellow|italic"
	linesStyle := "blue"
	titleStyle := "green|bold"
	valueStyle := "magenta"

	colors := [16]string{
		"black",
		"red",
		"lightRed",
		"green",
		"lightGreen",
		"yellow",
		"lightYellow",
		"blue",
		"lightBlue",
		"magenta",
		"lightMagenta",
		"cyan",
		"lightCyan",
		"white",
		"gray",
	}

	cfmt.Printf("{{╭───────··}}::%s {{%s}}::%s {{··───────\n}}::%s", linesStyle, relayUrl, relayStyle, linesStyle)
	if data.Name != "" {
		cfmt.Printf("{{├─·}}::%s {{Name:}}::%s {{%s}}::%s\n", linesStyle, titleStyle, data.Name, valueStyle)
	}
	if data.About != "" {
		cfmt.Printf("{{├─·}}::%s {{About:}}::%s {{%s}}::%s\n", linesStyle, titleStyle, data.About, valueStyle)
	}
	if data.Lud16 != "" {
		cfmt.Printf("{{├─·}}::%s {{Lud16:}}::%s {{%s}}::%s\n", linesStyle, titleStyle, data.Lud16, valueStyle)
	}
	if data.Lud06 != "" {
		cfmt.Printf("{{├─·}}::%s {{Lud06:}}::%s {{%s}}::%s\n", linesStyle, titleStyle, data.Lud06, valueStyle)
	}
	if data.Nip05 != "" {
		cfmt.Printf("{{├─·󰞑}}::%s {{Nip05:}}::%s {{%s}}::%s\n", linesStyle, titleStyle, data.Nip05, valueStyle)
	}
	cfmt.Printf("{{├─·󰌋}}::%s {{Npub:}}::%s {{%s}}::%s\n", linesStyle, titleStyle, npub, valueStyle)
	cfmt.Printf("{{╰──··}}::%s", linesStyle)

	for _, color := range colors {
		if color != "" {
			cfmt.Printf("{{·─}}::%s", color)
		}
	}
	cfmt.Print("\n\n")
}
