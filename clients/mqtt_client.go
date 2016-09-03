package clients

import (
	"fmt"
	"strings"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/mainflux/mainflux-lite/routes"
)

type MqttConn struct {
	Opts *mqtt.ClientOptions
	Client mqtt.Client
}

//define a function for the default message handler
var msgHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

	s := strings.Split(msg.Topic(), "/")
	chanId := s[len(s)-1]
	status := routes.WriteChannel(chanId, msg.Payload())

	fmt.Println(status)
}

func (mqc *MqttConn) MqttSub() {
	// Create a ClientOptions struct setting the broker address, clientid, turn
	// off trace output and set the default message handler
	mqc.Opts = mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
	mqc.Opts.SetClientID("mainflux")
	mqc.Opts.SetDefaultPublishHandler(msgHandler)

	//create and start a client using the above ClientOptions
	mqc.Client = mqtt.NewClient(mqc.Opts)
	if token := mqc.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to all channels of all the devices and request messages to be delivered
	// at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := mqc.Client.Subscribe("devices/+/channels/+", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}
}
