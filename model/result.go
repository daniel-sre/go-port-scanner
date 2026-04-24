package model

type Job struct{
	Type string
	Url string
	IP string
	Port int
}

type Result struct {
	Type string		`json:"type"`
	Host string		`json:"host,omitempty"`
	Port int		`json:"port,omitempty"`
	Url string		`json:"url,omitempty"`
	Status string	`json:"status"`
	Code int		`json:"code,omitempty"`
	Time float64	`json:"time"`
	Err  string		`json:"error,omitempty"`
}
