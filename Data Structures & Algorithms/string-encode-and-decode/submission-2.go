type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    var result string
    for _,str := range strs {
        // thêm length
        result += strconv.Itoa(len(str)) // độ dài str

        // thêm separator metadata
        result += "#"

        // thêm data
        result += str
    }

    // hello, world => 5#hello5#world

    return result
}

func (s *Solution) Decode(encoded string) []string {
    // 5#hello5#world => hello,world
    var result []string
    for len(encoded) > 0 {
        // tìm dấu #
		index := strings.Index(encoded, "#")

        length, _ := strconv.Atoi(encoded[:index]) // 5#hello

        encoded = encoded[index+1:] // hello 

        // lấy word
        word := encoded[:length]

        result = append(result, word)

        // cắt phần đã lấy
        encoded = encoded[length:] 
    }

    return result
}
