//		Input "go run "Program B"/main.go" into terminal to run file

/*
Demonstrate loops,  DONE
arrays,	DONE
slices, DONE
math operations, DONE
conditional statements, DONE
functions/methods, DONE
user input/output DONE
*/


package main

import (
	"fmt"
	"strconv"
	"math"
)

// type Thing interface {
// 	DoSomething()
// }

type Thing struct {
	thingNumber int
	thingSkill string
}

func (t Thing) DoSomething() {
	fmt.Println("I am Thing ", t.thingNumber, " and I am good at ", t.thingSkill)
}

func playWithLoops(){
	fmt.Println("Loop counts to 4")
	for i := 1; i < 5; i ++{
		fmt.Print(i)
		if (i < 4) {
			fmt.Print(", ")
		} 
	}
	fmt.Println()
}

func playWithArrays(){
	array1 := [4]int{}
	array2 := [4]int{1,2,3,4}
	strArray := [4]string{"Duck", "Duck", "Duck", "Goose"}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(strArray)

	fmt.Println("Changed an indice in an array:")
	array1[2] = 2
	fmt.Println(array1)

	fmt.Println("Length of that Array: " + strconv.Itoa(len(array1)))

	println("Loop through a different array")
	for /*leave as _, if you don't want to enumerate*/index, value := range array2{
		fmt.Println("Index ", index, ": ", value)
	}

}

func playWithSclices(){
	array1 := [4]int{1,2,3,4}
	slice1 := array1[:]
	fmt.Println(slice1)

	fmt.Println("append the int '5' to a slice")
	slice1 = append(slice1, 5)
	fmt.Println(slice1)
}

func playWithMath(){
	num := 2

	fmt.Println("Base number (BN): " , num)
	fmt.Println(num, " plus 10: ", (num + 10))
	fmt.Println(num, " minus 10: ", (num - 10))
	fmt.Println(num, " times 10: ", (num * 10))
	fmt.Println(num, " divided by 10: ", (num / 10))
	fmt.Println(num, " divided by 10 (as a float): ", (float64(num)/10)) //		NOTE: simply deviding by 10.0 does not work

	fmt.Println(num, " to the power of 10: ", (math.Pow(float64(num), 10)))
	fmt.Println("Log base 10 of ", num, ": ", (math.Log10(float64(num))))
	fmt.Println("Square root of ", num, ": ", (math.Sqrt(float64(num))))
	fmt.Println("Sin(", num, "): ", math.Sin(float64(num)), " (Radians)")
}

func playWithConditionals(){
	num := 2.0

	fmt.Println(num, " is equal to 2.0: ", (num == 2.0))
	fmt.Println(num, " is NOT equal to 2.0: ", (num != 2.0))
	fmt.Println(num, " is less than 2.0: ", (num < 2.0))
	fmt.Println(num, " is greater than 2.0: ", (num > 2.0))
	fmt.Println(num, " is a multiple of 2.0: ", (int64(num) % 2.0 == 0))
}

func playWithMethods(){

	T := Thing{thingNumber: 0, thingSkill: "nothing"}
	T.DoSomething()
}

func playWithInput() {
	var i int 
	fmt.Print("\nInput an integer: ")
	fmt.Scan(&i)
	fmt.Println(i, "was a poor choice")
}


func main(){
	fmt.Println("Program B")
	println("\n-----------------------\nLOOPS:\n-----------------------")
	playWithLoops()
	println("\n-----------------------\nARRAYS:\n-----------------------")
	playWithArrays()
	println("\n-----------------------\nSLICES:\n-----------------------")
	playWithSclices()
	println("\n-----------------------\nMATH:\n-----------------------")
	playWithMath()
	println("\n-----------------------\nCONDITIONALS:\n-----------------------")
	playWithConditionals()
	println("\n-----------------------\nMETHODS/STRUCTS:\n-----------------------")
	playWithMethods()
	println("\n-----------------------\nINPUTS:\n-----------------------")
	playWithInput()


}
