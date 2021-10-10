/*
@Time : 2021/6/23 19:09
@Author : ZhaoJunfeng
@File : producer
*/
package main

import (
    "fmt"
    "github.com/Shopify/sarama"
)

/*
	初始化NewConfig配置 sarama.NewConfig
	创建生产者sarama.NewSyncProducer
	创建消息sarama.ProducerMessage
	发送消息client.SendMessage
*/
func main() {

    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.Return.Successes = true

    config.Version = sarama.V0_11_0_2

    msg := &sarama.ProducerMessage{}
    msg.Topic = "gently_ai_face_property"
    //msg.Value = sarama.StringEncoder(`hostname"ali-loglib-119-45filenamepv_1000002{"pageid":"直播讲座","appid":"1000002","hostname":"ali-loglib-119-45","timestamp":1534759693455,"clits":1534759697839,"data":{"logorder":"14","currenthref":"http:\/\/biglive.xueersi.com\/Lecturelives\/newDetail\/64394\/e2bb2473367ec506eb237392c61fdaca\/","ajaxhref":"http:\/\/biglive.xueersi.com\/Lecturelives\/userOnline\/715-64394\/cccafc7b83544252f422e68c2669cf0b\/?time=1534759697729"},"xesid":"51c1484f4ff3f3e4698c7af91cc929ab","clientid":"8","prelogid":"a5a70dbeddfaf49485aec3e84d715854","userid":"3361143","logid":"c36e4ee5d15d87088762fdac3ec54147"}`)

    //msg.Value = sarama.StringEncoder(`{"clientid":14,"appid":"1001128","sign":"f8d10944364d888386abf9419a8830f4","clits":1535361573768,"data":{"sysinfo":{"model":"iPhone 6","pixelRatio":2,"windowWidth":375,"windowHeight":603,"system":"iOS 10.0.1","language":"zh_CN","version":"6.6.3","screenWidth":375,"screenHeight":667,"SDKVersion":"1.9.91","brand":"devtools","fontSizeSetting":16,"benchmarkLevel":1,"batteryLevel":100,"statusBarHeight":20,"platform":"devtools"},"urlparams":{},"scene":1001,"logtype":"click","appname":"学而思网校","enterpage":"pages/loading/index","url":"pages/index/index","referer":"","rsd":"","logorder":0,"siteid":"","absiteid":"","sourceid":"","pageuid":"","pagetype":"","ajaxhref":"","aid":"","traceid":"","log_from":"advert","client_type":"touch","event_name":"click","advert_info":{"order":0,"ad_id":"5","material_id":29,"grade_id":3,"extra":null,"dnf":{"Adslot_id":13,"grade":3}}},"xesid":"e97818e097eb5490c34c0bf680249fc3","userid":"2383743","sessid":"1bd52168604d2a83249f5482034ff095"}`)

    var logstring string = `{"service":"face","call_name":"face_property",
  "args":{
    "image_url":"https://ai-audit-prison-class.oss-cn-beijing.aliyuncs.com/xiaozuke/20210528/F96E79FF11174D318A17016C0B6762EA_379181037419__uid_s_21599056__uid_e_video_20210528104714164.jpg",
    "live_id":123456,
    "user_id":2345,
    "extra_info":"This is extra_info",
     "callback":"https://api.xueersi.com/gentlyapi/middle_platform/ai/resource/callback"
  }
}`
    msg.Value = sarama.StringEncoder(logstring)
    client, err := sarama.NewSyncProducer([]string{"10.90.73.26:9092"}, config)
    if err != nil {
        fmt.Println("producer close, err:", err)
        return
    }

    defer client.Close()

    pid, offset, err := client.SendMessage(msg)
    if err != nil {
        fmt.Println("send message failed,", err)
        return
    }

    fmt.Printf("pid:%v offset:%v\n", pid, offset)

}
