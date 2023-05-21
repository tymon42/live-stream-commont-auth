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
	Buid int `json:"buid"`
}

type Charge struct {
	Buid   int `json:"buid"`
	Amount int `json:"amount"`
}

func main() {
	// 主机名
	var api string
	// 端口号
	var roomid string
	// vcode 前缀
	var vcodePrefix string
	// vcode 后缀长度
	var vcodeSuffixLen string

	flag.StringVar(&api, "api", "http://127.0.0.1:8888", "bili-danmu-auth 服务 api,默认为 http://127.0.0.1:8888")
	flag.StringVar(&roomid, "r", "", "被侦听直播间房间号,默认为空")
	flag.StringVar(&vcodePrefix, "p", "vc-", "vcode 前缀,默认为 vc-")
	flag.StringVar(&vcodeSuffixLen, "l", "6", "vcode 后缀长度,默认为6")

	// 从arguments中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp。
	flag.Parse()

	log.SetLevel(log.DebugLevel)
	httpClient := resty.New()
	c := client.NewClient(roomid)
	reg1 := regexp.MustCompile(vcodePrefix + `\S{` + vcodeSuffixLen + `}`)
	regDev := regexp.MustCompile(`开发者登录或注册-` + `\S{11}`)
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	//弹幕事件
	c.OnDanmaku(func(danmaku *message.Danmaku) {
		// fmt.Println(danmaku.Content)
		if regDev.MatchString(danmaku.Content) {
			fmt.Println("dev login or signup")
			// 处理 vcode, 向 /api/v1/vcode/{vcode}/verify 发送请求
			res, err := httpClient.R().
				SetBody(SubmitVCode{Buid: danmaku.Sender.Uid}).
				Post(api + "/api/v1/vcode/" + regDev.FindString(danmaku.Content))
			if err != nil {
				log.Infoln("submit vcode failed, err: ", err)
			}
			if res.StatusCode() == 200 {
				fmt.Println("submit vcode success")
			} else {
				fmt.Println("submit vcode failed, status code: ", res.StatusCode())
				fmt.Printf("res: %v\n", res)
			}
		}

		result1 := reg1.FindAllStringSubmatch(danmaku.Content, -1)
		if len(result1) > 0 {
			fmt.Println("vcode = ", result1[0][0])
			// 处理 vcode, 向 /api/v1/vcode/{vcode}/verify 发送请求
			res, err := httpClient.R().
				SetBody(SubmitVCode{Buid: danmaku.Sender.Uid}).
				Post(api + "/api/v1/vcode/" + result1[0][0])
			if err != nil {
				log.Infoln("submit vcode failed, err: ", err)
			}
			if res.StatusCode() == 200 {
				fmt.Println("submit vcode success")
			} else {
				fmt.Println("submit vcode failed, status code: ", res.StatusCode())
				fmt.Printf("res: %v\n", res)
			}
		}
	})

	c.OnGift(func(gift *message.Gift) {
		fmt.Printf("gift: %+v\n", gift)

		if gift.CoinType != "gold" {
			return
		}

		res, err := httpClient.R().
			SetBody(Charge{Buid: gift.Uid, Amount: gift.Price / 10}).
			Post(api + "/api/v1/recharge/")
		if err != nil {
			log.Infoln("charge failed, err: ", err)
		}
		if res.StatusCode() == 200 {
			fmt.Println("charge success, buid: ", gift.Uid, " amount: ", gift.Price/10)
		} else {
			fmt.Println("charge failed, status code: ", res.StatusCode())
			fmt.Printf("res: %v\n", res)
		}

	})

	err := c.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("started")
	select {}
}
