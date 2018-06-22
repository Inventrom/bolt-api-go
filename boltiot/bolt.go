// Bolt Cloud API Framework.
// Author: Vimal Sheoran.
// Copyright Inventrom Pvt Ltd, 2018.
// Licensed under MIT License.
// Check the LICENSE.md file for license information.

package boltiot

// Base url for all API endpoints
var BaseURL = "http://cloud.boltiot.com/remote/"

// Bolt a struct for handling request authentication
type Bolt struct{
  APIKey string    // Store the user's API key
  DeviceID string  // Store the user's Device Id
}

// DigitalWrite method
func (bolt Bolt) DigitalWrite(pin, state string) string{

  /*
    Write a HIGH or a LOW value to a digital pin.
    :param str pin: pin  numberto write the value
    :param str state: value of pin
    :returns:  request status, value
    :example: {"success": "1", "value": "1"}
    :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/digitalWrite?pin="+pin+"&state="+state+"&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// DigitalRead method
func (bolt Bolt) DigitalRead(pin string) string{

  /*
    Reads the value from a specified digital pin.
    :param str: pin: number of the digital pin you want to read
    :returns:  request status, value
    :example: {"success": "1", "value": "1"}
    :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/digitalRead?pin="+pin+"&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// AnalogWrite method
func (bolt Bolt) AnalogWrite(pin, value string) string{

  /*
    Writes an analog value to a pin.
    :param str pin: pin  number to write the value
    :param str value: between 0 (always off) and 255 (always on).
    :returns:  request status, command status
    :example: {"success": "1", "value": "1"}
    :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/analogWrite?pin="+pin+"&value="+value+"&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// AnalogRead method
func (bolt Bolt) AnalogRead(pin string) string{

  /*
    Reads the value from the specified analog pin.
    :param str pin: number of the digital pin you want to read
    :returns:  request status, value(0 to 254)
    :example: {"success": "1", "value": "130"}
    :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/analogRead?pin="+pin+"&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// SerialBegin method
func (bolt Bolt) SerialBegin(baudRate string) string{

  /*
    Sets the data rate in bits per second (baud) for serial data transmission.
    :param str baud_rate: speed: in bits per second (baud)
    :returns:  request status, serialBegin Status
    :example: {"success": "1", "value": "Success"}\
    :rtype: JSON
  */
  url := BaseURL+bolt.APIKey+"/serialBegin?baud="+baudRate+"&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// SerialRead method
func (bolt Bolt) SerialRead(readTill string) string{

  /*
    Reads incoming serial data.
    :param str char_till: read the character upto specific index
    :returns:  request status, value
    :example: {"success": "1", "value": "inventrom"}
    :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/serialRead?till="+readTill+"&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// SerialWrite method
func (bolt Bolt) SerialWrite(data string) string{

  /*
        Writes the data to the serial port.
        :param str data: in bits per second (baud)
        :returns:  request status, serialWrite Status
        :example: {"success": "1", "value": "Serial write Successful"}
        :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/serialWrite?data="+data+"&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// Version method
func (bolt Bolt) Version() string{

  /*
     Check the Bolt hardware and firmware version
    :param None
    :returns:  request status, Bolt Hardware Version and Firmware Version
    :example: {"success": "1", "value": "{\"Bolt Hardware Version\":\"2\",
                                        \"Firmware Version\":\"1.0.1\"}"}
    :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/version?&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// Restart method
func (bolt Bolt) Restart() string{

  /*
    Restart the device
    :param None
    :returns:  request status, command status
    :example: {"success": "1", "value": "Restarted"}
    :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/restart?&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// IsAlive method
func (bolt Bolt) IsAlive() string{

  /*
    Check the device status
    :param None
    :returns:  request status, device status: dead,alive
    :example: {"success": "1", "value": "alive"}
    :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/isAlive?&deviceName="+bolt.DeviceID
  return makeRequest(url)
}

// IsOnline method
func (bolt Bolt) IsOnline() string{
  /*
    Check the connectivity of the device
    :param None
    :returns: request status, device connectivity, timestamp
    :example: {"success","1","value":"online","Sun 2018-05-06 08:14:43 UTC"}
    :rtype: JSON
  */

  url := BaseURL+bolt.APIKey+"/isOnline?&deviceName="+bolt.DeviceID
  return makeRequest(url)
}
