package squeezer

import (
	"errors"
	"log"
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
)

type Spec struct {
	Backoff     int
	Batch       int
	Begin       int
	ContentType string
	End         int
	Method      string
	Tolerance   int
	URLPattern  string
}

type Task struct {
	spec Spec
	// TODO: rename?
	client *http.Client
}

func NewTask(spec Spec) (*Task, error) {
	if spec.Backoff < 0 {
		return nil, errors.New("TODO")
	}

	if spec.Batch < 1 {
		return nil, errors.New("TODO")
	}

	if spec.Begin > spec.End {
		return nil, errors.New("TODO")
	}

	// TODO: validate ContentType?
	// TODO: validate Method
	// TODO: validate URLPattern?

	if spec.Tolerance < 0 {
		return nil, errors.New("TODO")
	}

	client := &http.Client{}

	return &Task{
		spec,
		client,
	}, nil
}

func (t *Task) Run() {
	log.Printf("TODO")

	// TODO: rename
	TODOerrors := 0

	for i := t.spec.Begin; i < t.spec.End; i++ {
		data, err := t.Request(i)

		if err == nil {
			TODOerrors = 0
		} else {
			TODOerrors++
		}

		if TODOerrors > t.spec.Tolerance {
			log.Printf("TODO %d", TODOerrors)
			break
		}

		if (i-t.spec.Begin+1)%t.spec.Batch == 0 {
			log.Printf("TODO %d seconds", t.spec.Backoff)
			time.Sleep(time.Duration(t.spec.Backoff) * time.Second)
		}
	}

	log.Printf("TODO")
}

func (t *Task) Request(index int) ([]byte, error) {
	url := fmt.Sprintf(t.spec.URLPattern, index)
	req, err := http.NewRequest(t.spec.Method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", t.spec.ContentType)
	res, err := t.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	// TODO: add custom error detection logic

	return data, nil
}
