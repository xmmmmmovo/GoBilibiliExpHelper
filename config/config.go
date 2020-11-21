package config

// Config 配置结构体
type Config struct {
	User struct {
		Mode  string `yaml:"mode"`
		Login struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}
		Token struct {
			DEDEUSERID string `yaml:"DEDEUSERID"`
			SESSDATA   string `yaml:"SESSDATA"`
			BILIJCT    string `yaml:"BILI_JCT"`
		}
	}
	Home struct {
		On bool `yaml:"on"`
	}
	Live struct {
		On bool `yaml:"on"`
	}
	Comic struct {
		On       bool   `yaml:"on"`
		Platform string `yaml:"platform"`
	}
	Mail struct {
		On   bool   `yaml:"on"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Ext struct {
		On    bool `yaml:"on"`
		Kafka struct {
			On bool `yaml:"on"`
		}
	}
}
