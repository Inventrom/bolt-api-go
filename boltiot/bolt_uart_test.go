// Bolt Cloud API Framework.
// Author: Vimal Sheoran.
// Copyright Inventrom Pvt Ltd, 2018.
// Licensed under MIT License.
// Check the LICENSE.md file for license information.

package boltiot

import (
  "fmt"
  "testing"
  "time"
)

// Test functions for UART functionality.

func TestSerialBeginSuccessFull(t *testing.T){
  // Testing a successfull SerialBegin.
  t.Skip("API times out")
  parseJSON(bolt.SerialBegin(UARTConfig["VALID_BAUD_RATE"]))
  if(responseStruct.Success != UARTConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UARTConfig["VALID_BAUD_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
  time.Sleep(3*time.Second)
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
  time.Sleep(3*time.Second)
}

func TestSerialReadSuccessfull(t *testing.T){
  parseJSON(bolt.SerialRead(UARTConfig["VALID_TILL"]))
  if(responseStruct.Success != UARTConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
  time.Sleep(3*time.Second)
}

func TestSerialReadFailedInvalidTill(t *testing.T){
  t.Skip("Will be implemented after API fixes")
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
  time.Sleep(3*time.Second)
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
  time.Sleep(3*time.Second)
}

func TestSerialWriteFailedInvalidData(t *testing.T){
  // Testing a failed SerialWrite.
  // In this case an invalid amount of data is passed.
  t.Skip("Unexpected behaviour. Was timing out previously")
  parseJSON(bolt.SerialWrite(UARTConfig["INVALID_WRITE_VALUE"]))
  fmt.Println(responseStruct.Success)
  fmt.Println(responseStruct.Value)
  if(responseStruct.Success != UARTConfig["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UARTConfig["INVALID_DATA_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
  time.Sleep(3*time.Second)
}
