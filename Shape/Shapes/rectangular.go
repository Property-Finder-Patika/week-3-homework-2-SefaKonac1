package shapes

/*Rectangular is subclass of Shape abstract class. It derives all attributes from super class via composition operation*/
type Rectangular struct {
	s Shape
}

/*calculates of rectangular area*/
func (r Rectangular) CalculateArea() float64 {
	return float64(r.s.width * r.s.length)
}

/*calculates of rectangular circumstances*/
func (r Rectangular) CalculateCircumstances() float64 {
	return float64(r.s.width*2 + r.s.length*2)
}

/*the help of this method, It could be creating new rectangular objects.*/
func NewRectangular(width float32, length float32) Rectangular {

	/*if width or length of rectangular smaller than zero It generates panic situation */
	if width <= 0 || length <= 0 {
		panic("Invalid Parameter Error : width or length cannot be smaller than 0 !\n")
	}
	/*otherwise it creates a new rectangular object and return it's value*/
	shape := Shape{width, length}
	rectangular := Rectangular{s: shape}

	return rectangular
}
