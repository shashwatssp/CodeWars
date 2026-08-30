package kata
​
func CountPositivesSumNegatives(numbers []int) []int {
  var res []int
​
  cntPos:=0
  cntNeg:=0
​
  for _,num:= range numbers{
    if(num>0){
    cntPos++
    } else if (num<0){
      cntNeg+= num
    }
  }
  
  res = append(res,cntPos)
  res = append(res, cntNeg)
 
  return res
}
​