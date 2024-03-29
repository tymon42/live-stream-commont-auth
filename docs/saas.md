# Saas service
English | [简体中文](saas_zh.md)

If you don't want to host your own live-stream-comment-auth service, you can use our saas service.

We provide a free tier for you to try out. And you can also pay for a higher tier to get more features.

## Pricing
| Tier             | Price                     | Features                      |
| ---------------- | ------------------------- | ----------------------------- |
| Free             | Free                      | 500 requests at begining      |
| Paid via Battery | 1 Baterry                 | 50 requests                   |
| Paid in cash     | ¥1.00                     | 1000 requests                 |
| Big Uploader     | ¥2.50/week or ¥5.99/month | Chose your own auth live room |

PS:
1. You can send 5 Danmu in any live room to get 1 battery as a gift from Bilibili each day. And that is enough for most of the your testing. 
2. It will be a subscription service in the futrue. But now, you can contact us if you want to pay for a higher tier in a much lower price. 

## How to use
You can also use our RESTFUL API `https://danmu-auth.fly.dev` to register or login.
Check our [API document](https://krzwk4bbxe.apifox.cn), Password: `UNbP8vcJ` , for more details.

### 1. Register or login as a developer

1. Request the login verification code. 
2. Send the verification code to the bilibili [live room](https://live.bilibili.com/12834880) as a Danmu.
3. Request the JWT token with the verification code.

PS: You need the JWT token to request all the developer API.

### 2. How to allow your users to login
Once you have the JWT token, you can the Get /vcode API to generate a verification code for your user by providing the devloper key.

1. Request the verification code for user.
2. User send the verification code to a bilibili live room as a Danmu.
3. Our service will watching the live room and log the verification code sending event.
4. User or your app request the JWT token with the verification code and client_id.
5. Our service will check the verification code sending event and return the JWT token if the verification code is valid.
6. Your app can use the JWT token to request the Get /jwt API to get the user info.
