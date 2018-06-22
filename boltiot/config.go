// Bolt Cloud API Framework.
// Author: Vimal Sheoran.
// Copyright Inventrom Pvt Ltd, 2018.
// Licensed under MIT License.
// Check the LICENSE.md file for license information.

package boltiot

import (
  "encoding/json"
)

// Here is a collection of configurations and important functions.
// These configurations aid in testing.

// Configurations for testing GPIO functions.
var GPIOConfig = map[string]string{
  "VALID_PIN": "2",
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

// Configurations for testing UART functions.
var UARTConfig = map[string]string{
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

// Configurations for testing Utilities.
var UtilityConfig = map[string]string{
  "SUCCESS_RESPONSE": "1",
  "FAILED_RESPONSE": "0",
  "RESTART_RESPONSE": "Restarted",
  "RESTART_ALTERNATIVE_RESPONSE": "Command timed out",
  "ONLINE_VALUE": "online",
}

// Configuration for user authentication.
var Credentials = map[string]string{
  "API_KEY": "7bc48b25-f5c8-4ef3-9477-c599835583da",
  "DEVICE_ID": "BOLT3729610",
}

// ResponseStruct a struct to hold the response from an API request.
// This struct holds the parsed json obtained after making an API request.
type ResponseStruct struct{
  Success string `json:"success"` // Holds the JSON "success" field.
  Value string `json:"value"`     // Holds the JSON "value" field.
}
var responseStruct ResponseStruct

// Initializing an instance of the Bolt struct.
var bolt = Bolt{Credentials["API_KEY"],Credentials["DEVICE_ID"]}

// This function is responsible for parsing the received JSON data from the API request.
func parseJSON(passedInResponse string){
  byteData := []byte(passedInResponse)     // Convert the response into Bytes
  json.Unmarshal(byteData, &responseStruct) // and store it into respective struct fields
}
