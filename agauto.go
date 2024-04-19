package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

// encodeCredentials encodes username and password into a Basic Auth string
func encodeCredentials(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// addDNSRewrite adds a new DNS rewrite rule
func addDNSRewrite(baseURL, domain, answer, credentials string) (string, error) {
	client := &http.Client{}
	dnsRewrite := DNSRewrite{Domain: domain, Answer: answer}
	jsonData, err := json.Marshal(dnsRewrite)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", baseURL+"/rewrite/add", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Basic "+credentials)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Handle non-200 status codes with the response body
	if response.StatusCode >= 300 {
		var respError map[string]interface{}
		json.Unmarshal(body, &respError) // Parse error response body
		return "", fmt.Errorf("API request failed (%d): %v", response.StatusCode, respError)
	}

	return string(body), nil
}

// removeDNSRewrite removes an existing DNS rewrite rule
func removeDNSRewrite(baseURL, domain, answer, credentials string) (string, error) {
	client := &http.Client{}
	dnsRewrite := DNSRewrite{Domain: domain, Answer: answer}
	jsonData, err := json.Marshal(dnsRewrite)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", baseURL+"/rewrite/delete", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Basic "+credentials)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Handle non-200 status codes with the response body
	if response.StatusCode >= 300 {
		var respError map[string]interface{}
		json.Unmarshal(body, &respError) // Parse error response body
		return "", fmt.Errorf("API request failed (%d): %v", response.StatusCode, respError)
	}

	return string(body), nil
}

func printUsage() {
	fmt.Println(`Usage:
  --add or -a to add a DNS rewrite rule
  --remove or -r to remove a DNS rewrite rule
  --url [URL] (required) The base URL of the AdGuard API
  --username or -u [Username] (required) Username for authentication
  --password or -p [Password] (required) Password for authentication
  --domain or -d [Domain] (required) The domain to add/remove in DNS rewrite
  --value or -v [Value] (required) The IP or domain to answer DNS queries with
Examples:
  agauto --add --url http://your_adguard_host/control --username admin --password secret --domain example.com --value 192.168.1.1
  agauto --remove --url http://your_adguard_host/control --username admin --password secret --domain example.com --value 192.168.1.1`)
}

func main() {
	var (
		addFlag    = flag.Bool("add", false, "Add a DNS rewrite rule")
		removeFlag = flag.Bool("remove", false, "Remove a DNS rewrite rule")
		url        = flag.String("url", "", "The base URL of the AdGuard API")
		username   = flag.String("username", "", "Username for authentication")
		password   = flag.String("password", "", "Password for authentication")
		domain     = flag.String("domain", "", "The domain to add/remove in DNS rewrite")
		answer     = flag.String("value", "", "The IP or domain to answer DNS queries with")
	)
	flag.Parse()

	if *addFlag == *removeFlag && !*addFlag {
		fmt.Println("Specify either --add or --remove.")
		printUsage()
		return
	}

	if *url == "" || *username == "" || *password == "" || *domain == "" || *answer == "" {
		fmt.Println("All fields are required.")
		printUsage()
		return
	}

	credentials := encodeCredentials(*username, *password)

	if *addFlag {
		response, err := addDNSRewrite(*url, *domain, *answer, credentials)
		if err != nil {
			fmt.Println("Error adding DNS rewrite:", err)
			return
		}
		if response == "{}" || response == "" { // Assuming the API returns {} or an empty string on success
			fmt.Println("OK")
		} else {
			fmt.Println("Add DNS Rewrite Response:", response)
		}
	} else if *removeFlag {
		response, err := removeDNSRewrite(*url, *domain, *answer, credentials)
		if err != nil {
			fmt.Println("Error removing DNS rewrite:", err)
			return
		}
		if response == "{}" || response == "" { // Assuming the API returns {} or an empty string on success
			fmt.Println("OK")
		} else {
			fmt.Println("Remove DNS Rewrite Response:", response)
		}
	}
}
