
func birthdayCakeCandles(candles []int32) int32 {
    // Write your code here
    hashMap := make(map[string]int)
    max := candles[0]
    for _, v := range candles {
        if v > max {
            max = v
        }

        value, exists := hashMap[strconv.Itoa(int(v))]
        if exists {
            hashMap[strconv.Itoa(int(v))] = value + 1
        } else {
            hashMap[strconv.Itoa(int(v))] = 1
        }
    }

    return int32(hashMap[strconv.Itoa(int(max))])
}