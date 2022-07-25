package main

import "fmt"

// ISP
// Interface Segregation Principle

type Document struct {
	content string
}

type Machine interface {
	Print(d *Document)
	Fax(d *Document)
	Scan(d *Document)
}

type MultiFunctionPrinter struct{}

func (m *MultiFunctionPrinter) Print(d *Document) {
}

func (m *MultiFunctionPrinter) Fax(d *Document) {
}

func (m *MultiFunctionPrinter) Scan(d *Document) {
}

// Somebody working with OldFashionedPrinter
type OldFashionedPrinter struct{}

func (o *OldFashionedPrinter) Print(d *Document) {
}

func (o *OldFashionedPrinter) Fax(d *Document) {
	panic("Operation not supported in this machine!")
}

func (o *OldFashionedPrinter) Scan(d *Document) {
	panic("Operation not supported in this machine!")
}

// ISP
// - Try to separate interface that people may need
type Printer interface {
	Print(d *Document)
}

type Scanner interface {
	Scan(d *Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d *Document) {
	fmt.Print(d.content, "\n")
}

type Photocopier struct{}

func (p Photocopier) Scan(d *Document) {
	panic("implement me")
}

func (p Photocopier) Print(d *Document) {
	panic("implement me")
}

// Decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m *MultiFunctionMachine) Print(d *Document) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Print("Printer not ready!. Please, initialize printer first!\n")
		}
	}()
	m.printer.Print(d)
}

func (m *MultiFunctionMachine) Scan(d *Document) {
	defer func() {
		if a := recover(); a != nil {
			fmt.Print("Scanner not ready!. Please, initialize scanner first!\n")
		}
	}()
	m.scanner.Scan(d)
}

func main() {
	doc := Document{content: "This Doc!"}
	//printer := MyPrinter{}
	multiFunctionMachines := MultiFunctionMachine{}

	multiFunctionMachines.Print(&doc)
	multiFunctionMachines.Scan(&doc)
}
