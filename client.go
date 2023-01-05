package main

import (
	"fmt"
	//导入对应MQTT包，没有的go get安装
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

// 定义函数回调
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	//创建ClientOptions struct 设置 the broker address, clientid,
	//关闭trace output and 设置 the default message handler
	//有用户名密码的还需要设置用户名密码
	//opts := MQTT.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883")
	opts := MQTT.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opts.SetClientID("go-demo")
	//opts.SetUsername("emqx")
	//opts.SetPassword("public")
	opts.SetUsername("admin")
	opts.SetPassword("zd5201314..")
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe("go-mqtt/demo", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
	//from the server after sending each message
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := c.Publish("go-mqtt/demo", 0, false, text)
		token.Wait()
	}

	time.Sleep(3 * time.Second)

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("go-mqtt/demo"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}
