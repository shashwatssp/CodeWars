package kata
import "strings"
​
func FindShort(s string) int {
​
  minLen:= 1000000
​
  words:= strings.Fields(s)
​
​
  for _,word:= range words{
  
   curLen:= len(string(word))
​
   if(curLen< minLen){
   minLen = curLen
   }
  }
  return minLen
}