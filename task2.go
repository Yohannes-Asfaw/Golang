package golang
import (
	"strings"
	"unicode"
)
// Task one
func wordFrequencyCount(text string) map[string]int{
	text=strings.ToLower(text)
	var cleanedText strings.Builder
    for _,char := range text{

		if unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char){
			cleanedText.WriteRune(char)

		}
	}
	words:=strings.Fields(cleanedText.String())
	frequency:=make(map[string]int)
	for _,word:=range words{
		frequency[word]++
	}
	return frequency
}

// Task two
func isPalidrome(text string ) bool{
	text=strings.ToLower(text)
	var cleanedText strings.Builder

	for _, char:= range text{
		if unicode.IsLetter(char) || unicode.IsDigit(char){
			cleanedText.WriteRune(char)
		}
	}
	cleanedstr:=cleanedText.String()
	n:=len(cleanedstr)
	 
	for i :=0;i<n/2;i++{
		if cleanedstr[i]!=cleanedstr[n-1-i]{
			return false
		}
		
	}
	return true

}