type Solution struct{}

// func (s *Solution) Encode(strs []string) string {
//     // strs ["hello","world"] => str hello,world
//     var result string
//     for index, str := range strs {
//         result += str
//         if index != len(strs) - 1 { 
//             result += ","
//         } 
//     }
//     return result
// }

func (s *Solution) Encode(strs []string) string {
    // strs ["hello","world"] => str hello,world
    var result string
    for i :=0 ; i < len(strs);i++{
        result += strs[i]
        if i != len(strs) - 1 { 
            result += ","
        } 
    }
    return result
}

func (s *Solution) Decode(encoded string) []string {
   // str hello,world => strs ["hello","world"]
   var result[]string
   var current string
   for i:=0;i<len(encoded);i++ {
        if encoded[i] == ',' {
            result = append(result,current) // append world 
            current = ""
        }else{
            current += string(encoded[i])
        }
    }

    result = append(result, current) // add last word 
    return result
}
