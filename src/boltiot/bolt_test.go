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
	
	returnedResponse := bolt.DigitalWrite("4","HIGH")
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)

	if(responseStruct.SuccessStatus != "1"){
		t.Error("Digital Write Failed")
	} else if(responseStruct.ReturnValue != "1"){
		t.Error("Digital Write Failed")
	} else{
		fmt.Println("Digital Write Successful!")
	}
}


func TestDigitalRead(t *testing.T){

	returnedResponse := bolt.DigitalRead("1")
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)

	if (responseStruct.SuccessStatus != "1") {
		t.Error("Digital Read Failed")
	} else if (responseStruct.ReturnValue != "1"){
		t.Error("Digital Read Failed")
	} else{
		fmt.Println("Digital Read Successful!")
	}
}

func TestAnalogWrite(t *testing.T){
	
	returnedResponse := bolt.AnalogWrite("0","100")
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)

	if(responseStruct.SuccessStatus != "1"){
		t.Error("Analog Write Failed")
	} else if(responseStruct.ReturnValue != "1"){
		t.Error("Analog Write Failed")
	} else{
		fmt.Println("Analog Write Successful!")
	}
}

func TestAnalogRead(t *testing.T){
	
	returnedResponse := bolt.AnalogRead("A0")
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)

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
	
	returnedResponse := bolt.SerialBegin("9600")
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)

	if(responseStruct.ReturnValue == "Success" || responseStruct.ReturnValue == "Command timed out"){ 
		fmt.Println("Serial Begin Successful!")
	} else{
		t.Error("Serial Begin Failed")
	}
	
}

func TestSerialWrite(t *testing.T){
	
	returnedResponse := bolt.SerialWrite("hello")
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)

	if(responseStruct.SuccessStatus != "1"){
		t.Error("Serial Write Failed")
	} else if(responseStruct.ReturnValue != "Serial write Successful"){
		t.Error("Serial Write Failed")
	} else{
		fmt.Println("Serial Write Successful!")
	}
}


func TestSerialRead(t *testing.T) {

	returnedResponse := bolt.SerialRead("5")
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)
	
	if(responseStruct.SuccessStatus != "1"){
		t.Error("Serial Read Failed")
	} else if(responseStruct.ReturnValue != "hello"){
		t.Error("Serial Read Failed")
	} else{
		fmt.Println("Serial Read Successful!")
	}
}

func TestIsAlive(t *testing.T){
	
	returnedResponse := bolt.IsAlive()
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)

	if(responseStruct.SuccessStatus != "1"){
		t.Error("Is Alive Failed")
	} else if(responseStruct.ReturnValue != "alive"){
		t.Error("Is Alive Failed")
	} else{
		fmt.Println("Is Alive Successful!")
	}
}

func TestIsOnline(t *testing.T){

	returnedResponse := bolt.IsOnline()
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)

	if (responseStruct.SuccessStatus != "1") {
		t.Error("Is Online Failed")
	} else if(responseStruct.ReturnValue != "online"){
		t.Error("Is Online Failed")
	} else{
		fmt.Println("Is Online Successful!")
	}
}

func TestRestart(t *testing.T){
	
	returnedResponse := bolt.Restart()
	byte_data := []byte(returnedResponse)
	json.Unmarshal(byte_data, &responseStruct)

	if (responseStruct.ReturnValue == "Restarted" || responseStruct.ReturnValue == "Command timed out"){
		fmt.Println("Restart Successful!")
		 
	} else{
		t.Error("Restart Failed")
	}
}