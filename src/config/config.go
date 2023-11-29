package config

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Directories []string `xml:"directory"`
}

func loadDefaultConfig() bool {
	defaultXml := []byte(`
		<config>
			<!--  填写需要覆盖的目录  -->
			<!--  例如以下会把commit的内容覆盖到plugins下，没有的目录自动构建
			<directory>D:/server1/plugins</directory>
			<directory>D:/server2/plugins</directory>
		</config>
	`)
	err := ioutil.WriteFile("config.xml", defaultXml, 0644)
	return err == nil
}

func parseConfig(filename string) (Config, error) {
	var config Config
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = xml.Unmarshal(bytes, &config)
	return config, err
}

func IsExistFile(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func IsExistCommitFloder(fileName string) bool {
	commitFolderPath := "commit/" + fileName
	return IsExistFile(commitFolderPath)
}

func GetOrCreateConfig() (Config, error) {
	_, err := os.Stat("config.xml")
	if err != nil {
		if !loadDefaultConfig() {
			fmt.Println("出错了,未能加载config.xml")
		} else {
			fmt.Println("请前往config.xml中配置替换的目标文件夹")
		}
	}
	return parseConfig("config.xml")
}
