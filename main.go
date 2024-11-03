package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func isValidIP_Naive(ip string) error {

	// Basic check
	if len(ip) < 7 || len(ip) > 15 {
		return errors.New("invalid IP")
	}

	octets := strings.Split(ip, ".")

	// Should only have 4 octets
	if len(octets) != 4 {
		return errors.New("invalid IP")

	}

	for _, octet := range octets {
		// 01, 012 are not a valid octet of ip address
		// 0 is a valid octet of IP address
		if (len(octet) > 1 && octet[0] == '0') || (len(octet) < 1 || len(octet) > 3) {
			return errors.New("invalid IP")
		}

		n, err := strconv.ParseInt(octet, 10, 0)

		if err != nil {
			return errors.New("invalid IP")
		}

		if n < 0 || n > 255 {
			return errors.New("invalid IP")
		}
	}

	return nil
}

func isValidIP_SinglePass(ip string) error {
	if len(ip) < 7 || len(ip) > 15 {
		return errors.New("invalid IP")
	}

	end := len(ip)
	octets := 0
	octet := 0
	octetLen := 0

	for i := 0; octets <= 4; i++ {
		if end == i || ip[i] == '.' {
			octets += 1
			if octetLen < 1 || octetLen > 3 {
				// Invalid octet length
				return errors.New("invalid IP")
			}

			if octet < 0 || octet > 255 {
				// Invalid octet value
				return errors.New("invalid IP")
			}

			if i == end {
				break
			}

			octetLen = 0
			octet = 0
			continue // No processing for '.' or reached end
		}

		// Must be a valid numeric digit
		if ip[i] < '0' || ip[i] > '9' {
			return errors.New("invalid IP")
		}

		// .001, .01 -> Invalid Octet
		// .0 -> Valid Octet
		if octetLen == 0 && ip[i] == '0' && i+1 < end && ip[i+1] != '.' {
			return errors.New("invalid IP")
		}

		// Basic logic to convert string to int
		octet = octet*10 + int(ip[i]-'0')
		octetLen += 1

		if octetLen > 3 {
			return errors.New("invalid Ip")
		}
	}

	if octets != 4 {
		return errors.New("invalid IP")
	}

	return nil
}

func main() {
	ip := "191.168.0.1"

	var err error

	err = isValidIP_Naive(ip)

	if err != nil {
		log.Fatal("Failed: Naive")
	}

	err = isValidIP_SinglePass(ip)

	if err != nil {
		log.Fatal("Failed: SinglePass")
	}

	validIp := net.ParseIP(ip)

	if validIp == nil {
		log.Fatal("Failed: net.ParseIP")
	}

	fmt.Println("All passed for more tests run `go test`")
}
