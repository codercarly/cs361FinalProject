package main

import (
	s "cs361FinalProject/svc"
)

// const (
// 	SERVER_HOST = "127.0.0.1"
// 	SERVER_PORT = "3610"
// 	SERVER_TYPE = "tcp"
// )

func main() {
	// Connect to microservice via socket
	// conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	// println("Connecting to the microservice...")
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()

	// file, err := os.Open("message.txt")
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }

	// bufferFileName := make([]byte, 64)
	// bufferFileSize := make([]byte, 10)
	// conn.Write([]byte(bufferFileSize))
	// conn.Write([]byte(bufferFileName))
	// sendBuffer := make([]byte, 1024)
	// for {
	// 	_, err = file.Read(sendBuffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	conn.Write(sendBuffer)
	// }
	// fmt.Println("File has been sent")

	s.AppTitleIntro()
	s.PickAnIngredient()
}
