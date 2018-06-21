// Bolt Cloud API Framework.
// Author: Vimal Sheoran.
// Copyright Inventrom Pvt Ltd, 2018.
// Licensed under MIT License.

// Check the LICENSE.md file for license information.
package boltiot

import (
	"testing"
	"fmt"
	"strconv"
)

// Test functions for GPIO functionality

func TestDigitalWriteSuccessfull(t *testing.T){
	// Testing successfull DigitalWrite.
  parseJson(bolt.DigitalWrite(GPIO_CONFIG["VALID_PIN"],
                              GPIO_CONFIG["VALID_DIGITAL_WRITE_VALUE"]))
  if(responseStruct.Success != GPIO_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIO_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestDigitalWriteFailedInvalidPin(t *testing.T){
	// Testing failed DigitalWrite.
	// In this case an invalid pin value is passed.
  parseJson(bolt.DigitalWrite(GPIO_CONFIG["INVALID_PIN"],
                              GPIO_CONFIG["VALID_DIGITAL_WRITE_VALUE"]))
  if(responseStruct.Success != GPIO_CONFIG["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIO_CONFIG["INVALID_PIN_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestDigitalWriteFailedInvalidState(t *testing.T){
	// Testing failed DigitalWrite.
	// In this case an invalid write/state value is passed.
  parseJson(bolt.DigitalWrite(GPIO_CONFIG["VALID_PIN"],
                              GPIO_CONFIG["INVALID_DIGITAL_WRITE_VALUE"]))
  if(responseStruct.Success != GPIO_CONFIG["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIO_CONFIG["INVALID_STATE_RESPONSE"]){
    t.Error("Failed")
  } else {
    fmt.Println("*****Passed*****")
  }
}

func TestAnalogWriteSuccessfull(t *testing.T){
	// Testing successfull AnalogWrite.
  parseJson(bolt.AnalogWrite(GPIO_CONFIG["ANALOG_WRITE_PIN"],
                             GPIO_CONFIG["ANALOG_WRITE_VALUE"]))
  if(responseStruct.Success != GPIO_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIO_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestAnalogWriteFailedInvalidPin(t *testing.T){
	// Testing failed AnalogWrite.
	// In this case an incorrect pin value is passed.
  parseJson(bolt.AnalogWrite(GPIO_CONFIG["INVALID_PIN"],
                             GPIO_CONFIG["ANALOG_WRITE_VALUE"]))
  if(responseStruct.Success != GPIO_CONFIG["FAILED_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != GPIO_CONFIG["INVALID_PIN_RESPONSE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestDigitalReadSuccessfull(t *testing.T){
	// Testing successfull DigitalRead.
  parseJson(bolt.DigitalRead(GPIO_CONFIG["VALID_PIN"]))
  if(responseStruct.Success != GPIO_CONFIG["SUCCESS_RESPONSE"]){
		t.Error("Failed")
	} else if(responseStruct.Value != GPIO_CONFIG["SUCCESS_RESPONSE"]){
		t.Error("Failed")
	} else{
		fmt.Println("*****Passed******")
	}
}

func TestDigitalReadFailedInvalidPin(t *testing.T){
	// Testing a failed DigitalRead.
	// In this case an invalid pin value is passed.
	parseJson(bolt.DigitalRead(GPIO_CONFIG["INVALID_PIN"]))
	if(responseStruct.Success != GPIO_CONFIG["FAILED_RESPONSE"]){
		t.Error("Failed")
	} else if(responseStruct.Value != GPIO_CONFIG["INVALID_PIN_RESPONSE"]){
		t.Error("Failed")
	} else{
		fmt.Println("******Passed*****")
	}
}

func TestAnalogReadSuccessfull(t *testing.T){
	// Testing successfull AnalogRead.
	parseJson(bolt.AnalogRead(GPIO_CONFIG["ANALOG_READ_PIN"]))
	value, _ := strconv.Atoi(responseStruct.Value)
	if(responseStruct.Success != GPIO_CONFIG["SUCCESS_RESPONSE"]){
		t.Error("Failed")
	} else if(!(value < 1024 && value > 0)){
		t.Error("Failed")
	} else{
		fmt.Println("*****Passed*****")
	}
}

func TestAnalogReadFailedInvalidPin(t *testing.T){
	// Testing a failed AnalogRead.
	// In this case an invalid pin value is passed.
	parseJson(bolt.AnalogRead(GPIO_CONFIG["INVALID_PIN"]))
	if(responseStruct.Success != GPIO_CONFIG["FAILED_RESPONSE"]){
		t.Error("Failed")
	} else if(responseStruct.Value != GPIO_CONFIG["INVALID_PIN_RESPONSE"]){
		t.Error("Failed")
	} else{
		fmt.Println("*****Passed*****")
	}
}
