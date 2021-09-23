package cmds

import (
	"fmt"
	ap "github.com/MikeModder/anpan"
	dG "github.com/bwmarrin/discordgo"
	"github.com/rumblefrog/go-a2s"
	"strconv"
	"time"
)

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
		var sName []string
		var sMap []string
		var sPlayers []uint8
		var sMaxPlayers []uint8
		var sVac []bool
		var sOs []a2s.ServerOS
		var sPass []bool
		var sGame []string
		var sGameID []uint16

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
				sName := append(sName, info.Name)
				sMap := append(sMap, info.Map)
				sPlayers := append(sPlayers, info.Players)
				sMaxPlayers := append(sMaxPlayers, info.MaxPlayers)
				sVac := append(sVac, info.VAC)
				sOs := append(sOs, info.ServerOS)
				sPass := append(sPass, info.Visibility)
				sGame := append(sGame, info.Game)
				sGameID := append(sGameID, info.ID)
				srvOs := ""

				if sOs[0] == a2s.ServerOS_Mac {
					srvOs = "Mac OSX"
				} else if sOs[0] == a2s.ServerOS_Linux {
					srvOs = "Linux"
				} else if sOs[0] == a2s.ServerOS_Windows {
					srvOs = "Windows"
				} else {
					srvOs = "Unknown OS"
				}

				//Edit embed
				ed := &dG.MessageEmbed{
					Title:       sName[0],
					Description: "Currently " + strconv.Itoa(int(sPlayers[0])) + "/" + strconv.Itoa(int(sMaxPlayers[0])) + " online.",
					Footer: &dG.MessageEmbedFooter{
						Text:    fmt.Sprintf("Took %.2fs to query server!", time.Since(ts).Seconds()),
						IconURL: ctx.Session.State.User.AvatarURL("512"),
					},
				}

				//Fields
				ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
					Name:  "Current map:",
					Value: sMap[0],
				})

				ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
					Name:  "Security:",
					Value: "VAC: " + strconv.FormatBool(sVac[0]) + " | Password: " + strconv.FormatBool(sPass[0]),
				})

				ed.Fields = append(ed.Fields, &dG.MessageEmbedField{
					Name:  "System Info",
					Value: "OS: " + srvOs + " | Game: " + sGame[0] + " | Game ID: " + strconv.Itoa(int(sGameID[0])),
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
