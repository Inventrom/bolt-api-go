// Bolt Cloud API Framework.
// Author: Vimal Sheoran.
// Copyright Inventrom Pvt Ltd, 2018.
// Licensed under MIT License.

// Check the LICENSE.md file for license information.
package boltiot

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// Test functions for GPIO functionality

func TestDigitalWriteSuccessfull(t *testing.T){
	// Testing successfull DigitalWrite.
  parseJSON(bolt.DigitalWrite(GPIOConfig["VALID_PIN"],
                              GPIOConfig["VALID_DIGITAL_WRITE_VALUE"]))
  if(responseStruct.Success != GPIOConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIOConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
	time.Sleep(3*time.Second)
}

func TestDigitalWriteFailedInvalidPin(t *testing.T){
	// Testing failed DigitalWrite.
	// In this case an invalid pin value is passed.
	t.Skip("Skipped...will be implemented after API bug fix")
  parseJSON(bolt.DigitalWrite(GPIOConfig["INVALID_PIN"],
                              GPIOConfig["VALID_DIGITAL_WRITE_VALUE"]))
  if(responseStruct.Success != GPIOConfig["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIOConfig["INVALID_PIN_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
	time.Sleep(3*time.Second)
}

func TestDigitalWriteFailedInvalidState(t *testing.T){
	// Testing failed DigitalWrite.
	// In this case an invalid write/state value is passed.
  parseJSON(bolt.DigitalWrite(GPIOConfig["VALID_PIN"],
                              GPIOConfig["INVALID_DIGITAL_WRITE_VALUE"]))
  if(responseStruct.Success != GPIOConfig["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIOConfig["INVALID_STATE_RESPONSE"]){
    t.Error("Failed")
  } else {
    fmt.Println("*****Passed*****")
  }
	time.Sleep(3*time.Second)
}

func TestAnalogWriteSuccessfull(t *testing.T){
	// Testing successfull AnalogWrite.
  parseJSON(bolt.AnalogWrite(GPIOConfig["ANALOG_WRITE_PIN"],
                             GPIOConfig["ANALOG_WRITE_VALUE"]))
  if(responseStruct.Success != GPIOConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIOConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
	time.Sleep(3*time.Second)
}

func TestAnalogWriteFailedInvalidPin(t *testing.T){
	// Testing failed AnalogWrite.
	// In this case an incorrect pin value is passed.
  parseJSON(bolt.AnalogWrite(GPIOConfig["INVALID_PIN"],
                             GPIOConfig["ANALOG_WRITE_VALUE"]))
  if(responseStruct.Success != GPIOConfig["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIOConfig["INVALID_PIN_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
	time.Sleep(3*time.Second)
}

func TestDigitalReadSuccessfull(t *testing.T){
	// Testing successfull DigitalRead.
  parseJSON(bolt.DigitalRead(GPIOConfig["VALID_PIN"]))
	if(responseStruct.Success != GPIOConfig["SUCCESS_RESPONSE"]){
		t.Error("Failed")
	} else if(responseStruct.Value != GPIOConfig["SUCCESS_RESPONSE"]){
		t.Error("Failed")
	} else{
		fmt.Println("*****Passed*****")
	}
	time.Sleep(3*time.Second)
}

func TestDigitalReadFailedInvalidPin(t *testing.T){
	// Testing a failed DigitalRead.
	// In this case an invalid pin value is passed.
	parseJSON(bolt.DigitalRead(GPIOConfig["INVALID_PIN"]))
	if(responseStruct.Success != GPIOConfig["FAILED_RESPONSE"]){
		t.Error("Failed")
	} else if(responseStruct.Value != GPIOConfig["INVALID_PIN_RESPONSE"]){
		t.Error("Failed")
	} else{
		fmt.Println("*****Passed*****")
	}
	time.Sleep(3*time.Second)
}

func TestAnalogReadSuccessfull(t *testing.T){
	// Testing successfull AnalogRead.
	parseJSON(bolt.AnalogRead(GPIOConfig["ANALOG_READ_PIN"]))
	value, _ := strconv.Atoi(responseStruct.Value)
	if(responseStruct.Success != GPIOConfig["SUCCESS_RESPONSE"]){
		t.Error("Failed")
	} else if(!(value < 1024 && value > 0)){
		t.Error("Failed")
	} else{
		fmt.Println("*****Passed*****")
	}
	time.Sleep(3*time.Second)
}

func TestAnalogReadFailedInvalidPin(t *testing.T){
	// Testing a failed AnalogRead.
	// In this case an invalid pin value is passed.
	parseJSON(bolt.AnalogRead(GPIOConfig["INVALID_PIN"]))
	if(responseStruct.Success != GPIOConfig["FAILED_RESPONSE"]){
		t.Error("Failed")
	} else if(responseStruct.Value != GPIOConfig["INVALID_PIN_RESPONSE"]){
		t.Error("Failed")
	} else{
		fmt.Println("*****Passed*****")
	}
	time.Sleep(3*time.Second)
}
