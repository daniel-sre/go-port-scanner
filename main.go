package main

import (
	"flag"
	"fmt"
	"go-port-scanner/model"
	"go-port-scanner/output"
	"go-port-scanner/scanner"
	"go-port-scanner/parser"
	"io"
	"os"
	"sync"

)


func run() error {
	target := flag.String("t","","(IP or CIDR)")
	ports := flag.String("p","80"," 80 or 80-100 or 80,443,22")
	url := flag.String("u","","url")
	fallback := flag.Bool("fallback",false,"try both http and https")
	workers := flag.Int("w",10,"worker number")
	jsonOutput := flag.Bool("json",false,"json output")
	pretty := flag.Bool("pretty",false,"pretty json")
	fileoutput := flag.String("o","","output file")
	flag.Parse()
	
	if *url != "" && *target !="" {
		return  fmt.Errorf("Error: -t and -u cannot be used together")
	}

	if *fallback && *url == "" {
    	return  fmt.Errorf("Error: --fallback can only be used with -u")
	}

	if *pretty && !*jsonOutput {
		return fmt.Errorf("-pretty must be used with -json")
	}

	jobs := make(chan model.Job,100)
	result := make(chan model.Result,100)
	var wg sync.WaitGroup

	wg.Add(*workers)
	for i := 0;i < *workers;i++{
		go func () {	
			defer wg.Done()
			for job := range jobs{
				if job.Type == "URL"{
					resList := scanner.ScanUrl(job.Url,job.Port,*fallback)
					for _,r := range resList{
						result <- r
					}
				}else{
					result <- scanner.ScanPort(job.IP,job.Port)
				}
			}
		}()
	}

	go func(){
		wg.Wait()
		close(result)
	}()

	switch{

	case *url != "":
		portList, err := parser.ParsePort(*ports)
		if err != nil{
			return err
		}
		for _,p := range portList{
			jobs <- model.Job {
				Type: "URL",
				Url: *url,
				Port: p,
			}
		}
		
	case *target != "":
		portList, err := parser.ParsePort(*ports)
		if err != nil {
			return err

		}
		ips, err:= scanner.GenerateIP(*target)
		if err != nil {
			return err

		}

		for _, ip := range ips{
			for _,port := range portList{
				jobs <- model.Job {
					Type: "PORT",
					IP: ip.String(),
					Port: port,
				}
			}
		}

	default:
		fmt.Println("Usage:")
    	fmt.Println("  scan -t <ip/cidr> -p <ports>")
    	fmt.Println("  scan -u <url>")
		return  fmt.Errorf("invalid arguments")

	}
	close(jobs)

	var writer io.Writer = os.Stdout
	if *fileoutput != "" {
		f, err := os.Create(*fileoutput)
		if err != nil {
			return err
		}
		defer f.Close()
		writer = f
	}
	output.Output(result,writer,*jsonOutput,*pretty)
	return nil
}


func main(){

	if err := run();err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}