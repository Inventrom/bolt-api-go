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
	// This function tests for the Restart function.
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
	// This function tests for the Version function.
  parseJson(bolt.Version())
  if(responseStruct.Success != UTILITY_CONFIG["SUCCESS_RESPONSE"]){
    t.Error("Failed")
  } else if(responseStruct.Value == ""){
    t.Error("Failed")
  } else{
    fmt.Println("*****Passed*****")
  }
}
