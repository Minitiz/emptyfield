package main

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/minitiz/emptyfield"
	"gopkg.in/yaml.v2"
)

type test struct {
	field1 string `fields:"omitempty"`
	field2 string
}

func main() {
	ParseConfigMap()

	MDR, err := emptyfield.Check(reflect.ValueOf(CFG), emptyfield.JSONOmitEmptyEnabled)
	fmt.Println(MDR)
	fmt.Println(err)

}

type ConfigMap struct {
	Usermgtdb struct {
		AuthSource string
		Svc        string
		Username   string
		Password   string
	}
	Projectmgtdb struct {
		AuthSource string
		Svc        string
		Username   string
		Password   string
	}
	Backofficedb struct {
		AuthSource string
		Svc        string
		Username   string
		Password   string
	}
	Eventfrontdb struct {
		AuthSource string
		Svc        string
		Username   string
		Password   string
	}
	Usermgt struct {
		Svc          string
		AddressMicro string
	}
	Projectmgt struct {
		Svc          string
		AddressMicro string
	}
	RedisLogmgt struct {
		Svc      string
		Password string
	}
	Minio struct {
		Server          string
		AccessKey       string
		SecretKey       string
		BucketLogs      string
		BucketArtifacts string
	}
	SMTP struct {
		Address  string
		Port     int
		User     string
		Password string
	}
	Admin struct {
		Username string
		Password string
	}
	Salt struct {
		Salt   string
		Length int
	}
	Bosession struct {
		Timeout int    `field:"omitempty"`
		Hash    string `json:"test,omitempty"`
	} `field:"omitempty"`
}

var CFG ConfigMap = ConfigMap{}

func ParseConfigMap() error {

	valsByte, err := ioutil.ReadFile("testmar.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(valsByte, &CFG)
	if err != nil {
		return err
	}
	return nil
}
