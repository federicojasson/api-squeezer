package serial

import "os"

type FileSerializer struct {
	path string
	file *os.File
}

func NewFileSerializer(path string) *FileSerializer {
	return &FileSerializer{
		path: path,
		file: nil,
	}
}

func (s *FileSerializer) Open() error {
	file, err := os.Create(s.path)

	if err != nil {
		return err
	}

	s.file = file
	return nil
}

func (s *FileSerializer) Close() error {
	defer s.file.Close()
	return s.file.Sync()
}

func (s *FileSerializer) Serialize(data []byte) error {
	// TODO: is this the most efficient?
	data = append(data, []byte("\n")...)
	_, err := s.file.Write(data)

	return err
}
