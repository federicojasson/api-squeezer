package squeezer

import (
	"io/ioutil"
	"net/http"
)

type Validator interface {
	Validate(data []byte) error
}

type Reader struct {
	validators []Validator
	client     *http.Client
}

func NewReader(validators []Validator) *Reader {
	return &Reader{
		validators: validators,
		client:     &http.Client{},
	}
}

func (r *Reader) Read(req *http.Request) ([]byte, error) {
	res, err := r.client.Do(req)

	if err != nil {
		return nil, err
	}

	return r.unmarshal(res)
}

func (r *Reader) unmarshal(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	for _, validator := range r.validators {
		err = validator.Validate(data)

		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
