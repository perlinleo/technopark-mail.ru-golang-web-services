package src

import (
	"fmt"
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

func Uniq(text []string,
		   count bool,
		   deleteUnrepeated bool,
		   unique bool,
		   caseInsensitive bool,
		   ignoreFirst int,
		   ignoreStartSymbols int) string {
			var result string;
			stringComparator := compareTwoString;
			
			if caseInsensitive {
				stringComparator = strings.EqualFold;
			}
			
			if deleteUnrepeated || unique {
				count = true;
			}
		

			previous := subSentence(text[0],ignoreFirst,ignoreStartSymbols);
			var counter int;
			for i := 1; i < len(text); i++ {
				current := subSentence(text[i],ignoreFirst,ignoreStartSymbols);
				// fmt.Println(current);
				if !stringComparator(previous, current) { 
					if count && !(unique || deleteUnrepeated) {
						result += fmt.Sprintf("%d ", counter+1);
						counter=0;
					}
					if !(unique || deleteUnrepeated) { 
						result += fmt.Sprintf("%s\n", text[i-1]);
					} else if unique {
						if(counter==0){
							result += fmt.Sprintf("%s\n", text[i-1]);
						}
						counter = 0;
					} else if deleteUnrepeated {
						// fmt.Printf("word %s added (counter: %d)", text[i-1], counter);
						if (counter>0){
							result += fmt.Sprintf("%s\n", text[i-1]);
						}
						counter = 0;
					}
				} else if count {
					counter++;
				}
				previous = current;
			}
			if(count) && !(unique || deleteUnrepeated) { 
				result+= fmt.Sprintf("%d ", counter+1);
				result += fmt.Sprintf("%s\n", text[len(text)-1]);
				// fmt.Printf("%d ", counter+1)
			} else if unique {
				if(counter==0){
					result += fmt.Sprintf("%s\n", text[len(text)-1]);
				}
			} else if deleteUnrepeated {
				if(counter>0) {
					result += fmt.Sprintf("%s\n", text[len(text)-1])
				}
			} else {
				result += fmt.Sprintf("%s\n", text[len(text)-1]);
			}
			return result;
		   }
