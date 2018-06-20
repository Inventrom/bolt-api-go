var GPIO_CONFIG = map[string]string

GPIO_CONFIG["VALID_PIN"] = "0"
GPIO_CONFIG["INVALID_PIN"] = "16"
GPIO_CONFIG["VALID_DIGITAL_WRITE_VALUE"] = "HIGH"
GPIO_CONFIG["INVALID_DIGITAL_WRITE_VALUE"] = "MEDIUM"
GPIO_CONFIG["ANALOG_WRITE_VALUE"] = "100"
GPIO_CONFIG["ANALOG_READ_PIN"] = "A0"
GPIO_CONFIG["ANALOG_WRITE_PIN"] = "0"
GPIO_CONFIG["SUCCESS_RESPONSE"] = "1"
GPIO_CONFIG["FAILED_RESPONSE"] = "0"
GPIO_CONFIG["INVALID_PIN_RESPONSE"] = "Invalid pin value"
GPIO_CONFIG["INVALID_STATE_RESPONSE"] = "Invalid state"

var UART_CONFIG = map[string]string

UART_CONFIG["SUCCESS_RESPONSE"] = "1"
UART_CONFIG["FAILED_RESPONSE"] = "0"
UART_CONFIG["VALID_BAUD_RATE"] = "9600"
UART_CONFIG["INVALID_BAUD_RATE"] = "10"
UART_CONFIG["VALID_BAUD_RESPONSE"] = "Success"
UART_CONFIG["INVALID_BAUD_RESPONSE"] = "Invalid baud value"
UART_CONFIG["VALID_TILL"] = "10"
UART_CONFIG["INVALID_TILL"] = "1000"
UART_CONFIG["VALID_TILL_RESPONSE"] = ""
UART_CONFIG["INVALID_TILL_RESPONSE"] = "Invalid till value"
UART_CONFIG["VALID_WRITE_VALUE"] = "hello"
UART_CONFIG["INVALID_WRITE_VALUE"] = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
UART_CONFIG["VALID_DATA_RESPONSE"] = "Serial write Successful"
UART_CONFIG["INVALID_DATA_RESPONSE"] = "Command timed out"

var CREDENTIALS = map[string]string

CREDENTIALS["API_KEY"] = "xxxx"
CREDENTIALS["DEVICE_ID"] = "xxxx"
