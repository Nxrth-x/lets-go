package main
import (
    "fmt"
)

func main() {
	functions()
  arrays_and_slices()
  conditionals()
	loops()
}

func loops() {
  i := 1
	
	for i <=10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i<=10; i++ {
		fmt.Printf("Number: %d\n", i)
	}

	// FizzBuzz
	for i:=0; i<100; i++ {
		switch(true) {
			case i%15==0:
				fmt.Println("FizzBuzz")
			case i%3==0:
				fmt.Println("Fizz")
			case i%5==0:
				fmt.Println("Buzz")
			default:
				fmt.Println(i)
		}
	}
}

func conditionals() {
  x := 5
  y := 10
  
  // Conditionals
  if x < y {
    fmt.Printf("%d is less than %d\n", x, y)
  } else if x > y {
    fmt.Printf("%d is greater than %d\n", x, y)
  } else {
    fmt.Printf("%d is equal to %d\n", x, y)
  }
    
  // Switch
  color := "red"
  switch color {
	  case "red":
	    fmt.Println("Color is red")
	  case "blue":
	    fmt.Println("Color is blue")
	  default:
	    fmt.Println("Color is not blue or red")
  }
}

func arrays_and_slices() {
  // Arrays
  var fruitArray[2]string
  
  fruitArray[0] = "Apple"
  fruitArray[1] = "Orange"
  
  fmt.Println(fruitArray) // Print the entire array
  fmt.Println(fruitArray[1]) // Print orange
  
  namesArray := [2]string{"Eder", "Lima"} // Declaring and assigning
  fmt.Println(namesArray)
  
  // Slices
  brandsArray := []string {"Mercedes", "Amazon", "GitHub"}
  fmt.Println("Brands:", brandsArray,"\nLength: ", len(brandsArray))
  fmt.Println("Slicing the slice (🍉):", brandsArray[1:2]) // Amazon
}

func functions() {
  fmt.Println(greeting("Eder"))
  fmt.Printf("%d + %d = %d\n", 3, 4, sum(3, 4))
}

func sum(a, b int) int {
  return a + b
}

func greeting(name string) string {
  return "Hello, "+ name + "!"
}