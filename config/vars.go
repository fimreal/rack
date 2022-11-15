package config

type ConfigFile struct {
	Port        string       `yaml:"port"`
	Workdir     string       `yaml:"workdir"`
	Tools       *Tools       `yaml:"tools"`
	Aliyun      *Aliyun      `yaml:"aliyun"`
	Email       *Email       `yaml:"email"`
	IP2Location *IP2Location `yaml:"ip2location"`
}

type Tools struct {
	Enable bool `yaml:"enable"`
}

type Aliyun struct {
	Enable   bool   `yaml:"enable"`
	AkID     string `yaml:"ak_id"`
	AkSecret string `yaml:"ak_secret"`
	RegionID string `yaml:"region_id"`
}

type Email struct {
	Enable         bool   `yaml:"enable"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	Smtpserver     string `yaml:"smtpserver"`
	Smtpserverport string `yaml:"smtpserverport"`
}

type IP2Location struct {
	Enable  bool   `yaml:"enable"`
	DbType  string `yaml:"db_type"`
	DbLevel string `yaml:"db_level"`
	Token   string `yaml:"token"`
}

var Config = &ConfigFile{}
