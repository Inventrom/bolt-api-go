// Bolt Cloud API Framework.
// Author: Vimal Sheoran.
// Copyright Inventrom Pvt Ltd, 2018.
// Licensed under MIT License.
// Check the LICENSE.md file for license information.

// This file can be used when required to do an end-to-end testing of the library.
// These tests are run to check wether the library and the bolt device are responding
// to each other.

// Before conducting these tests, make sure that your hardware is setup,
// as mentioned in the hardware_config.txt


package boltiot

import (
	"testing"
	"encoding/json"
	"strconv"
	"fmt"
)

// Declaring a struct to hold the parsed JSON from the response
// of the API request.
type BoltboltResponseStruct struct{
	SuccessStatus string `json:"success"`
	ReturnValue string `json:"value"`
	TimeStamp string `json:"time,omitempty"`
}

// Instantiating the response struct.
var boltResponseStruct BoltboltResponseStruct

func TestDigitalWrite(t *testing.T){
	// Test DigitalWrite.
	parseJsonResponse(bolt.DigitalWrite("4","HIGH"))
	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Digital Write Failed")
	} else if(boltResponseStruct.ReturnValue != "1"){
		t.Error("Digital Write Failed")
	} else{
		fmt.Println("Digital Write Successful!")
	}
}


func TestDigitalRead(t *testing.T){
	// Test DigitalRead.
	// The output of this test will depend upon the, response of the DigitalWrite,
	// function.
	parseJsonResponse(bolt.DigitalRead("1"))

	if (boltResponseStruct.SuccessStatus != "1") {
		t.Error("Digital Read Failed")
	} else if (boltResponseStruct.ReturnValue != "1"){
		t.Error("Digital Read Failed")
	} else{
		fmt.Println("Digital Read Successful!")
	}
}

func TestAnalogWrite(t *testing.T){
	// Test AnalogWrite.
	parseJsonResponse(bolt.AnalogWrite("0","100"))

	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Analog Write Failed")
	} else if(boltResponseStruct.ReturnValue != "1"){
		t.Error("Analog Write Failed")
	} else{
		fmt.Println("Analog Write Successful!")
	}
}

func TestAnalogRead(t *testing.T){
	// Test AnalogRead.
	parseJsonResponse(bolt.AnalogRead("A0"))

	analogValue, _ := strconv.Atoi(boltResponseStruct.ReturnValue)
	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Analog Read Failed")
	} else if(analogValue > 1024 || analogValue < 0){
		t.Error("Analog Read Failed")
	} else{
		fmt.Println("Analog Read Successful!")
	}
}


func TestSerialBegin(t *testing.T){
	// Test SerialBegin.
	parseJsonResponse(bolt.SerialBegin("9600"))

	if(boltResponseStruct.ReturnValue == "Success" || boltResponseStruct.ReturnValue == "Command timed out"){
		fmt.Println("Serial Begin Successful!")
	} else{
		t.Error("Serial Begin Failed")
	}

}

func TestSerialWrite(t *testing.T){
	// Test SerialWrite.
	parseJsonResponse(bolt.SerialWrite("hello"))

	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Serial Write Failed")
	} else if(boltResponseStruct.ReturnValue != "Serial write Successful"){
		t.Error("Serial Write Failed")
	} else{
		fmt.Println("Serial Write Successful!")
	}
}


func TestSerialRead(t *testing.T) {
	// Test SerialRead.
	// The output of this test will depend upon the output of the SerialWrite
	// function.
	parseJsonResponse(bolt.SerialRead("5"))

	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Serial Read Failed")
	} else if(boltResponseStruct.ReturnValue != "hello"){
		t.Error("Serial Read Failed")
	} else{
		fmt.Println("Serial Read Successful!")
	}
}

func TestIsAlive(t *testing.T){
	// Test the IsAlive function.
	parseJsonResponse(bolt.IsAlive())

	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Is Alive Failed")
	} else if(boltResponseStruct.ReturnValue != "alive"){
		t.Error("Is Alive Failed")
	} else{
		fmt.Println("Is Alive Successful!")
	}
}

func TestIsOnline(t *testing.T){
	// Test the IsOnline function.
	parseJsonResponse(bolt.IsOnline())

	if (boltResponseStruct.SuccessStatus != "1") {
		t.Error("Is Online Failed")
	} else if(boltResponseStruct.ReturnValue != "online"){
		t.Error("Is Online Failed")
	} else{
		fmt.Println("Is Online Successful!")
	}
}

func TestRestart(t *testing.T){
	// Test the Restart function.
	parseJsonResponse(bolt.Restart())

	if (boltResponseStruct.ReturnValue == "Restarted" || boltResponseStruct.ReturnValue == "Command timed out"){
		fmt.Println("Restart Successful!")

	} else{
		t.Error("Restart Failed")
	}
}

// Function to help parse the JSON data.
func parseJsonResponse(passedInResponse string){
	byte_data := []byte(passedInResponse)
	json.Unmarshal(byte_data, &boltResponseStruct)
}
