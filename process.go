package main

import (
	"fmt"
	dG "github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
	"github.com/rumblefrog/go-a2s"
	"time"
)

// StartStatusBarProcess = Starts the status bar service.
/*
 * session = *discordgo.Session
 * event = *discordgo.Ready
 */
func StartStatusBarProcess(session *dG.Session, _ *dG.Ready) {
	fmt.Println("Starting Status Bar Ticker")
	go func() {
		ticker := time.NewTicker(time.Second * 60)
		for ; true; <-ticker.C {
			for i := 0; i < len(sbs.Name); i++ {
				statusBarProcess(session, sbs.Name[i], sbs.Addr[i], sbs.Id[i])
			}
		}
	}()
}

// statusBarProcess = Updates status bar channel information every 60 seconds
/*
 * session = *discordgo.Session
 * name = server name
 * addr = server domain
 * id = channel id
 */
func statusBarProcess(session *dG.Session, name string, addr string, id string) {
	go func() {
		fmt.Println("working on " + name)
		client, err := a2s.NewClient(addr)
		if err != nil {
			// don't care
			fmt.Println("ruh roh")
		} else {
			defer func(client *a2s.Client) {
				err := client.Close()
				if err != nil {
					fmt.Println("mega ruh roh")
				}
			}(client)

			info, err := client.QueryInfo()
			if err != nil {
				// todo: do something here
				fmt.Printf("%s crapped itself\n", name)
			} else {
				r34l := float64(info.Players - info.Bots)
				status := fmt.Sprintf("%s%0.f", name, r34l)
				fmt.Println(status)
				
				same, err := session.Channel(id)
				if err != nil {
					fmt.Println("Cannot get channel name")
				}
				if same.Name != status {
					ch, err := session.ChannelEdit(id, &dG.ChannelEdit{Name: status})
					if err != nil {
						spew.Dump(ch)
						fmt.Println(err)
					}
				} else {
					fmt.Println(fmt.Sprintf("%s: Status hasn't changed, no need to update.", name))
				}
			}
		}
	}()
}
