[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** interview-qa/arrays
**Tags:** [go, arrays, functions]
]

# Go তে Array

## Array ঘোষণা করা

Go-তে নিচের সিনট্যাক্স ব্যবহার করে array ঘোষণা করা যায়:

```go
var arrayName [size]elementType
```

উদাহরণ:

```go
var numbers [5]int
```

## Arrays  ইনিশিয়ালাইজ করা

Array ডিক্লেয়ার করার সময় Array ইনিশিয়ালাইজ করা যেতে পারে:

```go
var numbers = [5]int{1, 2, 3, 4, 5}
```

অথবা শর্টহ্যান্ড নোটেশন ব্যবহার করতে পারেন:

```go
numbers := [5]int{1, 2, 3, 4, 5}
```

## Array উপাদানে অ্যাক্সেস করা

Array এর উপাদানগুলো ইনডেক্স ব্যবহার করে অ্যাক্সেস করা যায়, যেটি 0 থেকে শুরু হয়:

```go
fmt.Println(numbers[0]) // Output: 1
```

## Array তে লুপ চালানো

`for` লুপ ব্যবহার করে অ্যারের উপর লুপ চালানো যায়:

```go
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

অথবা `range` কীওয়ার্ড ব্যবহার করতে পারেন:

```go
for index, value := range numbers {
    fmt.Println(index, value)
}
```

## Multidimensional Arrays

Go মাল্টিডাইমেনশনাল অ্যারে সাপোর্ট করে। একটি টু-ডাইমেনশনাল অ্যারে নিচেরভাবে ঘোষণা করা হয়:

```go
var matrix [3][3]int
```

উদাহরণ:

```go
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
```

## অ্যারে অব অ্যারেস (Array of Arrays)

অ্যারে অব অ্যারেস (array of arrays) হলো এমন একটি ডেটা স্ট্রাকচার, যেখানে একটি অ্যারের প্রতিটি উপাদান আরেকটি অ্যারে বা স্লাইস। এটি সাধারণত ভ্যারিয়েবল সাইজের 2D ডেটা রাখার জন্য ব্যবহৃত হয়, যেমন প্রতিটি সারিতে ভিন্ন ভিন্ন উপাদান থাকতে পারে।


উদাহরণ:

```go
arrayOfArrays := [][]int{
    {1, 2},
    {3, 4, 5},
    {6},
}
```

## ফাংশনে Arrays পাঠানো

Go-তে অ্যারে ফাংশনে পাঠানো হয় ভ্যালু হিসেবে, অর্থাৎ ফাংশনটি অ্যারের একটি কপি পায়:

```go
func printArray(arr [5]int) {
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
}
```

মূল অ্যারে পরিবর্তন করতে হলে, আপনাকে অ্যারের পয়েন্টার পাঠাতে হবে:

```go
func modifyArray(arr *[5]int) {
    arr[0] = 100
}
```

## Frequently Asked Questions

### Q1: Go-তে কিভাবে অ্যারের দৈর্ঘ্য নির্ধারণ করব?

অ্যারের দৈর্ঘ্য নির্ধারণ করতে len() বিল্ট-ইন ফাংশন ব্যবহার করুন।

উদাহরণ:

```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Length of the array:", len(arr)) // Output: 5
}
```

### Q2: Go-তে কিভাবে একটি অ্যারে কপি করব?

Go-তে অ্যারে কপি করার জন্য আপনি সেটিকে একই টাইপ এবং সাইজের অন্য অ্যারেতে অ্যাসাইন করতে পারেন।

উদাহরণ:

```go
package main

import "fmt"

func main() {
    original := [3]int{1, 2, 3}
    copy := original
    fmt.Println("Original:", original) // Output: [1 2 3]
    fmt.Println("Copy:", copy)         // Output: [1 2 3]
}
```

### Q3: কিভাবে কপি না করে ফাংশনে array পাঠিয়ে সেটিকে পরিবর্তন করা যায়?

কপি না করে পরিবর্তন করতে চাইলে অ্যারের পয়েন্টার পাঠান।

উদাহরণ:

```go
package main

import "fmt"

func modifyArray(arr *[3]int) {
    arr[0] = 42
}

func main() {
    arr := [3]int{1, 2, 3}
    modifyArray(&arr)
    fmt.Println("Modified array:", arr) // Output: [42 2 3]
}
```

### test করার জন্য উদাহরণ: main.go

```go
package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var arr3 [5]int
	// fmt.Println(arr3) // this should print [0 0 0 0 0]
	// println("Length:", len(arr3))

	for i := 0; i < len(arr3); i++ {
	    fmt.Scan(&arr3[i])
	}
	fmt.Println(arr3)

	strArr := [3]string{"one", "two", "three"}

	for i := 0; i < len(strArr); i++ {
		fmt.Print(strArr[i]+ "")
	}


}



```
