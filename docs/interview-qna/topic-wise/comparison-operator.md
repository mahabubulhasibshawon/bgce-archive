[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** e.g., interview-qa/topic_name
**Tags:** [go, concurrency, channels]
]

---

# তুলনামূলক অপারেটর (Comparison Operators)

## তুলনামূলক অপারেটর দুটি মানকে একে অপরের সাথে তুলনা করার জন্য ব্যবহার করা হয়

| অপারেটর | নাম          | উদাহরণ |
| ------- | ------------ | ------ |
| ==      | সমান         | x == y |
| !=      | সমান নয়     | x != y |
| >       | এর চেয়ে বড় | x > y  |
| <       | এর চেয়ে ছোট | x < y  |
| >=      | সমান বা বড়  | x >= y |
| <=      | সমান বা ছোট  | x <= y |

```go
package main

import "fmt"

func main() {
	fmt.Println(2 > 2) // false
	fmt.Println(2 < 2) // false
	fmt.Println(2 >= 2) // true
	fmt.Println(2 <= 2) // true
	fmt.Println(2 == 2) // true
	fmt.Println(2 != 2) // false
}
```

## Frequently Asked Questions

### 1. Go-তে তুলনামূলক অপারেটরের উদ্দেশ্য কী?

**উত্তর:** তুলনামূলক অপারেটর দুটি মানকে তুলনা করে এবং একটি boolean ফলাফল (true বা false) রিটার্ন করে।

**উদাহরণ:**

```go
package main

import "fmt"

func main() {
	x := 10
	y := 20
	fmt.Println(x > y) // false
	fmt.Println(x < y) // true
}
```

### 2. Go-তে কি স্ট্রিং এর সাথে তুলনামূলক অপারেটর ব্যবহার করা যায়?

**উত্তর:** হ্যাঁ, স্ট্রিং এর অক্ষরক্রম (lexicographical order) অনুযায়ী তুলনা করতে এই অপারেটর ব্যবহার করা যায়।

**উদাহরণ:**

```go
package main

import "fmt"

func main() {
	fmt.Println("apple" > "banana") // false
	fmt.Println("apple" < "banana") // true
}
```

### 3. Go-তে struct এর সাথে `==` অপারেটর কীভাবে কাজ করে?

**উত্তর:** যদি struct-এর সব ফিল্ড তুলনাযোগ্য হয়, তাহলে `==` অপারেটর দিয়ে struct তুলনা করা যায়।

**উদাহরণ:**

```go
package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	fmt.Println(p1 == p2) // true
}
```

### 4. যদি দুইটি আলাদা টাইপের মান তুলনা করা হয়, তাহলে কী হয়?

**উত্তর:** ভিন্ন টাইপের মান তুলনা করলে compile-time error হয়।

**উদাহরণ:**

```go
package main

func main() {
	// নিচের লাইনটি আনকমেন্ট করলে compile-time error হবে
	// fmt.Println(10 == "10")
}
```

### 5. Go-তে কি pointer এর সাথে তুলনা করা যায়?

**উত্তর:** হ্যাঁ, pointer গুলোকে `==` এবং `!=` অপারেটর দিয়ে তুলনা করা যায়।

**উদাহরণ:**

```go
package main

import "fmt"

func main() {
	a := 10
	b := 10
	pa := &a
	pb := &b
	fmt.Println(pa == pb) // false
}
```

### 6. `!=` অপারেটর কীভাবে কাজ করে?

**উত্তর:** `!=` অপারেটর যাচাই করে যে দুইটি মান সমান নয় কিনা।

**উদাহরণ:**

```go
package main

import "fmt"

func main() {
	x := 5
	y := 10
	fmt.Println(x != y) // true
}
```

### 7. Go-তে কি array গুলো তুলনা করা যায়?

**উত্তর:** হ্যাঁ, যদি array-এর সব উপাদান তুলনাযোগ্য হয় তাহলে `==` এবং `!=` দিয়ে array তুলনা করা যায়।

**উদাহরণ:**

```go
package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	a2 := [3]int{1, 2, 3}
	fmt.Println(a1 == a2) // true
}
```

### 8. দুটি slice তুলনা করলে কী হয়?

**উত্তর:** slice গুলোকে সরাসরি `==` দিয়ে তুলনা করা যায় না, শুধুমাত্র `nil` এর সাথে তুলনা করা যায়।

**উদাহরণ:**

```go
package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println(s1 == nil) // false
	// নিচের লাইনটি আনকমেন্ট করলে compile-time error হবে
	// fmt.Println(s1 == s2)
}
```

### 9. Go-তে floating-point সংখ্যা কীভাবে তুলনা করা হয়?

**উত্তর:** floating-point সংখ্যা তুলনা করা যায়, কিন্তু নির্ভুলতা সংক্রান্ত সমস্যার কারণে সতর্ক থাকতে হয়।

**উদাহরণ:**

```go
package main

import "fmt"

func main() {
	x := 0.1 + 0.2
	y := 0.3
	fmt.Println(x == y) // false নির্ভুলতার কারণে
}
```

### 10. কাস্টম টাইপের মধ্যে তুলনা করা যায় কি?

**উত্তর:** কাস্টম টাইপের ভিত্তি যদি তুলনাযোগ্য হয় তাহলে সেগুলো তুলনা করা যায়।

**উদাহরণ:**

```go
package main

import "fmt"

type Age int

func main() {
	var a1 Age = 30
	var a2 Age = 25
	fmt.Println(a1 > a2) // true
}
```

---
