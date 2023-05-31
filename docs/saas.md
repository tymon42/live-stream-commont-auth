# Saas service
If you don't want to host your own live-stream-comment-auth service, you can use our saas service.

We provide a free tier for you to try out. And you can also pay for a higher tier to get more features.

## Pricing
| Tier | Price    | Features                 |
| ---- | -------- | ------------------------ |
| Free | Free     | 100 requests at begining |
| Paid | 1 Batery | 100 requests            |

PS:
1. You can send 5 Danmu in any live room to get 1 battery as a gift from Bilibili each day. 
2. It will be a subscription service in the futrue. But now, you can contact us if you want to pay for a higher tier in a much lower price. 

## How to use
You can also use our RESTFUL API https://danmu-auth.fly.dev to register or login.
Check our [API document](https://krzwk4bbxe.apifox.cn), Password: `UNbP8vcJ` , for more details.
链接: https://krzwk4bbxe.apifox.cn  访问密码: UNbP8vcJ

### 1. Register or login as a developer

1. Request the login verification code. 
2. Send the verification code to a bilibili live room as a Danmu.
3. Request the JWT token with the verification code.

PS: You need the JWT token to request all the developer related API.

### 2. How to allow your users to login
Once you have the JWT token, you can the Get /vcode API to generate a verification code for your user by providing the devloper key.

1. Request the verification code for user.
2. User send the verification code to a bilibili live room as a Danmu.
3. Our service will watching the live room and log the verification code sending event.
4. User or your app request the JWT token with the verification code and client_id.
5. Our service will check the verification code sending event and return the JWT token if the verification code is valid.
6. Your app can use the JWT token to request the Get /jwt API to get the user info.
