package cmds

import (
	ap "ahegao/handler"
	"fmt"
	dG "github.com/bwmarrin/discordgo"
	"github.com/cavaliergopher/grab/v3"
	"os"
	"strings"
	"time"
)

// DownloadMapCommand - WIP FUNCTION
/*
 NOTE: currently a bit of a mess and not implemented by default.
*/
func DownloadMapCommand(ctx ap.Context, args []string) error {
	// thetime
	ts := time.Now()

	//Declaring output vars
	op := "Initialising"

	//Reply with embed
	embed := &dG.MessageEmbed{
		Title:       "Downloading map",
		Description: op,
		Footer: &dG.MessageEmbedFooter{
			Text:    "Calculating time to run command",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	msg, err := ctx.ReplyEmbed(embed)
	if err != nil {
		return err
	}

	if len(args) >= 1 {
		if strings.Index(args[0], "https://gamebanana.com/mods/") == 0 {
			_, err = ctx.Session.ChannelMessageEdit(ctx.Message.ChannelID, msg.ID, fmt.Sprintf("currently wip"))
			return err
			// fun part, will code later
			// https://api.gamebanana.com/Core/Item/Data?itemtype=Map&itemid=206904&fields=Owner%28%29.name%2Cname%2CFiles%28%29.aFiles%28%29&help
			/*bhopUrl := strings.Split(args[0], "https://gamebanana.com/mods/")
			bhopMapID := bhopUrl[1]
			println(bhopUrl)
			op = fmt.Sprintf("Querying Gamebanana with request url: %s", args[0])
			ed := &dG.MessageEmbed{
				Title:       fmt.Sprintf("Downloading map: Gamebanana API"),
				Description: op,
				Footer: &dG.MessageEmbedFooter{
					Text:    fmt.Sprintf("Time taken: %.2fs", time.Since(ts).Seconds()),
					IconURL: ctx.Session.State.User.AvatarURL("512"),
				},
			}
			ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)

			// create request
			gbReq := "https://api.gamebanana.com/Core/Item/Data?itemtype=Mod&itemid="+bhopMapID+"&fields=Files().aFiles()&format=json&flags=JSON_UNESCAPED_SLASHES"
			gbResp, err := http.Get(gbReq)
			if err != nil {
				fmt.Println("Failed to get a response from the request. Error returned:\n ", err)
				op = fmt.Sprintf("Failed to get a response from the request: %v\n", err.Error())
				ed = &dG.MessageEmbed{
					Title:       fmt.Sprintf("Downloading map: Gamebanana API - [**FAILED**]"),
					Description: op,
					Footer: &dG.MessageEmbedFooter{
						Text:    fmt.Sprintf("Time taken: %.2fs", time.Since(ts).Seconds()),
						IconURL: ctx.Session.State.User.AvatarURL("512"),
					},
				}
				ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
				return nil
			}
			defer gbResp.Body.Close()
			body, err := ioutil.ReadAll(gbResp.Body)
			mapDir := srcds.MapDir*/

		} else {
			//print first shit
			bhopMap := args[0]
			println(bhopMap)
			op = fmt.Sprintf("Searching for: %s", bhopMap)
			ed := &dG.MessageEmbed{
				Title:       fmt.Sprintf("Downloading map: %s", bhopMap),
				Description: op,
				Footer: &dG.MessageEmbedFooter{
					Text:    fmt.Sprintf("Time taken: %.2fs", time.Since(ts).Seconds()),
					IconURL: ctx.Session.State.User.AvatarURL("512"),
				},
			}
			ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)

			// now we actually make the client and specify vars
			mapDir := srcds.MapDir
			fastDL := srcds.FastDL
			client := grab.NewClient()
			req, _ := grab.NewRequest(mapDir, fastDL+bhopMap+".bsp.bz2")

			op = fmt.Sprintf("Sent request for: %s", bhopMap)
			ed = &dG.MessageEmbed{
				Title:       fmt.Sprintf("Downloading map: %s", bhopMap),
				Description: op,
				Footer: &dG.MessageEmbedFooter{
					Text:    fmt.Sprintf("Time taken: %.2fs", time.Since(ts).Seconds()),
					IconURL: ctx.Session.State.User.AvatarURL("512"),
				},
			}
			ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)

			// start the download
			fmt.Printf("Downloading %v...\n", req.URL())
			resp := client.Do(req)
			fmt.Printf("  %v\n", resp.HTTPResponse.Status)

			// start ui loop
			t := time.NewTicker(500 * time.Millisecond)
			defer t.Stop()
		Loop:
			for {
				select {
				case <-t.C:
					fmt.Printf("  Transferred %v bytes - ETA: %v\n", byteString(resp.BytesComplete()), resp.ETA())
					op = fmt.Sprintf("  Transferred %v bytes - ETA: %v\n", byteString(resp.BytesComplete()), resp.ETA())
					ed = &dG.MessageEmbed{
						Title:       fmt.Sprintf("Downloading map: %s", bhopMap),
						Description: op,
						Footer: &dG.MessageEmbedFooter{
							Text:    fmt.Sprintf("Time taken: %.2fs", time.Since(ts).Seconds()),
							IconURL: ctx.Session.State.User.AvatarURL("512"),
						},
					}
					ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
				case <-resp.Done:
					//download finished
					break Loop
				}
			}

			// error check
			if err := resp.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
				op = fmt.Sprintf("Download failed: %v\n", err.Error())
				ed = &dG.MessageEmbed{
					Title:       fmt.Sprintf("Downloading map: %s - [**FAILED**]", bhopMap),
					Description: op,
					Footer: &dG.MessageEmbedFooter{
						Text:    fmt.Sprintf("Time taken: %.2fs", time.Since(ts).Seconds()),
						IconURL: ctx.Session.State.User.AvatarURL("512"),
					},
				}
				ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
				return nil
			}

			// finished download
			fmt.Printf("Download saved to ./%v \n", resp.Filename)
			op = fmt.Sprintf("Download finished successfully. Change to a different map if the !map %s command shows unknown map.\n Download saved to ./%v \n", bhopMap, resp.Filename)
			ed = &dG.MessageEmbed{
				Title:       fmt.Sprintf("Downloading map: %s - [**FINISHED**]", bhopMap),
				Description: op,
				Footer: &dG.MessageEmbedFooter{
					Text:    fmt.Sprintf("Time taken: %.2fs", time.Since(ts).Seconds()),
					IconURL: ctx.Session.State.User.AvatarURL("512"),
				},
			}
			ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, ed)
			return nil
		}
	} else {
		fmt.Printf("No args provided")
		embed = &dG.MessageEmbed{
			Title:       "Downloading map: No args provided",
			Description: "Please provide a map name or a gamebanana mod link like so:\n *.dl bhop_arcane_v1* or *.dl https://gamebanana.com/mods/126232*",
			Footer: &dG.MessageEmbedFooter{
				Text:    fmt.Sprintf("Time taken: %.2fs", time.Since(ts).Seconds()),
				IconURL: ctx.Session.State.User.AvatarURL("512"),
			},
		}
	}

	return nil
}

// ByteString = Prints download progress.
/*
 * n = Bytes downloaded so far, accepts the response for BytesComplete as an int64
 */
func byteString(n int64) string {
	if n < 1<<10 {
		return fmt.Sprintf("%dB", n)
	}
	if n < 1<<20 {
		return fmt.Sprintf("%dKB", n>>10)
	}
	if n < 1<<30 {
		return fmt.Sprintf("%dMB", n>>20)
	}
	if n < 1<<40 {
		return fmt.Sprintf("%dGB", n>>30)
	}
	return fmt.Sprintf("%dTB", n>>40)
}
