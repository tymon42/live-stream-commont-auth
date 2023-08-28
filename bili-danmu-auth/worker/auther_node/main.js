"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const blive_message_listener_1 = require("blive-message-listener");
const args = require('minimist')(process.argv.slice(2));
const baseUrl = "https://danmu-auth.fly.dev";
const roomid = args['roomid'];
const ApiKey = args['apiKey'];
const handler = {
    onIncomeDanmu: (msg) => {
        console.log(`[new msg]`);
        console.log(`[danmu] ${msg.body.content}`);
        // 定义两个正则表达式规则
        // 匹配 vc- 开头的字符串, 且后面跟着10位数字或字母
        const vcReg = /^vc-\S{10}$/;
        // 匹配 开发者登录或注册- 开头的字符串, 且后面跟着11位数字或字母
        const loginReg = /^开发者登录或注册-\S{11}$/;
        // 如果 msg.body.content 字符串匹配到了 vcReg 或 loginReg 规则
        if (vcReg.test(msg.body.content) || loginReg.test(msg.body.content)) {
            const vcode = msg.body.content;
            console.log(`[vcode] ${vcode}`);
            const uid = msg.body.user.uid;
            const url = `${baseUrl}/api/v1/vcode/${vcode}`;
            // console.log(`[url] ${url}`)
            const body = {
                buid: uid,
                api_key: ApiKey,
            };
            // console.log(`[body] ${JSON.stringify(body)}`);
            // 向 /api/v1/vcode/{vcode} 发送 POST 请求, Body 为 JSON 格式的字符串, 内容为 { "buid": msg.body.uid, "api_key": ApiKey }, vcode = msg.body.content
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
    },
    onOpen: () => {
        console.log('connected');
    },
    onStartListen: () => {
        console.log('start listen');
    },
    onError: (err) => {
        console.log(`error: ${err}`);
    }
};
const instance = (0, blive_message_listener_1.startListen)(roomid, handler, {
    ws: {
        host: 'wss://tx-bj-live-comet-03.chat.bilibili.com',
    }
});
console.log(`start listen room ${instance.roomId}`);
// instance.close()
