package cmd

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

type Log struct {
	LoomVersion  string
	KlangVersion string

	// Colors
	RESET_COLOR   string
	PRIMARY_COLOR string
	PRIMARY_LIGHT string
	PRIMARY_DARK  string
	ACCENT_COLOR  string
	GRAY_LIGHT    string
	GRAY_MEDIUM   string
	GRAY_DARK     string
	NEUTRAL_COLOR string
	SUCESS_COLOR  string
	WARNING_COLOR string
	ERROR_COLOR   string
}

func (l Log) Header() {
	width := l.getTerminalWidth()

	if width < 80 {
		// Modo vertical (9:16)
		fmt.Printf("%s╭──────────────────────────╮%s\n", l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s                          %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s            %s##%s            %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s              %s##%s          %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s        %s##  ######%s        %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s      %s##  ####    ##%s      %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s        %s######%s            %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s              %s##%s          %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s            %s##%s            %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s                          %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s├──────────────────────────┤%s\n", l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s loom - %s               %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.LoomVersion, l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Printf("%s│%s Klang Project Mgr        %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
		return
	}

	// Modo horizontal (16:9)
	fmt.Printf("%s╭─────────────────────────────────────────────────────────────────────────╮%s\n", l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s                                            				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s          %s##%s                                				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s            %s##%s                              				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s      %s##  ######%s      	   loom version %s                            %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.LoomVersion, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s    %s##  ####    ##%s    	   Klang Project Manager 			  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s      %s######%s                                				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s            %s##%s            	                     	                  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s          %s##%s                                				  %s│%s\n", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_LIGHT, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Printf("%s│%s                                            				  %s│%s", l.PRIMARY_DARK, l.RESET_COLOR, l.PRIMARY_DARK, l.RESET_COLOR)
}

func (l Log) newLine(s string) {
	fmt.Printf("%s│%s ", l.PRIMARY_DARK, l.RESET_COLOR)
	print(s)
	fmt.Printf("%s│%s \n", l.PRIMARY_DARK, l.RESET_COLOR)

}

func (l Log) addNewLineToHealder(text string) {
	width := l.getTerminalWidth()

	if width < 80 {
		fmt.Printf("%s│%s ", l.PRIMARY_DARK, l.RESET_COLOR)
		fmt.Print(text)
		fmt.Printf("%s│%s \n", l.PRIMARY_DARK, l.RESET_COLOR)

		return
	}

	fmt.Printf("\n%s│%s ", l.PRIMARY_DARK, l.RESET_COLOR)
	fmt.Print(text)
	fmt.Printf("%s│%s", l.PRIMARY_DARK, l.RESET_COLOR)

}

func (l Log) finalizeBottomheader() {
	width := l.getTerminalWidth()

	if width < 80 {
		fmt.Printf("%s╰──────────────────────────╯%s\n", l.PRIMARY_DARK, l.RESET_COLOR)

		return
	}

	fmt.Printf("%s\n╰─────────────────────────────────────────────────────────────────────────╯%s\n", l.PRIMARY_DARK, l.RESET_COLOR)
}

func (l Log) Line() {
	width := l.getTerminalWidth()
	fmt.Printf("%s%s%s\n", l.PRIMARY_LIGHT, strings.Repeat("─", width), l.RESET_COLOR)
}

func (l Log) getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || width < 20 {
		return 80 // fallback
	}
	return width
}

func (l Log) padCenter(text string, width int) string {
	if len(text) >= width {
		return text
	}
	left := (width - len(text)) / 2
	right := width - len(text) - left
	return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}

func NewLog() Log {
	return Log{
		LoomVersion:  "v0.8.3",
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
