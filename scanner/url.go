package scanner

import (

	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"go-port-scanner/model"
	"go-port-scanner/util"

)

func getSchemes(port int, fallback bool) []string {

	if fallback {
		return []string{"https","http"}
	}

	switch port {
		case 443:
			return []string{"https"}
		case 80:
			return []string{"http"}
		default:
			return []string{"https","http"}
		}

}


func ScanUrl(host string, port int, fallback bool) []model.Result {
	var results []model.Result

	input := host

	if !strings.HasPrefix(host,"http://") && !strings.HasPrefix(host, "https://") {
		input = "https://" + input
	}

	u, err := url.Parse(input)
	if err != nil {
		return []model.Result {
			{
			Type: "url",
			Url: host,
			Status: "FAIL",
			Err: util.ErrString(err),
			},
		}
	}
	hostname := u.Hostname()

	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	
	var lastErr error
	success := false

	schemes := getSchemes(port, fallback)

	for _,s := range schemes {
		newUrl := fmt.Sprintf("%s://%s:%d", s,hostname,port)

		start := time.Now()
		resp, err := client.Get(newUrl)
		end := time.Since(start)
		if err != nil {
			lastErr = err 
			continue
		}
		success = true
		resp.Body.Close()
		results = append(results, model.Result{
			Type: "URL",
			Url: newUrl,
			Status: "OK",
			Code: resp.StatusCode,
			Time: end.Seconds(),
			Err: util.ErrString(err),
		})

		break
	}
	if !success{
		results = append(results, model.Result{
			Type: "URL",
			Url: fmt.Sprintf("%s:%d", hostname, port),
			Status: "FAIL",
			Err: util.ErrString(lastErr),
		})
	}
	
	return results
}