# Bolt API Framework (BETA VERSION)

# Introduction

This library provides easy to use interface of the Bolt Cloud APIs. The functions are similar in syntax to standard Arduino or Rasberry Pi functions.

# Installation

The library can be easily installed using `go get`. Paste the following command into your terminal to install the library.

`go get -v github.com/Inventrom/bolt-api-go/boltiot`

# Usage Guide

## GPIO Functions

1. digitalWrite Command
```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.DigitalWrite("2","HIGH")
  fmt.Println(response)
}
```
2. digitalRead Command
```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.DigitalRead("2")
  fmt.Println(response)
}
```

3. analogWrite Command
```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.AnalogWrite("0","100")
  fmt.Println(response)
}
```

4. analogRead Command
```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.AnalogRead("A0")
  fmt.Println(response)
}
```

## UART Functions

1. serialBegin Command
```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.SerialBegin("9600")
  fmt.Println(response)
}
```

2. serialWrite Command

```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.SerialWrite("hello")
  fmt.Println(response)
}
```

3. serialRead Command

```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.SerialRead("5")
  fmt.Println(response)
}
```

## Utility Functions

1. isOnline Command - Checks the connectivity status of the Bolt Module to the cloud.

```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.IsOnline()
  fmt.Println(response)
}
```

2. restart Command - Restart your device
```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.Restart()
  fmt.Println(response)
}
```
3. version Command - Check the device version
```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.Version()
  fmt.Println(response)
}
```

4. isAlive Command (Deprecated) - Check the connectivity status of your Bolt device

```go
package main

import "fmt"
import boltiot "github.com/Inventrom/bolt-api-go/boltiot"

var bolt = boltiot.Bolt{"YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE"}

func main(){
  response := bolt.IsAlive()
  fmt.Println(response)
}
```

## Bolt API documenation

You can find the Bolt API documentation here http://cloud.boltiot.com/api_credentials.

# Testing

To perform a unit test and functionality test follow these instructions.

1. Navigate to **YOUR_GO_PATH/src/github.com/Inventrom/bolt-api-go/**

2. Now setup your hardware according to the instructions provided in the **hardware_config.txt**

3. Make sure that your Bolt device is connected to the Bolt Cloud.

3. Navigate inside the **boltiot/** folder.

4. Open up your terminal in this location and run the following command `go test`

# Contributing

Your contributions are always welcome! Please refer to the contribution guidelines.

## Guidelines
* Fork the repository on GitHub.
* First checkout to dev branch.
* Create a feature branch only when you are working on a new feature.
* Write a test which shows that the bug was fixed or that the feature works as expected.
* Never work on master branch
* Send a pull request and wait until it gets merged and published. :)
* Check your spelling and grammar.
* Remove any trailing whitespaces.
