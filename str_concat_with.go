package sprint


func StrConcatWith(strs []string, sep string) string {
    var result string

    if len(strs) > 0 {
        result = strs[0]

        for i := 1; i < len(strs); i++ {
            result += sep + strs[i]
        }
    }

    return result
}