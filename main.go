package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	count := flag.Bool("c", false, "Count repeating strings?");
	deleteUnrepeated := flag.Bool("d", false, "Delete repeating strings?");
	unique := flag.Bool("u", false, "Show unrepeated strings only?");
	caseInsensitive := flag.Bool("i", false, "case-insensitive");
	ignoreFirst := flag.Int("f", 0, "Ignore first {num} lines");
	ignoreEnd := flag.Int("s",0, "Ignore last {num} lines");

	flag.Parse();

	// Проверка на то,что более одного из флагов -c -d -u - true
	if *count && (*deleteUnrepeated || *unique) || (*deleteUnrepeated && *unique) {
		log.Fatal("Can`t use -c,-d,-u together");
	}
    file, err := os.Open("Input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()


    buf := make([]byte, 32*1024) 

    for {
        n, err := file.Read(buf)

        if n > 0 {
            s := strings.Split(string(buf[:n]),"\n")
			previous := "";
			var counter int;
			for i := *ignoreFirst; i < (len(s)-*ignoreEnd); i++ {
				if !*caseInsensitive {
					if previous!=s[i] {
						if *count {
							fmt.Printf("%d ", counter);
							counter = 0;
						} else {
							fmt.Println(s[i]);
						}
					} else {
						counter++;
					}
				} else {
					if !strings.EqualFold(previous, s[i]) {
						fmt.Println(s[i]);
					} else {
						counter++;
					}
				}
				previous = s[i];
			}
        }

        if err == io.EOF {
            break
        }
        if err != nil {
            log.Printf("read %d bytes: %v", n, err)
            break
        }
    }

}