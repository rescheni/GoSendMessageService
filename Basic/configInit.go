package basic

import (
	log "GoMessageService/log"
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
	ServerJiang struct {
		ServerJiangKey  string `yaml:"server_jiang_key"`
		ServerJiangDesp string `yaml:"server_jiang_desp"`
	} `yaml:"server_jiang"`
	Feishu struct {
		FeishuAppId     string `yaml:"feishu_app_id"`
		FeishuAppSecret string `yaml:"feishu_app_secret"`
		FeishuUserId    string `yaml:"feishu_user_id"`
	} `yaml:"feishu"`
	DingDing struct {
		AccessToken string `yaml:"access_token"`
	} `yaml:"dingding"`
	WxPusher struct {
		WxPusherKey string `yaml:"wx_push_key"`
		DefaultUid  string `yaml:"default_uid"`
	} `yaml:"wx_push"`
	Napcat struct {
		NapcatUrl   string `yaml:"napcat_url"`
		NapcatToken string `yaml:"napcat_token"`
		NapcatQQ    string `yaml:"napcat_qq"`
	} `yaml:"napcat"`
	Api struct {
		ApiPort int    `yaml:"api_port"`
		ApiHost string `yaml:"api_host"`
		ApiPath string `yaml:"api_path"`
		ApiKey  string `yaml:"api_key"`
	} `yaml:"api"`
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
			log.Logger.Fatal("Error opening config file: ", err)
		}
		defer file.Close()

		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(configInstance)
		if err != nil {
			log.Logger.Fatal("Error decoding config file: ", err)
		}
	})

	return configInstance
}
