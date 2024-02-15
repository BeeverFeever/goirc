package main

import (
	"bufio"
	"fmt"
	"io"
	"time"
    "log"

	"net"
	"os"

    "goterm"
    "github.com/muesli/termenv"

	// tea "github.com/charmbracelet/bubbletea"
	// "github.com/muesli/termenv"
)

const (
	ip   = "irc.libera.chat"
	port = "6667"
)

func main() {
	conn, err := net.Dial("tcp", "irc.libera.chat:6667")
	// conn, err := net.Dial("tcp", "127.0.0.1:6667")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Dialing error: ", err)
		return
	}
	defer conn.Close()

	var input chan string = make(chan string)

	go func() {
	    reader := bufio.NewReader(os.Stdin)
	    for {
	        msg, err := reader.ReadString('\n')
	        if err != nil {
	            fmt.Println(err)
	            return
	        }
	        input <- msg
	    }
	}()

    var replies chan string = make(chan string)

	reader := bufio.NewReader(conn)
    go func() {
        for {
            msg, err := reader.ReadString('\n')
            if err != nil {
                log.Fatal(err)
                return
            }
            replies <- msg
        }
    }()

	// msg := <- input
	// conn.Write(msg)
	// conn.Write([]byte("MSG NickServ IDENTIFY quiples blackcanvasmat\r\n"))
    conn.Write([]byte("PASS thisisapass\n"))
	// time.Sleep(time.Second * 5)
	conn.Write([]byte("NICK boinkus\n"))
	// time.Sleep(time.Second * 2)
    conn.Write([]byte("USER boinkus 0 * :beever"))
    // time.Sleep(time.Second * 2)
	// conn.Write([]byte("JOIN #test\n"))
    // conn.Write([]byte("MOTD\n"))
    conn.Write([]byte("JOIN #test\n"))
    for {
        fmt.Print(<- replies)
        time.Sleep(time.Millisecond * 30)
        // prompt, ok := <- input
        // if ok {
        //     conn.Write([]byte(prompt))
        // }
    }

	for {
		var str []byte
		n, err := reader.Read(str)
		if err == io.EOF {
			fmt.Println("Connection closed")
			return
		}
		if n > 0 {
			fmt.Println(n, str)
		}
	}
}
