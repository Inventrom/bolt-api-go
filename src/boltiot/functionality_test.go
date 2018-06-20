package boltiot

import (
	"testing"
	"encoding/json"
	"strconv"
	"fmt"
)

type BoltResponseStruct struct{
	SuccessStatus string `json:"success"`
	ReturnValue string `json:"value"`
	TimeStamp string `json:"time,omitempty"`
}

var responseStruct BoltResponseStruct

var bolt = Bolt{"7bc48b25-f5c8-4ef3-9477-c599835583da","BOLT3729610"}

func TestDigitalWrite(t *testing.T){
	
	parseJsonResponse(bolt.DigitalWrite("4","HIGH"))

	if(responseStruct.SuccessStatus != "1"){
		t.Error("Digital Write Failed")
	} else if(responseStruct.ReturnValue != "1"){
		t.Error("Digital Write Failed")
	} else{
		fmt.Println("Digital Write Successful!")
	}
}


func TestDigitalRead(t *testing.T){

	parseJsonResponse(bolt.DigitalRead("1"))

	if (responseStruct.SuccessStatus != "1") {
		t.Error("Digital Read Failed")
	} else if (responseStruct.ReturnValue != "1"){
		t.Error("Digital Read Failed")
	} else{
		fmt.Println("Digital Read Successful!")
	}
}

func TestAnalogWrite(t *testing.T){
	
	parseJsonResponse(bolt.AnalogWrite("0","100"))

	if(responseStruct.SuccessStatus != "1"){
		t.Error("Analog Write Failed")
	} else if(responseStruct.ReturnValue != "1"){
		t.Error("Analog Write Failed")
	} else{
		fmt.Println("Analog Write Successful!")
	}
}

func TestAnalogRead(t *testing.T){
	
	parseJsonResponse(bolt.AnalogRead("A0"))

	analogValue, _ := strconv.Atoi(responseStruct.ReturnValue)
	if(responseStruct.SuccessStatus != "1"){
		t.Error("Analog Read Failed")
	} else if(analogValue > 1024 || analogValue < 0){
		t.Error("Analog Read Failed")
	} else{
		fmt.Println("Analog Read Successful!")
	}
}


func TestSerialBegin(t *testing.T){
	
	parseJsonResponse(bolt.SerialBegin("9600"))

	if(responseStruct.ReturnValue == "Success" || responseStruct.ReturnValue == "Command timed out"){ 
		fmt.Println("Serial Begin Successful!")
	} else{
		t.Error("Serial Begin Failed")
	}
	
}

func TestSerialWrite(t *testing.T){
	
	parseJsonResponse(bolt.SerialWrite("hello"))

	if(responseStruct.SuccessStatus != "1"){
		t.Error("Serial Write Failed")
	} else if(responseStruct.ReturnValue != "Serial write Successful"){
		t.Error("Serial Write Failed")
	} else{
		fmt.Println("Serial Write Successful!")
	}
}


func TestSerialRead(t *testing.T) {

	parseJsonResponse(bolt.SerialRead("5"))
	
	if(responseStruct.SuccessStatus != "1"){
		t.Error("Serial Read Failed")
	} else if(responseStruct.ReturnValue != "hello"){
		t.Error("Serial Read Failed")
	} else{
		fmt.Println("Serial Read Successful!")
	}
}

func TestIsAlive(t *testing.T){
	
	parseJsonResponse(bolt.IsAlive())

	if(responseStruct.SuccessStatus != "1"){
		t.Error("Is Alive Failed")
	} else if(responseStruct.ReturnValue != "alive"){
		t.Error("Is Alive Failed")
	} else{
		fmt.Println("Is Alive Successful!")
	}
}

func TestIsOnline(t *testing.T){

	parseJsonResponse(bolt.IsOnline())

	if (responseStruct.SuccessStatus != "1") {
		t.Error("Is Online Failed")
	} else if(responseStruct.ReturnValue != "online"){
		t.Error("Is Online Failed")
	} else{
		fmt.Println("Is Online Successful!")
	}
}

func TestRestart(t *testing.T){
	
	parseJsonResponse(bolt.Restart())

	if (responseStruct.ReturnValue == "Restarted" || responseStruct.ReturnValue == "Command timed out"){
		fmt.Println("Restart Successful!")
		 
	} else{
		t.Error("Restart Failed")
	}
}

func parseJsonResponse(passedInResponse string){
	byte_data := []byte(passedInResponse)
	json.Unmarshal(byte_data, &responseStruct)
}