// Bolt Cloud API Framework.
// Author: Vimal Sheoran.
// Copyright Inventrom Pvt Ltd, 2018.
// Licensed under MIT License.
// Check the LICENSE.md file for license information.

package boltiot

import (
	"testing"
	"fmt"
)

// Test functions for UART functionality.

func TestSerialBeginSuccessFull(t *testing.T){
	// Testing a successfull SerialBegin.
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
	// Testing a failed SerialBegin.
	// In this case an invalid baud rate is passed.
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
	// Testing a successfull SerialRead.
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
	// Testing a failed SerialRead.
	// In this case an invalid till value is passed.
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
	// Testing a successfull SerialWrite.
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
	// Testing a failed SerialWrite.
	// In this case an invalid amount of data is passed.
  parseJson(bolt.SerialWrite(UART_CONFIG["INVALID_WRITE_VALUE"]))
  if(responseStruct.Success != UART_CONFIG["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UART_CONFIG["INVALID_DATA_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}
