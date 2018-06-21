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
}

func TestVersionSuccessfull(t *testing.T){
	// This function tests for the Version function.
  parseJSON(bolt.Version())
  if(responseStruct.Success != UtilityConfig["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value == ""){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}
