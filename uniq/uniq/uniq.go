package uniq

import (
	"fmt"
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


// Находит уникальные/неуникальные строки по параметрам из
// массива строк и возвращает строку
func Uniq(text []string,
		   count bool,
		   deleteUnrepeated bool,
		   unique bool,
		   caseInsensitive bool,
		   ignoreFirst int,
		   ignoreStartSymbols int) string {
			var result string;
			
			stringComparator := compareTwoStrings;
			
			if caseInsensitive {
				stringComparator = strings.EqualFold;
			}
			
			text = append(text, "\n");
			previous := subSentence(text[0],ignoreFirst,ignoreStartSymbols);
			
			var counter int;

			for i := 1; i < len(text); i++ {
				current := subSentence(text[i],ignoreFirst,ignoreStartSymbols);
				
				if !stringComparator(previous, current) { 
					if count {
						result += fmt.Sprintf("%d ", counter+1);
						counter=0;
					}
					if !(unique || deleteUnrepeated) { 
						result += fmt.Sprintf("%s\n", text[i-1]);
					} else {
						if(counter==0 && unique){
							result += fmt.Sprintf("%s\n", text[i-1]);
						}
						if (counter>0 && deleteUnrepeated){
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