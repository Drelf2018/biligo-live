package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SpiData struct {
	Code int `json:"code"`
	Data struct {
		B3 string `json:"b_3"`
		B4 string `json:"b_4"`
	} `json:"data"`
	Message string `json:"message"`
}

func GetBuvid3() string {
	resp, err := http.Get("https://api.bilibili.com/x/frontend/finger/spi")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var data SpiData
	if json.NewDecoder(resp.Body).Decode(&data) != nil {
		return ""
	}

	return data.Data.B3
}

type DanmuInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Group            string  `json:"group"`
		BusinessID       int     `json:"business_id"`
		RefreshRowFactor float64 `json:"refresh_row_factor"`
		RefreshRate      int     `json:"refresh_rate"`
		MaxDelay         int     `json:"max_delay"`
		Token            string  `json:"token"`
		HostList         []struct {
			Host    string `json:"host"`
			Port    int    `json:"port"`
			WssPort int    `json:"wss_port"`
			WsPort  int    `json:"ws_port"`
		} `json:"host_list"`
	} `json:"data"`
}

func GetKey(room int) string {
	resp, err := http.Get(fmt.Sprintf("https://api.live.bilibili.com/xlive/web-room/v1/index/getDanmuInfo?id=%d", room))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var data DanmuInfo
	if json.NewDecoder(resp.Body).Decode(&data) != nil {
		return ""
	}

	return data.Data.Token
}
