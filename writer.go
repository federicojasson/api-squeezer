package squeezer

type Serializer interface {
	Open() error
	Close() error
	Serialize(data []byte) error
}

type Writer struct {
	serializer Serializer
	channel    chan []byte
}

func NewWriter(serializer Serializer, buffer int) *Writer {
	return &Writer{
		serializer: serializer,
		channel:    make(chan []byte, buffer),
	}
}

func (w *Writer) Open() error {
	return w.serializer.Open()
}

func (w *Writer) Close() error {
	close(w.channel)
	return w.serializer.Close()
}

func (w *Writer) Listen() error {
	for data := range w.channel {
		err := w.serializer.Serialize(data)

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *Writer) Write(data []byte) {
	w.channel <- data
}
