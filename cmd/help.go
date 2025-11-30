package cmd

import (
	"fmt"
)

func showHelp() {	
	log := NewLog()

	log.Header()
	fmt.Println()
	
	fmt.Printf("%sCommands%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %snew%s   <project>     Create a new Klang project\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %slex%s   <file.k>      Lexicalize a Klang file\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %supdate%s	      Update loom (and in the future Klang)\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %suninstall%s	      Uninstall loom (and in the future Klang and Klang dependences)\n", log.PRIMARY_COLOR, log.RESET_COLOR)

	fmt.Println()
	
	fmt.Printf("%sOptions%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %s-V%s, %s--version%s   Show versions\n", log.PRIMARY_COLOR, log.RESET_COLOR, log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  %s-h%s, %s--help%s      Show this help\n",    log.PRIMARY_COLOR, log.RESET_COLOR, log.PRIMARY_COLOR, log.RESET_COLOR)
	
	fmt.Println()
	
	fmt.Printf("%sVersions%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("  loom:       %s-dev\n", log.LoomVersion)
	fmt.Printf("  Klang Core: %s-dev\n", log.KlangVersion)
	
	log.Line()
	
}