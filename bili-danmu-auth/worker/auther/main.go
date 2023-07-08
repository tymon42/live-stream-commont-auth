package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/Akegarasu/blivedm-go/client"
	"github.com/Akegarasu/blivedm-go/message"
	_ "github.com/Akegarasu/blivedm-go/utils"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type SubmitVCode struct {
	ApiKey string `json:"api_key"`
	Buid   int    `json:"buid"`
}

type Charge struct {
	ApiKey string `json:"api_key"`
	Buid   int    `json:"buid"`
	Amount int    `json:"amount"`
}

func main() {
	// 主机名
	var api string
	// 端口号
	var roomid string
	// uid
	var uid int
	// vcode 前缀
	var vcodePrefix string
	// vcode 后缀长度
	var vcodeSuffixLen string
	// worker 服务 api_key
	var apiKey string

	flag.StringVar(&api, "api", "http://127.0.0.1:8888", "bili-danmu-auth 服务 api,默认为 http://127.0.0.1:8888")
	flag.StringVar(&roomid, "r", "", "被侦听直播间房间号,默认为空")
	flag.IntVar(&uid, "u", 0, "被侦听直播间用户uid,默认为0")
	flag.StringVar(&vcodePrefix, "p", "vc-", "vcode 前缀,默认为 vc-")
	flag.StringVar(&vcodeSuffixLen, "l", "6", "vcode 后缀长度,默认为6")
	flag.StringVar(&apiKey, "k", "", "worker 服务 api_key,默认为空")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	log.SetLevel(log.DebugLevel)
	httpClient := resty.New()
	c := client.NewClient(roomid, uid)
	reg1 := regexp.MustCompile(vcodePrefix + `\S{` + vcodeSuffixLen + `}`)
	regDev := regexp.MustCompile(`开发者登录或注册-` + `\S{11}`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	//弹幕事件
	c.OnDanmaku(func(danmaku *message.Danmaku) {
		if !regDev.MatchString(danmaku.Content) && !reg1.MatchString(danmaku.Content) {
			fmt.Printf("[%s] %s\n", danmaku.Sender.Uname, danmaku.Content)
			return
		}

		// fmt.Println(danmaku.Content)
		if regDev.MatchString(danmaku.Content) {
			fmt.Println("dev login or signup")
			// 处理 vcode, 向 /api/v1/vcode/{vcode}/verify 发送请求
			res, err := httpClient.R().
				SetBody(SubmitVCode{Buid: danmaku.Sender.Uid, ApiKey: apiKey}).
				Post(api + "/api/v1/vcode/" + regDev.FindString(danmaku.Content))
			if err != nil {
				log.Infoln("submit vcode failed, err: ", err)
			}
			if res.StatusCode() == 200 {
				fmt.Println("submit vcode success, buid: ", danmaku.Sender.Uid, " vcode: ", regDev.FindString(danmaku.Content))
			} else {
				fmt.Println("submit vcode failed, status code: ", res.StatusCode())
				fmt.Printf("res: %v\n", res)
			}
		}

		// result1 := reg1.FindAllStringSubmatch(danmaku.Content, -1)
		// if len(result1) > 0 {
		// 	// 处理 vcode, 向 /api/v1/vcode/{vcode}/verify 发送请求
		// 	res, err := httpClient.R().
		// 		SetBody(SubmitVCode{Buid: danmaku.Sender.Uid}).
		// 		Post(api + "/api/v1/vcode/" + result1[0][0])
		// 	if err != nil {
		// 		log.Infoln("submit vcode failed, err: ", err)
		// 	} else if res.StatusCode() == 200 {
		// 		fmt.Println("submit vcode success")
		// 	}
		// }

		if reg1.MatchString(danmaku.Content) {
			// 处理 vcode, 向 /api/v1/vcode/{vcode}/verify 发送请求
			res, err := httpClient.R().
				SetBody(SubmitVCode{Buid: danmaku.Sender.Uid, ApiKey: apiKey}).
				Post(api + "/api/v1/vcode/" + reg1.FindString(danmaku.Content))
			if err != nil {
				log.Infoln("submit vcode failed, err: ", err)
			} else if res.StatusCode() == 200 {
				fmt.Println("submit vcode success")
			}
		}

	})

	c.OnGift(func(gift *message.Gift) {
		if gift.CoinType != "gold" {
			return
		}

		res, err := httpClient.R().
			SetBody(Charge{Buid: gift.Uid, Amount: (gift.Price / 10) / 2, ApiKey: apiKey}). // 每 1 毛钱兑换 50 余额
			Post(api + "/api/v1/recharge/")
		if err != nil {
			log.Infoln("charge failed, err: ", err)
		} else if res.StatusCode() == 200 {
			fmt.Println("charge success, buid: ", gift.Uid, " amount: ", gift.Price/10)
		}

	})

	err := c.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("started")
	select {}
}
