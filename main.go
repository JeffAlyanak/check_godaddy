package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jeffalyanak/check_godaddy/logger"
	"github.com/jeffalyanak/check_godaddy/model"
)

func main() {
	logger := logger.Get()

	// Struct for holding data
	var d model.GoDaddyDomain

	// Handle cli arguments
	domain := flag.String("domain", "", "domain to search")
	key := flag.String("key", "", "API Key")
	secret := flag.String("secret", "", "API Secret")

	warn := flag.Int64("warn", 15, "days until warning")
	crit := flag.Int64("crit", 7, "days until critical")

	// Create warn and crit durations
	warning := time.Duration(int64(time.Hour) * int64(24**warn))
	critical := time.Duration(int64(time.Hour) * int64(24**crit))

	flag.Parse()

	if *domain == "" {
		logger.Println("No domain provided")
		fmt.Println("No domain provided")
		os.Exit(3)
	}
	if *key == "" {
		logger.Println("No API key provided")
		fmt.Println("No API key provided")
		os.Exit(3)
	}
	if *secret == "" {
		logger.Println("No API secret provided")
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
		logger.Println("Error!")
		fmt.Println("Error!")
		os.Exit(3)
	}
	defer resp.Body.Close()

	// Check for rate limiting
	if resp.StatusCode == 429 {
		retry, _ := strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			logger.Println(err)
			fmt.Println(err)
			os.Exit(3)
		}

		// Wait for a bit
		delay := time.Duration(int64(time.Second)*int64(retry) + 1)
		time.Sleep(time.Duration(delay))

		logger.Println("Rate limit reached, waiting " + strconv.Itoa(retry) + "s")
	}

	// Marshal json data into struct
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &d); err != nil {
		logger.Println(err)
		fmt.Println(err)
		os.Exit(3)
	}

	// Differential between now and expiry
	diff := d.Expires.Sub(time.Now())

	// Exit status and string
	exit_status := 0
	exit_string := ""

	// Determine status
	if d.RenewAuto {
		exit_string += "OK - [" + *domain + "] Autorenewal enabled. Expires "
	} else {
		if diff < 0 {
			exit_status = 2
			exit_string += "CRITICAL - [" + *domain + "] Expired "
		} else if diff < warning {
			exit_status = 2
			exit_string += "CRITICAL - [" + *domain + "] Expires "
		} else if diff < critical {
			exit_status = 1
			exit_string += "WARNING - [" + *domain + "] Expires "
		} else {
			exit_string += "OK - [" + *domain + "] Expires "
		}
	}
	exit_string += "in " + durationDays(diff) + ", at " + d.Expires.String() + " | expiry=" + strconv.FormatInt(d.Expires.Unix(), 10) + ", autorenew=" + boolToString(d.RenewAuto)

	logger.Println(exit_string)
	fmt.Println(exit_string)
	os.Exit(exit_status)
}

func durationDays(diff time.Duration) string {
	if float64(diff) < 86400000000000 {
		return durationHours(diff)
	}
	return strconv.FormatFloat(float64(diff)/86400000000000, 'f', 0, 64) + " day(s)"
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func durationHours(diff time.Duration) string {
	return strconv.FormatFloat(float64(diff)/3600000000000, 'f', 0, 64) + " hours(s)"
}
