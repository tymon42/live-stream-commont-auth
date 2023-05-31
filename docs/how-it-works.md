# How live-stream-comment-auth works?

# Main idea
The main idea of the live-stream-comment-auth is capturing infomation that post by registered users on live stream platform. The infomation must has the uid that represents the particular user.  

For example, Danmu in bilibili live has an excellent data structure for the identity authorzing. It includes bili user's uid and a 20 chars space to convey a message. If a user wants to login to a third party platform, we could sent the user a verification code in 20 chars and the user should send it to a chose bilibili live room as a Danmu. So we shall get a Danmu message via a bili's Danmu WS connection. And the message contains both the buid and the verification code.  

# How user login or signup?
1. User request a verification code from the third party platform, we call it "vcode".
2. User opens a bilibili live room and send the vcode as a Danmu.
3. The live-stream-comment-auth worker program captures the Danmu and check if the vcode is valid.
4. If the vcode is valid, the worker program will send a JWT token to the third party platform.
5. User can use the JWT token to login to the third party platform.
