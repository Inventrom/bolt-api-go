package boltiot
 
var BASE_URL string = "http://cloud.boltiot.com/remote/"

type Bolt struct{

	apiKey, deviceId string
}

func (bolt Bolt) digitalWrite string(pin, state string){

	/*
		Write a HIGH or a LOW value to a digital pin.

        :param str pin: pin  numberto write the value
        :param str state: value of pin

        :returns:  request status, value
        :example: {"success": "1", "value": "1"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/digitalWrite?pin="+pin+"&state="+state+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) digitalRead string(pin string){

	/*
		Reads the value from a specified digital pin.

        :param str: pin: number of the digital pin you want to read

        :returns:  request status, value
        :example: {"success": "1", "value": "1"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/digitalRead?=pin"+pin+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) analogWrite string(pin, value string){

	/*
		Writes an analog value to a pin.

        :param str pin: pin  number to write the value
        :param str value: between 0 (always off) and 255 (always on).

        :returns:  request status, command status
        :example: {"success": "1", "value": "1"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/analogWrite?pin="+pin+"&value="+value+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) analogRead string(pin string){

	/*
		Reads the value from the specified analog pin.

        :param str pin: number of the digital pin you want to read

        :returns:  request status, value(0 to 254)
        :example: {"success": "1", "value": "130"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/analogRead?pin="+pin+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) serialBegin string(baudRate string){

	/*
		Sets the data rate in bits per second (baud) for serial data transmission.

        :param str baud_rate: speed: in bits per second (baud)

        :returns:  request status, serialBegin Status
        :example: {"success": "1", "value": "Success"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/serialBegin?baud="+baudRate+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) serialRead string(readTill string){

	/*
		Reads incoming serial data.

        :param str char_till: read the character upto specific index

        :returns:  request status, value
        :example: {"success": "1", "value": "inventrom"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/serialRead?till="+readTill+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) serialWrite string(data string){

	/*
		 Writes the data to the serial port.

        :param str data: in bits per second (baud)

        :returns:  request status, serialWrite Status
        :example: {"success": "1", "value": "Serial write Successful"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/serialWrite?data="+data+"&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) version string(){

	/*
		 Check the Bolt hardware and firmware version

        :param None

        :returns:  request status, Bolt Hardware Version and Firmware Version
        :example: {"success": "1", "value": "{\"Bolt Hardware Version\":\"2\",
                                            \"Firmware Version\":\"1.0.1\"}"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/version?&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) restart string(){

	/*
		Restart the device

        :param None

        :returns:  request status, command status
        :example: {"success": "1", "value": "Restarted"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/restart?&deviceName="+bolt.deviceId
	return makeRequest(url)
}

func (bolt Bolt) isAlive string(){

	/*
		Check the device status

        :param None

        :returns:  request status, device status: dead,alive
        :example: {"success": "1", "value": "alive"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/isAlive?&deviceName="+bolt.deviceId
}

func (bolt Bolt) isOnline string(){
	/*
		Check the connectivity of the device

		:param None
		:returns: request status, device connectivity, timestamp
		:example: {"success","1","value":"online","Sun 2018-05-06 08:14:43 UTC"}
		:rtype: JSON
	*/

	url := BASE_URL+bolt.apiKey+"/isOnline?&deviceName="+bolt.deviceId
}