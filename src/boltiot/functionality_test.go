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
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
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
	parseJSONResponse(bolt.DigitalWrite("4","HIGH"))
	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Digital Write Failed")
	} else if(boltResponseStruct.ReturnValue != "1"){
		t.Error("Digital Write Failed")
	} else{
		fmt.Println("Digital Write Successful!")
	}
	time.Sleep(3*time.Second)
}


func TestDigitalRead(t *testing.T){
	// Test DigitalRead.
	// The output of this test will depend upon the, response of the DigitalWrite,
	// function.
	parseJSONResponse(bolt.DigitalRead("1"))

	if (boltResponseStruct.SuccessStatus != "1") {
		t.Error("Digital Read Failed")
	} else if (boltResponseStruct.ReturnValue != "1"){
		t.Error("Digital Read Failed")
	} else{
		fmt.Println("Digital Read Successful!")
	}
	time.Sleep(3*time.Second)
}

func TestAnalogWrite(t *testing.T){
	// Test AnalogWrite.
	parseJSONResponse(bolt.AnalogWrite("0","100"))

	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Analog Write Failed")
	} else if(boltResponseStruct.ReturnValue != "1"){
		t.Error("Analog Write Failed")
	} else{
		fmt.Println("Analog Write Successful!")
	}
	time.Sleep(3*time.Second)
}

func TestAnalogRead(t *testing.T){
	// Test AnalogRead.
	parseJSONResponse(bolt.AnalogRead("A0"))

	analogValue, _ := strconv.Atoi(boltResponseStruct.ReturnValue)
	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Analog Read Failed")
	} else if(analogValue > 1024 || analogValue < 0){
		t.Error("Analog Read Failed")
	} else{
		fmt.Println("Analog Read Successful!")
	}
	time.Sleep(3*time.Second)
}


func TestSerialBegin(t *testing.T){
	// Test SerialBegin.
	parseJSONResponse(bolt.SerialBegin("9600"))

	if(boltResponseStruct.ReturnValue == "Success" || boltResponseStruct.ReturnValue == "Command timed out"){
		fmt.Println("Serial Begin Successful!")
	} else{
		t.Error("Serial Begin Failed")
	}
	time.Sleep(3*time.Second)
}

func TestSerialWrite(t *testing.T){
	// Test SerialWrite.
	parseJSONResponse(bolt.SerialWrite("hello"))

	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Serial Write Failed")
	} else if(boltResponseStruct.ReturnValue != "Serial write Successful"){
		t.Error("Serial Write Failed")
	} else{
		fmt.Println("Serial Write Successful!")
	}
	time.Sleep(3*time.Second)
}


func TestSerialRead(t *testing.T) {
	// Test SerialRead.
	// The output of this test will depend upon the output of the SerialWrite
	// function.
	parseJSONResponse(bolt.SerialRead("5"))
	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Serial Read Failed")
	} else if(boltResponseStruct.ReturnValue != "hello"){
		t.Error("Serial Read Failed")
	} else{
		fmt.Println("Serial Read Successful!")
	}
	time.Sleep(3*time.Second)
}

func TestIsAlive(t *testing.T){
	// Test the IsAlive function.
	parseJSONResponse(bolt.IsAlive())

	if(boltResponseStruct.SuccessStatus != "1"){
		t.Error("Is Alive Failed")
	} else if(boltResponseStruct.ReturnValue != "alive"){
		t.Error("Is Alive Failed")
	} else{
		fmt.Println("Is Alive Successful!")
	}
	time.Sleep(3*time.Second)
}

func TestIsOnline(t *testing.T){
	// Test the IsOnline function.
	parseJSONResponse(bolt.IsOnline())

	if (boltResponseStruct.SuccessStatus != "1") {
		t.Error("Is Online Failed")
	} else if(boltResponseStruct.ReturnValue != "online"){
		t.Error("Is Online Failed")
	} else{
		fmt.Println("Is Online Successful!")
	}
	time.Sleep(3*time.Second)
}

func TestRestart(t *testing.T){
	// Test the Restart function.
	parseJSONResponse(bolt.Restart())

	if (boltResponseStruct.ReturnValue == "Restarted" || boltResponseStruct.ReturnValue == "Command timed out"){
		fmt.Println("Restart Successful!")

	} else{
		t.Error("Restart Failed")
	}
	time.Sleep(3*time.Second)
}

// Function to help parse the JSON data.
func parseJSONResponse(passedInResponse string){
	byteData := []byte(passedInResponse)
	json.Unmarshal(byteData, &boltResponseStruct)
}
