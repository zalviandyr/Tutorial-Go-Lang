package main

import "errors"

func main() {
	result, err := Divide(5, 2)
	if err != nil {
		println(err.Error())
	} else {
		println(result)
	}
}

func Divide(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("must greater then 0")
	} else {
		return a / b, nil
	}
}
