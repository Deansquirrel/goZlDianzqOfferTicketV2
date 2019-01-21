package object

type SysConfig struct {
	Total   total   `toml:"total"`
	PeiZhDb peiZhDb `toml:"peiZhiDb"`
}

type total struct {
	IsDebug bool `toml:"isDebug"`
	Port    int  `toml:"port"`

	MaxTicketNum int    `toml:"maxTicketNum"`
	SnoWorkerId  int    `toml:"snoWorkerId"`
	AppId        string `toml:"appid"`
	JPeiZh       string `toml:"jpeizh"`
}

type peiZhDb struct {
	Server   string `toml:"server"`
	Port     int    `toml:"port"`
	DbName   string `toml:"dbName"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}
