package cmds

import (
	"fmt"
	str "strings"
	"sync"
	"time"

	ap "ahegao/handler"
	dG "github.com/bwmarrin/discordgo"
	"github.com/rumblefrog/go-a2s"
)

func CurrentMapCommand(ctx ap.Context, _ []string) error {
	// thetime
	ts := time.Now()

	//Declaring output var
	op := ""

	//Reply with embed
	embed := &dG.MessageEmbed{
		Title:       "Current maps.",
		Description: "Please wait while we query the servers for maps",
		Footer: &dG.MessageEmbedFooter{
			Text:    "Calculating time to query servers.",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	msg, err := ctx.ReplyEmbed(embed)
	if err != nil {
		return err
	}

	//Wait group shit
	var wg sync.WaitGroup
	wg.Add(len(srv.Name))

	//Loop through each server and get current map
	for i := 0; i < len(srv.Name); i++ {
		cum := i
		go func(cum int) {
			defer wg.Done()
			var cm []string
			if client, err := a2s.NewClient(srv.Addr[cum]); err != nil {
				fmt.Println("Error creating new A2S client. Error: ", err)
			} else {
				defer client.Close()
				if maaps, err := client.QueryInfo(); err != nil {
					fmt.Println("Error querying map. Error: ", err)
				} else {
					cm = append(cm, maaps.Map)
				}
				if len(cm) > 0 {
					op += fmt.Sprintf("%s**%s**\n", srv.Name[cum], str.Join(cm, "."))
				}
			}
		}(cum)
	}
	wg.Wait()

	fmt.Println("Waitgroup Done")

	//Edit embed
	ed := &dG.MessageEmbed{
		Title:       "Current maps.",
		Description: op,
		Footer: &dG.MessageEmbedFooter{
			Text:    fmt.Sprintf("Took %.2fs to query servers!", time.Since(ts).Seconds()),
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	//Return embed
	_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
	return nil
}
