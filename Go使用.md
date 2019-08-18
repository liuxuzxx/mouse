## Go 探索

#### Go的项目性质管理
因为凡是项目，肯定有依赖，go的项目使用go mod作为依赖管理操作
```
使用go mod初始化一个项目
go mod init mouse  ---mouse是项目名字
```

#### Packages

```
Every Go program is made up of packages.

翻译过来就是:每个Go程序都是由package组成
```

#### Variable
```
https://blog.golang.org/gos-declaration-syntax
这篇文章Go作者解释了，为什么要使用这种形式的声明语法了
```

###### 声明变量的方式
1. var形式
```cgo
var c,python,java,cpp bool

感觉必须写var，一直对var就是不怎么感觉熟悉，感觉怪怪的
```
2. implicit type

其实就是var的一种隐式的行为