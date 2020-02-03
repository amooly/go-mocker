package main

type ServerConfig struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type RequestConfig struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

type ResponseConfig struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

type RouteConfig struct {
	Request  RequestConfig  `json:"request"`
	Response ResponseConfig `json:"response"`
}

type Config struct {
	Server ServerConfig  `json:"server"`
	Routes []RouteConfig `json:"routes"`
}
