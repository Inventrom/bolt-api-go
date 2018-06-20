package boltiot

import (
	"testing"
	"encoding/json"
	"strconv"
	"fmt"
)

type ResponseStruct struct{
  Success string `json:"success"`
  Value string `json:"value"`
}

var response ResponseStruct
var bolt = Bolt{CREDENTIALS["API_KEY"],CREDENTIALS["DEVICE_ID"]}

func TestSerialBeginSuccessFull(t *testing.T){
  parseJson(bolt.SerialBegin(UART_CONFIG["VALID_BAUD_RATE"]))
  if(responseStruct.Success != UART_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UART_CONFIG["VALID_BAUD_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialBeginFailedInvalidBaudRate(t *testing.T){
  parseJson(bolt.SerialBegin(UART_CONFIG["INVALID_BAUD_RATE"]))
  if(responseStruct.Success != UART_CONFIG["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UART_CONFIG["INVALID_BAUD_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialReadSuccessfull(t *testing.T){
  parseJson(bolt.SerialRead(UART_CONFIG["VALID_TILL"]))
  if(responseStruct.Success != UART_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Success != UART_CONFIG["VALID_TILL_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialReadFailedInvalidTill(t *testing.T){
  parseJson(bolt.SerialRead(UART_CONFIG["INVALID_TILL"]))
  if(responseStruct.Success != UART_CONFIG["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Success != UART_CONFIG["INVALID_TILL_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialWriteSuccessfull(t *testing.T){
  parseJson(bolt.SerialWrite(UART_CONFIG["VALID_WRITE_VALUE"]))
  if(responseStruct.Success != UART_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UART_CONFIG["VALID_DATA_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialWriteFailedInvalidData(t *testing.T){
  parseJson(bolt.SerialWrite(UART_CONFIG["INVALID_WRITE_VALUE"]))
  if(responseStruct.Success != UART_CONFIG["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UART_CONFIG["INVALID_DATA_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func parseJson(passedInResponse string){
	byte_data := []byte(passedInResponse)
	json.Unmarshal(byte_data, &responseStruct)
}
