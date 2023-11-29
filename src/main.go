package main

import (
	"bufio"
	"fmt"
	"os"
	"replacetools/src/commit"
	"replacetools/src/config"
	"strings"
)

var out config.Config
var reader *bufio.Reader

func init() {
	if !config.IsExistFile("commit") {
		os.Mkdir("commit", 0755)
	}
	if !config.IsExistFile("logs") {
		os.Mkdir("logs", 0755)
	}
}

func main() {

	out, _ = config.GetOrCreateConfig()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入 命令 或 按Ctrl+C 退出 (输入help查看帮助)")
	for {
		readString, _ := reader.ReadString('\n')
		readString = strings.TrimSpace(readString)
		if strings.HasPrefix(readString, "update ") {
			submit := strings.TrimPrefix(readString, "update ")
			if !config.IsExistCommitFloder(submit) {
				fmt.Println("commit文件夹下没有名为 " + submit + " 的待覆盖文件")
				continue
			}
			commit.Commit(out, submit)
			fmt.Println("commit success!")
		} else if strings.HasPrefix(readString, "help") {
			fmt.Println("-------")
			fmt.Println("使用方法: ")
			fmt.Println("在config.xml中配置替换目录,commit底下新建文件夹作为更新文件")
			fmt.Println("例如在config.xml中配置如下:")
			fmt.Println("\t\t<config>\n\t\t\t<!--  填写需要覆盖的目录  -->\n\t\t\t<directory>D:/server1/plugins</directory>\n\t\t\t<directory>D:/server2/plugins</directory>\n\t\t</config>\n\t")
			fmt.Println("commit文件夹下建立一个 '活动更新'")
			fmt.Println("存放目录")
			fmt.Println("CMI:")
			fmt.Println("  config.yml")
			fmt.Println("输入update 活动更新")
			fmt.Println("将会把CMI底下的config.yml复制到上方两个服务端中")
			fmt.Println("-------")
		}
	}
}
