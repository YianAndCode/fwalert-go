package fwalert

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type fwalert struct {
	channels map[string]string
}

func New() *fwalert {
	return &fwalert{
		channels: make(map[string]string),
	}
}

func (f *fwalert) AddChannel(channel, webhookUrl string) {
	f.channels[channel] = webhookUrl
}

func (f *fwalert) GetChannel(channel string) string {
	return f.channels[channel]
}

func (f *fwalert) RemoveChannel(channel string) {
	delete(f.channels, channel)
}

func (f *fwalert) Send(ctx context.Context, channel string, data interface{}) error {
	if url, ok := f.channels[channel]; ok {
		return f.SendAlert(ctx, url, data)
	}
	return fmt.Errorf("channel not found: %s", channel)
}

func (f *fwalert) SendAlert(ctx context.Context, webhookUrl string, params interface{}) error {
	datas, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", webhookUrl, bytes.NewBuffer(datas))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("%s", resp.Status)
	}

	return nil
}
