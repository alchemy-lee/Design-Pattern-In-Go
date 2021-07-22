package observer

import "fmt"

type Subject struct {
	observers []Observer
}

func (s *Subject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) NotifyObservers(message string) {
	for _, o := range s.observers {
		o.Update(message)
	}
}

// Observer 观察者接口
type Observer interface {
	Update(message string)
}

type ConcreteObserverOne struct {
}

func (o ConcreteObserverOne) Update(message string) {
	fmt.Printf("Observer One receive message: %s\n", message)
}

type ConcreteObserverTwo struct {
}

func (o ConcreteObserverTwo) Update(message string) {
	fmt.Printf("Observer Two receive message: %s\n", message)
}
