package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	//Kullanıcıdan ip aralığı alma
	fmt.Print("Taranacak IP Aralığını Giriniz/(Başlangıç ve Bitiş IP Aralığını '-' ile Ayırınız): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ipRange := scanner.Text()

	//IP aralığını '-' ile ayırma
	rangeParts := strings.Split(ipRange, "-")
	if len(rangeParts) != 2 {
		fmt.Println("Geçersiz IP aralığı!")
		return
	}

	startIP := rangeParts[0]
	endIP := rangeParts[1]

	fmt.Printf("Başlangıç IP: %s\nBitiş IP: %s\n", startIP, endIP)
	/*
		startIP := "192.168.1.1"
		endIP := "192.168.1.10"*/

	//ip aralığındaki tüm ip adresleri tara
	var wg sync.WaitGroup
	for ip := startIP; ip <= endIP; ip = getNextIP(ip) {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			if isIPActive(ip) {
				fmt.Printf("Aktif IP: %s\n", ip)
			}
		}(ip)
	}
	wg.Wait()
}

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
