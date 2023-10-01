// Auther: Sam Li
import { KeepLiveWS , toMessageData, getLongRoomId } from 'tiny-bilibili-ws'
import minimist from 'minimist'
import fetch from 'node-fetch'

const baseUrl = "https://danmu-auth.fly.dev"
const argv = process.argv.slice(2)
const argv_minimist = minimist(argv)
const roomid = argv_minimist['roomid']
const uid = argv_minimist['uid']
const key = argv_minimist['key']
const ApiKey = argv_minimist['apiKey']
const { data } = await getLongRoomId(Number(roomid))
const room = data.room_id
const option = {
  uid: parseInt(uid),
  key: key,
}
const ws = new KeepLiveWS(room,option)
  
ws.runWhenConnected(() => {
  console.log(`正在监听 ${room}`)
})
  
  
ws.on('DANMU_MSG', (danmu) => {
  
  
  const content = toMessageData(danmu).info[1]
  const userinfo = toMessageData(danmu).info[2]
  console.log(`uid:${userinfo[0]}   名字:${userinfo[1]}   msg:${content}`)
  
  const vcReg = /^vc-\S{10}$/;
    // 匹配 开发者登录或注册- 开头的字符串, 且后面跟着11位数字或字母
  const loginReg = /^开发者登录或注册-\S{11}$/;
    // 如果 msg.body.content 字符串匹配到了 vcReg 或 loginReg 规则
  if (vcReg.test(content) || loginReg.test(content)) {
      const vcode = content;
      console.log(`[vcode] ${vcode}`);
      const uid = userinfo[0];
      const url = `${baseUrl}/api/v1/vcode/${vcode}`;
      // console.log(`[url] ${url}`)
      const body = {
          buid: uid,
          api_key: ApiKey,
      };
  
        
      fetch(url, {
          method: "POST",
          body: JSON.stringify(body),
          headers: {
              'Content-Type': 'application/json'
          }
      }).then(res => res.json()).then(res => {
            console.log(res);
      });
  }
})
  
ws.on('error', (e) => {
  console.error('错误: ', e)
})
  
ws.on('close', () => {
  console.log(`退出监听 ${room}`)
})
