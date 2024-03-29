package handler

import (
	"github.com/bwmarrin/discordgo"
)

/*CommandFunc defines a normal command's function.
 * Context	-> The context supplied by the command handler. Refer to Context for help.
 * []string	-> The arguments sent along with the command, basically the rest of the message after the command and the prefix. Note that this is split by spaces.
 */
type CommandFunc func(Context, []string) error

/*DebugFunc is used for debugging output.
 * string	-> The returned message.
 */
type DebugFunc func(string)

/*HelpCommandFunc defines a help command's function.
 * Context		-> The context supplied by the command handler. Refer to Context for help.
 * []string		-> The arguments sent along with the command, basically the rest of the message after the command and the prefix. Note that this is split by spaces. This can be used to show help for a specific command.
 * []*Command	-> The command slice, containing all commands.
 * []string		-> The prefixes used by the command handler.
 */
type HelpCommandFunc func(Context, []string, []*Command, []string) error

/*
PrerunFunc is the type for the function that can be run before command execution.
It is executed after any permission checks and the likes but run before the command itself.
If all goes well, return true. otherwise, false.
  - Parameters:
  - Context				-> The supplied content.
  - *Command				-> The command that is about to be executed.
  - []string				-> The arguments sent along with the command, basically the rest of the message after the command and the prefix. Note that this is split by spaces. This can be used to show help for a specific command.

Notes:
  - This is executed before the actual command, unless the guild object is not nil, then it's run before the permission check.
*/
type PrerunFunc func(Context, *Command, []string) bool

/*OnErrorFunc is the type for the function that can be run.
 * Context	-> The context supplied by the command handler. Refer to Context for help.
 * *Command	-> The command in question.
 * []string		-> The arguments sent along with the command, basically the rest of the message after the command and the prefix. Note that this is split by spaces. This can be used to show help for a specific command.
 * error	-> The returned error.
 */
type OnErrorFunc func(Context, *Command, []string, error)

// CommandType defines where commands can be used.
type CommandType int

// HelpCommand defines a help command.
// Refer to SetHelpCommand for help.
type HelpCommand struct {
	Aliases         []string
	Name            string
	SelfPermissions int64
	UserPermissions int64
	Run             HelpCommandFunc
}

// CommandHandler contains all the data needed for the handler to function.
// Anything inside here must be controlled with the appropriate Get/Set/Remove function.
type CommandHandler struct {
	enabled          bool
	checkPermissions bool
	ignoreBots       bool
	respondToPings   bool
	useState         bool

	owners   []string
	prefixes []string

	commands    []*Command
	helpCommand *HelpCommand

	debugFunc   DebugFunc
	onErrorFunc OnErrorFunc
	prerunFunc  PrerunFunc
}

// Command represents a command.
// Refer to AddCommand for help.
type Command struct {
	Aliases     []string
	Description string
	Name        string

	Hidden    bool
	OwnerOnly bool

	SelfPermissions int64
	UserPermissions int64

	Run CommandFunc

	Type CommandType
}

// Context holds the data required for command execution.
type Context struct {
	// Handler represents the handler on which this command was registered on.
	Handler *CommandHandler

	// Channel defines the channel in which the command has been executed.
	Channel *discordgo.Channel

	// Guild defines the guild in which the command has been executed.
	// Note that this may be nil under certain circumstances.
	Guild *discordgo.Guild

	// Member defines the member in the guild in which the command has been executed.
	// Note that, if guild is nil, this will be nil too.
	Member *discordgo.Member

	// Message defines the message that has executed this command.
	Message *discordgo.Message

	// Session defines the session that this command handler run on top of.
	Session *discordgo.Session

	// User defines the user that has executed the command.
	User *discordgo.User
}

const (
	// CommandTypeEverywhere defines a command that can be executed anywhere.
	CommandTypeEverywhere CommandType = iota

	// CommandTypeGuild defines a command that cannot be executed in direct messages.
	CommandTypeGuild

	// CommandTypePrivate defines a command that cannot be executed in a guild.
	CommandTypePrivate
)
