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
	var log *logger.Logger

	// Struct for holding data
	var d model.GoDaddyDomain

	// Handle cli arguments
	logging := flag.Bool("v", false, "enable logging")
	domain := flag.String("domain", "", "domain to search")
	key := flag.String("key", "", "API Key")
	secret := flag.String("secret", "", "API Secret")

	warn := flag.Int64("warn", 15, "days until warning")
	crit := flag.Int64("crit", 7, "days until critical")

	// Create warn and crit durations
	warning := time.Duration(int64(time.Hour) * int64(24**warn))
	critical := time.Duration(int64(time.Hour) * int64(24**crit))

	flag.Parse()

	if *logging {
		log = logger.Get()
	}

	if *domain == "" {
		if *logging {
			log.Println("No domain provided")
		}
		fmt.Println("No domain provided")
		os.Exit(3)
	}
	if *key == "" {
		if *logging {
			log.Println("No API key provided")
		}
		fmt.Println("No API key provided")
		os.Exit(3)
	}
	if *secret == "" {
		if *logging {
			log.Println("No API secret provided")
		}
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
		if *logging {
			log.Println("Error!")
		}
		fmt.Println("Error!")
		os.Exit(3)
	}
	defer resp.Body.Close()

	// Check for rate limiting
	if resp.StatusCode == 429 {
		retry, _ := strconv.Atoi(resp.Header.Get("Retry-After"))
		if err != nil {
			if *logging {
				log.Println(err)
			}
			fmt.Println(err)
			os.Exit(3)
		}

		// Wait for a bit
		delay := time.Duration(int64(time.Second)*int64(retry) + 1)
		time.Sleep(time.Duration(delay))

		if *logging {
			log.Println("Rate limit reached, waiting " + strconv.Itoa(retry) + "s")
		}
	}

	// Marshal json data into struct
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &d); err != nil {
		if *logging {
			log.Println(err)
		}
		fmt.Println(err)
		os.Exit(3)
	}

	// Differential between now and expiry
	diff := d.Expires.Sub(time.Now())

	// Exit status and string
	exit_status := 0
	exit_string := ""

	if d.Code == "ERROR_INTERNAL" { // Check for internal errors on Godaddy's side.
		exit_status = 3
		exit_string = "GODADDY ERROR_INTERNAL - " + d.Message
	} else if d.Expires == time.Unix(0, 0) { // Check if the expiry is otherwise not found.
		exit_status = 3
		exit_string = "UNKNOWN - No expiry time returned by GoDaddy."
	} else {
		if d.RenewAuto { // Check if autorenewal is enabled.
			exit_string += "OK - [" + *domain + "] Autorenewal enabled. Expires "
		} else { // Check if the expiration falls within the warning or critical levels
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
	}
	if *logging {
		log.Println(exit_string)
	}
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
