package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// sync.Map is used since we do access the map from multiple goroutines
type Cache struct {
	m sync.Map
}

// return cached data if exists, otherwise fetches from api and caches locally
func (cache *Cache) proxy(w http.ResponseWriter, r *http.Request) {
	key := r.RequestURI
	log.Print(key)

	// ignore favicon requests if from browser
	if strings.Index(key, "/favicon") == 0 {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	cached, ok := cache.m.Load(key)
	if ok {
		w.Write(cached.([]byte))
	} else {
		log.Print("updating weather data from api")
		resp, e := http.Get("http://api.openweathermap.org" + key)

		defer resp.Body.Close()
		if e != nil {
			log.Fatal(e)
		} else {

			if resp.StatusCode != http.StatusOK {
				log.Print("invalid status code:", resp.StatusCode)
				io.WriteString(w, "{}")
				return
			}

			b, e := ioutil.ReadAll(resp.Body)
			if e != nil {
				log.Fatal(e)
			} else {
				// cache response for future use
				cache.m.Store(key, b)
				log.Print("weather data cached")

				// response also to client
				w.Write(b)
			}
		}
	}
}

func main() {
	cache := Cache{}
	http.HandleFunc("/", cache.proxy)
	go func() {
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	// keep cached for an hour
	for range time.Tick(time.Hour * 1) {
		log.Print("deleting ranges")
		cache.m.Range(func(k, v interface{}) bool {
			cache.m.Delete(k)
			return true
		})
	}
}
