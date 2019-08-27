package config

// Config 全局配置示例
var Config *config

// config 全部配置
type config struct {

	// 数据库配置
	DB db `yaml:"db"`

	// 应用配置
	App app `yaml:"app"`

	// 端口号(1024~65535)
	Port int `yaml:"port"`
}

// 数据库配置
type db struct {

	// 连接字符串
	Dsn string `yaml:"dsn"`

	// 数据库名称
	DbName string `yaml:"db_name"`

	// 最大空闲连接数
	MaxIdleConn int `yaml:"maxidle_conn"`

	// 最大连接数
	MaxOpenConn int `yaml:"maxopen_conn"`

	// 密码
	Pwd string `yaml:"pwd"`
}

// 应用配置
type app struct {

	// jwt配置
	Jwt jwt `yaml:"jwt"`
}

// jwt配置
type jwt struct {

	// jwt密钥
	Secret string `yaml:"secret"`
}
