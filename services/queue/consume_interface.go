package queue

type ConsumeInterface interface {
	Run()
	// Handle该方法中必须捕获异常,返回false表示处理发生错误，需要noack,返回true，表示正常处理，可以ack
	Handle(body []byte, consume ConsumeInterface) bool
	HandleExec() bool
}
