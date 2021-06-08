package messaging

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type IMessagingClient interface {
	// used to send an OffchainProcessMessage
	// message - message payload
	// recipient - address of the recipient (i.e. NATS server url) taken from workgroup
	// token - token used to authenticate (i.e. NATS server token) taken from the workgroup
	SendMessage(message string, recipient string, token string)

	// used to receive messages sent by other participants to our nats server
	// serverUrl - local server url
	// topic - listening topic
	// onMessageReceived - callback function
	Subscribe(serverUrl string, topic string, onMessageReceived func(string, string))
}

type NatsMessagingClient struct {
}

func (client *NatsMessagingClient) SendMessage(message string, recipient string, token string) {
	// https://docs.nats.io/developing-with-nats/security/token
	nc, err := nats.Connect(token + "@" + recipient)

	if err != nil {
		// TODO: Add logging
	}

	defer nc.Close()

	// TODO: https://docs.nats.io/developing-with-nats/sending/replyto
	nc.Publish("TODO: subject", []byte(message))
}

func (client *NatsMessagingClient) Subscribe(serverUrl string, topic string, onMessageReceived func(string, string)) {
	// https://docs.nats.io/developing-with-nats/security/token
	nc, err := nats.Connect(serverUrl)

	if err != nil {
		// TODO: Add logging
	}

	defer nc.Close()

	nc.Subscribe(topic, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		onMessageReceived(string("TODO: m.Sender"), string(m.Data))
	})
}
