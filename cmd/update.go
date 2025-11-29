package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func updateCommand(){
	if len(os.Args) < 3{
		fmt.Println("Usage: loom new <project_name>")
		return
	}

	cmd := exec.Command("")

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command failed with error: %v", err) 
		return 
	} 

	fmt.Println("Command executed successfully!")
}