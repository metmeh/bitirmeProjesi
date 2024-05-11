package main

/*
import (
	"net"
	"strconv"
	"strings"
)

func getNextIP(ip string) string {
	oktet := strings.Split(ip, ".")
	sonOktet, _ := strconv.Atoi(oktet[3])
	sonOktet++
	oktet[3] = strconv.Itoa(sonOktet)
	return strings.Join(oktet, ".")
}

func isIPActive(ip string) bool {
	_, err := net.Dial("ip4:icmp", ip)
	return err == nil
}
*/
