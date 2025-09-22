package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				c <- fmt.Sprintf("erro ao acessar %s: %v", url, err)
				return
			}
			defer resp.Body.Close()

			html, err := io.ReadAll(resp.Body)
			if err != nil {
				c <- fmt.Sprintf("erro ao ler %s: %v", url, err)
				return
			}

			r := regexp.MustCompile("(?i)<title>(.*?)</title>")
			matches := r.FindStringSubmatch(string(html))
			if len(matches) > 1 {
				c <- matches[1]
			} else {
				c <- fmt.Sprintf("sem título em %s", url)
			}
		}(url)
	}
	return c
}

func main() {
	t1 := titulo("https://www.cod3r.com.br", "https://www.google.com")
	t2 := titulo("https://www.amazon.com", "https://www.youtube.com")

	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}
