package svc

import (
	"fmt"
	"os/exec"
)

func CallMicroserviceClient(fileName string) error {
	println("Setting up microservice client ...")
	cmd := exec.Command("python3", "/Users/carly/cs361FinalProject/microservice/setUpClient.py", fileName)
	
	_, err := cmd.CombinedOutput()
	if err != nil {
		print("ERROR")
		fmt.Println(err)
	}

	return nil
}
