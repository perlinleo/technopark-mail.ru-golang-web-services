package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"myuniq/uniq"
	"os"
	"strings"
)


func main() {
	parsedParams:= uniq.ParseParams();

	var s []string;

	if parsedParams.InputFile != "" { 
		file, err := os.Open(parsedParams.InputFile)
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
				s = strings.Split(string(buf[:n]),"\n")
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("read %d bytes: %v", n, err)
				break
			}
		}
	} else {
		reader:= bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-->");
			line, err := reader.ReadString('\n');
			line = strings.TrimSuffix(line,"\n");
			
			if line == "END" {
				break;
			}

			
			if(err != nil) {
				log.Fatal(err);
			}

			s = append(s, line);

		}
	}

	output := fmt.Sprintf("%s", uniq.Uniq(s,parsedParams));

	if parsedParams.OutputFile=="" {
		fmt.Printf("%s",output);
	} else {
		f, err := os.Create(parsedParams.OutputFile)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(output)

		if err2 != nil {
			log.Fatal(err2)
		}
	}

}