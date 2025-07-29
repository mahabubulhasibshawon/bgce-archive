[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-22
**Category:** interview-qa/Higher-Order
**Tags:** [go, First-Order, Higher-Order]
]

## First‑Order Function এবং Higher‑Order Function

### First‑Order Function

**First‑order function** হলো এমন একটি ফাংশন যা শুধুমাত্র মৌলিক ডেটা টাইপের (integer, string ইত্যাদি) উপর কাজ করে এবং এটি অন্য কোনো ফাংশনকে আর্গুমেন্ট হিসেবে নেয় না এবং কোনো ফাংশন রিটার্ন করেও না।

### Higher‑Order Function

**Higher‑order function** এমন একটি ফাংশন, যা অন্য ফাংশনকে আর্গুমেন্ট হিসেবে নিতে পারে এবং/অথবা একটি ফাংশন রিটার্ন করতে পারে। Higher‑order functions functional programming-এর অন্যতম মূল ভিত্তি।

### উদাহরণ কোড

#### First‑Order Function

```go
package main

import "fmt"

// First‑order function: does not take or return another function
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println("Result:", result) // Output: Result: 8
}
```

#### Higher‑Order Function

```go
package main

import "fmt"

// Higher‑order function: takes a function as an argument
func applyOperation(a int, b int, operation func(int, int) int) int {
    return operation(a, b)
}

// Function to be passed as an argument
func multiply(a int, b int) int {
    return a * b
}

func main() {
    result := applyOperation(5, 3, multiply)
    fmt.Println("Result:", result) // Output: Result: 15
}
```

---

### গণিতে লজিক (Logic in Mathematics)

Discrete mathematics-এ logic ব্যবহার করা হয় বস্তুর (objects) বৈশিষ্ট্য ও সম্পর্ক নির্ধারণ এবং বিশ্লেষণ করার জন্য।

1. **Object**: কোন সত্তা যার বাস্তব অস্তিত্ব আছে (যেমন, মানুষ, প্রাণী)
2. **Property**: বস্তুর বৈশিষ্ট্য (যেমন রঙ, উচ্চতা)
3. **Relation**: বস্তুর মধ্যে সম্পর্ক নির্ধারণ করে (যেমন, “সকল গ্রাহককে তাদের বিল পরিশোধ করতে হবে”)

**উদাহরণ**

* Object: Customer

* Property: Has a bill

* Relation: Must pay the bill

* **First‑Order Logic**: object, property এবং relation এর উপর কাজ করে।

* **Higher‑Order Logic**: ফাংশন এবং operation-এর সম্পর্ক নিয়ে কাজ করে।

ফাংশনের প্রেক্ষাপটে:

* **First‑Order Function**: সরাসরি object এবং তার property নিয়ে কাজ করে।
* **Higher‑Order Function**: ফাংশনের সম্পর্ক নিয়ে কাজ করে, বেশি abstraction এবং flexibility দেয়।

---

### Functional Paradigms

Functional programming একটি programming paradigm যেখানে প্রোগ্রামগুলো ফাংশনের মাধ্যমে গঠিত হয়। এটি গুরুত্ব দেয়:

* **Pure Functions**: একই ইনপুটে সবসময় একই আউটপুট এবং কোন side-effect থাকে না।
* **Immutability**: একটি ডেটা একবার তৈরি হলে পরিবর্তন করা যায় না; পরিবর্তে নতুন ডেটা তৈরি হয়।
* **First‑Class Functions**: ফাংশনগুলোকে variable এ assign করা যায়, function হিসেবে পাঠানো যায়, আর রিটার্ন করতেও পারে।
* **Higher‑Order Functions**: অন্য ফাংশনকে আর্গুমেন্ট হিসেবে নেয় বা ফাংশন রিটার্ন করে।

Functional ভাষা যেমন **Haskell**, **Racket**, **Lisp** এই abstraction-এর জন্য খুব শক্তিশালী।

---

### আরও উদাহরণ কোড

#### Higher‑Order Function যা অন্য ফাংশন রিটার্ন করে

```go
package main

import "fmt"

// Higher‑order function: returns another function
func call() func(int, int) {
    return add
}

func add(a, b int) {
    z := a + b
    fmt.Println(z)
}

func main() {
    // call() একটি higher‑order function যা add function রিটার্ন করে।
    // রিটার্ন হওয়া ফাংশনটি f-এ assign করে পরে কল করা হলো।
    f := call()
    f(10, 20) // Output: 30
}
```

#### First‑Class Functions ও Higher‑Order Function একসাথে

```go
package main

import "fmt"

// Higher‑order function: accepts another function as an argument
func applyAndReturn(fn func(int, int) int, x int, y int) int {
    return fn(x, y)
}

// Function to be passed as an argument
func subtract(a int, b int) int {
    return a - b
}

func main() {
    result := applyAndReturn(subtract, 10, 5)
    fmt.Println("Result:", result) // Output: Result: 5
}
```

---

### Interview Q\&A (Code Examples)

#### 1. **What is a higher‑order function?**

**Question**: What is a higher‑order function, and how does it work in Go?
**Answer**: Higher‑order function হলো এমন একটি ফাংশন যা অন্য ফাংশনকে প্যারামিটার হিসেবে নেয় বা অন্য ফাংশন রিটার্ন করে।

উদাহরণ:

```go
package main

import "fmt"

func applyOperation(a int, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func add(a, b int) int {
    return a + b
}

func main() {
    result := applyOperation(3, 4, add)
    fmt.Println("Result:", result) // Output: Result: 7
}
```

---

#### 2. **What is a first‑order function?**

**Question**: Explain a first‑order function in Go.
**Answer**: First‑order function হলো এমন একটি ফাংশন যা মৌলিক ডেটা টাইপ নিয়ে কাজ করে, অন্য ফাংশন নেয় না বা রিটার্নও করে না।

উদাহরণ:

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 5)
    fmt.Println("Result:", result) // Output: Result: 8
}
```

---

#### 3. **Can you create a function that returns another function?**

**Question**: Write a function that returns another function and demonstrates its usage.
**Answer**: হ্যাঁ, আপনি higher‑order function তৈরি করতে পারেন যা অন্য ফাংশন রিটার্ন করে। নিচে উদাহরণ:

```go
package main

import "fmt"

func multiply(a int) func(int) int {
    return func(b int) int {
        return a * b
    }
}

func main() {
    multiplyBy2 := multiply(2)
    fmt.Println("Result:", multiplyBy2(5)) // Output: Result: 10
}
```

---

#### 4. **What is an anonymous function in Go?**

**Question**: Show an example of an anonymous function in Go.
**Answer**: Anonymous function হলো একটি নামবিহীন ফাংশন, সাধারণত শর্ট-লিভেড কাজের জন্য ব্যবহৃত হয়।

উদাহরণ:

```go
package main

import "fmt"

func main() {
    func(a int, b int) {
        fmt.Println("Sum:", a+b)
    }(3, 4) // Output: Sum: 7
}
```

---

#### 5. **What is an Immediately Invoked Function Expression (IIFE) in Go?**

**Question**: Write a code example for an Immediately Invoked Function Expression (IIFE) in Go.
**Answer**: IIFE হলো এমন একটি ফাংশন যা ডিফাইন করেই সঙ্গে সঙ্গে কল করা হয়।

উদাহরণ:

```go
package main

import "fmt"

func main() {
    result := func(a int, b int) int {
        return a + b
    }(3, 4)

    fmt.Println("Result:", result) // Output: Result: 7
}
```
