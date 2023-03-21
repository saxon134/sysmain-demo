package conf

import (
	"gopkg.in/yaml.v2"
	"os"
)

var Conf *ModelConf

type ModelConf struct {
	Name string
	Http struct {
		Host string //一般填写内网IP，空的话会自动获取本机的IPv4地址
		Port int    //http端口，必选
	}

	Sysmain struct {
		Url        string //sysmain接口地址，完整地址如：http://192.168.1.12/sysmain
		ClientRoot string //根地址，需要同sysmain项目配置中clientroot保持一致，建议是：/sysmain
		Secret     string //如果配置了秘钥，所有接口都需要加密
		PingSecond int    //SDP client ping间隔秒数，默认5秒
	}

	Mq struct{}

	Redis struct {
		Host string
		Pass string
	}

	MySql struct {
		Host string
		Pass string
		User string
		Db   string
	}

	Feishu struct {
		Webhookurl string
		Secret     string
	}
}

func Init() *ModelConf {
	if Conf == nil {
		Conf = new(ModelConf)

		f_n := "./config.yaml"
		yamlData, err := os.ReadFile(f_n)
		if err != nil {
			panic("配置文件路径有误")
		}

		err = yaml.Unmarshal(yamlData, Conf)
		if err != nil {
			panic("配置文件信息有误")
		}
	}
	return Conf
}
