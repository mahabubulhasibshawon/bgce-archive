[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** interview-qa/conditional-statements
**Tags:** [go, conditional-statements, if-else]
]

---

## Go শর্ত (Conditions)

### শর্তভিত্তিক স্টেটমেন্ট আমাদের প্রোগ্রামের গঠন নিয়ন্ত্রণ করতে সাহায্য করে।

### প্রোগ্রামের প্রবাহ নিয়ন্ত্রণ করার বিভিন্ন উপায় আছে, (if, else if, else) তাদের অন্যতম।

### (if, else if, else) স্টেটমেন্টগুলো আমাদের প্রোগ্রাম চলাকালীন "সিদ্ধান্ত" নিতে সাহায্য করে, প্রোগ্রামিং-এ এগুলোকে (conditional statements) ও বলা হয়।

```go
    // Sudo Syntax
     if condition { <code> }
     else if condition { <code> }
     else { <code> }
```

```go
// Actual Code
package main
import "fmt"

func main() {
	password := "12345678"
	if len(password) > 7 {
		fmt.Println("Valid Password")
	} else {
		fmt.Println("Invalid Password")
	}
}
```

## Frequently Asked Questions (FAQs)

### 1. **Go-তে if-else স্টেটমেন্টের সিনট্যাক্স কী?**

**উত্তর:**
Go-তে if-else স্টেটমেন্টের সিনট্যাক্স নিচের মতো:

```go
if condition {
    // যদি condition সত্য হয় তবে এই কোড চলবে
} else {
    // যদি condition মিথ্যা হয় তবে এই কোড চলবে
}
```

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
	number := 10
	if number%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}
```

---

### 2. **আমরা কি else ছাড়াই if স্টেটমেন্ট ব্যবহার করতে পারি?**

**উত্তর:**
হ্যাঁ, শুধুমাত্র if স্টেটমেন্টও ব্যবহার করা যায়।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
	number := 5
	if number > 0 {
		fmt.Println("Positive Number")
	}
}
```

---

### 3. **Go-তে else-if ladder কী?**

**উত্তর:**
else-if ladder ব্যবহার করা হয় একাধিক condition একে একে যাচাই করতে।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}
}
```

---

### 4. **if স্টেটমেন্টে short statement কীভাবে ব্যবহার করা হয়?**

**উত্তর:**
short statement ব্যবহার করে আমরা একই লাইনে ভ্যারিয়েবল ডিক্লেয়ার ও condition চেক করতে পারি।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
	if num := 10; num%2 == 0 {
		fmt.Println("Even Number")
	}
}
```

---

### 5. **Go-তে nested if-else স্টেটমেন্ট ব্যবহার করা যায় কি?**

**উত্তর:**
হ্যাঁ, একটির ভিতরে আরেকটি if-else ব্যবহার করা যায়।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
	number := 15
	if number > 0 {
		if number%2 == 0 {
			fmt.Println("Positive Even Number")
		} else {
			fmt.Println("Positive Odd Number")
		}
	} else {
		fmt.Println("Non-Positive Number")
	}
}
```

---

### 6. **if-else এবং switch স্টেটমেন্টের মধ্যে পার্থক্য কী?**

**উত্তর:**
if-else ব্যবহার করা হয় শর্তভিত্তিক সিদ্ধান্তের জন্য, আর switch ব্যবহার করা হয় একাধিক ব্লকের মধ্যে একটি নির্দিষ্ট ব্লক নির্বাচন করার জন্য।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid Day")
	}
}
```

---

### 7. **একই if স্টেটমেন্টে একাধিক condition কীভাবে চেক করা যায়?**

**উত্তর:**
`&&` (AND) এবং `||` (OR) লজিক্যাল অপারেটর ব্যবহার করে একাধিক condition চেক করা যায়।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
	number := 15
	if number > 0 && number%3 == 0 {
		fmt.Println("Positive and Divisible by 3")
	}
}
```

---

### 8. **যদি if স্টেটমেন্টের condition boolean না হয়, তাহলে কী হয়?**

**উত্তর:**
Go-তে if condition অবশ্যই একটি boolean হতে হবে। না হলে compile-time error হবে।

**উদাহরণ:**

```go
// এটি compile error দেবে
// if 10 {
//     fmt.Println("Invalid Condition")
// }
```

---

### 9. **if condition-এ কি ফাংশন কল ব্যবহার করা যায়?**

**উত্তর:**
হ্যাঁ, একটি ফাংশন কল if স্টেটমেন্টে ব্যবহার করা যায়।

**উদাহরণ:**

```go
package main
import "fmt"

func isEven(num int) bool {
	return num%2 == 0
}

func main() {
	if isEven(10) {
		fmt.Println("Even Number")
	}
}
```

---

### 10. **User input-এর সাথে if-else কীভাবে ব্যবহার করব?**

**উত্তর:**
`fmt.Scan` ব্যবহার করে ইউজারের ইনপুট নিয়ে if-else স্টেটমেন্টে তা ব্যবহার করা যায়।

**উদাহরণ:**

```go
package main
import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}
}
```

---
