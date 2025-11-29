package cmd

import (
	"fmt"
	"os"
)

func Execute(){
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	switch os.Args[1]{
	case "new":
        newCommand()

	case "update":
		updateCommand()
		
    case "lex":
        lexCommand()
    case "-h", "--help":
        showHelp()
    case "-V", "--version":
		showVersion()
        
	default:
		fmt.Println("Unknown command: '", os.Args[1], "'");
		showHelp()
	}
}