package kata
​
func Enough (cap, on, wait int) int {
res:= cap-on-wait
​
if(res>=0){
  return 0
}
​
return -1*res
​
}