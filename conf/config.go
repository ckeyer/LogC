package conf

type Config struct {
	Name string
}

var config *Config

// GetConfig 获取应用配置
func GetConfig() (c *Config, err error) {
	if config == nil {
		return loadConfig()
	}
	return
}

// loadConfig 加载配置信息
func loadConfig() (c *Config, err error) {
	c = &Config{
		Name: "LogCentre",
	}
	return
}
