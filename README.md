
<div align=center><img src="https://github.com/meteorOSS/GoReplaceTools/assets/61687266/117b58ec-39fc-4107-9ee3-3f802b25fb6a"></div>

# GoReplaceTool

平时服务器更新的时候，写完的配置文件需要一个个拷贝到各个服务端。过程重复度高且容易出问题
于是这个工具就诞生了 (go写这些小脚本是真方便)

## 使用方式:
1. 在发布页根据服务器类型下载软件
2. 启动软件
3. 启动后会在当前目录下生成以下文件夹
   config.xml: 存放待覆盖目标路径
   commit: 存放覆盖文件
   ![image](https://github.com/meteorOSS/GoReplaceTools/assets/61687266/3a4f6085-1a5e-4240-a1fa-a2fbbea91a8b)
5. 一个示例
配置config.xml
``` xml
		<config>
			<!--  填写需要覆盖的目录  -->
			<!--  例如以下会把commit的内容覆盖到plugins下，没有的目录自动构建 --!>
			<directory>D:/server1/plugins</directory>
			<directory>D:/server2/plugins</directory>
		</config>
```
接下来，此次我将更新所有服务端CMI的配置文件. 在commit下新建文件夹，起一个名字，比如"CMI配置更新"
![image](https://github.com/meteorOSS/GoReplaceTools/assets/61687266/81cd21ef-436e-4d3c-8731-2de92aa54a9f)
我们需要将setting.yml覆盖到目标路径的CMI文件夹下 (D:/server1/plugins/CMI/)。
由于在上方的软件config.xml配置中，目标文件夹(服务端)仅配置到了plugins文件夹，于是"CMI配置更新"文件夹下还需要新建一个"CMI"文件夹
![image](https://github.com/meteorOSS/GoReplaceTools/assets/61687266/2b5aeb38-979c-4889-8727-4a0cebea6b7f)

打开软件 输入 update CMI配置更新
![image](https://github.com/meteorOSS/GoReplaceTools/assets/61687266/ff755265-458a-4748-86b6-fb1f8cb5edac)

### 在linux平台使用

输入chmod +x GoReplaceTools添加可执行权限

![image](https://github.com/meteorOSS/GoReplaceTools/assets/61687266/1f977ba1-161c-4def-b027-858e9604ef7d)

使用./GoReplaceTools，操作方式与上面一致

杀毒报告 [https://www.virscan.org/report/637d5cbabb88960c519dce1bd5afc9f63952f0115d6954b22dfa115dc22d417b](https://www.virscan.org/report/637d5cbabb88960c519dce1bd5afc9f63952f0115d6954b22dfa115dc22d417b)
