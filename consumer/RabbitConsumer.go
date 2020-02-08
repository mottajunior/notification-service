package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)


//procurando corrida => consume  "all_motoristas_devices_tokens/mensagem "corrida com origem X e destino Y para cliente rafael" => send notification for all motoristas.
//corrida encontrada => consume "cliente_token/mensagem: "motorista bruno esta a caminho" => send notification  for specific client
//objeto entregue => consume "cliente_token/mensagem: "objeto entregue, corrida finalizada" => send notification for specific client

func ConsumeQueue(){

	fmt.Println("Consume has started")
	url := os.Getenv("AMQP_URL")

	if url == "" {
		url = "amqp://guest:guest@localhost:5672"
	}

	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()

	msgs, err := channel.Consume("test", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	for msg := range msgs {
		//token,msg := formatMessage(msg.body)
		//sendNotification(token,msg)
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(false)
	}

	defer connection.Close()

}














//firebase-example:
// Create the message to be sent.
//msg := &fcm.Message{
//	To: "sample_device_token",
//	Data: map[string]interface{}{
//		"foo": "bar",
//	},
//}
//// Create a FCM client to send the message.
//client, err := fcm.NewClient("sample_api_key")
//if err != nil {
//	log.Fatalln(err)
//}
//// Send the message and receive the response without retries.
//response, err := client.Send(msg)
//if err != nil {
//	log.Fatalln(err)
//}
//
//log.Printf("%#v\n", response)

