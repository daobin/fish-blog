package db

import (
	"errors"
	"github.com/daobin/gotools"
	"log"
	"os"
	"path/filepath"
)

type mongo struct {
	cFile string // 配置文件
}

func (m *mongo) loadFile() string {
	if m.cFile != "" {
		return m.cFile
	}

	cFile, err := filepath.Abs("./conf/db/mongo.yml")
	log.Println("mongo配置文件: ", cFile)
	if err != nil {
		log.Println("mongo配置文件获取错误：", err.Error())
		return ""
	}

	_, err = os.Stat(cFile)
	if err != nil {
		log.Println("mongo配置文件获取错误：", err.Error())
		return "false"
	}

	m.cFile = cFile
	return cFile
}

func (m *mongo) InitFile() error {
	cFile := m.loadFile()
	if cFile == "" {
		return errors.New("mongo配置文件获取错误")
	}

	err := gotools.DB.Mongo.Init(cFile)
	if err != nil {
		log.Println("mongo配置初始化错误：", err.Error())
		return errors.New("mongo配置初始化错误")
	}

	return nil
}

// 校验是否实现相关接口
var _ iDb = (*mongo)(nil)
