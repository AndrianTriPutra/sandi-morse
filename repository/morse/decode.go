package morse

import (
	"context"
	"strings"
)

func (r repo) Decode(ctx context.Context, req string) []string {
	var pcs []string
	for strings.Contains(req, "|") {
		data := req[0:strings.Index(req, "|")]
		pcs = append(pcs, data)

		last := len(req)
		req = req[strings.Index(req, "|")+1 : last]

		if len(req) > 0 && !strings.Contains(req, "|") {
			pcs = append(pcs, req)
		}
	}

	var response []string
	for _, value := range pcs {
		for key, val := range Morse {
			if value == val {
				response = append(response, key)
				break
			}
		}
	}

	return response
}
