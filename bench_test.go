package main

import (
	"net"
	"regexp"
	"testing"
)

var defaultGateway = "192.168.0.1"
var bigString = "aslduasodusadiusoaiudowaiuoiuasdlsakjdlsakhdaklsjhdajsdkajsgdkasdk"
var validLengthNumber = "000000000"
var maxValidIP = "255.255.255.255"

func BenchmarkIPParser_DefaultGateway_Naive(b * testing.B) {
	for range b.N {
		isValidIP_Naive(defaultGateway)
	}
}

func BenchmarkIPParser_DefaultGateway_SinglePass(b * testing.B) {
	for range b.N {
		isValidIP_SinglePass(defaultGateway)
	}
}

func BenchmarkIPParser_DefaultGateway_BuiltIn(b *testing.B) {
	for range b.N {
		net.ParseIP(defaultGateway)
	}
}

func BenchmarkIPParser_DefaultGateway_Regex(b *testing.B) {
	re := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	for range b.N {
		re.MatchString(defaultGateway)
	}
}

func BenchmarkIPParser_BigString_Naive(b *testing.B) {
	for range b.N {
		isValidIP_Naive(bigString)
	}
}

func BenchmarkIPParser_BigString_SinglePass(b *testing.B) {
	for range b.N {
		isValidIP_SinglePass(bigString)
	}
}

func BenchmarkIPParser_BigString_BuiltIn(b *testing.B) {
	for range b.N {
		net.ParseIP(bigString)
	}
}

func BenchmarkIPParser_BigString_Regex(b *testing.B) {
	re := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	for range b.N {
		re.MatchString(bigString)
	}
}

func BenchmarkIPParser_ValidLengthNumber_Naive(b *testing.B) {
	for range b.N {
		isValidIP_Naive(validLengthNumber)
	}
}

func BenchmarkIPParser_ValidLengthNumber_SinglePass(b *testing.B) {
	for range b.N {
		isValidIP_SinglePass(validLengthNumber)
	}
}

func BenchmarkIPParser_ValidLengthNumber_BuiltIn(b *testing.B) {
	for range b.N {
		net.ParseIP(validLengthNumber)
	}
}

func BenchmarkIPParser_ValidLengthNumber_Regex(b *testing.B) {
	re := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	for range b.N {
		re.MatchString(validLengthNumber)
	}
}

func BenchmarkIPParser_MaxValid_Naive(b *testing.B) {
	for range b.N {
		isValidIP_Naive(maxValidIP)
	}
}

func BenchmarkIPParser_MaxValid_SinglePass(b *testing.B) {
	for range b.N {
		isValidIP_SinglePass(maxValidIP)
	}
}

func BenchmarkIPParser_MaxValid_BuiltIn(b *testing.B) {
	for range b.N {
		net.ParseIP(maxValidIP)
	}
}	

func BenchmarkIPParser_MaxValid_Regex(b *testing.B) {
	re := regexp.MustCompile(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`)
	for range b.N {
		re.MatchString(maxValidIP)
	}
}		
