package visitor

import "fmt"

type CarElementVisitor interface {
	VisitWheel(wheel Wheel)
	VisitEngine(engine Engine)
	VisitBody(body Body)
	VisitCar(car Car)
}

type Acceptor interface {
	Accept(visitor CarElementVisitor)
}

type Wheel string

func (w Wheel) Name() string {
	return string(w)
}

func (w Wheel) Accept(visitor CarElementVisitor) {
	visitor.VisitWheel(w)
}

type Engine struct{}

func (e Engine) Accept(visitor CarElementVisitor) {
	visitor.VisitEngine(e)
}

type Body struct{}

func (b Body) Accept(visitor CarElementVisitor) {
	visitor.VisitBody(b)
}

type Car []Acceptor

func (c Car) Accept(visitor CarElementVisitor) {
	for _, e := range c {
		e.Accept(visitor)
	}
	visitor.VisitCar(c)
}

type CarElementPrintVisitor struct{}

func (CarElementPrintVisitor) VisitWheel(wheel Wheel) {
	fmt.Println("Visiting " + wheel.Name() + " wheel.")
}

func (CarElementPrintVisitor) VisitEngine(engine Engine) {
	fmt.Println("Visiting engine")
}

func (CarElementPrintVisitor) VisitBody(body Body) {
	fmt.Println("Visiting body")
}

func (CarElementPrintVisitor) VisitCar(car Car) {
	fmt.Println("Visiting car")
}

type CarElementDoVisitor struct{}

func (CarElementDoVisitor) VisitWheel(wheel Wheel) {
	fmt.Println("Kicking my " + wheel.Name() + " wheel.")
}

func (CarElementDoVisitor) VisitEngine(engine Engine) {
	fmt.Println("Starting my engine")
}

func (CarElementDoVisitor) VisitBody(body Body) {
	fmt.Println("Moving my body")
}

func (CarElementDoVisitor) VisitCar(car Car) {
	fmt.Println("Starting my car")
}
