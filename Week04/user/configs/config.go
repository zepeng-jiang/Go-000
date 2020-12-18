package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Server   *ServerSetting
	App      *AppSetting
	Database *DatabaseSetting
}

type ServerSetting struct {
	RunMode      string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	LogSavePath           string
	LogFileName           string
	LogFileExt            string
	ContextTimeout        time.Duration
	UploadSavePath        string
	UploadServerUrl       string
	UploadImageMaxSize    int
	UploadImageAllowExts  []string
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func InitConfig() (*Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	Conf = &Config{}
	if err := vp.Unmarshal(Conf); err != nil {
		return nil, err
	}

	return Conf, nil
}

func NewDBEngine(conf *Config) (*gorm.DB, error) {
	db, err := gorm.Open(conf.Database.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		conf.Database.UserName,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.DBName,
		conf.Database.Charset,
		conf.Database.ParseTime,
	))
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}
