这里的go-net-test.go和gorm-demo.go两个文件都直接使用了package main和func main
需要在cmd环境中使用go run运行

config是配置文件，存放访问数据库所需信息
Init.go读取配置文件
table.go是表格的表头

############################################################
###   go-net-test.go原先是自己测试网络访问使用，其中包含pong,print,query    ###
###   gorm-demo.go原先是测试gorm框架操作数据库测试用			   ###
############################################################

tttt.txt是原用于调试的go文件,可忽略