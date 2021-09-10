# Fw: Alert 是什么

[饭碗警告（Fw: Alert）](https://fwalert.com/115)可以轻松将 webhook（HTTP 请求）、邮件转发为短信、电话等报警，内置强大的模板变量提取功能，既可以轻松与任意监控系统组合使用，也可以快速接入自研监控系统。

## Fw: Alert Go SDK

本 SDK 对 Fw: Alert 的 webhook 模式进行了封装，让你可以无需关注 HTTP 请求，只需几行代码即可快速在你的 Go 项目中接入 `饭碗警告`。  
*\*如需在 PHP 项目中使用，请移步：[https://github.com/YianAndCode/fwalert-php-sdk](https://github.com/YianAndCode/fwalert-php-sdk)*

在开始之前，请确保你已经注册好了[饭碗警告](https://fwalert.com/115)（点击左侧链接直达注册页）。

## 使用方式

首先在你的项目目录中执行：

```bash
go get -u github.com/YianAndCode/fwalert-go
```

然后在代码中引入包：

```golang
import "github.com/YianAndCode/fwalert-go"
```

接下来只需要：

```golang
fw := fwalert.New()
fw.SendAlert(
    context.Background(), // 如果你是在诸如 gin 之类的框架中使用，则传入 gin.Context 也可
    "这里替换成在饭碗警告后台拿到的 webhook url",
    map[string]string{
		"hello": "world", // 这里可以传任意你想要在告警文案中展示的字段，key-value 格式即可
	},
)
```

## 进阶用法

本 SDK 除了封装 HTTP 请求外，还增加了“频道”的概念：当你设置了多个告警规则的时候，不需要在你的代码中 New 一堆 fwalert 出来，只需要：

```golang
fw := fwalert.New()

// 提前注册好“频道”
fw.AddChannel("ch1", "webhook_url1")
fw.AddChannel("ch2", "webhook_url2")
// ...

fw.Send(
    context.Background(),
    "ch1", // 后续只需要使用频道别名就可以发送到指定的告警规则了
    map[string]string{
        "hello": "world"
    }
)
```
