package config

type ConfigModel struct {
	Database   Database
	Redis      Redis
	Dispatcher Dispatcher
	Jwt        Jwt
	Log        Log
}
type Database struct {
	DSN           string `json:"DSN"`
	SlowThreshold int    `json:"SlowThreshold"`
	LogLevel      int    `json:"LogLevel"`
	Colorful      bool   `json:"Colorful"`
}
type Redis struct {
	Addr         string `json:"Addr"`
	DB           int    `json:"DB"`
	PoolSize     int    `json:"PoolSize"`
	MinIdleConns int    `json:"MinIdleConns"`
	ExpiredTime  int    `json:"ExpiredTime"`
}
type Dispatcher struct {
	Size int `json:"Size"`
}
type Jwt struct {
	ExpiredTime int `json:"ExpiredTime"`
}
type Log struct {
	Location string `json:"Location"`
}
