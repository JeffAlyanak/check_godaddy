package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jeffalyanak/godaddy-check/model"
)

func main() {
	// Struct for holding data
	var d model.GoDaddyDomain

	// Handle cli arguments
	domain := flag.String("domain", "", "domain to search")
	key := flag.String("key", "", "API Key")
	secret := flag.String("secret", "", "API Secret")

	warn := flag.Int64("warn", 15, "days until warning (default: 15)")
	crit := flag.Int64("crit", 7, "days until critical (default: 7)")

	// Create warn and crit durations
	warning := time.Duration(int64(time.Hour) * int64(24**warn))
	critical := time.Duration(int64(time.Hour) * int64(24**crit))

	flag.Parse()

	if *domain == "" {
		fmt.Println("No domain provided")
		os.Exit(3)
	}
	if *key == "" {
		fmt.Println("No API key provided")
		os.Exit(3)
	}
	if *secret == "" {
		fmt.Println("No API secret provided")
		os.Exit(3)
	}

	// Build strings for request
	api_call := "https://api.godaddy.com/v1/domains/" + *domain
	auth := *key + ":" + *secret

	// Build request
	client := &http.Client{}
	req, _ := http.NewRequest("GET", api_call, nil)
	req.Header.Add("Authorization", "sso-key "+auth)

	// Make Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error!")
		os.Exit(3)
	}
	defer resp.Body.Close()

	// Check for rate limiting
	if resp.StatusCode == 429 {
		retry, _ := strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

		WriteRetryAfter(retry)
	}

	print(ReadRetryAfter())

	// Marshal json data into struct
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &d); err != nil {
		log.Fatalf("Parse response failed, reason: %v \n", err)
	}

	// Differential between now and expiry
	diff := d.Expires.Sub(time.Now())

	// Exit status and string
	exit_status := 0
	exit_string := ""

	if d.RenewAuto {
		exit_string += "OK - Autorenewal enabled. Expires "
	} else {
		if diff < 0 {
			exit_status = 2
			exit_string += "CRITICAL - Expired "
		} else if diff < warning {
			exit_status = 2
			exit_string += "CRITICAL - Expires "
		} else if diff < critical {
			exit_status = 1
			exit_string += "WARNING - Expires "
		} else {
			exit_string += "OK - Expires "
		}
	}
	exit_string += d.Expires.String() + ", in " + diff.String()

	fmt.Println(exit_string)
	os.Exit(exit_status)
}

func WriteRetryAfter(retry int) {
	f, err := os.Create("retry_after")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	r := time.Duration(int64(time.Second) * int64(retry))

	_, err = f.WriteString(time.Now().Add(r).String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		os.Exit(3)
	}
}

func ReadRetryAfter() string {
	str := ""

	f, err := os.Open("retry_after")
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	return str
}
