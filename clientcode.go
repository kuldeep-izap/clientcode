package main

import (
"fmt"
"github.com/kuldeep-izap/izapRestclient"
"encoding/json")

func main(){
	
	var header = make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
	
	input := izapRestclient.InputStruct{
		Url:"http://httpbin.org/headers",	
		Headers:header,
		}
	
	resp,err := izapRestclient.GetResponse(input)
	if err != nil {
		fmt.Println(err);
	}
	
	var output struct{
		Header map[string]string `json:"headers"`
	}

	
	json.Unmarshal(resp,&output)
	
	for name, val := range output.Header {
		fmt.Println(name+"  :  "+val)
	}
	
	fmt.Println("---------------------------------------------------------")
	respheader,err := izapRestclient.GetHeaders(input)
	if err != nil {
		fmt.Println(err);
	}
	fmt.Println(respheader)
	for name,val := range respheader{
		fmt.Println(name+"  :  "+val[0])
	}
}	