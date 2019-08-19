package Models
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v2"
	"io/ioutil"
)
type conf struct {
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pwd string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}
func main() {
	var c conf
	conf:=c.GetConf()
	fmt.Println(conf)
	db, _ := gorm.Open("mysql", ":@tcp()/?charset=utf8&parseTime=True&loc=Local")
	//
	defer db.Close()
}
func Get() *conf {
	return  &conf{}
}
func (c *conf) GetConf() *conf {
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
