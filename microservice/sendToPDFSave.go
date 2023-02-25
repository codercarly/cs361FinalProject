package svc

import (
	"fmt"
	"io"
	"net"
	"os"
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "3610"
	SERVER_TYPE = "tcp"
)

func SendToPDF(fileName string) error {
	// println("DEBUG: Send to PDF")
	// cmd := exec.Command("python3", "/Users/carly/cs361FinalProject/microservice/testmain.py", fileName)
	// log.Printf("Running Send to PDF command and waiting for it to finish...")
	// err := cmd.Run()
	// log.Printf("Command finished with error: %v", err)


	// Connect to PDF Save microservice via socket
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	println("Connecting to the microservice...")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	os.Chdir("/Users/carly/cs361FinalProject/microservice")

	println("Opening file to send...")
	file, err := os.Open(fileName)
	if err != nil {
		println("ERROR: cannot open file")
		fmt.Println(err)
		return err
	}

	bufferFileName := make([]byte, 64)
	bufferFileSize := make([]byte, 10)
	conn.Write([]byte(bufferFileSize))
	conn.Write([]byte(bufferFileName))
	sendBuffer := make([]byte, 1024)
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		conn.Write(sendBuffer)
	}
	fmt.Printf("%s has been sent to PDF Save Microservice\n", fileName)

	return nil
}
