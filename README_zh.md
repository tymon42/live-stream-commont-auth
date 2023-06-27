# live-stream-comment-auth 直播评论身份认证登录系统
[English](README.md) | 简体中文

我想建立一个Bilibili第三方的直播互动工具时，想到了这个想法。我和[@Cunoe](https://github.com/CUNOE)已经实现了一个草案版本，并且在我们的工具上运行良好。现在，我正在开发一个更为通用的开源版本，并将提供需要付费的服务（但当然也会有免费的试用范围）。

# live-stream-comment-auth 这是什么？

这是直播平台（如哔哩哔哩或YouTube）第三方应用程序的授权解决方案设计。您可以轻松使用RESTFUL API服务进行授权并为您自己的应用程序（或弹幕游戏）生成JWT令牌。只有在特定直播平台上注册的用户才能通过授权流程。

我们相信，此类消息对于登录操作已足够安全。

请查看[如何运作](docs/how-it-works_zh.md)以获取更多详细信息。

# 使用官方Saas服务

更多详情，请参阅[Saas文档](docs/saas_zh.md)。
