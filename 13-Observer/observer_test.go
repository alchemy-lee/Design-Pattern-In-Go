package observer

import "testing"

func TestObserver(t *testing.T) {
	subject := Subject{observers: []Observer{}}
	observerOne := ConcreteObserverOne{}
	observerTwo := ConcreteObserverTwo{}

	t.Log("Register observers.")
	subject.RegisterObserver(observerOne)
	subject.RegisterObserver(observerTwo)
	t.Log("Start to notify observers.")
	subject.NotifyObservers("test")

	t.Log("Remove observerTwo.")
	subject.RemoveObserver(observerTwo)
	t.Log("Start to notify observers.")
	subject.NotifyObservers("after remove")
}
