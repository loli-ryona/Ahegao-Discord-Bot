package cmds

import (
	"fmt"
	"sort"
	"strings"

	ap "ahegao/handler"
	dG "github.com/bwmarrin/discordgo"
)

// This is ripped straight from anpan bc its nice
func HelpCommand(context ap.Context, args []string, commands []*ap.Command, prefixes []string) error {
	// This is a check for a command to make sure only appropriate commands are shown.
	// TODO: Add a permission check
	typeCheck := func(chn dG.ChannelType, cmd ap.CommandType) bool {
		switch cmd {
		case ap.CommandTypeEverywhere:
			return true
		case ap.CommandTypePrivate:
			if chn == dG.ChannelTypeDM {
				return true
			}
		case ap.CommandTypeGuild:
			if chn == dG.ChannelTypeGuildText {
				return true
			}
		}

		return false
	}

	// Check if my g is checking for a specific command
	if len(args) >= 1 {
		var (
			owneronlystring = "No"
			typestring      = "Anywhere"
		)

		for _, command := range commands {
			if args[0] != command.Name {
				continue
			}
			if command.Hidden || !typeCheck(context.Channel.Type, command.Type) {
				return nil
			}
			if command.OwnerOnly {
				owneronlystring = "No"
			}

			switch command.Type {
			case ap.CommandTypePrivate:
				typestring = "Private"
			case ap.CommandTypeGuild:
				typestring = "Guild-only"
			}

			prefixesBuilder := strings.Builder{}
			if len(prefixes) == 1 {
				prefixesBuilder.WriteString(fmt.Sprintf("The prefix is %s", prefixes[0]))
			} else {
				prefixesBuilder.WriteString("The prefixes are ")
				for i, prefix := range prefixes {
					if i+1 == len(prefixes) {
						prefixesBuilder.WriteString(fmt.Sprintf("and %s", prefix))
					} else {
						prefixesBuilder.WriteString(fmt.Sprintf("%s, ", prefix))
					}
				}
			}

			// Check for aliases
			aliases := "**None.**"
			if len(command.Aliases) > 0 {
				aliases = strings.Join(command.Aliases, "`, `")
				aliases = "`" + aliases + "`"
			}

			// Return info for the command
			_, err := context.ReplyEmbed(&dG.MessageEmbed{
				Title:       "Help",
				Color:       0x08a4ff,
				Description: fmt.Sprintf("**%s**\nAliases: %s\nOwner only: **%s**\nUsable: **%s**", command.Description, aliases, owneronlystring, typestring),
				Footer: &dG.MessageEmbedFooter{
					Text: fmt.Sprintf(" %s.", prefixesBuilder.String()),
				},
			})

			return err
		}

		// We've not found anything :(
		_, err := context.Reply("Command `" + args[0] + "` doesn't exist.")
		return err
	}

	// Show the g what commands we actually have if not looking for specific
	var (
		count          int
		commandsSorted = make([]*ap.Command, len(commands))
		embed          = &dG.MessageEmbed{
			Title: "Commands",
			Color: 0x08a4ff,
		}
		names = make([]string, len(commands))
	)

	// Get all names, sort a-z, and arrange accordingly
	for i, cmd := range commands {
		names[i] = cmd.Name
	}
	sort.Strings(names)
	for i, v := range names {
		for _, v2 := range commands {
			if v2.Name == v {
				commandsSorted[i] = v2
				break
			}
		}
		if commandsSorted[i] == nil {
			return fmt.Errorf("sort failure")
		}
	}

	// Now that we've sorted the commands, we can show them to the user.
	for _, cmd := range commandsSorted {
		if !cmd.Hidden && typeCheck(context.Channel.Type, cmd.Type) {
			embed.Fields = append(embed.Fields, &dG.MessageEmbedField{
				Name:   cmd.Name,
				Value:  cmd.Description,
				Inline: count%2 == 0,
			})
			count++
		}
	}

	// Footer for additional information.
	var footer strings.Builder

	// Check total commands and print
	if count == 1 {
		footer.WriteString("There is 1 command.")
	} else {
		footer.WriteString(fmt.Sprintf("There are %d commands.", count))
	}

	footer.WriteString(" | ")

	// print prefixes
	if len(prefixes) == 1 {
		footer.WriteString(fmt.Sprintf("The prefix is %s.", prefixes[0]))
	} else {
		prefixesBuilder := strings.Builder{}
		for i, prefix := range prefixes {
			if i+1 == len(prefixes) {
				prefixesBuilder.WriteString(fmt.Sprintf("and %s", prefix))
			} else {
				prefixesBuilder.WriteString(fmt.Sprintf("%s, ", prefix))
			}
		}
		footer.WriteString(fmt.Sprintf("The prefixes are %s.", prefixesBuilder.String()))
	}
	embed.Footer = &dG.MessageEmbedFooter{Text: footer.String()}
	_, err := context.ReplyEmbed(embed)
	return err
}
