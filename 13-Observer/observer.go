package observer

import "fmt"

// Subject 被观察者接口
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message string)
}

// Observer 观察者接口
type Observer interface {
	Update(message string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers(message string) {
	for _, o := range s.observers {
		o.Update(message)
	}
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
