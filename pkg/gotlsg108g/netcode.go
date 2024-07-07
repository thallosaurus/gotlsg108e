package gotlsg108g

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	v8 "rogchap.com/v8go"
)

// Used to interact with the CGI Scripts
func DataRequest(client Client, jsname string, endpoint string, v url.Values) []byte {
	uri := fmt.Sprintf("http://%s/%s", client.host, endpoint)
	res, err := http.Post(uri, "application/x-www-urlencoded", strings.NewReader(v.Encode()))

	if nil != err {
		log.Fatal(err)
	}

	//log.Println("status", res.StatusCode)
	body, io_err := io.ReadAll(res.Body)

	if nil != io_err {
		log.Fatal(io_err)
	}

	if jsname != "" {
		return parse(jsname, bytes.NewReader(body))
	} else {
		return body
	}
}

// Used to request files from the frontend.
// Skips Javascript parsing if `jsname` is `nil`
func Request(client Client, jsname string, endpoint string, params *url.Values) []byte {
	uri := fmt.Sprintf("http://%s/%s", client.host, endpoint)

	if nil != params {
		uri = fmt.Sprintf("%s?%s", uri, params.Encode())
	}

	// make http request
	resp, err := http.Get(uri)
	if nil != err {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// read body response into buffer
	body, err := io.ReadAll(resp.Body)

	if nil != err {
		log.Fatal(err)
	}

	if jsname != "" {
		return parse(jsname, bytes.NewReader(body))
	} else {
		return body
	}
}

// Takes a bytes Reader that should read HTML and extracts the JavaScript Value of the given Variable Name as JSON bytes
func parse(jsname string, r *bytes.Reader) []byte {
	z, err := html.Parse(r)

	if nil != err {
		log.Fatal(err)
	}

	g := extract(jsname, z)
	g_b, err := g.MarshalJSON()
	if nil != err {
		log.Fatalln(err)
		return nil
	}

	return g_b
}

func extract(jsname string, n *html.Node) *v8.Value {
	if n.Type == html.ElementNode && n.Data == "script" {
		if n.FirstChild != nil && n.FirstChild.Data != "" {
			ctx := v8.NewContext()
			//log.Println("n.FirstChild.Data: ", n.FirstChild.Data)

			ctx.RunScript(n.FirstChild.Data, "parser.js")

			// try to return the value
			v, err := ctx.RunScript(jsname, "parser.js")

			if nil != err {
				log.Println(n.FirstChild.Data)
				log.Println(err)
				return nil
			}
			//ctx.Close()
			return v
		}
	} else {
		var v *v8.Value
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			v = extract(jsname, c)
			if nil != v {
				return v
			}
		}
		return v
	}
	return nil
}
