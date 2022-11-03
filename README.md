# [WIP] live-stream-commont-auth
I came up with this idea when I want to build a Bilibili third party live interactive tool. And [@Cunoe](https://github.com/CUNOE) and I have already implement a draf version and it is working well on our tool.  
Now I am working on open source version which gonna be more Universal.   
# What's this?  
It's a kind of authorize solution design for 3rd party app in live streaming platform like Bilibili or Youtube and so on.  

You can easy setup a self-host auth service and auth and generate JWT token for your own app. And only registied user on the sepcific live streaming platform can pass thourgh the auth process.  

# How it work?
The main idea of live-stream-commont-auth is capture on infomation that post by regiested user on platform(like bilibili or youtube) only. The infomation must has the uid that represent the perticular user.  

For example, danmu in bilibili live has an excelent data struct for the identity authrozing. It includes bili user's uid and a 20 chars space to convey a message. If we could send the user want login in our platform, we could sent the user a verification code in 20 chars and the user send it to a bilibili live room as a danmu. So we shall get a danmu message via a bili's danmu ws connection. And the message contain both the buid and verification code.  

We trust that kind of message is safe enough for login action.  

## Example  
### Bilibili
WIP