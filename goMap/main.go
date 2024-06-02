package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// Yaygın kullanılan portlar ve servisler
var portServices = map[int]string{
	21:   "FTP",
	22:   "SSH",
	23:   "Telnet",
	25:   "SMTP",
	53:   "DNS",
	80:   "HTTP",
	110:  "POP3",
	143:  "IMAP",
	443:  "HTTPS",
	3306: "MySQL",
	3389: "RDP",
	5900: "VNC",
	2221: "FTP",
}

func scanPort(ip string, port int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		// Port kapalı veya bağlantı kurulamadı
		return
	}
	conn.Close()

	service, exists := portServices[port]
	if exists {
		results <- fmt.Sprintf("Port %d (%s) is open", port, service)
	} else {
		results <- fmt.Sprintf("Port %d is open", port)
	}
}

func main() {
	var ip string
	fmt.Print("Taranacak IP adresini girin: ")
	fmt.Scanln(&ip)

	startPort := 1
	endPort := 5000
	var wg sync.WaitGroup
	results := make(chan string)

	// Tarama sonuçlarını toplamak için bir goroutine başlat
	go func() {
		for result := range results {
			fmt.Println(result)
		}
	}()

	// Portları taramak için goroutine'ler başlat
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go scanPort(ip, port, &wg, results)
	}

	// Tüm goroutine'lerin bitmesini bekle
	wg.Wait()
	close(results)
}
