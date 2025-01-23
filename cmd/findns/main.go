package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func queryDNS(domain, recordType string) {
	switch strings.ToUpper(recordType) {
	case "A":
		records, err := net.LookupHost(domain)
		if err != nil {
			fmt.Printf("Error querying A records: %v\n", err)
			return
		}
		for _, record := range records {
			fmt.Println(record)
		}
	case "AAAA":
		records, err := net.LookupIP(domain)
		if err != nil {
			fmt.Printf("Error querying AAAA records: %v\n", err)
			return
		}
		for _, record := range records {
			if record.To4() == nil { // Filters only IPv6 addresses
				fmt.Println(record)
			}
		}
	case "CNAME":
		record, err := net.LookupCNAME(domain)
		if err != nil {
			fmt.Printf("Error querying CNAME record: %v\n", err)
			return
		}
		fmt.Println(record)
	case "MX":
		records, err := net.LookupMX(domain)
		if err != nil {
			fmt.Printf("Error querying MX records: %v\n", err)
			return
		}
		for _, mx := range records {
			fmt.Printf("%s (Priority: %d)\n", mx.Host, mx.Pref)
		}
	case "TXT":
		records, err := net.LookupTXT(domain)
		if err != nil {
			fmt.Printf("Error querying TXT records: %v\n", err)
			return
		}
		for _, record := range records {
			fmt.Println(record)
		}
	default:
		fmt.Printf("Unsupported record type: %s\n", recordType)
	}
}

func main() {
	domain := flag.String("domain", "", "Domain name to query (required)")
	recordType := flag.String("type", "A", "DNS record type to query (A, AAAA, CNAME, MX, TXT)")
	flag.Parse()

	if *domain == "" {
		args := flag.Args()
		*domain = args[0]
		if *domain == "" {
			fmt.Println("Error: domain is required")
			flag.Usage()
			os.Exit(1)
		}
	}

	queryDNS(*domain, *recordType)
}
