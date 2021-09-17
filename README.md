# micro-mall-trolley

#### 介绍
微商城--购物车

#### 软件架构
grpc

#### 框架，库依赖
kelvins框架支持（gRPC，cron，queue，web支持）：https://gitee.com/kelvins-io/kelvins   
g2cache缓存库支持（两级缓存）：https://gitee.com/kelvins-io/g2cache   

#### 安装教程

1.仅构建  sh build.sh   
2 运行  sh build-run.sh   
3 停止 sh stop.sh

#### 使用说明

```toml
[kelvins-server]
Environment = "dev"

[kelvins-logger]
RootPath = "./logs"
Level = "debug"

[kelvins-auth]
Token = "c9VW6ForlmzdeDkZE2i8"
TransportSecurity = false

[kelvins-mysql]
Host = "127.0.0.1:3306"
UserName = "root"
Password = "yertyert"
DBName = "micro_mall_trolley"
Charset = "utf8"
PoolNum =  10
MaxIdleConns = 5
ConnMaxLifeSecond = 3600
MultiStatements = true
ParseTime = true

[kelvins-redis]
Host = "127.0.0.1:6379"
Password = "dfasdf"
DB = 1
PoolNum = 10

```

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request
