package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)


func subSentence(s string, ignoreStart int, ignoreStartSymbols int) string {
	current := s[ignoreStartSymbols:];
	
	words := strings.Fields(current);
	
	if(len(words)>0) {
	 	current = strings.Join(words[ignoreStart:]," ");
	}
	return current; 
}

func compareTwoString(s string, t string) bool {
	return s==t;
}

func uniq(text []string,
		   count bool,
		   deleteUnrepeated bool,
		   caseInsensitive bool,
		   ignoreFirst int,
		   ignoreStartSymbols int) string {
			var result string;
			stringComparator := compareTwoString;
			if caseInsensitive {
				stringComparator = strings.EqualFold;
			}
			previous := subSentence(text[0],ignoreFirst,ignoreStartSymbols);
			var counter int;
			for i := 1; i < len(text); i++ {
				current := subSentence(text[i],ignoreFirst,ignoreStartSymbols);
				// fmt.Println(current);
				if !stringComparator(previous, current) { 
					if count {
						result += fmt.Sprintf("%d ", counter+1);
						counter=0;
					}
					result += fmt.Sprintf("%s\n", text[i-1]);
				} else if count {
					counter++;
				}
				previous = current;
			}
			if(count) { 
				result+= fmt.Sprintf("%d ", counter+1);
				// fmt.Printf("%d ", counter+1)
			} 
			result += fmt.Sprintf("%s\n", text[len(text)-1]);

			return result;
		   }

func main() {
	count := flag.Bool("c", false, "Count repeating strings?");
	deleteUnrepeated := flag.Bool("d", false, "Delete repeating strings?");
	unique := flag.Bool("u", false, "Show unrepeated strings only?");
	caseInsensitive := flag.Bool("i", false, "case-insensitive");
	ignoreFirst := flag.Int("f", 0, "Ignore first {num} lines");
	ignoreStartSymbols := flag.Int("s",0, "Ignore last {num} lines");

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
            fmt.Printf("%s", uniq(s,*count,*deleteUnrepeated,*caseInsensitive,*ignoreFirst,*ignoreStartSymbols));
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