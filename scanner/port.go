package scanner

import (
	"time"
	"net"
	"strconv"
	"strings"
	"go-port-scanner/model"
	"go-port-scanner/util"
)



func ScanPort(ip string, port int) model.Result {
	start := time.Now()
	address := net.JoinHostPort(ip, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp",address, 3 * time.Second)

	end := time.Since(start)
	if err != nil{

		if nerr, ok := err.(net.Error); ok && nerr.Timeout(){
			return model.Result{
				Type: "PORT",
				Host: ip,
				Port: port,
				Status: "FILTERED",
				Time: end.Seconds(),
				Err: util.ErrString(err),
			}
					
		}

		if strings.Contains(err.Error(),"connection refused"){
			return model.Result{
				Type: "PORT",
				Host: ip,
				Port: port,
				Status: "CLOSED",
				Time: end.Seconds(),
				Err: util.ErrString(err),
			}
		}
			
		return model.Result{
			Type: "PORT",
			Host: ip,
			Port: port,
			Status: "UNKNOWN",
			Time: end.Seconds(),
			Err: util.ErrString(err),
		}
	}
	defer conn.Close()
	return model.Result{
			Type: "PORT",
			Host: ip,
			Port: port,
			Status: "OPEN",
			Time: end.Seconds(),
			Err: util.ErrString(err),
		}
}