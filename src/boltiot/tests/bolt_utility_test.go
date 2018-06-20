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

func TestIsOnlineSuccessfull(t *testing.T){
  parseJson(bolt.IsOnline())
  if(responseStruct.Success != UTILITY_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UTILITY_CONFIG["ONLINE_VALUE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestRestartSuccessfull(t *testing.T){
  parseJson(bolt.Restart())
  if(responseStruct.Value != UTILITY_CONFIG["RESTART_RESPONSE"]){
    if(responseStruct.Value != UTILITY_CONFIG["RESTART_ALTERNATIVE_RESPONSE"]){
      t.Error("Failed")
    }
  } else{
    fmt.Println("*****Passed*****")
  }
}

func TestVersionSuccessfull(t *testing.T){
  parseJson(bolt.Version())
  if(responseStruct.Success != UTILITY_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value == nil){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}

func parseJson(passedInResponse string){
	byte_data := []byte(passedInResponse)
	json.Unmarshal(byte_data, &responseStruct)
}
