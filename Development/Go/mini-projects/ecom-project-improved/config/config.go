package config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JWTSecretKey string
}
