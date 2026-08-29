package kata
​
func Digitize(n int) []int {
if n == 0 {
return []int{0};
}
​
result:= []int{}
​
for n>0{
num := n%10
n = n/10
result = append(result,num)
}
return result
}
​