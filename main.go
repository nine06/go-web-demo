package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

func httpsrv() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("from", r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto)
		r.Header["CLIENT-INFO"] = []string{r.RemoteAddr, r.Method, r.URL.RequestURI(), r.Proto}
		fmt.Printf("%#v", r.Header)
		resp, _ := json.MarshalIndent(r.Header, "", "  ")
		fmt.Fprintf(w, string(resp))
	})

	log.Println("Listening on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	go httpsrv()

	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val

		}
		return items

	}
	cnt := 0
	for {
		environment := getenvironment(os.Environ(), func(item string) (key, val string) {
			splits := strings.Split(item, "=")
			key = splits[0]
			val = splits[1]
			return

		})
		keys := []string{}
		for k := range environment {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println(key, "=", environment[key])
		}
		cnt += 1
		fmt.Println("hello#", cnt)
		/*
			for k, v := range environment {
				fmt.Println(k, "=", v)
			}
		*/
		time.Sleep(time.Second * 3000)
	}
}
