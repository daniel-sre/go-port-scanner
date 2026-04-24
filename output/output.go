package output

import (

	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"go-port-scanner/model"
)

func Output(result <-chan model.Result,writer io.Writer,jsonOutput,pretty bool){

	bufWrite := bufio.NewWriter(writer)
	defer bufWrite.Flush()

	if jsonOutput{
		encoder := json.NewEncoder(bufWrite)
		encoder.SetEscapeHTML(false)
		if pretty {
			encoder.SetIndent("","  ")
		}

		for res := range result{
			encoder.Encode(res)
		}
	} else {
		for res := range result{
			if res.Type == "URL" {
				fmt.Fprintf(bufWrite,"%-30s %-10s %-10d %.2fs\n",
					res.Url,
					res.Status,
					res.Code,
					res.Time,
				)
			} else {
				fmt.Fprintf(bufWrite,"%-15s %-5d %-10s %.2f\n",
					res.Host,
					res.Port,
					res.Status,
					res.Time,
				)	
			}
		}
	}
}