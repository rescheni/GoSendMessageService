package main

import (
	basic "GoMessageService/Basic"
	"GoMessageService/sendserver"
)

func main() {

	basic.LoadConfig()
	sendserver.EmailSend([]string{"1413024330@qq.com"})

}
