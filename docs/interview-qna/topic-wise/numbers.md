[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** interview-qa/operators
**Tags:** [go, operators, arithmetic, assignment]
]

# Operators

**Operators** ব্যবহার করে ভেরিয়েবল ও মানের উপর গণনা বা কাজ করা হয়।

---

## Arithmetic Operators (গাণিতিক অপারেটরস)

### 1. Addition (`+`): দুইটি মান যোগ করে। যেমন: `x + y`

### 2. Subtraction (`-`): একটি মান থেকে অন্যটি বিয়োগ করে। যেমন: `x - y`

### 3. Multiplication (`*`): দুইটি মান গুণ করে। যেমন: `x * y`

### 4. Division (`/`): প্রথম মানটিকে দ্বিতীয় মান দিয়ে ভাগ করে। যেমন: `x / y`

### 5. Modulus (`%`): ভাগফলের অবশিষ্ট মান রিটার্ন করে।

### 6. Increment (`++`): ভেরিয়েবলের মান ১ করে বাড়ায়।

### 7. Decrement (`--`): ভেরিয়েবলের মান ১ করে কমায়।

```go
package main
import "fmt"

func main() {
	fmt.Println(2 + 2) // 4
	fmt.Println(2 - 2) // 0
	fmt.Println(2 * 2) // 4
	fmt.Println(2 / 2) // 1
	fmt.Println(2 % 2) // 0
}
```

### Increment (বৃদ্ধি করা)

```go
package main
import ("fmt")

func main() {
  x := 10
  x++ // ১ যোগ করা হলো
  fmt.Println(x)
}
```

### Decrement (হ্রাস করা)

```go
package main
import ("fmt")

func main() {
  x := 10
  x-- // ১ কমানো হলো
  fmt.Println(x)
}
```

---

## Assignment Operators (অ্যাসাইনমেন্ট অপারেটরস)

ভেরিয়েবলে মান অ্যাসাইন বা সেট করতে ব্যবহার হয়।

| Operator | Example | অর্থ একই   |
| -------- | ------- | ---------- |
| =        | x = 5   | x কে ৫ দাও |
| +=       | x += 3  | x = x + 3  |
| -=       | x -= 3  | x = x - 3  |
| \*=      | x \*= 3 | x = x \* 3 |
| /=       | x /= 3  | x = x / 3  |
| %=       | x %= 3  | x = x % 3  |

---

## প্রশ্নোত্তর ও উদাহরণ

### 1. `lgNumber` নামের একটি ভেরিয়েবল তৈরি করে 1000 সেট করে বিভিন্ন অপারেশন দেখাও।

```go
package main
import ("fmt")

func main() {
  lgNumber := 1000

  // Add that variable with itself
  sum := lgNumber + lgNumber
  fmt.Println("Addition:", sum) // 2000

  // Subtract that variable by itself
  difference := lgNumber - lgNumber
  fmt.Println("Subtraction:", difference) // 0

  // Multiply that variable with itself
  product := lgNumber * lgNumber
  fmt.Println("Multiplication:", product) // 1000000

  // Divide that variable with itself
  quotient := lgNumber / lgNumber
  fmt.Println("Division:", quotient) // 1
}
```

---

## Frequently Asked Questions 

---

### 1. Go-তে কীভাবে যোগ ও বিয়োগ করবো?

**উত্তর:** `+` যোগ করার জন্য, `-` বিয়োগ করার জন্য।

```go
package main
import "fmt"

func main() {
    a, b := 10, 5
    fmt.Println("Addition:", a + b)      // 15
    fmt.Println("Subtraction:", a - b)   // 5
}
```

---

### 2. Go-তে ভাগফলের বাকি (remainder) কিভাবে পাওয়া যায়?

**উত্তর:** `%` অপারেটর ব্যবহার করে।

```go
package main
import "fmt"

func main() {
    fmt.Println("Remainder:", 10 % 3) // 1
}
```

---

### 3. Go-তে কীভাবে ভেরিয়েবল ইনক্রিমেন্ট করব?

**উত্তর:** `++` ব্যবহার করে মান ১ করে বাড়ানো যায়।

```go
package main
import "fmt"

func main() {
    x := 5
    x++
    fmt.Println("Incremented Value:", x) // 6
}
```

---

### 4. কোনো ভেরিয়েবলে মান যোগ করে কীভাবে সেট করব?

**উত্তর:** `+=` ব্যবহার করা যায়।

```go
package main
import "fmt"

func main() {
    x := 10
    x += 5
    fmt.Println("Updated Value:", x) // 15
}
```

---

### 5. গুণ করে কীভাবে মান সেট করব?

**উত্তর:** `*=` ব্যবহার করে।

```go
package main
import "fmt"

func main() {
    x := 10
    x *= 2
    fmt.Println("Updated Value:", x) // 20
}
```

---

### 6. ভাগ করে মান অ্যাসাইন করবো কীভাবে?

**উত্তর:** `/=` ব্যবহার করতে হবে।

```go
package main
import "fmt"

func main() {
    x := 10
    x /= 2
    fmt.Println("Updated Value:", x) // 5
}
```

---

### 7. কীভাবে ডিক্রিমেন্ট করব?

**উত্তর:** `--` দিয়ে ১ করে কমানো যায়।

```go
package main
import "fmt"

func main() {
    x := 5
    x--
    fmt.Println("Decremented Value:", x) // 4
}
```

---

### 8. এক লাইনে একাধিক গাণিতিক অপারেশন কীভাবে করব?

**উত্তর:** এক্সপ্রেশন ব্যবহার করেই করা যায়।

```go
package main
import "fmt"

func main() {
    result := (10 + 5) * 2 - 3
    fmt.Println("Result:", result) // 27
}
```

---

### 9. কোনো সংখ্যা জোড় না বিজোড় তা কীভাবে বের করব?

**উত্তর:** `% 2 == 0` হলে জোড়, না হলে বিজোড়।

```go
package main
import "fmt"

func main() {
    num := 7
    if num % 2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}
```

---

### 10. তৃতীয় কোনো ভেরিয়েবল ছাড়াই দুটি সংখ্যা কীভাবে swap করব?

**উত্তর:** গাণিতিক অপারেশন ব্যবহার করে।

```go
package main
import "fmt"

func main() {
    a, b := 5, 10
    a = a + b
    b = a - b
    a = a - b
    fmt.Println("Swapped Values:", a, b) // 10, 5
}
```
