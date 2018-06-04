package boltiot

import ("fmt"
		"net/http"
		"io/ioutil")

func makeRequest(url string) string{
	resp,err := http.Get(url)
	if(err != nil){
		fmt.Printf("%s",err)
	}else{
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Printf("%s",err)

		}
		return string(contents)
	}
	return ""
}
