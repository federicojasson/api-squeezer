package squeezer

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Spec struct {
	URLPattern  string
	Method      string
	ContentType string
	Begin       int
	End         int
	Batch       int
	Tolerance   int
	Backoff     int
}

func (s *Spec) Validate() error {
	// TODO: validate URLPattern?
	// TODO: validate Method
	// TODO: validate ContentType?

	if s.Begin > s.End {
		return errors.New("TODO")
	}

	if s.Batch < 1 {
		return errors.New("TODO")
	}

	if s.Tolerance < 0 {
		return errors.New("TODO")
	}

	if s.Backoff < 0 {
		return errors.New("TODO")
	}

	return nil
}

type Task struct {
	spec    Spec
	reader  *Reader
	writers []*Writer
	errors  int
}

func NewTask(spec Spec, reader *Reader, writers []*Writer) *Task {
	return &Task{
		spec:    spec,
		reader:  reader,
		writers: writers,
		errors:  0,
	}
}

func (t *Task) Run() error {
	for i := t.spec.Begin; i < t.spec.End; i++ {
		req, err := t.prepare(i)

		if err != nil {
			return err
		}

		t.execute(req)

		if t.errors > t.spec.Tolerance {
			break
		}

		if (i-t.spec.Begin+1)%t.spec.Batch == 0 {
			t.sleep()
		}
	}

	return nil
}

func (t *Task) prepare(index int) (*http.Request, error) {
	url := fmt.Sprintf(t.spec.URLPattern, index)
	req, err := http.NewRequest(t.spec.Method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", t.spec.ContentType)
	return req, nil
}

func (t *Task) execute(req *http.Request) {
	data, err := t.reader.Read(req)

	if err == nil {
		t.errors = 0
	} else {
		t.errors++
	}

	if err != nil {
		return
	}

	for _, writer := range t.writers {
		writer.Write(data)
	}
}

func (t *Task) sleep() {
	duration := time.Duration(t.spec.Backoff) * time.Second
	time.Sleep(duration)
}
