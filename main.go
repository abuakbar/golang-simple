package main

import (
	"fmt"
	"os"
)

func main() {

	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	file.Read(buffer)

	if len(buffer) >= 3 {
		if buffer[0] == 'b' {
			if buffer[1] == 'u' {
				if buffer[2] == 'g' {
					panic("This is a bug 1")
				}
			}
		}

		if buffer[0] == 'b' {
			if buffer[1] == 'g' {
				if buffer[2] == 'u' {
					panic("This is a bug 20")
				}
			}
		}

		if buffer[0] == 'u' {
			if buffer[1] == 'b' {
				if buffer[2] == 'g' {
					panic("This is a bug 31")
				}
			}
		}

		if buffer[0] == 'u' {
			if buffer[1] == 'g' {
				if buffer[2] == 'b' {
					panic("This is a bug 41")
				}
			}
		}

		if buffer[0] == 'g' {
			if buffer[1] == 'u' {
				if buffer[2] == 'b' {
					panic("This is a bug 5")
				}
			}
		}
	}
}
