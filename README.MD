# 快速指南

## 后端：

打开后端代码文件SmartLinkProject，然后先配置数据源，配置文件路径为config/settings.dev.yml，

```yml
database:
  # 数据库类型 mysql，sqlite3， postgres
  driver: mysql
  # 数据库连接字符串 mysql 缺省信息 charset=utf8&parseTime=True&loc=Local&timeout=1000ms
  source: user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
```



#### WARNING

Mysql 版本 8.0+ ，在此版本下最优；
其他低版本的会出现`Error 1071: Specified key was too long; max key length is 1000 bytes`等这类问题；请根据本地数据库版本进行对应修改；

报错原因：

`mysql`在创建单列索引的时候对列的长度是有限制的`myisam`和`innodb`存储引擎下长度限制分别为`1000 bytes`和`767 bytes`。

解决方法：



```sh
# 编辑配置文件vim /etc/my.cnf
# 在[mysqld] 下面添加MySQL默认的引擎设置default-storage-engine=InnoDB
# 重启服务service mysqld restart
```

删除库中迁移出的表，再次执行迁移命令，即可成功



数据初始化：先执行单独给出sql文件admin6.19.sql，在本地创建数据库。

接着启动服务，打开终端，输入以下代码：

go mod tidy

go build 

go run main.go  server -c config\settings.dev.yml



如果看到以下数据内容，请检查以下数据库配置；

```bash
2024-06-19 16:09:41.989 [INFO] Logger init success!
2024-06-19 16:09:41.990 [INFO] mysql-drive.go:20: user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local&timeout=1000ms
2024-06-19 16:09:44.350 [FATA] mysql-drive.go:23: mysql connect error : dial tcp 127.0.0.1:3306: connect: connection refused
```





## 前端部分（智联云文件和编辑器都需要进行以下两个操作）：

## 打开终端，先安装依赖：



```bash
$ npm install --registry=https://registry.npmmirror.com
# 或者使用$ cnpm install
# 上述两种安装报错或者node16+使用yarn进行安装$ yarn
```



#### INFO

这里还原包是需要一些时间的请耐心等待一下...

看到类似下面输出内容说明已经安装好了



```bash
Binary found at /Users/zhangwenjian/Code/go-test/go-admin-ui/node_modules/node-sass/vendor/darwin-x64-64/binding.node
Testing binary
Binary is fine
added 2033 packages from 1953 contributors in 40.229s
```

## 然后启动启动项目，使用`npm run dev`命令就好了。

```bash
# 启动页面
$ npm run dev
```

输出内容：

```bash
DONE  Compiled successfully in 22188ms                                                                                                         12:47:40 AM

  App running at:  - Local:   http://localhost:9527/  - Network: http://192.168.3.12:9527/
  Note that the development build is not optimized.  To create a production build, run npm run build.
```



#### INFO

此时前端项目已经启动了，但是有一点请注意：检查后端是否也启动了，否则页面会提示错误的哦。如果想要查看定时任务和日志的功能，需要额外再运行文件中的xxl-job和SLC_actuator项目。如果想连接实际设备，可以参考文件中的“设备端最终代码”文件。