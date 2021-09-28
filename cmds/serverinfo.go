package cmds

import (
	"fmt"
	ap "github.com/MikeModder/anpan"
	dG "github.com/bwmarrin/discordgo"
	"github.com/rumblefrog/go-a2s"
	"strconv"
	str "strings"
	"time"
)

func queryPlayers(addr string) (string, bool) {
	onlinePlayers := ""
	embedOnlinePlayers := false
	client, errClient := a2s.NewClient(addr)
	if errClient != nil {
		onlinePlayers = fmt.Sprintln("Error creating client. Error: ", errClient)
		return onlinePlayers, false
	}

	players, errPlayer := client.QueryPlayer()
	if errPlayer != nil {
		onlinePlayers = fmt.Sprintln("Error querying players. Error: ", errPlayer)
		return onlinePlayers, false
	}

	var realPlayers []string
	for _, player := range players.Players {
		if str.Index(player.Name, "!replay") == -1 &&
			str.Index(player.Name, "WR") == -1 &&
			str.Index(player.Name, "Main") == -1 &&
			str.Index(player.Name, "Bonus") == -1 &&
			str.Index(player.Name, "GOTV") == -1 {
			realPlayers = append(realPlayers, player.Name)
		}
	}
	if len(realPlayers) > 0 {
		embedOnlinePlayers = true
	}
	onlinePlayers += fmt.Sprintf("%s\n", str.Join(realPlayers, ", "))
	return onlinePlayers, embedOnlinePlayers
}

func ServerInfoCommand(ctx ap.Context, args []string) error {
	// thetime
	ts := time.Now()

	//Reply with embed
	embed := &dG.MessageEmbed{
		Title:       "Server info.",
		Description: "Please wait while we query the server.",
		Footer: &dG.MessageEmbedFooter{
			Text:    "Calculating time to query server.",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	msg, err := ctx.ReplyEmbed(embed)
	if err != nil {
		return err
	}

	//Check server was supplied
	if len(args) >= 1 {
		//Get server from command
		server := args[0]
		if client, err := a2s.NewClient(server); err != nil {
			fmt.Println("Error creating new A2S client. Error: ", err)
			embedErr := &dG.MessageEmbed{
				Title:       "Server info.",
				Description: fmt.Sprintln("Error creating new A2S client. Error: ", err),
				Footer: &dG.MessageEmbedFooter{
					Text:    fmt.Sprintf("Took %.2fs to query server!", time.Since(ts).Seconds()),
					IconURL: ctx.Session.State.User.AvatarURL("512"),
				},
			}
			_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, embedErr)
			return nil
		} else {
			defer client.Close()

			//Query server for player names
			onlinePlayers, hasPlayers := queryPlayers(server)

			//Query server for info
			if info, err := client.QueryInfo(); err != nil {
				fmt.Println("Error querying server. Error: ", err)
				embedErr := &dG.MessageEmbed{
					Title:       "Server info.",
					Description: fmt.Sprintln("Error querying server. Error: ", err),
					Footer: &dG.MessageEmbedFooter{
						Text:    fmt.Sprintf("Took %.2fs to query server!", time.Since(ts).Seconds()),
						IconURL: ctx.Session.State.User.AvatarURL("512"),
					},
				}
				_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, embedErr)
				return nil
			} else {
				//Assign vars
				sPlayers := info.Players - info.Bots
				sOs := info.ServerOS
				srvOs := ""

				if sOs == a2s.ServerOS_Mac {
					srvOs = "Mac OSX"
				} else if sOs == a2s.ServerOS_Linux {
					srvOs = "Linux"
				} else if sOs == a2s.ServerOS_Windows {
					srvOs = "Windows"
				} else {
					srvOs = "Unknown OS"
				}

				//Edit embed
				ed := &dG.MessageEmbed{
					Title:       info.Name,
					Description: "Currently **" + strconv.Itoa(int(sPlayers)) + "** players and **" + strconv.Itoa(int(info.Bots)) + "** bots out of **" + strconv.Itoa(int(info.MaxPlayers)) + "** Total online.",
					Footer: &dG.MessageEmbedFooter{
						Text:    fmt.Sprintf("Took %.2fs to query server!", time.Since(ts).Seconds()),
						IconURL: ctx.Session.State.User.AvatarURL("512"),
					},
				}

				//Fields
				if hasPlayers == true {
					ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
						Name:  "Current players: ",
						Value: onlinePlayers,
					})
				}

				ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
					Name:  "Current map:",
					Value: info.Map,
				})

				ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
					Name:  "Security:",
					Value: "VAC: " + strconv.FormatBool(info.VAC) + " | Password: " + strconv.FormatBool(info.Visibility),
				})

				ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
					Name:  "System Info",
					Value: "OS: " + srvOs + " | Game: " + info.Game + " | Game ID: " + strconv.Itoa(int(info.ID)),
				})
				//Return embed
				_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
				return nil
			}
		}
	} else {
		fmt.Println("No server provided")
		noSrv := &dG.MessageEmbed{
			Title:       "Server info.",
			Description: fmt.Sprintln("No server provided. Usage `_serverinfo <url.com/ip/ip:port>`"),
			Footer: &dG.MessageEmbedFooter{
				Text:    fmt.Sprintf("Took %.2fs to query server!", time.Since(ts).Seconds()),
				IconURL: ctx.Session.State.User.AvatarURL("512"),
			},
		}
		_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, noSrv)
		return nil
	}
}
