package boltiot

import (
  "encoding/json"
)

var GPIO_CONFIG = map[string]string{
  "VALID_PIN": "0",
  "INVALID_PIN": "16",
  "VALID_DIGITAL_WRITE_VALUE": "HIGH",
  "INVALID_DIGITAL_WRITE_VALUE": "MEDIUM",
  "ANALOG_WRITE_VALUE": "100",
  "ANALOG_READ_PIN": "A0",
  "ANALOG_WRITE_PIN": "0",
  "SUCCESS_RESPONSE": "1",
  "FAILED_RESPONSE": "0",
  "INVALID_PIN_RESPONSE": "Invalid pin value",
  "INVALID_STATE_RESPONSE": "Invalid state",
}

var UART_CONFIG = map[string]string{
  "SUCCESS_RESPONSE": "1",
  "FAILED_RESPONSE": "0",
  "VALID_BAUD_RATE": "9600",
  "INVALID_BAUD_RATE": "10",
  "VALID_BAUD_RESPONSE": "Success",
  "INVALID_BAUD_RESPONSE": "Invalid baud value",
  "VALID_TILL": "10",
  "INVALID_TILL": "1000",
  "VALID_TILL_RESPONSE": "",
  "INVALID_TILL_RESPONSE": "Invalid till value",
  "VALID_WRITE_VALUE": "hello",
  "INVALID_WRITE_VALUE": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
  "VALID_DATA_RESPONSE": "Serial write Successful",
  "INVALID_DATA_RESPONSE": "Command timed out",
}

var UTILITY_CONFIG = map[string]string{
  "SUCCESS_RESPONSE": "1",
  "FAILED_RESPONSE": "0",
  "RESTART_RESPONSE": "Restarted",
  "RESTART_ALTERNATIVE_RESPONSE": "Command timed out",
  "ONLINE_VALUE": "online",
}

var CREDENTIALS = map[string]string{
  "API_KEY": "xxxx",
  "DEVICE_ID": "xxxx",
}

type ResponseStruct struct{
  Success string `json:"success"`
  Value string `json:"value"`
}

var responseStruct ResponseStruct
var bolt = Bolt{CREDENTIALS["API_KEY"],CREDENTIALS["DEVICE_ID"]}

func parseJson(passedInResponse string){
	byte_data := []byte(passedInResponse)
	json.Unmarshal(byte_data, &responseStruct)
}
