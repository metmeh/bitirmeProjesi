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
	// Kullanıcıdan ip aralığı alma
	fmt.Print("Taranacak IP Aralığını Giriniz/(Başlangıç ve Bitiş IP Aralığını '-' ile Ayırınız): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	ipRange := scanner.Text()

	// IP aralığını '-' ile ayırma
	rangeParts := strings.Split(ipRange, "-")
	if len(rangeParts) != 2 {
		fmt.Println("Geçersiz IP aralığı!")
		return
	}

	startIP := rangeParts[0]
	endIP := rangeParts[1]

	fmt.Printf("Başlangıç IP: %s\nBitiş IP: %s\n", startIP, endIP)

	// IP aralığındaki tüm ip adreslerini tara
	var wg sync.WaitGroup
	for ip := startIP; compareIPs(ip, endIP) <= 0; ip = getNextIP(ip) {
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
	for i := len(oktet) - 1; i >= 0; i-- {
		val, _ := strconv.Atoi(oktet[i])
		if val < 255 {
			oktet[i] = strconv.Itoa(val + 1)
			break
		} else {
			oktet[i] = "0"
		}
	}
	return strings.Join(oktet, ".")
}

func compareIPs(ip1, ip2 string) int {
	oktet1 := strings.Split(ip1, ".")
	oktet2 := strings.Split(ip2, ".")
	for i := 0; i < len(oktet1); i++ {
		val1, _ := strconv.Atoi(oktet1[i])
		val2, _ := strconv.Atoi(oktet2[i])
		if val1 < val2 {
			return -1
		} else if val1 > val2 {
			return 1
		}
	}
	return 0
}

func isIPActive(ip string) bool {
	conn, err := net.Dial("tcp", ip+":80")
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
