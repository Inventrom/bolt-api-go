// Bolt Cloud API Framework.
// Author: Vimal Sheoran.
// Copyright Inventrom Pvt Ltd, 2018.
// Licensed under MIT License.
// Check the LICENSE.md file for license information.

package boltiot

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func makeRequest(url string) string{

  /*
    Retrieve a response from a specified API.
    :param string url, url of the API
    :return string contents, contents of the response
    :rtype string
  */
  resp,err := http.Get(url)
  if(err != nil){
    fmt.Printf("%s",err)
  }else{
    defer resp.Body.Close()
    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil{
      fmt.Printf("%s",err)
    }
    return string(contents)
  }
  return ""
}
