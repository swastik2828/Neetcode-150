func lengthOfLongestSubstring(s string) int {
	ls := make([]int, 256)
	for i := range ls {
    	ls[i] = -1
	}

	 l := 0
	 r := 0
	 maxlen := 0;
    n :=  len(s);

        for r < n {
            if(ls[s[r]] != -1){
                if(ls[s[r]] >= l){
                    l = ls[s[r]] + 1;
                }
            }
            len := r-l+1;
            maxlen = max(len,maxlen);

            ls[s[r]] = r;
            r++;
        }
        return maxlen;

}
