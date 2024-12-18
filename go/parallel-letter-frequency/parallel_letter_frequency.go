package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	freqCh := make(chan FreqMap);
    for _,str := range texts{
			go func(s string){
                freqCh <- Frequency(s)
            }(str)
    }
    result := FreqMap{}
    for range texts{
    for l,v := range <-freqCh{
        result[l]+=v
    }    
    }
    
    fmt.Println(result);
    return result
}
