package core

// Struct for config flags
type Config struct {
	ServerIP string
	HTTP     HTTPConfig
}

type HTTPConfig struct {
	Timeout int
}
