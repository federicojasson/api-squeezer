package squeezer

import (
	"log"
	"sync"
)

type Runner struct {
	barrier *sync.WaitGroup
}

func NewRunner() *Runner {
	return &Runner{
		barrier: &sync.WaitGroup{},
	}
}

func (r *Runner) Run(spec Spec, validators []Validator, serializers []Serializer) error {
	err := spec.Validate()

	if err != nil {
		return err
	}

	reader := NewReader(validators)
	writers, err := r.prepare(serializers, spec.Batch)
	task := NewTask(spec, reader, writers)

	if err != nil {
		r.stop(writers)
		return err
	}

	r.start(writers)
	err = task.Run()
	r.stop(writers)

	if err != nil {
		return err
	}

	r.barrier.Wait()
	return nil
}

func (r *Runner) prepare(serializers []Serializer, buffer int) ([]*Writer, error) {
	var writers []*Writer

	for _, serializer := range serializers {
		writer := NewWriter(serializer, buffer)
		err := writer.Open()

		if err != nil {
			return writers, err
		}

		writers = append(writers, writer)
	}

	return writers, nil
}

func (r *Runner) start(writers []*Writer) {
	for _, writer := range writers {
		r.barrier.Add(1)

		go func(writer *Writer) {
			err := writer.Listen()

			if err != nil {
				log.Println(err)
			}

			r.barrier.Done()
		}(writer)
	}
}

func (r *Runner) stop(writers []*Writer) {
	for _, writer := range writers {
		err := writer.Close()

		if err != nil {
			log.Println(err)
		}
	}
}
