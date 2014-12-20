package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const layout = "2006-01-02"

func ensureSchema(s string) string {
	if !strings.HasPrefix(s, "http://") {
		return fmt.Sprintf("http://%s", s)
	}
	return s
}

type Response struct {
	date    string `responseDate`
	request string `request`
}

type Identify struct {
	Response
}

func main() {
	identify := flag.Bool("identify", false, "identify a repository")
	rawurl := flag.String("url", "http://doabooks.org/oai", "oai endpoint")
	set := flag.String("set", "", "set to download")
	prefix := flag.String("prefix", "", "prefix to download")
	from := flag.String("from", "1970-01-01", "harvest range")
	to := flag.String("to", time.Now().Format(layout), "harvest range")

	flag.Parse()

	endpoint := ensureSchema(*rawurl)

	if *identify {
		if endpoint == "" {
			log.Fatal("specify endpoint with -url")
		}

		v := url.Values{}
		v.Set("verb", "Identify")

		u := fmt.Sprintf("%s?%s", endpoint, v.Encode())
		log.Println(u)
		resp, err := http.Get(u)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
		os.Exit(0)
	}

	if endpoint == "" {
		log.Fatal("specify endpoint with -url")
	}

	log.Printf("%s %s--%s (set=%s, prefix=%s)", endpoint, *from, *to, *set, *prefix)
}
