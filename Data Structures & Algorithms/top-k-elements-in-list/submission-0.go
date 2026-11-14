
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

    // Count frequencies
    for _, num := range nums {
        freq[num]++
    }

    // Extract keys
    keys := make([]int, 0, len(freq))

    for key := range freq {
        keys = append(keys, key)
    }

    // Sort by frequency (descending)
    sort.Slice(keys, func(i, j int) bool {
        return freq[keys[i]] > freq[keys[j]]
    })

    // Return top k
    return keys[:k]
}
