package main

import "fmt"

func main(){
	 myArray:= [...] int {11,22,33}
	 mySlice:=myArray[:]
	 fmt.Println("Slice: ",mySlice)
	 mySlice[0]=99
	 fmt.Println("Slice: ",mySlice)

	 myColors:=[...]string{"Red","Yellow","Green","Blue","Pink"}
	 myColorsSlice:=myColors[1:len(myColors)-1]
	 fmt.Println("Color Slice: ",myColorsSlice)

	 nums:=make([]int,5)
	 nums[1]=11
	 fmt.Println(nums)
	 nums=append(nums, 112)
	 fmt.Println(nums)

	 fmt.Println("LENGTH: ",len(nums))
	 fmt.Println("CAPACITY: ",cap(nums))

	 myArray2:=[10]int{11,22,33,44,55,66,77,88,99,111}
	 mySlice2:=myArray2[:]
	 mySlice3:=myArray2[:]

	 fmt.Println("O'ZAIRSHSIZ")
	 fmt.Println("MY ARRAY 2: ",myArray2)
	 fmt.Println("MY SLICE 2: ",mySlice2)
	 fmt.Println("MY SLICE 3: ",mySlice3)

	 mySlice2[0] = 123
	 fmt.Println("O'ZGARISHI BILAN")
	 fmt.Println("MY ARRAY 2: ",myArray2)
	 fmt.Println("MY SLICE 2: ",mySlice2)
	 fmt.Println("MY SLICE 3: ",mySlice3)

}