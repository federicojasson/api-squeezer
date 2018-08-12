package main

import (
	"flag"
	"github.com/federicojasson/api-squeezer"
	"log"
	"math"
)

func main() {
	var backoff int
	var batch int
	var begin int
	var contentType string
	var end int
	var method string
	var tolerance int
	var urlPattern string

	// TODO: add descriptions
	flag.IntVar(&backoff, "backoff", 1, "TODO")
	flag.IntVar(&batch, "batch", 10, "TODO")
	flag.IntVar(&begin, "begin", 1, "TODO")
	flag.StringVar(&contentType, "content-type", "application/json", "TODO")
	flag.IntVar(&end, "end", math.MaxInt32, "TODO")
	flag.StringVar(&method, "method", "GET", "TODO")
	flag.IntVar(&tolerance, "tolerance", 10, "TODO")
	flag.StringVar(&urlPattern, "url-pattern", "", "TODO")
	flag.Parse()

	task, err := squeezer.NewTask(squeezer.Spec{
		Backoff:     backoff,
		Batch:       batch,
		Begin:       begin,
		ContentType: contentType,
		End:         end,
		Method:      method,
		Tolerance:   tolerance,
		URLPattern:  urlPattern,
	})

	if err != nil {
		log.Fatal(err)
	}

	task.Run()
}
