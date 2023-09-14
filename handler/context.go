package handler

import (
	"io"

	"github.com/bwmarrin/discordgo"
)

/*Reply directly replies with a message.
 * message	- The message content.
 */
func (c *Context) Reply(message string) (*discordgo.Message, error) {
	return c.Session.ChannelMessageSend(c.Channel.ID, message)
}

/*ReplyComplex combines Reply, ReplyEmbed and ReplyFile as a way to send a message with, for example, Text and an Embed together.
 * message	- The message content.
 * tts		- Whether the client should read the message out or not.
 * embed	- The embed for this message. Refer to discordgo.MessageEmbed for more info.
 * files	- The files to send across. These (collectively) cannot pass more than 8 Megabytes. Refer to discordgo.File for information.
 */
func (c *Context) ReplyComplex(message string, tts bool, embed *discordgo.MessageEmbed, files []*discordgo.File) (*discordgo.Message, error) {
	return c.Session.ChannelMessageSendComplex(c.Channel.ID, &discordgo.MessageSend{
		Content: message,
		Embed:   embed,
		TTS:     tts,
		Files:   files,
	})
}

/*ReplyEmbed directly replies with a embed, but not with a message.
 * embed	- The embed for this message. Refer to discordgo.MessageEmbed for more info.
 */
func (c *Context) ReplyEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return c.Session.ChannelMessageSendEmbed(c.Channel.ID, embed)
}

/*ReplyFile directly replies with a file, but not with a message.
 * files	- The files to send across. These (collectively) cannot pass more than 8 Megabytes.
 */
func (c *Context) ReplyFile(filename string, file io.Reader) (*discordgo.Message, error) {
	return c.Session.ChannelFileSend(c.Channel.ID, filename, file)
}
