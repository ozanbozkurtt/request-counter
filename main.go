package main

import (
	"fmt"
	"net/http"
	"sync"
)

// IstekSayaci, HTTP isteklerini saymak için kullanılacak tür.
type IstekSayaci struct {
	mu    sync.Mutex
	sayac int
}

// Yeni bir istek sayacı oluşturur.
func YeniIstekSayaci() *IstekSayaci {
	return &IstekSayaci{}
}

// HTTP isteklerini sayar ve istemciye sayacı döndürür.
func (c *IstekSayaci) Sayac(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sayac++
	fmt.Fprintf(w, "Toplam HTTP İstek Sayısı: %d", c.sayac)
}

func main() {
	// Yeni bir istek sayacı oluştur.
	counter := YeniIstekSayaci()

	// "/istek-sayaci" adresine istek geldiğinde counter'ı kullan.
	http.HandleFunc("/istek-sayaci", counter.Sayac)

	// Sunucuyu 8080 portunda başlat.
	fmt.Println("HTTP Sunucusu başlatıldı. http://localhost:8080/istek-sayaci adresine gidin.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP sunucusunda hata oluştu:", err)
	}
}
