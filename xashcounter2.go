package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {

	Version := "v.25.3.16"

	if len(os.Args) == 2 && os.Args[1] == "-h" {
		println("XashCounter2 (API Server) " + Version)
		println("Usage: xashcounter2 [port]")
		println("")
		os.Exit(3)
	}

	HTTPPort := "81"
	if len(os.Args) == 2 {
		HTTPPort = os.Args[2]
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if strings.Index(r.RequestURI, ":") <= 1 {
			w.WriteHeader(400)
			fmt.Fprintf(w, `Xash API Server!`)
		} else { //if RURI has colon
			println(r.RequestURI)
			RURI := strings.Split(r.RequestURI, "/")[1]
			XashHost := strings.Split(RURI, ":")[0]
			XashPort := strings.Split(RURI, ":")[1]
			//fmt.Fprintf(w, "Host: "+XashHost+"<br>")
			//fmt.Fprintf(w, "Port: "+XashPort+"<br>")

			// connect to Xash server
			//ffffffff 69 6e 66 6f 20 34 39
			Request := []byte{0xff, 0xff, 0xff, 0xff, 0x69, 0x6e, 0x66, 0x6f, 0x20, 0x34, 0x39}
			conn, err := net.Dial("udp", XashHost+":"+XashPort)
			if err != nil {
				println("UDP Error")
				return
			}

			// send to server
			fmt.Fprintf(conn, string(Request))
			// wait for reply

			Received1 := make([]byte, 256)
			_, err = bufio.NewReader(conn).Read(Received1)
			if err == nil {
				PlayerCounterString := (strings.SplitN(string(Received1[:]), "\\", 20)[12])
				println(PlayerCounterString)

				if err == nil {
					fmt.Fprintf(w, PlayerCounterString)
				}
			}
		}
	})
	println("Listen on X port..")
	http.ListenAndServe(":"+HTTPPort, nil)

}
