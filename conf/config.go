//
package conf

var BrokerConfig = struct {
	PORT   string `default:"8001"`
	Logger struct {
		LogLevel      string `default:"trace"` // trace/info/debug/warn/error/fatal
		MaxRotateSize int32  `default:"10"`    // 10M
	}
	MongoDB struct {
		Url            string
		Name           string
		TimeOut        int
		PoolLimit      int
		SessionTimeout int
	}
	Server struct {
		MusicPath string
	}
}{}
