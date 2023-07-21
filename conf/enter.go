package conf

type iConf interface {
	loadFile() string
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
}

var App = new(app)
