# Introduction

This library provides easy to use interface of the Bolt Cloud APIs. The functions are similar in syntax to standard Arduino or Rasberry Pi functions.

# Installation

Follow these instructions to install the Bolt API library directly to your GOPATH.

1. Clone or Download the repository from Github

2. Navigate to the **/src** folder and copy the **/boltiot** folder from the **/src** folder and paste it into your **GOPATH**

3. Open up the terminal in the following location **<YOUR GOPATH>/boltiot** and run the following command `go install`

4. Now you are all setup to start working with the APIs

# Usage Guide

## GPIO Functions

1. digitalWrite Command
```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.DigitalWrite("2","HIGH")
	fmt.Println(response)
}
```
2. digitalRead Command
```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.DigitalRead("2")
	fmt.Println(response)
}
```

3. analogWrite Command
```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.AnalogWrite("0","100")
	fmt.Println(response)
}
```

4. analogRead Command
```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.AnalogRead("A0")
	fmt.Println(response)
}
```

## UART Functions

1. serialBegin Command
```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.SerialBegin("9600")
	fmt.Println(response)
}
```

2. serialWrite Command

```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.SerialWrite("hello")
	fmt.Println(response)
}
```

3. serialRead Command

```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.SerialRead("5")
	fmt.Println(response)
}
```

## Utility Functions

1. isOnline Command - Checks the connectivity status of the Bolt Module to the cloud.

```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.IsOnline()
	fmt.Println(response)
}
```

2. restart Command - Restart your device
```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.Restart()
	fmt.Println(response)
}
```
3. version Command - Check the device version
```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.Version()
	fmt.Println(response)
}
```

4. isAlive Command (Deprecated) - Check the connectivity status of your Bolt device

```
package main

import (
	"fmt"
	"boltiot"
	)

var bolt = boltiot.Bolt("YOUR-API_KEY-GOES-HERE","YOUR-DEVICE_ID-GOES-HERE")

func main(){
	response := bolt.IsAlive()
	fmt.Println(response)
}
```

## Bolt API documenation

You can find the Bolt API documentation here http://cloud.boltiot.com/api_credentials.

# Testing

To perform a unit test follow these instructions.

1. Setup your Bolt according to the instructions provided in the **test_configurations.txt** file provided under the **/src** directory.

2. Make sure that your Bolt device is connected to the Bolt Cloud.

3. Navigate to your package folder, **<YOUR-GOPATH>/boltiot**, there is a test file in there named **bolt_test.go**

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
