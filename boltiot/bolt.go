package boltiot
 
var BASE_URL string = "http://cloud.boltiot.com/remote/"

type Bolt struct{

	ApiKey string
	DeviceId string
}

func (bolt Bolt) DigitalWrite(pin, state string) string{

	/*
		Write a HIGH or a LOW value to a digital pin.

        :param str pin: pin  numberto write the value
        :param str state: value of pin

        :returns:  request status, value
        :example: {"success": "1", "value": "1"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/digitalWrite?pin="+pin+"&state="+state+"&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) DigitalRead(pin string) string{

	/*
		Reads the value from a specified digital pin.

        :param str: pin: number of the digital pin you want to read

        :returns:  request status, value
        :example: {"success": "1", "value": "1"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/digitalRead?=pin"+pin+"&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) AnalogWrite(pin, value string) string{

	/*
		Writes an analog value to a pin.

        :param str pin: pin  number to write the value
        :param str value: between 0 (always off) and 255 (always on).

        :returns:  request status, command status
        :example: {"success": "1", "value": "1"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/analogWrite?pin="+pin+"&value="+value+"&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) AnalogRead(pin string) string{

	/*
		Reads the value from the specified analog pin.

        :param str pin: number of the digital pin you want to read

        :returns:  request status, value(0 to 254)
        :example: {"success": "1", "value": "130"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/analogRead?pin="+pin+"&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) SerialBegin(baudRate string) string{

	/*
		Sets the data rate in bits per second (baud) for serial data transmission.

        :param str baud_rate: speed: in bits per second (baud)

        :returns:  request status, serialBegin Status
        :example: {"success": "1", "value": "Success"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/serialBegin?baud="+baudRate+"&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) SerialRead(readTill string) string{

	/*
		Reads incoming serial data.

        :param str char_till: read the character upto specific index

        :returns:  request status, value
        :example: {"success": "1", "value": "inventrom"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/serialRead?till="+readTill+"&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) SerialWrite(data string) string{

	/*
		 Writes the data to the serial port.

        :param str data: in bits per second (baud)

        :returns:  request status, serialWrite Status
        :example: {"success": "1", "value": "Serial write Successful"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/serialWrite?data="+data+"&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) Version() string{

	/*
		 Check the Bolt hardware and firmware version

        :param None

        :returns:  request status, Bolt Hardware Version and Firmware Version
        :example: {"success": "1", "value": "{\"Bolt Hardware Version\":\"2\",
                                            \"Firmware Version\":\"1.0.1\"}"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/version?&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) Restart() string{

	/*
		Restart the device

        :param None

        :returns:  request status, command status
        :example: {"success": "1", "value": "Restarted"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/restart?&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) isAlive() string{

	/*
		Check the device status

        :param None

        :returns:  request status, device status: dead,alive
        :example: {"success": "1", "value": "alive"}

        :rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/isAlive?&deviceName="+bolt.DeviceId
	return makeRequest(url)
}

func (bolt Bolt) isOnline() string{
	/*
		Check the connectivity of the device

		:param None
		:returns: request status, device connectivity, timestamp
		:example: {"success","1","value":"online","Sun 2018-05-06 08:14:43 UTC"}
		:rtype: JSON
	*/

	url := BASE_URL+bolt.ApiKey+"/isOnline?&deviceName="+bolt.DeviceId
	return makeRequest(url)
}