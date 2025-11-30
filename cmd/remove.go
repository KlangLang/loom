package cmd

import (
	"fmt"
	"os"
	"os/user"
) 


func removeCommand() { 
	l := NewLog() 
	const FILE = "loom" 
	pathRoot := "/usr/local/bin/" 
	defaultPath := pathRoot + FILE 

	if verifyIfExists(defaultPath) { 
		unInstall(defaultPath, l) 
		return 
	} 

	fmt.Printf("The binary %s does not exist in %s\n", FILE, defaultPath) 

	// Try to find the home directory 
	probablePath, err := user.Current() 
		if err != nil { 
		fmt.Printf("Unable to get your home path: %v\n", err) 
		return 
	} 

	localpath := probablePath.HomeDir + "/.local/bin/" + FILE 
	fmt.Printf("Trying to find it locally in %s\n", localpath) 

	if verifyIfExists(localpath) { 
		unInstall(localpath, l) 
		return 
	} 

	fmt.Printf("Binary %s does not exist in path %s\n", FILE, localpath) 
	fmt.Printf("Probably %s is not installed on your system.\n", FILE) 
	} 

	func verifyIfExists(file string) bool { 
	if _, err := os.Stat(file); os.IsNotExist(err) { 
	return false 
	} 

	return true 
	} 

	func unInstall(filePath string, l Log){ 
	fmt.Printf("%s◉%s Trying to remove loom in %s\n", l.PRIMARY_LIGHT, l.RESET_COLOR, filePath) 
	err := os.Remove(filePath) 

	if err != nil { 
	fmt.Printf("%s✖%s Error uninstalling loom: %v", l.ERROR_COLOR, l.RESET_COLOR, err) 
	return 
	} 

	fmt.Printf("%s✔%s loom %s uninstalled successfully!\n", l.SUCESS_COLOR, l.RESET_COLOR, l.LoomVersion) 
}