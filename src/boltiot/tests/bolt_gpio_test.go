package boltiot_test

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
  } else if(resposneStruct.Value != GPIO_CONFIG["INVALID_PIN_RESPONSE"]){
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
	if(responseStruct.Success != GPIO_CONFIG["SUCCESS_RESPONSE"]){
		t.Error("Failed")
	} else if(!(int(responseStruct.Value) < 1024 && int(responseStruct.Value) > 0)){
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

func parseJson(passedInResponse string){
	byte_data := []byte(passedInResponse)
	json.Unmarshal(byte_data, &responseStruct)
}
