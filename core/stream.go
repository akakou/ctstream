package core

type CtStream interface {
	Init() error
	Start(Callback)
	Next(callback Callback)
	Run(Callback)
	Stop()
	Await()
}
