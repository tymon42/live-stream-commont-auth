# How live-stream-comment-auth works?
English | [简体中文](how-it-works_zh.md)
## Main idea
`live-stream-comment-auth` principle is to capture the information sent by registered users on the live platform. The information must have a unique user ID (uid) representing a specific user.

For example, the danmu in bilibili live has a well-designed authentication data structure. It includes the uid of bili users and a 20-character space to pass messages. If the user wants to log in to a third-party platform, we can send a 20-character verification code to the user, and the user should send it to the selected bilibili live room as danmu. Therefore, we will capture the danmu message through the danmu machine. The message contains buid and the verification code.

## How user login or signup?  
1. User request a verification code from the third party platform, we call it "vcode".  
2. User opens a bilibili live room and send the vcode as a Danmu.  
3. The live-stream-comment-auth worker program captures the Danmu and check if the vcode is valid.  
4. If the vcode is valid, the worker program will send a JWT token to the third party platform.  
5. User can use the JWT token to login to the third party platform.  
