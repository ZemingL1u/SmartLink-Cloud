package main

import (
	"SmartLinkProject/cmd"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

// @title go-admin API
// @version 2.0.0
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description 添加qq群: 521386980 进入技术交流群 请先star，谢谢！
// @license.name MIT
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// 初始化 MQTT 客户端
	//mqtt.InitMqttClient("tcp://p1ec8c41.ala.dedicated.aliyun.emqxcloud.cn:1883")

	cmd.Execute()

	// // 程序结束前断开连接
	// if mqtt.MqttClient.IsConnected() {
	// 	mqtt.MqttClient.Disconnect(250)
	// }
}
