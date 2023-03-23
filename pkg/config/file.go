package config

// type ConfigFile struct {
// 	Port        string       `yaml:"port"`
// 	Workdir     string       `yaml:"workdir"`
// 	Tools       *Tools       `yaml:"tools"`
// 	Aliyun      *Aliyun      `yaml:"aliyun"`
// 	Email       *Email       `yaml:"email"`
// 	IP2Location *IP2Location `yaml:"ip2location"`
// }

// type Tools struct {
// 	Enable bool `yaml:"enable"`
// }

// type Aliyun struct {
// 	Enable   bool   `yaml:"enable"`
// 	AkID     string `yaml:"ak_id"`
// 	AkSecret string `yaml:"ak_secret"`
// 	RegionID string `yaml:"region_id"`
// }

// type Email struct {
// 	Enable         bool   `yaml:"enable"`
// 	Username       string `yaml:"username"`
// 	Password       string `yaml:"password"`
// 	Smtpserver     string `yaml:"smtpserver"`
// 	Smtpserverport string `yaml:"smtpserverport"`
// }

// type IP2Location struct {
// 	Enable  bool   `yaml:"enable"`
// 	DbType  string `yaml:"db_type"`
// 	DbLevel string `yaml:"db_level"`
// 	Token   string `yaml:"token"`
// }

// var Config = &ConfigFile{}

// func LoadConfigFile() {
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yaml")
// 	viper.AddConfigPath("./")
// 	viper.AddConfigPath(viper.GetString(os.Getenv("HOME") + "/.rack/"))
// 	if err := viper.ReadInConfig(); err != nil {
// 		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
// 			// 如果是因为找不到文件，则忽略该错误
// 			ezap.Warn(err.error())
// 		} else {
// 			ezap.Fatalf("Loading config file failed: %v\n", err)
// 		}
// 	}
// 	// 监听文件修改，热加载配置
// 	viper.WatchConfig()
// 	viper.OnConfigChange(func(in fsnotify.Event) {
// 		ezap.Warnf("Config file changed: %s, %s", in.Name, in.Op)
// 	})
// }
