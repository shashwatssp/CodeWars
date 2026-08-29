package kata
​
func FindMultiples(integer, limit int) []int {
  result:= []int{};
  
  i:=1
​
  for integer*i<= limit{
  result= append(result, integer*i)
  i++
  }
  return result
}
​