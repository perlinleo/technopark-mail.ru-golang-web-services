package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"myuniq/uniq"
	"os"
	"strings"
)


func main() {
	count := flag.Bool("c", false, "Count repeating strings?");
	deleteUnrepeated := flag.Bool("d", false, "Delete repeating strings?");
	unique := flag.Bool("u", false, "Show unrepeated strings only?");
	caseInsensitive := flag.Bool("i", false, "case-insensitive");
	ignoreFirst := flag.Int("f", 0, "Ignore first {num} lines");
	ignoreStartSymbols := flag.Int("s",0, "Ignore last {num} lines");
	inputFile := flag.String("input_file", "", "File to read");
	outputFile := flag.String("output_file", "", "File to write");

	flag.Parse();

	
	if uniq.СheckParams(*count, *deleteUnrepeated, *unique) {
		log.Fatal("Can`t use -c,-d,-u together");
	}

	var s []string;

	if *inputFile != "" { 
		file, err := os.Open(*inputFile)
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

	output := fmt.Sprintf("%s", uniq.Uniq(s,*count,
		*deleteUnrepeated,
		*unique,
		*caseInsensitive,
		*ignoreFirst,
		*ignoreStartSymbols));

	if *outputFile=="" {
		fmt.Printf("%s",output);
	} else {
		f, err := os.Create(*outputFile)

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