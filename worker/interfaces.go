package worker

type ConsumerInterface interface {
	Start()
	Close()
}
