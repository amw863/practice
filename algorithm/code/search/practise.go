package search

func bs(v int, s[]int) bool {
    l := 0
    h := len(s) -1
    
    for l <= h {
        m := l+(h-l)>>1
        if s[m] == v {
            return true
        } else if s[m] > v {
            h = m - 1 
        }else {
            l = m +1
        }
    }

    return false
}
