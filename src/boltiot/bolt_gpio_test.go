package boltiot

import (
	"testing"
	"fmt"
	"strconv"
)

func TestDigitalWriteSuccessfull(t *testing.T){
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
	parseJson(bolt.AnalogRead(GPIO_CONFIG["INVALID_PIN"]))
	if(responseStruct.Success != GPIO_CONFIG["FAILED_RESPONSE"]){
		t.Error("Failed")
	} else if(responseStruct.Value != GPIO_CONFIG["INVALID_PIN_RESPONSE"]){
		t.Error("Failed")
	} else{
		fmt.Println("*****Passed*****")
	}
}
