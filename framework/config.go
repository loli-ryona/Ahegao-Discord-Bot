package framework

import (
	hnd "ahegao/handler"
	"fmt"
	dG "github.com/bwmarrin/discordgo"
)

// New = Returns a new interface containing the specified configs
/*
 * *Configs = Use to specify which configs to add to the interface
 */
func New(o *Configs) *Configs {
	return &Configs{
		MainConfig:      o.MainConfig,
		ServersConfig:   o.ServersConfig,
		StatusBarConfig: o.StatusBarConfig,
		LennyConfig:     o.LennyConfig,
		SrcdsConfig:     o.SrcdsConfig,
	}
}

// CreateCommandHandler = Creates a new command handler
/*
 * session = *discordgo.Session
 * configs = Configs struct, needed to give handler.CommandHandler necessary information
 */
func CreateCommandHandler(session *dG.Session, configs *Configs) *hnd.CommandHandler {
	fmt.Println("Creating command handler")
	cmdHandler := hnd.New(configs.MainConfig.Prefixes, configs.MainConfig.Owner, session.StateEnabled, configs.MainConfig.IgnoreBots, configs.MainConfig.RespondToPings, configs.MainConfig.CheckPermissions, CmdPrerun, OnError, HandlerDebug)
	return &cmdHandler
}
