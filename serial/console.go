package serial

import "fmt"

type ConsoleSerializer struct {
}

func NewConsoleSerializer() *ConsoleSerializer {
	return &ConsoleSerializer{}
}

func (s *ConsoleSerializer) Open() error {
	return nil
}

func (s *ConsoleSerializer) Close() error {
	return nil
}

func (s *ConsoleSerializer) Serialize(data []byte) error {
	line := string(data)
	_, err := fmt.Println(line)

	return err
}
