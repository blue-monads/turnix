package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/adrg/xdg"
)

var (
	// "http://localhost:7703/z/project/simplerat/1/device/finish-setup?token=xyz"
	urlBase = ""

	config = &Config{}
)

type Config struct {
	BaseURL          string `json:"base_url"`
	AuthSessionToken string `json:"auth_session_token"`
	AuthRefreshToken string `json:"auth_refresh_token"`
}

type DeviceRegisterResponse struct {
	RefreshToken string `json:"refreshToken"`
	SessionToken string `json:"sessionToken"`
}

func main() {

	if os.Getenv("CNC_SERVER_URL") != "" {
		urlBase = os.Getenv("CNC_SERVER_URL")
	}

	// parse flag and get url as fallback

	if urlBase == "" {
		flag.StringVar(&urlBase, "url", urlBase, "CNC server URL")
		flag.Parse()
	}

	purl, err := url.Parse(urlBase)
	if err != nil {
		panic(err)
	}

	uquery := purl.Query()
	registerToken := uquery.Get("token")
	uquery.Del("token")
	purl.RawQuery = ""

	purl.Path = strings.Replace(purl.Path, "finish-setup", "", 1)

	configFolder := xdg.DataHome + "/simplerat"
	configFile := configFolder + "/config.json"

	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		os.MkdirAll(configFolder, 0755)
	}

	if _, err := os.Stat(configFile); os.IsNotExist(err) {

		if urlBase == "" {
			slog.Error("No Register URL provided and no config file found")
			return
		}

		tryCount := 0

		for {

			if tryCount > 5 {
				slog.Error("Failed to register device")
				break
			}

			if tryCount > 0 {
				time.Sleep(5 * time.Second)
			}

			tryCount++

			reqData := map[string]any{
				"token": registerToken,
			}

			payloadBytes, err := json.Marshal(reqData)
			if err != nil {
				slog.Error("Error encoding request", slog.String("error", err.Error()))
				continue
			}

			req, err := http.NewRequest("POST", purl.String()+"finish-setup", bytes.NewReader(payloadBytes))
			if err != nil {
				slog.Error("Error creating request", slog.String("error", err.Error()))
				continue
			}

			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)

			if err != nil {
				slog.Error("Error sending request", slog.String("error", err.Error()))
				continue
			}

			if resp.StatusCode != 200 {
				slog.Error("Error response", slog.Int("status", resp.StatusCode))
				break
			}

			rpayload := &DeviceRegisterResponse{}

			err = json.NewDecoder(resp.Body).Decode(rpayload)
			if err != nil {
				slog.Error("Error decoding response", slog.String("error", err.Error()))
				break
			}

			rconfig := &Config{
				BaseURL:          purl.String(),
				AuthSessionToken: rpayload.SessionToken,
				AuthRefreshToken: rpayload.RefreshToken,
			}

			configBytes, err := json.Marshal(rconfig)
			if err != nil {
				slog.Error("Error encoding config", slog.String("error", err.Error()))
				break
			}

			err = os.WriteFile(configFile, configBytes, 0644)
			if err != nil {
				slog.Error("Error writing config", slog.String("error", err.Error()))
				break
			}

		}

	}

	configBytes, err := os.ReadFile(configFile)
	if err != nil {
		slog.Error("Error reading config", slog.String("error", err.Error()))
		return
	}

	err = json.Unmarshal(configBytes, config)
	if err != nil {
		slog.Error("Error decoding config", slog.String("error", err.Error()))
		return
	}

	agentSvc := &AgentService{
		configHome:       configFolder,
		configFile:       configFile,
		baseURL:          config.BaseURL,
		registerToken:    registerToken,
		authSessionToken: config.AuthSessionToken,
		authRefreshToken: config.AuthRefreshToken,
	}

	err = agentSvc.Run()
	if err != nil {
		slog.Error("Error running agent service", slog.String("error", err.Error()))
		return
	}

}
