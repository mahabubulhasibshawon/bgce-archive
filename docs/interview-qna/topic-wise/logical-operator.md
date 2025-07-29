[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** interview-qa/logical_operators
**Tags:** [go, logical_operators]
]

# Logical Operators (লজিকাল অপারেটরস)

## Logical operators ব্যবহার করা হয় ভেরিয়েবল বা মানের (values) মধ্যে লজিক বোঝার জন্য।

---

## 1. Logical and (`&&`)

### যখন দুটি শর্তই true হয়, তখন `true` রিটার্ন করে।

---

## 2. Logical or (`||`)

### যখন যেকোনো একটি শর্ত true হয়, তখন `true` রিটার্ন করে।

---

## 3. Logical not (`!`)

### ফলাফল উল্টে দেয়। যদি result `true` হয়, তাহলে এটি `false` রিটার্ন করবে।

---

```go
package main

import "fmt"

func main() {
	fmt.Println(true && true)   // true
	fmt.Println(true && false)  // false

	fmt.Println(true || true)   // true
	fmt.Println(false || false) // false

	fmt.Println(!true)  // false
	fmt.Println(!false) // true
}
```

---

## প্রায় জিজ্ঞাসা করা প্রশ্ন (FAQ)

---

### 1. Go-তে `&&` এবং `||` এর মধ্যে পার্থক্য কী?

**উত্তর:**

* `&&` (Logical AND) তখনই `true` রিটার্ন করে যখন দুইটি শর্তই `true` হয়।
* `||` (Logical OR) তখনই `true` রিটার্ন করে যখন অন্তত একটি শর্ত `true` হয়।

**উদাহরণ কোড:**

```go
package main

import "fmt"

func main() {
	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
}
```

---

### 2. Go-তে `!` অপারেটর কীভাবে কাজ করে?

**উত্তর:**

* `!` (Logical NOT) অপারেটর একটি boolean মান উল্টে দেয়।

**উদাহরণ কোড:**

```go
package main

import "fmt"

func main() {
	fmt.Println(!true)  // false
	fmt.Println(!false) // true
}
```

---

### 3. Go-তে logical operators কি non-boolean value-এর সঙ্গে ব্যবহার করা যায়?

**উত্তর:**

* না, Go-তে logical operators শুধুমাত্র boolean মানের সাথেই কাজ করে।

---

### 4. একাধিক logical operators একসাথে কীভাবে ব্যবহার করব?

**উত্তর:**

* Parentheses `()` ব্যবহার করে precedence (অগ্রাধিকার) নিয়ন্ত্রণ করতে পারেন।

**উদাহরণ কোড:**

```go
package main

import "fmt"

func main() {
	fmt.Println((true && false) || true) // true
}
```

---

### 5. Go-তে logical operators-এর precedence কী?

**উত্তর:**

* `!` এর precedence সবচেয়ে বেশি, এরপর `&&`, তারপর `||`।

**উদাহরণ কোড:**

```go
package main

import "fmt"

func main() {
	fmt.Println(!true || false && true) // false
}
```

---

### 6. কীভাবে conditional statement-এ logical operators ব্যবহার করব?

**উত্তর:**

* `if` statement-এর মধ্যে শর্তগুলো মিলিয়ে দেখার জন্য logical operators ব্যবহার করা হয়।

**উদাহরণ কোড:**

```go
package main

import "fmt"

func main() {
	if true && !false {
		fmt.Println("Condition met")
	}
}
```

---

### 7. Go-তে logical operators কি short-circuit করে?

**উত্তর:**

* হ্যাঁ, `&&` এবং `||` উভয়ই short-circuit করে। যেমন `&&` প্রথম শর্ত `false` হলে দ্বিতীয়টা আর দেখে না।

**উদাহরণ কোড:**

```go
package main

import "fmt"

func main() {
	fmt.Println(false && (5 > 3)) // false (5 > 3 evaluate হয় না)
}
```

---

### 8. Logical expression কীভাবে debug করব?

**উত্তর:**

* `fmt.Println` ব্যবহার করে মাঝখানের intermediate ফলাফলগুলো দেখতে পারেন।

**উদাহরণ কোড:**

```go
package main

import "fmt"

func main() {
	condition1 := true
	condition2 := false
	fmt.Println(condition1 && condition2) // false
}
```

---

### 9. কী loop-এর মধ্যে logical operators ব্যবহার করা যায়?

**উত্তর:**

* হ্যাঁ, loop-এর condition-এ logical operators ব্যবহার করা যায়।

**উদাহরণ কোড:**

```go
package main

import "fmt"

func main() {
	for i := 0; i < 10 && i%2 == 0; i++ {
		fmt.Println(i)
	}
}
```

---

### 10. `nil` value-এর সাথে logical operators ব্যবহার করলে কী হয়?

**উত্তর:**

* সরাসরি `nil` এর সাথে logical operators ব্যবহার করা যায় না। আগে `nil` check করতে হয়।

**উদাহরণ কোড:**

```go
package main

import "fmt"

func main() {
	var ptr *int = nil
	fmt.Println(ptr == nil || true) // true
}
```
