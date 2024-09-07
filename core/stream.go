package core

type CtStream interface {
	Init() error
	Run(Callback)
	Stop()
	Await()
}
