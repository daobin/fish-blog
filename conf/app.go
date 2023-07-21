package conf

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type app struct {
	cFile string // 配置文件
}

func (a *app) loadFile() string {
	if a.cFile != "" {
		return a.cFile
	}

	cFile, err := filepath.Abs("./conf/app.yml")
	log.Println("app配置文件: ", cFile)
	if err != nil {
		log.Println("app配置文件获取错误：", err.Error())
		return ""
	}

	_, err = os.Stat(cFile)
	if err != nil {
		log.Println("app配置文件获取错误：", err.Error())
		return "false"
	}

	a.cFile = cFile
	return cFile
}

func (a *app) GetString(key string) string {
	cFile := a.loadFile()
	if cFile == "" {
		return ""
	}

	ko := koanf.New(".")
	err := ko.Load(file.Provider(cFile), yaml.Parser())
	if err != nil {
		log.Println("app配置文件加载错误：", err.Error())
		return ""
	}

	return ko.String("app." + key)
}

func (a *app) GetInt(key string) int {
	val := a.GetString(key)
	intVal, _ := strconv.Atoi(val)

	return intVal
}

func (a *app) GetBool(key string) bool {
	val := a.GetString(key)
	boolVal, _ := strconv.ParseBool(val)

	return boolVal
}

// 校验是否实现相关接口
var _ iConf = (*app)(nil)
