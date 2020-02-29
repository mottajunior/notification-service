package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"strings"
	"log"
	"github.com/appleboy/go-fcm"
	// "github.com/NaySoftware/go-fcm"
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
		notificationBody := string(msg.Body)
		fmt.Println("message received: " + notificationBody)
		SplitedNotification := strings.Split(notificationBody, "/")
		tokens := SplitedNotification[0]
		message := SplitedNotification[1]
		raceId := SplitedNotification[2]

		
		AllTokens :=  strings.Split(tokens, "#!#")

		if len(AllTokens) > 1{
			SendNotificationForManyTokens(AllTokens,message,raceId)
		} else{
			SendNotificationForOneToken(AllTokens[0],message)
		}

		msg.Ack(false)
	}
	defer connection.Close()
}

func SendNotificationForManyTokens(tokens []string, message string, raceId string){
	fmt.Println("ENVIANDO NOTIFICACAO PARA TODOS MOTORISTAS")
	fmt.Println("TOKEN ==>",tokens)
	fmt.Println("MENSAGEM ==>", message)
	fmt.Println("raceId ==>",raceId)

	//Create the message to be sent.
	msg := &fcm.Message{
		To: tokens[0],
		Data: map[string]interface{}{
			"message": message,
			"title": "Corrida Encontrada",
			"body": raceId	,		
		},
	}	
	// Create a FCM client to send the message.
	client, err := fcm.NewClient("AAAAXSrd1t0:APA91bEveSugmY7A9ismkVFE9GKPRzmYb0CoGmf57kv2SwwNW4VfyiYAApLuTCr-A2haHKx3A5XneTo2b97vssBObZ9Yiko1Akd1ZOyMdjeVhF4mGXhAuPSdFt7djJpahDur68o8CD1y")
	if err != nil {
		log.Fatalln(err)
	}
	
	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("vai printar a resposta da notificação")
	log.Printf("%#v\n", response)
}

func SendNotificationForOneToken(token string, message string){
	fmt.Println("ENVIANDO NOTIFICACAO PARA UMA PESSOA")
	fmt.Println("TOKEN ==>",token)
	fmt.Println("MENSAGEM ==>", message)
}
