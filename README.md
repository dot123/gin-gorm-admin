# gin-gorm-admin

基于 GIN + GORM + WIRE 实现web脚手架，目标是提供一套轻量的中后台开发框架，方便、快速的完成业务需求的开发。

> 账号：admin 密码：123456

> 账号：test 密码：123456

## 特性

* 遵循 `RESTful API` 设计规范 & 基于接口的编程规范
* 基于 `GIN` 框架，提供了丰富的中间件支持（JWTAuth、CORS、RequestRateLimiter、Recover、GZIP）
* 基于[jwt](https://github.com/appleboy/gin-jwt) 认证
* 基于[go-playground/validator](https://github.com/go-playground/validator)开源库简化gin的请求校验
* 用Docker上云
* 在token过期后的一个小时内，用户再次操作会要求重新登陆
* 基于[swaggo](https://github.com/swaggo)为Go工程生成自动化接口文档
* 基于[wire](https://github.com/google/wire)依赖注入
* 基于[gorm](https://gorm.io/zh_CN/)全功能ORM
* 基于[air](https://github.com/cosmtrek/air)自动编译，重启程序
* 基于redis限制请求频率

### 项目结构

<pre><code>
├─api
├─cmd
├─configs
├─docs
├─internal
│  ├─config
│  ├─contextx
│  ├─errors
│  ├─ginx
│  ├─middleware
│  ├─models
│  ├─schema
│  ├─service
│  └─validate
└─pkg
    ├─fileStore
    ├─gormx
    ├─hash
    ├─helper
    ├─logger
    ├─monitor
    ├─rabbitMQ
    ├─redisHelper
    ├─store
    ├─timer
    ├─types
    └─validate
</code></pre>

### 下载依赖

<pre><code>depend.cmd</code></pre>

### 代码生成与运行

##### 生成

<pre><code>generate.cmd</code></pre>

##### 数据库

<pre><code>gin-admin.sql</code></pre>

##### 运行

<pre><code>run.cmd 或go run ./cmd/gin-gorm-admin/ web -c ./configs/config.toml</code></pre>

##### docker部署

<pre><code>deploy.cmd</code></pre>

## 前端工程

基于 [vue](https://github.com/vuejs/vue) 和 [element-ui](https://github.com/ElemeFE/element)实现：[gin-admin-vue](https://github.com/dot123/gin-admin-vue)
