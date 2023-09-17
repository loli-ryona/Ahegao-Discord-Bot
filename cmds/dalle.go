package cmds

import (
	hnd "ahegao/handler"
	"context"
	"fmt"
	dG "github.com/bwmarrin/discordgo"
	"github.com/sashabaranov/go-openai"
	"strings"
	"time"
)

func DalleCommand(ctx hnd.Context, args []string) error {
	//thetime
	ts := time.Now()

	//initial response
	embed := &dG.MessageEmbed{
		Title:       fmt.Sprintf("DALLE: Generating..."),
		Description: "Please wait while the image is generated",
		Footer: &dG.MessageEmbedFooter{
			Text:    "Calculating time to generate image.",
			IconURL: ctx.Session.State.User.AvatarURL("512"),
		},
	}

	msg, err := ctx.ReplyEmbed(embed)
	if err != nil {
		return err
	}

	if len(args) >= 1 {
		q := strings.Join(args, " ")
		s := openai.CreateImageSize256x256
		f := openai.CreateImageResponseFormatURL
		c := openai.NewClient(cfg.OpenAIAPI)
		o := false

		//command size options
		switch {
		case strings.Contains(q, "--size=256"):
			s = openai.CreateImageSize256x256
			o = true
		case strings.Contains(q, "--size=512"):
			s = openai.CreateImageSize512x512
			o = true
		case strings.Contains(q, "--size=1024"):
			s = openai.CreateImageSize1024x1024
			o = true
		}

		r, err := c.CreateImage(
			context.Background(),
			openai.ImageRequest{
				Prompt:         q,
				Size:           s,
				ResponseFormat: f,
				N:              1,
			})

		if err != nil {
			genErr := &dG.MessageEmbed{
				Title:       fmt.Sprintf("DALLE: Error"),
				Description: fmt.Sprintf("Image generation failed: %v\n", err),
				Footer: &dG.MessageEmbedFooter{
					Text:    fmt.Sprintf("Took %.2fs to return!", time.Since(ts).Seconds()),
					IconURL: ctx.Session.State.User.AvatarURL("512"),
				},
			}

			_, err := ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, genErr)
			if err != nil {
				return err
			}
			return err
		}

		embed = &dG.MessageEmbed{
			Title:       fmt.Sprintf("DALLE: %s's Generated Image", ctx.User.Username),
			Description: "Image successfully generated!",
			Image: &dG.MessageEmbedImage{
				URL: r.Data[0].URL,
			},
			Footer: &dG.MessageEmbedFooter{
				Text:    fmt.Sprintf("Took %.2fs to generate image!", time.Since(ts).Seconds()),
				IconURL: ctx.Session.State.User.AvatarURL("512"),
			},
		}

		embed.Fields = append(embed.Fields, &dG.MessageEmbedField{
			Name:  "Prompt",
			Value: q,
		})

		if o == true {
			embed.Fields = append(embed.Fields, &dG.MessageEmbedField{
				Name:  "Size",
				Value: s,
			})
		}

		_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, embed)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("No prompt provided")
		embed = &dG.MessageEmbed{
			Title:       "DALLE",
			Description: fmt.Sprintf("Please provide a prompt to generate the image from. You can also optionally specify a size."),
			Footer: &dG.MessageEmbedFooter{
				Text:    fmt.Sprintf("Took %.2fs to return!", time.Since(ts).Seconds()),
				IconURL: ctx.Session.State.User.AvatarURL("512"),
			},
		}

		embed.Fields = append(embed.Fields, &dG.MessageEmbedField{
			Name:  "Size",
			Value: "If you wish to generate in a larger size you can do so by appending --size=xxx to your prompt.\n You can choose from 3 sizes, **256**,**512**, and **1024**.\n *Please note the larger the size, the longer image generation takes.*",
		})

		_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, msg.ID, embed)
		if err != nil {
			return err
		}
	}

	return nil
}
