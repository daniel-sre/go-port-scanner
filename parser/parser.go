package parser

import (
	"strings"
	"fmt"
	"strconv"
)

func ParsePort(p string) ([]int, error) {
	var ports []int
	for _, part := range strings.Split(p,","){
		part = strings.TrimSpace(part)

		if strings.Contains(part,"-"){

			if strings.Count(part,"-") != 1{
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			rangeParts := strings.Split(part,"-")
			if len(rangeParts) != 2 || rangeParts[0] == "" || rangeParts[1] == ""{
				return nil,fmt.Errorf("invalid range format '%s' (expected start-end)", part)
			}

			portStart, err1 := strconv.Atoi(rangeParts[0])
			portEnd, err2 := strconv.Atoi(rangeParts[1])

			if err1 != nil || err2 != nil  {
				return nil, fmt.Errorf("invalid range '%s' (must be numbers)", part)
			}
			if portStart < 1 || portEnd > 65535 {
				return nil, fmt.Errorf("port out of range '%s' (1-65535)", part)
			}

			if portStart > portEnd {
				return nil, fmt.Errorf("invalid range '%s' (start > end)", part)
			}

			for i := portStart;i <= portEnd; i++{
				ports = append(ports,i)
			}
			continue
		}
		port, err := strconv.Atoi(part)
		if err != nil {
			return nil,fmt.Errorf("invalid single port: %s", part)
		}
		ports = append(ports, port)
	}
	return ports,nil
}