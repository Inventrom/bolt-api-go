package boltiot

import "net/http"

func makeRequest string(url string){

	resp, err := http.Get(url)
	if(err){
		return "An Error Has Occured"
	}
	else{
		return resp(string)
	}
}