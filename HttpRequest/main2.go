// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	resp, err := http.Get("http://google.com")
// 	if err != nil {
// 		// panic(err)
// 		fmt.Println("Error: ", err)
// 		os.Exit(1)
// 	}
// 	// defer resp.Body.Close()
// 	fmt.Println("hello")

//		fmt.Println(resp)
//	}
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main2() {
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	// io.Copy(os.Stdout, resp.Body)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// // fmt.Println(bs)
	// fmt.Println(string(bs))

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
