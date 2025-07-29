[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** e.g., interview-qa/switch-statement
**Tags:** [go, switch-statement]
]

# The switch Statement

## switch স্টেটমেন্ট আমাদের অনেক বিকল্পের মধ্যে একটি কোড ব্লক চালানোর অনুমতি দেয়।

```go
package main
import ("fmt")

func main() {
  day := 8

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
    fmt.Println("Not a weekday")
  }
}
```

## Frequently Asked Questions

### ১. Go তে `switch` স্টেটমেন্ট কী?

**উত্তর:** `switch` স্টেটমেন্ট অনেক কোড ব্লকের মধ্যে থেকে একটি বেছে নিয়ে চালাতে ব্যবহৃত হয়। এটি একাধিক `if-else` স্টেটমেন্টের বিকল্প।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
    num := 2
    switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    case 3:
        fmt.Println("Three")
    default:
        fmt.Println("Other number")
    }
}
```

---

### ২. একক `case`-এ কি একাধিক মান ব্যবহার করা যায়?

**উত্তর:** হ্যাঁ, একাধিক মান একই `case` এ গ্রুপ করা যায়।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
    char := 'a'
    switch char {
    case 'a', 'e', 'i', 'o', 'u':
        fmt.Println("Vowel")
    default:
        fmt.Println("Consonant")
    }
}
```

---

### ৩. Go তে `switch`-এ `break` প্রয়োজন কি?

**উত্তর:** না, Go স্বয়ংক্রিয়ভাবে প্রতিটি `case` এর পরে ব্রেক করে, তাই `break` দিতে হয় না।

---

### ৪. `fallthrough` কীভাবে ব্যবহার করবেন?

**উত্তর:** `fallthrough` ব্যবহার করলে পরবর্তী `case` কোডও নির্বাহ হবে, যদিও শর্ত মিলবে না।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
    num := 1
    switch num {
    case 1:
        fmt.Println("One")
        fallthrough
    case 2:
        fmt.Println("Two")
    default:
        fmt.Println("Default")
    }
}
```

---

### ৫. `switch` কি এক্সপ্রেশন ছাড়াও ব্যবহার করা যায়?

**উত্তর:** হ্যাঁ, `switch` এক্সপ্রেশন ছাড়াই ব্যবহার করা যায়, যা একাধিক `if-else` এর মত কাজ করে।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
    num := 10
    switch {
    case num < 0:
        fmt.Println("Negative")
    case num == 0:
        fmt.Println("Zero")
    case num > 0:
        fmt.Println("Positive")
    }
}
```

---

### ৬. `switch`-এ কি string ব্যবহার করা যায়?

**উত্তর:** হ্যাঁ, `switch` স্টেটমেন্টে string ব্যবহার করা যায়।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
    color := "red"
    switch color {
    case "red":
        fmt.Println("Stop")
    case "yellow":
        fmt.Println("Caution")
    case "green":
        fmt.Println("Go")
    default:
        fmt.Println("Unknown color")
    }
}
```

---

### ৭. `switch` কি টাইপ অ্যাসার্শন হ্যান্ডেল করতে পারে?

**উত্তর:** হ্যাঁ, `switch` টাইপ অ্যাসার্শন হ্যান্ডেল করতে পারে।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
    var x interface{} = 42
    switch v := x.(type) {
    case int:
        fmt.Printf("%d is an int\n", v)
    case string:
        fmt.Printf("%s is a string\n", v)
    default:
        fmt.Println("Unknown type")
    }
}
```

---

### ৮. `switch`-এ কি ফাংশনের রিটার্ন ভ্যালু ব্যবহার করা যায়?

**উত্তর:** হ্যাঁ, আপনি ফাংশনের রিটার্ন ভ্যালুকে `switch` এ ব্যবহার করতে পারেন।

**উদাহরণ:**

```go
package main
import "fmt"

func getNumber() int {
    return 3
}

func main() {
    switch getNumber() {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    case 3:
        fmt.Println("Three")
    default:
        fmt.Println("Other number")
    }
}
```

---

### ৯. `switch` কি নেস্টেড (অন্দরভুক্ত) হতে পারে?

**উত্তর:** হ্যাঁ, `switch` স্টেটমেন্ট নেস্টেড হতে পারে।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
    num := 2
    switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
        switch num {
        case 2:
            fmt.Println("Nested switch")
        }
    default:
        fmt.Println("Other number")
    }
}
```

---

### ১০. যদি কোনো `case` ম্যাচ না করে এবং `default` না থাকে তাহলে কী হয়?

**উত্তর:** কোনো `case` ম্যাচ না করলে এবং `default` না থাকলে `switch` কিছুই করে না।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
    num := 5
    switch num {
    case 1:
        fmt.Println("One")
    case 2:
        fmt.Println("Two")
    }
    // কোন আউটপুট হবে না
}
```
