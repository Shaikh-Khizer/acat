package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)

const (
	RED     = "\033[31m"
	GREEN   = "\033[32m"
	YELLOW  = "\033[33m"
	BLUE    = "\033[34m"
	MAGENTA = "\033[35m"
	RESET   = "\033[0m"
)

func colorize(s, color string, useColor bool) string {
	if !useColor {
		return s
	}
	return color + s + RESET
}

func printHelp() {
	fmt.Println("acat - visualize invisible characters")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  acat [options] [file]")
	fmt.Println("  cat file | acat")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -f <file>   read from file")
	fmt.Println("  -nc         disable color")
	fmt.Println("  -l          show line numbers")
	fmt.Println("  --raw       raw mode (no transformation)")
	fmt.Println("  --only      show only control characters")
	fmt.Println("  -c <char>   custom space character (default: .)")
}

func main() {
	// Flags
	noColor := flag.Bool("nc", false, "disable color")
	file := flag.String("f", "", "input file")
	showLineNum := flag.Bool("l", false, "show line numbers")
	raw := flag.Bool("raw", false, "raw mode")
	only := flag.Bool("only", false, "only control chars")
	spaceChar := flag.String("c", ".", "space char")

	flag.Parse()

	// Detect input source
	stdinIsTTY := term.IsTerminal(int(os.Stdin.Fd()))

	if *file == "" && flag.NArg() == 0 && stdinIsTTY {
		printHelp()
		return
	}

	// Color logic (AUTO always)
	stdoutIsTTY := term.IsTerminal(int(os.Stdout.Fd()))
	useColor := stdoutIsTTY && !*noColor

	// Input source
	var reader io.Reader

	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		defer f.Close()
		reader = f
	} else if flag.NArg() > 0 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		defer f.Close()
		reader = f
	} else {
		reader = os.Stdin
	}

	buf := bufio.NewReader(reader)

	lineNum := 1
	printLinePrefix := *showLineNum

	for {
		r, _, err := buf.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "Read error:", err)
			os.Exit(1)
		}

		// Line number
		if printLinePrefix && r != '\n' {
			fmt.Printf("%d | ", lineNum)
			printLinePrefix = false
		}

		if *raw {
			fmt.Printf("%c", r)
			if r == '\n' {
				lineNum++
				printLinePrefix = *showLineNum
			}
			continue
		}

		switch r {

		case ' ':
			if !*only {
				fmt.Print(colorize(*spaceChar, RED, useColor))
			}

		case '\n':
			fmt.Print(colorize(`\n`, GREEN, useColor))
			fmt.Print("\n")

			lineNum++
			printLinePrefix = *showLineNum

		case '\t':
			fmt.Print(colorize(`\t`, YELLOW, useColor))

		case '\r':
			fmt.Print(colorize(`\r`, BLUE, useColor))

		case 0:
			fmt.Print(colorize(`\0`, MAGENTA, useColor))

		default:
			if !*only {
				fmt.Printf("%c", r)
			}
		}
	}
}