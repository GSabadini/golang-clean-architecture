package queue

type (
	// Producer port
	Producer interface {
		Publish([]byte) error
	}

	// Consumer port
	Consumer interface {
		Consume() error
	}
)
