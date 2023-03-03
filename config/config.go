package config

import (
	"errors"
	"github.com/Unknwon/goconfig"
	"log"
	"os"
)

/*
读取静态配置文件，为其创建结构体
*/

// SYNTAX（Go程序的执行次序）：import -> const -> var -> init() -> main()
const pathFromRoot = "/conf/conf.ini"

// File 是文件名（public），*goconfig.ConfigFile 是才类型
// SYNTAX（哲学｜错觉）：Go语言中，首字母大小写决定可见性，无需通过额外关键字修饰；并非"大写*类、小写*对象"
var File *goconfig.ConfigFile

func init() {
	// 1.1.1 以项目根目录作为头
	dirRoot, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pathAbsolute := dirRoot + pathFromRoot

	// 1.1.2 用户自定义目录路径
	if len(os.Args) > 1 {
		dirRoot = os.Args[1]
		if dirRoot != "" {
			pathAbsolute = os.Args[1] + pathFromRoot
		}
	}

	// 1.2 健壮
	if !fileExist(pathAbsolute) {
		panic(errors.New("配置文件不存在"))
	}

	// 2.1 读取对应路径的文件
	File, err = goconfig.LoadConfigFile(pathAbsolute)
	if err != nil {
		log.Fatal("读取配置文件出错", err)
	}

	// build 没有经过测试的代码肯定是错的
	// fmt.Println(File)
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
