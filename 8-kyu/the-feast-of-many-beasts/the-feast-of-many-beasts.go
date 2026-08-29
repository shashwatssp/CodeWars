package kata
​
func Feast(beast string, dish string) bool {
  n:= len(dish)
  m:= len(beast)
  return  beast[0]==dish[0] && beast[m-1]==dish[n-1]
}