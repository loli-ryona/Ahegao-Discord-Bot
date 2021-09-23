package framework

//config.json
type Config struct {
	AuthID           string   `json:"AuthID"`
	Prefixes         []string `json:"prefixes"`
	Owner            []string `json:"owner"`
	IgnoreBots       bool     `json:"ignorebots"`
	RespondToPings   bool     `json:"respondtopings"`
	CheckPermissions bool     `json:"checkpermissions"`
}

//servers.json
type Servers struct {
	Name []string `json:"name"`
	Addr []string `json:"addr"`
}

//statusbar.json
type StatusBarServers struct {
	Name []string `json:"name"`
	Addr []string `json:"addr"`
	Id   []string `json:"id"`
}
