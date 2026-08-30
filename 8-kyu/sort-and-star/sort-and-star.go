package kata
​
func TwoSort(arr []string) string {
smallest:= arr[0]
​
for _,word:= range arr{
  if word<smallest{
    smallest = word
  }
}
​
result:=""
​
for i,char:= range smallest{
​
  if(i>0){
    result+= "***"
  }
​
  result+= string(char)
}
​
return result
}