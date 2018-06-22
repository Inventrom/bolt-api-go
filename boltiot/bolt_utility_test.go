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

// Test functions for Utility functions.

func TestIsOnlineSuccessfull(t *testing.T){
  // This function tests for IsOnline function.
  parseJSON(bolt.IsOnline())
  if(responseStruct.Success != UtilityConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value != UtilityConfig["ONLINE_VALUE"]){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
  time.Sleep(3*time.Second)
}

func TestRestartSuccessfull(t *testing.T){
  // This function tests for the Restart function.
  parseJSON(bolt.Restart())
  if(responseStruct.Value != UtilityConfig["RESTART_RESPONSE"]){
    if(responseStruct.Value != UtilityConfig["RESTART_ALTERNATIVE_RESPONSE"]){
      t.Error("Failed")
    }
  } else{
    fmt.Println("*****Passed*****")
  }
  time.Sleep(15*time.Second)
}

func TestVersionSuccessfull(t *testing.T){
  // This function tests for the Version function.
  t.Skip("Skipped. Will be implemented after the API patches.")
  parseJSON(bolt.Version())
  if(responseStruct.Success != UtilityConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value == ""){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
  time.Sleep(3*time.Second)
}
