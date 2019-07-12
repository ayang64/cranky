package cranky

import "log"

func gen(start uint64, count uint64, depth int) <-chan uint64 {
	rc := make(chan uint64, depth)
	go func() {
		for i, end := start, start+count; i < end; i++ {
			rc <- i
		}
		close(rc)
	}()
	return rc
}

func pow10(e uint64) uint64 {
	rc := uint64(1)
	for i := uint64(0); i < e; i++ {
		rc *= 10
	}
	return rc
}

func undigitize(s []uint64) uint64 {
	rc := uint64(0)
	for cur := s; len(cur) > 0; cur = cur[1:] {
		rc += cur[0] * pow10(uint64(len(cur)-1))
	}
	return rc
}

func digitize(v uint64) []uint64 {
	rc := []uint64{}
	for i := v; i != 0; i /= 10 {
		rc = append([]uint64{i % 10}, rc...)
	}
	return rc
}

func permute(d []uint64) <-chan [2]uint64 {
	rc := make(chan [2]uint64)
	go func() {
		for i := 1; i < len(d); i++ {
			rc <- [2]uint64{undigitize(d[:i]), undigitize(d[i:])}
		}
		close(rc)
	}()
	return rc
}

func Sum(j int, start uint64, count uint64) uint64 {
	reverse := func(s []uint64) []uint64 {
		rc := make([]uint64, len(s), len(s))
		for i := 0; i < len(s); i++ {
			rc[i] = s[len(s)-i-1]
		}
		return rc
	}

	iscranky := func(v1 uint64, v2 uint64, a []uint64) bool {
		h := digitize(v1 * v2)
		if len(h) != len(a) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if h[i] != a[i] {
				return false
			}
		}
		return true
	}

	rc := uint64(0)

	sem := make(chan struct{}, j)
	for v := range gen(start, count, j) {
		sem <- struct{}{}
		go func(v uint64) {
			d := digitize(v)
			answer := reverse(d)
			for pair := range permute(d) {
				if iscranky(pair[0], pair[1], answer) {
					log.Printf("%v is cranky (%v)", v, pair)
					rc += v
				}
			}
			<-sem
		}(v)
	}
	return rc
}
