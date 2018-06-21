package boltiot

import (
	"testing"
	"fmt"
)

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
  } else if(responseStruct.Value == ""){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}
