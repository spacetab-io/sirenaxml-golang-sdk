package publisher

type Publisher interface {
	PublishLogs(logAttributes map[string]string, request, response []byte) error
}