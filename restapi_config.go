package gobe

type RestApiBaseConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port string `mapstructure:"port" json:"port"`
}
