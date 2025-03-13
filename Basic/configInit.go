package basic

import (
	"log"
	"os"
	"sync"

	yaml "sigs.k8s.io/yaml/goyaml.v3"
)

// Config 结构体定义了所有需要的配置字段
type Config struct {
	Email struct {
		EmailAddress string `yaml:"email_address"`
		Username     string `yaml:"username"`
		SmtpServer   string `yaml:"smtp_server"`
		SmtpPort     int    `yaml:"smtp_port"`
		AuthCode     string `yaml:"auth_code"`
		Encryption   string `yaml:"encryption"`
		SenderName   string `yaml:"sender_name"`
	} `yaml:"email"`
}

var (
	configInstance *Config
	once           sync.Once
)

// LoadConfig 负责读取配置文件并加载配置
func LoadConfig() *Config {
	once.Do(func() {
		configInstance = &Config{}
		file, err := os.Open("Message_main_config.yaml")
		if err != nil {
			log.Fatal("Error opening config file: ", err)
		}
		defer file.Close()

		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(configInstance)
		if err != nil {
			log.Fatal("Error decoding config file: ", err)
		}
	})

	return configInstance
}
