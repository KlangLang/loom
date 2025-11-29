package cmd

import "fmt"

type Log struct {
	LoomVersion string
	KlangVersion string
	
	// Colors
	RESET_COLOR string
	PRIMARY_COLOR string
	PRIMARY_LIGHT string
	PRIMARY_DARK string
	ACCENT_COLOR string
	GRAY_LIGHT string
	GRAY_MEDIUM string
	GRAY_DARK string
	NEUTRAL_COLOR string
	SUCESS_COLOR string
	WARNING_COLOR string
	ERROR_COLOR string
}

func (l Log) Header() {
	fmt.Printf("%s╭─────────────────────────────────────────────────────────────────────────╮%s\n", l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s                                            				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s          %s##%s                                				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s            %s##%s                              				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s      %s##  ######%s      	   loom                  			  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s    %s##  ####    ##%s    	   Klang Project Manager 			  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s      %s######%s                                				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s            %s##%s            version %s  •  Klang Core %s	  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.LoomVersion, l.KlangVersion, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s          %s##%s                                				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s                                      			           	  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s╰─────────────────────────────────────────────────────────────────────────╯%s\n", l.PRIMARY_DARK, l.RESET_COLOR)
}

func (l Log) Line() {
    fmt.Printf("%s───────────────────────────────────────────────────────────────────────────%s\n",
        l.PRIMARY_LIGHT, l.RESET_COLOR)
}

func NewLog() Log {
	return Log{
		LoomVersion: "v0.1.6",
		KlangVersion: "v0.1.10",
		
		RESET_COLOR:   "\033[0m",
		PRIMARY_COLOR: "\033[38;2;127;0;31m",
		PRIMARY_DARK:  "\033[38;2;90;0;22m",
		PRIMARY_LIGHT: "\033[38;2;179;0;45m",
		ACCENT_COLOR:  "\033[38;2;212;0;58m",	
		GRAY_LIGHT:    "\033[38;2;191;191;191m",
		GRAY_MEDIUM:   "\033[38;2;138;138;138m",
		GRAY_DARK:     "\033[38;2;43;43;43m",
		SUCESS_COLOR:  "\033[38;2;76;175;80m",
		WARNING_COLOR: "\033[38;2;255;193;7m",
		ERROR_COLOR:   "\033[38;2;255;82;82m",
	}
}