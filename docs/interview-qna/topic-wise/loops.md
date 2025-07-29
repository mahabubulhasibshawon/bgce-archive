[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** e.g., interview-qa/topic_name
**Tags:** [go, concurrency, channels]
]


# While Loop 

## অন্যান্য প্রোগ্রামিং ভাষার মতো Go-তে `while` লুপের জন্য আলাদা কোনো কীওয়ার্ড নেই। তবে আমরা `for` লুপ ব্যবহার করে `while` লুপের কাজ করতে পারি।

```go
// 0 থেকে 10 পর্যন্ত সংখ্যা প্রিন্ট করার প্রোগ্রাম
package main
import ("fmt")

func main() {
  number := 0

  for number <= 10 {
    fmt.Println(number)
    number++
  }
}
```

## Frequently Asked Questions

### 1. Go-তে কীভাবে infinite loop তৈরি করা যায়?

**উত্তর:** Go-তে শর্ত ছাড়া `for` লুপ ব্যবহার করে একটি infinite loop তৈরি করা যায়।

```go
package main
import ("fmt")

func main() {
  for {
    fmt.Println("This is an infinite loop")
  }
}
```

---

### 2. Go-তে কীভাবে loop থেকে বের হওয়া যায়?

**উত্তর:** `break` স্টেটমেন্ট ব্যবহার করে loop মাঝপথে থামানো যায়।

```go
package main
import ("fmt")

func main() {
  for i := 0; i < 10; i++ {
    if i == 5 {
      break
    }
    fmt.Println(i)
  }
}
```

---

### 3. Go-তে কীভাবে কোনো iteration স্কিপ করা যায়?

**উত্তর:** `continue` স্টেটমেন্ট ব্যবহার করে বর্তমান iteration বাদ দিয়ে পরের iteration-এ যাওয়া যায়।

```go
package main
import ("fmt")

func main() {
  for i := 0; i < 10; i++ {
    if i%2 == 0 {
      continue
    }
    fmt.Println(i)
  }
}
```

---

### 4. Go-তে কী loop-এর সাথে label ব্যবহার করা যায়?

**উত্তর:** হ্যাঁ, nested loops নিয়ন্ত্রণ করতে label ব্যবহার করা যায়।

```go
package main
import ("fmt")

func main() {
OuterLoop:
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      if i == 1 && j == 1 {
        break OuterLoop
      }
      fmt.Println(i, j)
    }
  }
}
```

---

### 5. Go-তে `do-while` লুপ কীভাবে তৈরি করব?

**উত্তর:** Go-তে `do-while` নেই, তবে `for` লুপ ব্যবহার করে একই রকম আচরণ তৈরি করা যায়।

```go
package main
import ("fmt")

func main() {
  number := 0
  for {
    fmt.Println(number)
    number++
    if number > 5 {
      break
    }
  }
}
```

---

### 6. Go-তে কীভাবে slice-এর উপর iterate করা যায়?

**উত্তর:** `range` কীওয়ার্ড ব্যবহার করে slice-এর উপর loop চালানো যায়।

```go
package main
import ("fmt")

func main() {
  numbers := []int{1, 2, 3, 4, 5}
  for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
  }
}
```

---

### 7. Go-তে কীভাবে map-এর উপর iterate করব?

**উত্তর:** `range` কীওয়ার্ড ব্যবহার করে map-এর key ও value-তে loop চালানো যায়।

```go
package main
import ("fmt")

func main() {
  myMap := map[string]int{"a": 1, "b": 2, "c": 3}
  for key, value := range myMap {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
  }
}
```

---

### 8. Go-তে কীভাবে string-এর উপর iterate করব?

**উত্তর:** `range` ব্যবহার করে string-এর প্রতিটি character-এর উপর loop চালানো যায়।

```go
package main
import ("fmt")

func main() {
  str := "hello"
  for index, char := range str {
    fmt.Printf("Index: %d, Character: %c\n", index, char)
  }
}
```

---

### 9. Go-তে লুপ ব্যবহার করে কীভাবে একটি সংখ্যার factorial বের করব?

**উত্তর:** `for` লুপ ব্যবহার করে সহজেই factorial হিসাব করা যায়।

```go
package main
import ("fmt")

func main() {
  number := 5
  factorial := 1
  for i := 1; i <= number; i++ {
    factorial *= i
  }
  fmt.Println("Factorial:", factorial)
}
```

---

### 10. Go-তে কীভাবে লুপ ব্যবহার করে একটি slice রিভার্স করব?

**উত্তর:** `for` লুপ ব্যবহার করে slice-এর উপাদানগুলো swap করে রিভার্স করা যায়।

```go
package main
import ("fmt")

func main() {
  numbers := []int{1, 2, 3, 4, 5}
  for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
    numbers[i], numbers[j] = numbers[j], numbers[i]
  }
  fmt.Println("Reversed Slice:", numbers)
}
```
