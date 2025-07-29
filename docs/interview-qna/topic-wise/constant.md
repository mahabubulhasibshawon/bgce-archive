[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** interview-qa/constants
**Tags:** [go, constants, beginner]
]

# `const` কীওয়ার্ড একটি ভেরিয়েবলকে "ধ্রুবক" হিসেবে ঘোষণা করে, অর্থাৎ এটি অপরিবর্তনীয় এবং শুধুমাত্র পড়া যায়।

```go
package main
import ("fmt")

const user = "admin" // এটি পরিবর্তন করা যাবে না

func main() {
  fmt.Println("admin")
}
```

# ধ্রুবক (Constant) সম্পর্কিত নিয়মাবলী

### 1. `const` এর নামকরণ ভেরিয়েবলের মতো একই নিয়ম অনুসরণ করে

### 2. সাধারণত `const` এর নাম বড় অক্ষরে লেখা হয়

### 3. `const` ফাংশনের ভেতরেও এবং ফাংশনের বাইরেও ঘোষণা করা যায়

---

# Frequently Asked Questions

### 1. **Go-তে ধ্রুবক (constant) কী?**

**উত্তর:** `const` হচ্ছে একটি ভেরিয়েবল যার মান একবার নির্ধারণ করলে আর পরিবর্তন করা যায় না। `const` কীওয়ার্ড ব্যবহার করে ধ্রুবক ঘোষণা করা হয়।

**কোড উদাহরণ:**

```go
package main
import "fmt"

const PI = 3.14

func main() {
    fmt.Println("The value of PI is:", PI)
}
```

---

### 2. **const কি ফাংশনের ভিতরে ঘোষণা করা যায়?**

**উত্তর:** হ্যাঁ, const ফাংশনের ভিতরেও এবং ফাংশনের বাইরেও ঘোষণা করা যায়।

**কোড উদাহরণ:**

```go
package main
import "fmt"

func main() {
    const GREETING = "Hello, World!"
    fmt.Println(GREETING)
}
```

---

### 3. **const কি শুধুমাত্র সংখ্যার মান ধারণ করতে পারে?**

**উত্তর:** না, `const` string, boolean বা character এর মতো মানও ধারণ করতে পারে।

**কোড উদাহরণ:**

```go
package main
import "fmt"

const IS_ACTIVE = true
const MESSAGE = "Welcome to Go!"

func main() {
    fmt.Println("Is Active:", IS_ACTIVE)
    fmt.Println("Message:", MESSAGE)
}
```

---

### 4. **const এর মান কি রানটাইমে গণনা করা যায়?**

**উত্তর:** না, const এর মান কম্পাইল টাইমেই নির্ধারণযোগ্য হতে হবে।

**কোড উদাহরণ:**

```go
package main
import "fmt"

const VALUE = 10 * 2 // Valid

func main() {
    fmt.Println("Value:", VALUE)
}
```

---

### 5. **যদি আপনি একটি const এর মান পরিবর্তন করতে চান, তাহলে কী হবে?**

**উত্তর:** কম্পাইলার একটি ত্রুটি (error) দেবে যদি আপনি const এর মান পরিবর্তন করতে চান।

**কোড উদাহরণ:**

```go
package main
import "fmt"

const NAME = "John"

func main() {
    // NAME = "Doe" // এই লাইনটি আনকমেন্ট করলে কম্পাইলেশন ত্রুটি হবে
    fmt.Println(NAME)
}
```

---

### 6. **const কি expression এ ব্যবহার করা যায়?**

**উত্তর:** হ্যাঁ, ধ্রুবক অন্য const এর মান গণনার জন্য expression এ ব্যবহার করা যায়।

**কোড উদাহরণ:**

```go
package main
import "fmt"

const A = 5
const B = 10
const SUM = A + B

func main() {
    fmt.Println("Sum:", SUM)
}
```

---

### 7. **Go-তে `const` এবং `var` এর মধ্যে পার্থক্য কী?**

**উত্তর:** `const` এমন মানের জন্য ব্যবহৃত হয় যা কখনো পরিবর্তন হবে না, আর `var` এমন মানের জন্য ব্যবহৃত হয় যা পরিবর্তনশীল।

**কোড উদাহরণ:**

```go
package main
import "fmt"

const FIXED = 100
var changeable = 200

func main() {
    fmt.Println("Fixed:", FIXED)
    fmt.Println("Changeable:", changeable)
    changeable = 300
    fmt.Println("Updated Changeable:", changeable)
}
```

---

### 8. **const কি array বা slice টাইপ হতে পারে?**

**উত্তর:** না, const array, slice বা map টাইপ হতে পারে না।

**কোড উদাহরণ:**

```go
package main
import "fmt"

func main() {
    // const ARR = [3]int{1, 2, 3} // এটি কম্পাইলেশন ত্রুটি তৈরি করবে
    fmt.Println("Constants cannot be arrays or slices.")
}
```

---

### 9. **Go-তে const কি এক্সপোর্ট (export) করা যায়?**

**উত্তর:** হ্যাঁ, যদি const এর নাম বড় হাতের অক্ষরে শুরু হয়, তাহলে তা এক্সপোর্ট করা যায়।

**কোড উদাহরণ:**

```go
package main
import "fmt"

const ExportedConstant = "I am exported!"

func main() {
    fmt.Println(ExportedConstant)
}
```

---

### 10. **Go-তে untyped constants কী?**

**উত্তর:** Untyped constants এর নির্দিষ্ট কোনো টাইপ থাকে না যতক্ষণ না তা কোনো ভেরিয়েবলে অ্যাসাইন করা হয়।

**কোড উদাহরণ:**

```go
package main
import "fmt"

const VALUE = 42

func main() {
    var x int = VALUE
    var y float64 = VALUE
    fmt.Println("x:", x, "y:", y)
}
```

---
