package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
	"net"
	"os"
	"strings"
)

var serveType, host, port, addr string

const (
	MLeftButton = 1
	MRightButton = 3
	MMidButton = 2
)

// hookType is used to help server judge which key type is send form the client
type hookType struct {
	Htype int
	KeyChar rune
	KeyCode uint16
	Button uint16
	KeySting string
	Ord coOrd
}

type coOrd struct {
	X, Y int16
}

func init() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: client|server host port")
		os.Exit(1)
	}
	serveType = strings.ToLower(os.Args[1])
	host = os.Args[2]
	port = os.Args[3]
	addr = fmt.Sprintf("%s:%s", host, port)
}

func main() {
	if serveType == "server" {
		server()
	} else if serveType == "client" {
		client()
	}
}

func server() {
	servAddr := fmt.Sprintf(":%s", port)
	udpAddr, err := net.ResolveUDPAddr("udp4", servAddr)
	checkErr(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkErr(err)
	defer conn.Close()
	htype := new(hookType)
	for {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkErr(err)
		if n != 0 {
			json.Unmarshal(buf[0:n], htype)
			switch htype.Htype {
			case hook.MouseMove:
				fmt.Printf("\rx: %d, y: %d", htype.Ord.X, htype.Ord.Y)
				robotgo.MoveMouse(int(htype.Ord.X), int(htype.Ord.Y))
			case hook.MouseDown:
				fmt.Printf("%s\r", )
				switch htype.Button {
				case MLeftButton:
					robotgo.MouseClick("left")
				case MRightButton:
					robotgo.MouseClick("right")
				case MMidButton:
					robotgo.MouseClick("center")
				}
			}
			if htype.Htype == hook.MouseMove {
			}
		}
	}
}

// Todo: set a 0 opacity dialog to block local input from keyboard and mouse
func client() {
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkErr(err)
	defer conn.Close()
	htype := new(hookType)
	robotgo.EventHook(hook.MouseMove, []string{}, func(e hook.Event) {
		htype.Htype = hook.MouseMove
		htype.Ord.X = e.X
		htype.Ord.Y = e.Y
		data, _ := json.Marshal(htype)
		fmt.Printf("%s\r", data)
		_, err = conn.Write([]byte(data))
		checkErr(err)
	})

	robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
		htype.Htype = hook.MouseDown
		htype.KeyChar = e.Keychar
		htype.KeyCode = e.Keycode
		htype.Button = e.Button
		htype.KeySting = e.String()
		data, _ := json.Marshal(htype)
		fmt.Printf("%s\r", data)
		_, err = conn.Write([]byte(data))
		checkErr(err)
	})
	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func checkErr(err error) {
	if err != nil {
		log.Printf("error: %s", err.Error())
		fmt.Println(err)
		os.Exit(1)
	}
}
