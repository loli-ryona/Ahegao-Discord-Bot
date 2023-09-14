package handler

import "errors"

var (
	// ErrBotBlocked is thrown when the message handler encounters a bot, but ignoring bots was set to true.
	ErrBotBlocked = errors.New("command Handler: The given author was a bot and the IgnoreBots setting is true")

	// ErrCommandAlreadyRegistered is thrown when a command by the same name was registered previously.
	ErrCommandAlreadyRegistered = errors.New("command Handler: Another command was already registered by this name")

	// ErrCommandNotFound is thrown when a message tries to invoke an unknown command, or when an attempt at removing an unregistered command was made.
	ErrCommandNotFound = errors.New("command Handler: Command not found")

	// ErrDataUnavailable is thrown when data is unavailable, like channels, users or something else.
	ErrDataUnavailable = errors.New("command Handler: Necessary data couldn't be fetched")

	// ErrDMOnly is thrown when a DM-only command is executed on a guild.
	ErrDMOnly = errors.New("command Handler: DM-Only command on guild")

	// ErrGuildOnly is thrown when a guild-only command is executed in direct messages.
	ErrGuildOnly = errors.New("command Handler: Guild-Only command in DMs")

	// ErrOwnerOnly is thrown when an owner-only command is executed.
	ErrOwnerOnly = errors.New("command Handler: Owner-Only command")

	// ErrSelfInsufficientPermissions is thrown when the bot itself does not have enough permissions.
	ErrSelfInsufficientPermissions = errors.New("command Handler: Insufficient permissions for the bot")

	// ErrUserInsufficientPermissions is thrown when the user doesn't meet the required permissions.
	ErrUserInsufficientPermissions = errors.New("command Handler: Insufficient permissions for the user")
)
