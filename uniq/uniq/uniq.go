package uniq

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

// Возвращает подстроку без заданного количества начальных символов
// либо без определенного количества первых слов в строке
func subSentence(s string, ignoreStart int, ignoreStartSymbols int) string {

	if(ignoreStartSymbols>=len(s)) {
		return s;
	}
	current := s[ignoreStartSymbols:];
	
	words := strings.Fields(current);
	
	if(len(words)>0) {
	 	current = strings.Join(words[ignoreStart:]," ");
	}
	
	return current; 
}
 
// Сравнивает две строки и возвращает true, если они
// одинаковые (Необходимо для того,чтобы обычная функция сравнения
// была совместима с strings.EqualFold и ее можно было поместить в
// StringComparator uniq.go 43 строка) ,что спасает от огомного числа if-ов
func compareTwoStrings(s string, t string) bool {
	return s==t;
}


type Params struct {
	Count bool;
	DeleteUnrepeated bool;
	Unique bool;
	CaseInsensitive bool;
	IgnoreFirst int;
	IgnoreStartSymbols int;
	InputFile string;
	OutputFile string;
}


// Разбирает параметры и складывает в структуру с параметрами
func ParseParams() Params {
	count := flag.Bool("c", false, "Count repeating strings?");
	deleteUnrepeated := flag.Bool("d", false, "Delete repeating strings?");
	unique := flag.Bool("u", false, "Show unrepeated strings only?");
	caseInsensitive := flag.Bool("i", false, "case-insensitive");
	ignoreFirst := flag.Int("f", 0, "Ignore first {num} lines");
	ignoreStartSymbols := flag.Int("s",0, "Ignore last {num} lines");
	inputFile := flag.String("input_file", "", "File to read");
	outputFile := flag.String("output_file", "", "File to write");
	flag.Parse();
	if СheckParams(*count, *deleteUnrepeated, *unique) {
		log.Fatal("Can`t use -c,-d,-u together");
	}
	return Params{
		*count,
		*deleteUnrepeated,
		*unique,
		*caseInsensitive,
		*ignoreFirst,
		*ignoreStartSymbols,
		*inputFile,
		*outputFile,
	}
}


// Находит уникальные/неуникальные строки по параметрам из
// массива строк и возвращает строку
func Uniq(text []string,parsedParams Params) string {var result string;
			
			stringComparator := compareTwoStrings;
			
			if parsedParams.CaseInsensitive {
				stringComparator = strings.EqualFold;
			}
			
			text = append(text, "\n");
			previous := subSentence(text[0],parsedParams.IgnoreFirst,parsedParams.IgnoreStartSymbols);
			
			var counter int;

			for i := 1; i < len(text); i++ {
				current := subSentence(text[i],parsedParams.IgnoreFirst,parsedParams.IgnoreStartSymbols);
				
				if !stringComparator(previous, current) { 
					if parsedParams.Count {
						result += fmt.Sprintf("%d ", counter+1);
						counter=0;
					}
					if !(parsedParams.Unique || parsedParams.DeleteUnrepeated) { 
						result += fmt.Sprintf("%s\n", text[i-1]);
					} else {
						if(counter==0 && parsedParams.Unique){
							result += fmt.Sprintf("%s\n", text[i-1]);
						}
						if (counter>0 && parsedParams.DeleteUnrepeated){
							result += fmt.Sprintf("%s\n", text[i-1]);
						}
						counter = 0;
					}
				} else {
					counter++;
				}
				previous = current;
			}

			return result;
}

// Проверка совместимости флагов
func СheckParams(count , deleteUnrepeated, unique bool) (bool) {
	return count && (deleteUnrepeated || unique) || (deleteUnrepeated && unique)
}