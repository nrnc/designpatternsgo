package observer

type Subject interface {
	Register(ob Observer)
	Deregister(ob Observer)
	NotifyAll()
}
