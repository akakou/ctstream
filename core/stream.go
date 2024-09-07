package core

type CtStream interface {
	Init() error
	Start(Callback)
	Run(Callback)
	Stop()
	Await()
}
