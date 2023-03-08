package main

type Product struct {
	name string
	color Color 
	size Size
}

type Color int 

const (
	green Color = iota
	red
	blue
)

type Size int 
const (
	small Size = iota
	large 
	medium
)

type Filter struct {

}


func (f Filter) FilterByColor (color Color,products []Product) []*Product {
	result := make([]*Product,0)
	for i,v := range products {
		if v.color == color {
			result = append(result,&products[i])
		}
	}
	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type AndSpecification struct {
	first,second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}

type ColorSpec struct {
	color Color
}
func (c ColorSpec) IsSatisfied(p *Product) bool {
	return  c.color == p.color
}

type SizeSpec struct {
	size Size
}
func (s SizeSpec) IsSatisfied(p *Product) bool {
	return  s.size == p.size
}

type BetterFilter struct {}

func (bf BetterFilter) Filter (s Specification,products []Product) []*Product {

	result := make([]*Product, 0)

	for i, v := range products {
		if s.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result

}
