package handler

import (
	"os"
	"os/signal"
	"syscall"
)

/*New creates a new command handler.
 * prefixes 		- The prefixes to use for the command handler.
 * owners   		- The owners of this application; these are used for Owner-Only commands.
 * useState 		- Whether to use the session's state th fetch data or not. The state will be ignored if the State field of the session used in the message handler is set false.
 * ignoreBots 		- Whether to ignore users marked as bots or not.
 * checkPermissions	- Whether to check permissions or not.
 * useRoutines 		- Whether to execute commands outside the event's routine.
Notes:
Refer to MessageHandler to properly activate the command handler.
*/
func New(prefixes []string, owners []string, useState, ignoreBots, respondToPings, checkPermssions bool, prerunFunc PrerunFunc, errorFunc OnErrorFunc, debugFunc DebugFunc) CommandHandler {
	return CommandHandler{
		checkPermissions: checkPermssions,
		debugFunc:        debugFunc,
		enabled:          true,
		ignoreBots:       ignoreBots,
		onErrorFunc:      errorFunc,
		owners:           owners,
		prefixes:         prefixes,
		prerunFunc:       prerunFunc,
		respondToPings:   respondToPings,
		useState:         useState,
	}
}

// WaitForInterrupt makes your application wait for an interrupt.
// A SIGINT, SIGTERM or a console interrupt will make this function stop.
// Note that the Exit function in the os package will make this function stop, too.
func WaitForInterrupt() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
