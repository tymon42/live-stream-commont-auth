# [WIP] live-stream-commont-auth
I came up with this idea when I wanted to build a Bilibili third party live-interactive tool. And [@Cunoe](https://github.com/CUNOE) and I have already implemented a draft version and it is working well on our tool.  
Now I am working on open source version which is gonna be more universal.   
# What's this?  
It's a kind of authorization solution design for 3rd party app in live streaming platform like Bilibili or Youtube and so on.  

You can easily setup a self-host auth service and auth and generate JWT token for your own app. Only registered user on the specific live streaming platform can pass thourgh the auth process.  

# How it work?
The main idea of the live-stream-commont-auth is capturing on infomation that post by registered user on platform(like bilibili or youtube) only. The infomation must has the uid that represents the perticular user.  

For example, danmu in bilibili live has an excellent data structure for the identity authorzing. It includes bili user's uid and a 20 chars space to convey a message. If a user wants to login to a thired party platform, we could sent the user a verification code in 20 chars and the user should send it to a chose bilibili live room as a danmu. So we shall get a danmu message via a bili's danmu WS connection. And the message contains both the buid and the verification code.  

We believe that kind of message is safe enough for login action.  

## Example  
### Bilibili
WIP