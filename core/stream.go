package core

type CtStream interface {
	Init() error
	Next(Callback)
	Run(Callback)
	Stop()
	Await()
}
