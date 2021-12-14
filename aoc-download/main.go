package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/pflag"
)

func myUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	pflag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "Environment:\n  AOC_SESSION Session Cookie for Advent of Code\n")
}

func main() {
	now := time.Now()
	year := pflag.IntP("year", "y", now.Year(), "Advent of Code Year")
	day := pflag.IntP("day", "d", now.Day(), "Advent of Code Day")
	help_flag := pflag.BoolP("help", "h", false, "show help")

	pflag.Usage = myUsage
	pflag.Parse()

	if *help_flag {
		myUsage()
		os.Exit(0)
	}

	cookie := os.Getenv("AOC_SESSION")

	if cookie == "" {
		fmt.Fprintf(os.Stderr, "Error: You must set the AOC_SESSION environment variable to the session cookie.\n")
		myUsage()
		os.Exit(1)
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", *year, *day)

	c := &http.Cookie{
		Name:  "session",
		Value: cookie,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating request: %s\n", err)
		os.Exit(4)
	}
	req.AddCookie(c)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving input: %s\n", err)
		os.Exit(2)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output: %s\n", err)
		os.Exit(3)
	}
}
