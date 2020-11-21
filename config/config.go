package config

// 配置结构体
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
			BILIJCT    string `yaml:"BILIJCT"`
		}
	}
	Home struct {
		On bool `yaml:"on"`
	}
	Live struct {
		On bool `yaml:"on"`
	}
	Comic struct {
		On bool `yaml:"on"`
	}
	Mail struct {
		On bool `yaml:"on"`
	}
	Ext struct {
		On    bool `yaml:"on"`
		Kafka struct {
			On bool `yaml:"on"`
		}
	}
}
