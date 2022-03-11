package conf

import (
	"fmt"
	"strings"
	"task/model"

	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RabbitMQ         string
	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQUser     string
	RabbitMQPassWord string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，检查文件路径", err)
	}
	LoadMysqlData(file)
	pathMysql := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(pathMysql)
	LoadRabbitMQData(file)
	pathRabbitMQ := strings.Join([]string{RabbitMQ, "://", RabbitMQUser, ":", RabbitMQPassWord, "@", RabbitMQHost, ":", RabbitMQPort, "/"}, "")
	model.RabbitMQ(pathRabbitMQ)
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRabbitMQData(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
}
