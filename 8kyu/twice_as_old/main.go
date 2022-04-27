package main
import "fmt"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
   for i := 0; i < sonYearsOld; i++ {
	   decreaseDadYear := dadYearsOld - i
	   decreaseSonYear := sonYearsOld - i
	   if decreaseSonYear == 0 {
		   break
	   }
	   if decreaseDadYear % decreaseSonYear == 0 && decreaseDadYear / decreaseSonYear == 2 {
		   return i
	   }
   }
   for i := 0; i < 150; i++ {
       increaseDadYear := dadYearsOld + i
	   increaseSonYear := sonYearsOld + i
	   if increaseSonYear == 0 {
		   continue
	   }
	   if increaseDadYear % increaseSonYear == 0 && increaseDadYear / increaseSonYear == 2 {
		   return i
	   }
   }
   return -1
}

func main() {
	fmt.Println(TwiceAsOld(44, 22))
}