# 昆朋新能内部使用项目预算API

#### 项目介绍
- 采用 iris 框架目后台api [IrisApiProject](https://github.com/snowlyg/IrisApiProject.gits)
- 采用了 gorm 数据库模块
- 使用了 sqlite3 数据库
- 用 vue 写了一个前端 [IrisApiVueAdmin](https://github.com/snowlyg/IrisApiVueAdmin)
---

#### 项目目录结构
- apidoc 接口文档目录
- caches redis缓存目录
- config 项目配置文件目录
- controllers 控制器文件目录
- database 数据库文件目录
- middleware 中间件文件目录
- models 模型文件目录
- tmp 测试数据库 sqlite3 文件目录
- tools 其他公用方法目录
---

#### api项目初始化

>拉取项目

```
git clone https://github.com/snowlyg/IrisApiProject.git
```

>加载依赖管理包

```
本来是用 godep 管理的，使用后发现还是是有问题。暂时不使用依赖管理包，依赖要自行下载。
```

>项目配置文件 /config/config.toml

```
cp config.toml.example config.toml
```

>运行项目 

```
gowatch //安装 gowatch 后才可以使用这个命令，不然只能使用

go run main.go // go 命令
```

---
##### 单元测试 
>http test

```
 go test -v  //所有测试
 
 go test -run TestUserCreate -v //单个测试
 
```

---

##### api 文档使用
自动生成文档 (访问过接口就会自动成功)
因为原生的 jquery.min.js 里面的 cdn 是使用国外的，访问很慢。
有条件的可以开个 vpn ,如果没有可以根据下面的方法修改一下，访问就很快了
>打开 apidoc/index.html 修改里面的

```
https://ajax.googleapis.com/ajax/libs/jquery/2.1.3/jquery.min.js

国内的 cdn


https://cdn.bootcss.com/jquery/2.1.3/jquery.min.js
```

>访问文档，从浏览器直接打开 apidoc/index.html 文件

---

##### 全局财务配置
- 包含管理费用、物业水电、差旅、业务招待、工资、社保、公积金作为全局预算配置
- 上述配置每月自动更新，取前面的加权平均，具体算法为：

##### 项目
- 
