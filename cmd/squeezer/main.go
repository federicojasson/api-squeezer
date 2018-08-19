package main

import (
	"flag"
	"github.com/federicojasson/api-squeezer"
	"github.com/federicojasson/api-squeezer/serial"
	"github.com/federicojasson/api-squeezer/validation"
	"log"
	"math"
)

func main() {
	var file string
	var verbose bool
	var spec squeezer.Spec
	var validators []squeezer.Validator
	var serializers []squeezer.Serializer

	// TODO: add usage message
	flag.StringVar(&file, "file", "", "")
	flag.BoolVar(&verbose, "verbose", false, "")
	flag.StringVar(&spec.URLPattern, "url-pattern", "", "")
	flag.StringVar(&spec.Method, "method", "GET", "")
	flag.StringVar(&spec.ContentType, "content-type", "application/json", "")
	flag.IntVar(&spec.Begin, "begin", 1, "")
	flag.IntVar(&spec.End, "end", math.MaxInt32, "")
	flag.IntVar(&spec.Batch, "batch", 10, "")
	flag.IntVar(&spec.Tolerance, "tolerance", 10, "")
	flag.IntVar(&spec.Backoff, "backoff", 5, "")
	flag.Parse()

	validator := validation.NewNoOpValidator()
	validators = append(validators, validator)

	if file != "" {
		serializer := serial.NewFileSerializer(file)
		serializers = append(serializers, serializer)
	}

	if verbose {
		serializer := serial.NewConsoleSerializer()
		serializers = append(serializers, serializer)
	}

	runner := squeezer.NewRunner()
	err := runner.Run(spec, validators, serializers)

	if err != nil {
		log.Fatal(err)
	}
}
