package main

import (
	"errors"
	"fmt"
	"math"
)

func circleAreaNewError(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}

	return math.Pi * radius * radius, nil
}

func circleAreaErrorF(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}

	return math.Pi * radius * radius, nil
}

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleAreaCustomError(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}

	return math.Pi * radius * radius, nil
}

type areaError2 struct {
	err    string  //error description
	length float64 //length which caused the error
	width  float64 //width which caused the error
}

func (e *areaError2) Error() string {
	return e.err
}

func (e *areaError2) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError2) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError2{err, length, width}
	}
	return length * width, nil
}

func main() {
	radius := -20.0
	// area, err := circleAreaNewError(radius) // USING THE NEW FUNCTION
	area, err := circleAreaErrorF(radius) // ADD MORE INFORMATION

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Area of circle %0.2f", area)
	}

	// ADD MORE INFORMATION USING STRUCT TYPE AND FIELDS
	area, err = circleAreaCustomError(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero\n", err.radius)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("Area of rectangle1 %0.2f\n", area)
	}

	// ADD MORE INFORMATION USING METHODS ON STRUCT TYPES
	length, width := -5.0, -9.0
	area, err = rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError2); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}
