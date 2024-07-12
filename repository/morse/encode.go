package morse

import (
	"context"
)

func (r repo) Encode(ctx context.Context, req string) []string {
	var pcs []string
	var i int
	for len(pcs) < len(req) {
		a := req[i : i+1]
		//log.Printf("a[%d] %s", i, a)

		i++
		for key, val := range Morse {
			if a == key {
				pcs = append(pcs, val)
				break
			}
		}
	}

	return pcs
}
