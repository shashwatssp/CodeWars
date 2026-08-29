package kata
​
func DNAStrand(dna string) string {
  result:= ""
​
  for _,char := range dna{
    if char == 'T'{
    result+= string("A")
    } else if char =='A'{
   result+= string("T")
    } else if char == 'G'{
    result+= string("C")
    } else if char =='C'{
   result+= string("G")
    } else{
      result += string(char)
    }
  }
  return result
}
​