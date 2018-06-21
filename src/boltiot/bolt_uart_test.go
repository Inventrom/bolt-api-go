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
  parseJSON(bolt.SerialBegin(UARTConfig["VALID_BAUD_RATE"]))
  if(responseStruct.Success != UARTConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UARTConfig["VALID_BAUD_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialBeginFailedInvalidBaudRate(t *testing.T){
	// Testing a failed SerialBegin.
	// In this case an invalid baud rate is passed.
  parseJSON(bolt.SerialBegin(UARTConfig["INVALID_BAUD_RATE"]))
  if(responseStruct.Success != UARTConfig["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UARTConfig["INVALID_BAUD_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialReadSuccessfull(t *testing.T){
	// Testing a successfull SerialRead.
  parseJSON(bolt.SerialRead(UARTConfig["VALID_TILL"]))
  if(responseStruct.Success != UARTConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Success != UARTConfig["VALID_TILL_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialReadFailedInvalidTill(t *testing.T){
	// Testing a failed SerialRead.
	// In this case an invalid till value is passed.
  parseJSON(bolt.SerialRead(UARTConfig["INVALID_TILL"]))
  if(responseStruct.Success != UARTConfig["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Success != UARTConfig["INVALID_TILL_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialWriteSuccessfull(t *testing.T){
	// Testing a successfull SerialWrite.
  parseJSON(bolt.SerialWrite(UARTConfig["VALID_WRITE_VALUE"]))
  if(responseStruct.Success != UARTConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UARTConfig["VALID_DATA_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestSerialWriteFailedInvalidData(t *testing.T){
	// Testing a failed SerialWrite.
	// In this case an invalid amount of data is passed.
  parseJSON(bolt.SerialWrite(UARTConfig["INVALID_WRITE_VALUE"]))
  if(responseStruct.Success != UARTConfig["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UARTConfig["INVALID_DATA_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}
