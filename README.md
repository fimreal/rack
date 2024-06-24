# rack
练习使用 gin 做的简易工具，方便作为容器服务部署到云平台 SaaS，提供简单有用的接口服务，方便平时测试使用。

#### 项目结构
```bash
.
├── Dockerfile
├── Dockerfile-alpine-tzdata
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── main.go
├── mods.go
├── module  # 模块目录(更多见 https://github.com/rack-plugins/)
│   ├── ...
├── pkg   # 通用包
│   ├── cmd         # cobra 命令行参数
│   ├── components  # 第三方包
│   ├── config      # viper 配置文件
│   ├── serve       # gin 服务
│   ├── service     # 旧的 gin 服务
│   └── utils
```

#### 项目开发
```bash
air
```

#### 项目构建
```bash
make build
```

#### 项目运行
```bash
docker run -it --rm epurs/rack:latest
```


# github

[rack.git](https://github.com/fimreal/rack)