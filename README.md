golang后端开发框架 
===================
## 总体说明
框架集成了http服务，定时任务，消息队列，延时消息队列，数据库，缓存，环境配置自动化处理，
能满足大多数场景下互联网应用的后端开发，利用golang语言的轻量级和高并发的特点，
能轻松实现大并发，多用户访问的互联网后端项目。
建议采用Kubernetes进行分布式的自动化部署发布。


## 系统采用的主要三方组件库
gin-gonic/gin                   路由，参数验证
jinzhu/gorm                     数据库
go-redis/redis                  缓存
fluent/fluent-logger-golang     日志系统
timest/env                      配置文件
streadway/amqp                  消息队列

## 目录结构

初始的目录结构如下：

~~~
www  WEB部署目录（或者子目录）
├─controllers           应用目录
│  ├─admin
│  │   └─v1_1         版本v1.1目录
│  ├─home
│  │   └─v1_1         版本v1.1目录
│  ├─base.go          控制器基础文件，token验证等公共预处理在这里
│
├─models              模型目录，里面包含各个表的模型文件
│  ├─user_info.go     用户模型
│  └─...        
│
├─rountes              框架系统目录
│  ├─router.go         主路由文件
│  ├─...               
│
├─services              公共服务组件目录，
│  ├─db                 数据库公共调用目录
│  ├─cache              缓存公共调用目录
│  ├─log                日志公共调用目录
│  ├─queue              消息队列公共调用目录
│  ├─atom               公共系统常量，比如code返回码等等
│
├─consumer              消息队列的消费处理程序，
│  ├─main.go            消费进程的启动进程入口
│  ├─msg_push.go        消息推送逻辑处理示例
│  ├─...                
│
├─crontab               定时任务目录
│  ├─main.go            定时任务的启动进程入口
│  ├user
│  │ ├─rask_user_lvl.go  定时处理会员等级变更逻辑示例
│  ├─...    
│
├─config                配置目录，
│  ├─env.go             环境配置处理
│  ├─...    
│
├─entry                 http服务入口目录
│  ├─main.go            http服务入口文件 
│  ├─...              
│
├─README.md             项目说明文件
├─Makefile              编译文件
├─build.sh              编译执行文件
├─go.mod                gomod配置文件
├─.env                  环境变量配置示例文件

~~~

## 命名规范

### 目录和文件

*   目录小写+下划线模式；
*   文件名统一以`.go`为后缀；
*   文件名统一小写+下划线；

### 函数和类、属性命名
*   类的命名采用驼峰法，并且首字母大写，例如 `User`、`UserType`；
*   函数，方法，变量属性的命名均使用驼峰法的方式，例如 `GetIp` 或 `getIp`；

### 常量和配置
*   常量以大写字母和下划线命名，例如 `APP_PATH`和 `IMAGE_PATH`；
*   配置参数以驼峰法或小写字母和下划线命名均可，例如 `url_route_on` 和`UrlRouteOn`；

### 数据表和字段
*   数据表和字段采用小写加下划线方式命名，并注意字段名不要以下划线开头，例如 `gp_user` 表和 `user_name`字段，不建议使用驼峰和中文作为数据表字段命名。

## 路由
*    控制器采用gin的mvc的简易路由，参考routers

## 控制器
*    参考controllers目录，必须采用 Namespace模式

## model
*    参考models目录

## 数据库调用
*    数据库调用采用gorm库

## 缓存调用
*    数据库调用采用go-redis库

## 日志调用
*    数据库调用采用fluent-logger-golang库

## 单元测试

单元测试采用原生的golang单元测试即可

## 部署
*   本地安装 go 环境，设置 $GOPATH 环境变量，如果不设置GOPATH所有的三方包将放置到本地pkg目录下。


*   下载项目代码，
*	git clone http:***.git

	 
### 编译	
*   执行 ./build.sh  编译最终的可执行文件，编译好的可执行文件保存在build目录中。
*
*   部署的时候需要将.env环境配置文件和可执行文件放置在同一目录下。
*   启动web进程:  ./web-go-easy-frame
*
*   部署编译完毕
