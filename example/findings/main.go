// findings demonstrate how to retrieve the findings from DefectDojo.
// Options are defined to specify the size of the page retrieved, the offset and what information to prefetch with the request.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/blackaichi/go-defectdojo/defectdojo"
)

func main() {
	url := os.Getenv("DOJO_URI")
	token := os.Getenv("DOJO_APIKEY")

	dj, err := defectdojo.NewDojoClient(url, token, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()

	opts := &defectdojo.FindingsOptions{
		Limit:    20,
		Offset:   5,
		Prefetch: "duplicate_finding",
	}

	resp, err := dj.Findings.List(ctx, opts)
	if err != nil {
		fmt.Println("main:", err)
		return
	}

	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("main:", err)
		return
	}

	fmt.Println(string(b))
}
