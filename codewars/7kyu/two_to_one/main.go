package main
import ( 
  "fmt"
  "sort"
)

func TwoToOne(s1 string, s2 string) string {
  s            := []rune(s1+s2)
  resultString := []rune{}
  stringMap    := make(map[rune]bool)

  sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
  baseString := string(s)
  for _, char := range baseString  {
    if _, value := stringMap[char]; !value {
		stringMap[char] = true
		resultString = append(resultString, char)
	}
  }
  return string(resultString)
}

func main() {
	fmt.Println(TwoToOne("azzzaatttffggbbaallla", "bxxxxxxb"))
}