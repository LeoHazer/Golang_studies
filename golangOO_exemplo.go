package main

import "fmt"

type Animal struct{
    
}

func (a Animal) Comer() {
    fmt.Println("Comendo")
}

type MembroFamilia struct{
    
}
func (fm MembroFamilia) Nome() {
    fmt.Println("Meu nome não é Johnny")
}
type Cachorro struct {
    Animal // Struct incorporada
    MembroFamilia // Struct incorporada
}
func main() {
    d := Cachorro{}
    d.Comer() // Imprime "Comendo"
    d.Nome() // Imprime "Meu nome não é Johnny"
}
