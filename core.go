package ctstream

type CtStream interface {
	Init() error
	Start(callback Callback)
	Await()
	Run()
	Stop()
}
