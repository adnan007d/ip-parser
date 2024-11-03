package main

import (
	"net"
	"regexp"
	"testing"
)

type IpCheck struct {
	Ip    string
	Valid bool
}

var ips = [...]IpCheck{
	// Valid IPv4 addresses
	{Ip: "192.168.0.1", Valid: true},
	{Ip: "10.0.0.1", Valid: true},
	{Ip: "172.16.0.1", Valid: true},
	{Ip: "8.8.8.8", Valid: true},
	{Ip: "1.1.1.1", Valid: true},
	{Ip: "255.255.255.255", Valid: true},

	{Ip: "127.0.0.1", Valid: true},
	{Ip: "192.168.1.255", Valid: true},
	{Ip: "203.0.113.0", Valid: true},
	{Ip: "198.51.100.14", Valid: true},
	{Ip: "192.0.2.0", Valid: true},

	{Ip: "203.0.113.255", Valid: true},
	{Ip: "0.0.0.0", Valid: true},
	{Ip: "192.168.100.100", Valid: true},
	{Ip: "10.10.10.10", Valid: true},

	// Invalid IPv4 addresses
	{Ip: "256.256.256.256", Valid: false},
	{Ip: "123.456.789.0", Valid: false},
	{Ip: "192.168.1.256", Valid: false},
	{Ip: "999.999.999.999", Valid: false},

	{Ip: "300.1.1.1", Valid: false},
	{Ip: "192.168.1", Valid: false},
	{Ip: "192.168.0.0.1", Valid: false},
	{Ip: "192.168.-1.1", Valid: false},
	{Ip: "123..123.123", Valid: false},

	{Ip: "123.123.123", Valid: false},
	{Ip: "0.0.0.256", Valid: false},
	{Ip: "255.255.255.2555", Valid: false},
	{Ip: "192.168.1.1 extra", Valid: false},

	{Ip: "192.168.0.-5", Valid: false},
	{Ip: "10.0.0.0.0", Valid: false},
	{Ip: "8.8.8", Valid: false},
	{Ip: "192.168.0.0/24", Valid: false},
	{Ip: "123.123..123", Valid: false},

	{Ip: "999.888.777.666", Valid: false},
	{Ip: "300.300.300.300", Valid: false},
	{Ip: "1111.222.333.444", Valid: false},
	{Ip: "10.0.0.256", Valid: false},

	{Ip: "255.255.255.256", Valid: false},
	{Ip: "127.0.0.256", Valid: false},
	{Ip: "192.0.2.999", Valid: false},
	{Ip: "192.168.1.1.1", Valid: false},

	{Ip: "1234.567.89.0", Valid: false},
	{Ip: "1.1.1.-1", Valid: false},
	{Ip: "192.168.0.", Valid: false},
	{Ip: "0.0.0.", Valid: false},
	{Ip: "192.168.1000.1000", Valid: false},

	// Random invalid strings
	{Ip: "not.an.ip", Valid: false},
	{Ip: "abcd", Valid: false},
	{Ip: "192.168.one.one", Valid: false},
	{Ip: "IP_ADDRESS", Valid: false},
	{Ip: "localhost", Valid: false},

	{Ip: "abc.def.ghi.jkl", Valid: false},
	{Ip: "123.456.78.ab", Valid: false},
	{Ip: "abc.def.ghi", Valid: false},
	{Ip: "example.com", Valid: false},

	{Ip: "some_random_string", Valid: false},
	{Ip: "192_168_1_1", Valid: false},
	{Ip: "abc@xyz.com", Valid: false},
	{Ip: "192-168-1-1", Valid: false},

	{Ip: "255:255:255:255", Valid: false},
	{Ip: "192,168,1,1", Valid: false},
	{Ip: "192.168.1.1 extra", Valid: false},
	{Ip: "just_a_string", Valid: false},

	{Ip: "another_bad_input", Valid: false},
	{Ip: "string.with.dots", Valid: false},
	{Ip: "123...456", Valid: false},
	{Ip: "1.2.3.four", Valid: false},

	{Ip: "123.123.123.1234", Valid: false},
	{Ip: "0.0.0.0.0", Valid: false},
	{Ip: "255-255-255-255", Valid: false},
	{Ip: "no-ip-here", Valid: false},

	{Ip: "192.168", Valid: false},
	{Ip: "some.invalid.ip", Valid: false},
	{Ip: "this.is.not.ip", Valid: false},
	{Ip: "more.random.text", Valid: false},

	{Ip: "not_a_valid_ip", Valid: false},
	{Ip: "plain-text", Valid: false},
	{Ip: "some string", Valid: false},
	{Ip: "yet_another_string", Valid: false},

	{Ip: "localhost.localdomain", Valid: false},
	{Ip: "example.ip", Valid: false},
	{Ip: "fake-ip-123", Valid: false},
	{Ip: "string@domain", Valid: false},

	{Ip: "not-really-an-ip", Valid: false},
	{Ip: "192168011", Valid: false},
	{Ip: "192.168.1.1.1.1", Valid: false},
	{Ip: "127.0.0.1 extra", Valid: false},

	{Ip: "0..0.0", Valid: false},
	{Ip: "1..1.1.1", Valid: false},
	{Ip: ".123.123.123", Valid: false},
	{Ip: "255.255.255.256", Valid: false},
	{Ip: "192.168.000.001", Valid: false},
}

func TestIpParserNaive(t *testing.T) {
	for _, ip := range ips {
		err := isValidIP_Naive(ip.Ip)

		if ip.Valid && err != nil {
			t.Errorf("Test failed for IP: %s\nExpected valid, found %s", ip.Ip, err.Error())
		}

		if !ip.Valid && err == nil {
			t.Errorf("Test failed for IP: %s\nExpected invalid found to be valid", ip.Ip)
		}
	}
}

func TestIpParserSinglePass(t *testing.T) {
	for _, ip := range ips {
		err := isValidIP_SinglePass(ip.Ip)

		if ip.Valid && err != nil {
			t.Errorf("Test failed for IP: %s\nExpected valid, found %s", ip.Ip, err.Error())
		}

		if !ip.Valid && err == nil {
			t.Errorf("Test failed for IP: %s\nExpected invalid found to be valid", ip.Ip)
		}
	}
}

func TestIpParserNetPackage(t *testing.T) {
	for _, ip := range ips {
		err := net.ParseIP(ip.Ip)

		if ip.Valid && err == nil {
			t.Errorf("Test failed for IP: %s\nExpected valid, found %s", ip.Ip, err)
		}

		if !ip.Valid && err != nil {
			t.Errorf("Test failed for IP: %s\nExpected invalid found to be valid", ip.Ip)
		}
	}
}

func TestIPParser_Regex(t *testing.T) {
	re := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	for _, ip := range ips {
		matched := re.MatchString(ip.Ip)

		if ip.Valid && matched == false {
			t.Errorf("Test failed for IP: %s\nExpected valid, found %t", ip.Ip, matched)
		}

		if !ip.Valid && matched == true {
			t.Errorf("Test failed for IP: %s\nExpected invalid found to be valid", ip.Ip)
		}
	}
}
