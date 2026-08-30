package kata
​
func isVowel(char rune) bool{
​
return char=='a' || char=='e' || char=='i'   || char=='o' || char=='u' ||  char=='A' || char=='E' || char=='I'   || char=='O' || char=='U'
}
func Disemvowel(comment string) string {
result:=""
​
for _,char:= range comment{
  if !isVowel(char){
result+= string(char)
 }
}
​
return result
}