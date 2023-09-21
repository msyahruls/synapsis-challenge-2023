package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"synapsis-challange/model/domain"

	_ "github.com/joho/godotenv/autoload"
)

var (
	WsHost           = os.Getenv("WSS_HOST")
	WsType           = os.Getenv("WSS_TYPE")
	WsApiKey         = os.Getenv("WSS_API_KEY")
	WsChannelStart   = os.Getenv("WSS_CHANNEL_START")
	WsChannelEndChat = os.Getenv("WSS_CHANNEL_END_CHAT")
	WsMethod         = os.Getenv("WSS_METHOD")
	WsStatus         = os.Getenv("WSS_STATUS")
)

func PublishToWebSocket(dataWS interface{}, channel, method string) error {
	payload := domain.DataWS{
		Method: method,
		Params: domain.ParamsWS{
			Channel: channel,
			Data:    dataWS,
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		PanicIfError(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", WsHost, body)
	if err != nil {
		PanicIfError(err)
	}
	req.Header.Set("Content-Type", WsType)
	req.Header.Set("Authorization", WsApiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		PanicIfError(err)
	}

	logWs := fmt.Sprintf("publish to websocket success, channel : %s", channel)

	log.Println(logWs)

	defer resp.Body.Close()

	return nil
}
