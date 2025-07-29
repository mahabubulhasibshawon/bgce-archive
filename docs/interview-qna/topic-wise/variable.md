[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** e.g., interview-qa/topic_name
**Tags:** [go, varibales]
]

# Variables: ডেটা মান সংরক্ষণের জন্য কন্টেইনার

## Go তে ভেরিয়েবল ঘোষণা করার দুইটি পদ্ধতি আছে

### ১. `var` কিওয়ার্ড ব্যবহার করে, তারপর ভেরিয়েবলের নাম এবং টাইপ

```go
package main
import ("fmt")

var name = "John Doe" // ভেরিয়েবল এক

func main() {
  var fruit = "Apple" // ভেরিয়েবল দুই
  fmt.Println(name)
  fmt.Println(fruit)
}
```

### ২. `:=` ব্যবহার করে ভেরিয়েবলের মান দিয়ে ডিক্লেয়ার করা

```go
package main
import ("fmt")

func main() {
  varOne := 100
  varTwo := 2
  fmt.Println(varOne)
  fmt.Println(varTwo)
}
```

## Frequently Asked Questions (FAQs)

### ১. Go তে কিভাবে কনস্ট্যান্ট ভেরিয়েবল ডিক্লেয়ার করবেন?

**উত্তর:** `const` কিওয়ার্ড ব্যবহার করে কনস্ট্যান্ট ঘোষণা করুন। একবার ঘোষণা করলে মান পরিবর্তন করা যায় না।

```go
package main
import ("fmt")

func main() {
  const pi = 3.14
  fmt.Println(pi)
}
```

---

### ২. এক লাইনে একাধিক ভেরিয়েবল ডিক্লেয়ার করা যায়?

**উত্তর:** হ্যাঁ, `var` অথবা `:=` ব্যবহার করে এক লাইনে একাধিক ভেরিয়েবল ডিক্লেয়ার করা যায়।

```go
package main
import ("fmt")

func main() {
  var x, y, z int = 1, 2, 3
  fmt.Println(x, y, z)
}
```

---

### ৩. Go তে ভেরিয়েবলের শূন্য (zero) মান কী?

**উত্তর:** ভেরিয়েবল না দিলে Go অটোমেটিক শূন্য মান (zero value) দেয়।

```go
package main
import ("fmt")

func main() {
  var number int
  fmt.Println(number) // আউটপুট: 0
}
```

---

### ৪. গ্লোবাল ভেরিয়েবল কিভাবে ডিক্লেয়ার করবেন?

**উত্তর:** ফাংশনের বাইরে ভেরিয়েবল ডিক্লেয়ার করুন।

```go
package main
import ("fmt")

var globalVar = "I am global"

func main() {
  fmt.Println(globalVar)
}
```

---

### ৫. `var` দিয়ে ডিক্লেয়ার করা ভেরিয়েবলের মান পরিবর্তন করা যায়?

**উত্তর:** হ্যাঁ, `var` ভেরিয়েবল পুনরায় মান পেতে পারে।

```go
package main
import ("fmt")

func main() {
  var name = "John"
  name = "Doe"
  fmt.Println(name)
}
```

---

### ৬. যদি ভেরিয়েবল ডিক্লেয়ার করি কিন্তু ব্যবহার না করি তাহলে কী হয়?

**উত্তর:** Go অপ্রয়োজনীয় ভেরিয়েবল অনুমোদন করে না, কম্পাইল টাইমে এরর দিবে।

```go
package main
func main() {
  var unusedVar = 10 // এরর হবে কারণ ব্যবহার হয়নি
}
```

---

### ৭. নির্দিষ্ট টাইপ দিয়ে ভেরিয়েবল কিভাবে ডিক্লেয়ার করবেন?

**উত্তর:** `var` কিওয়ার্ডের পরে টাইপ উল্লেখ করুন।

```go
package main
import ("fmt")

func main() {
  var age int = 25
  fmt.Println(age)
}
```

---

### ৮. ভেরিয়েবল ডিক্লেয়ার করে কিন্তু ইনিশিয়ালাইজ না করলে কী হয়?

**উত্তর:** হবে, কিন্তু ভেরিয়েবল শূন্য মান পাবে।

```go
package main
import ("fmt")

func main() {
  var name string
  fmt.Println(name) // আউটপুট: ""
}
```

---

### ৯. ব্লকে ভেরিয়েবল কিভাবে ডিক্লেয়ার করবেন?

**উত্তর:** ব্লকের মধ্যে `var` ব্যবহার করুন।

```go
package main
import ("fmt")

func main() {
  var (
    x int = 10
    y string = "Hello"
  )
  fmt.Println(x, y)
}
```

---

### ১০. `var` এবং `:=` এর মধ্যে পার্থক্য কী?

**উত্তর:**

* `var` গ্লোবাল বা লোকাল উভয় ক্ষেত্রেই ব্যবহার হয় এবং স্পষ্ট টাইপ উল্লেখ করা যায়।
* `:=` শুধুমাত্র লোকাল ভেরিয়েবল জন্য শর্টহ্যান্ড, টাইপ ইনফারেন্স হয়।

```go
package main
import ("fmt")

func main() {
  var name string = "John" // স্পষ্ট টাইপ
  age := 30 // টাইপ ইনফারেন্স
  fmt.Println(name, age)
}
```
