package main

import (
	"fmt"
	dG "github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
	"github.com/rumblefrog/go-a2s"
)

/*Updates status bar channel information every 60 seconds
 * session = *discordgo.Session
 * name = server name
 * addr = server domain
 * id = channel id
 */
func process(session *dG.Session, name string, addr string, id string) {
	go func() {
		fmt.Println("working on " + name)
		client, err := a2s.NewClient(addr)
		if err != nil {
			// don't care
			fmt.Println("ruh roh")
		} else {
			defer client.Close()
			info, err := client.QueryInfo()
			if err != nil {
				// todo: do something here
				fmt.Printf("%s crapped itself\n", name)
			} else {
				r34l := float64(info.Players - info.Bots)
				status := fmt.Sprintf("%s%0.f", name, r34l)
				fmt.Println(status)

				// todo: do not update same value to save rate limits

				ch, err := session.ChannelEdit(id, &dG.ChannelEdit{Name: status})
				if err != nil {
					spew.Dump(ch)
					fmt.Println(err)
				}
			}
		}
	}()
}
