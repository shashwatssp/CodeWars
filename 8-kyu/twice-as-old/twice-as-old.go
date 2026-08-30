package kata
​
//import "math"
​
func TwiceAsOld(dadYearsOld, sonYearsOld int) int { 
   years := 2* sonYearsOld - dadYearsOld
​
   if(years<0){
     years*=-1
   }
​
   return years
}