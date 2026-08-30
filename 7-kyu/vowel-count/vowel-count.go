package kata
func isVowel (char rune) bool{
  return char=='A' || char=='E' || char=='I' || char=='O' || char=='U' || char=='a' || char=='e' || char=='i' || char=='o' || char=='u' 
}
func GetCount(str string) (count int) {
  for _,char:= range str{
    if(isVowel(char)){
      count++
    }
  }
  return count
}