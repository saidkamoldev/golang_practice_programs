package main

import (
	"fmt"
)

// Printer interfeysi
type Printer interface {
	Print() string
}

// Copier interfeysi, Printer interfeysini ham o'z ichiga oladi
type Copier interface {
	Printer
	Copy() string
}

// MultiFunctionPrinter struct
type MultiFunctionPrinter struct {
	brand string
}

func (mfp MultiFunctionPrinter) Print() string {
	return "Printing document from" + mfp.brand
}

func (mfp MultiFunctionPrinter) Copy() string {
	return "Copying document from" + mfp.brand
}

func main() {
	var device Copier = MultiFunctionPrinter{brand: "Canon"}

	fmt.Println(device.Print())
	fmt.Println(device.Copy())
}
