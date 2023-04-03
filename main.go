package main

import (
	"context"

	"github.com/casimir/xdg-go"
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

func main() {
	relayUrl, npub := readConfig()
  event := getProfile(relayUrl, npub)
	cfmt.Println(event)
}
