package boltiot
 
var BASE_URL string = "http://cloud.boltiot.com/remote/"

type Bolt struct{

	apiKey, deviceId string
}

func (bolt Bolt) digitalWrite string(pin, state string){

	url := BASE_URL+bolt.apiKey+"/digitalWrite?pin="+pin+"&state="+state+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) digitalRead string(pin string){

	url := BASE_URL+bolt.apiKey+"/digitalRead?=pin"+pin+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) analogWrite string(pin, value string){

	url := BASE_URL+bolt.apiKey+"/analogWrite?pin="+pin+"&value="+value+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) analogRead string(pin string){

	url := BASE_URL+bolt.apiKey+"/analogRead?pin="+pin+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) serialBegin string(baudRate string){

	url := BASE_URL+bolt.apiKey+"/serialBegin?baud="+baudRate+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) serialRead string(readTill string){

	url := BASE_URL+bolt.apiKey+"/serialRead?till="+readTill+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) serialWrite string(data string){

	url := BASE_URL+bolt.apiKey+"/serialWrite?data="+data+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) version string(){

	url := BASE_URL+bolt.apiKey+"/version?&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) restart string(){

	url := BASE_URL+bolt.apiKey+"/restart?&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) isAlive string(){

	url := BASE_URL+bolt.apiKey+"/isAlive?&deviceName="+bolt.deviceId
}

