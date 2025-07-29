[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-22
**Category:** interview-qa/Variadic
**Tags:** [go, Variadic Function]
]


# Go তে Variadic Functions

**Variadic function** হলো এমন একটি ফাংশন যা যেকোনো সংখ্যক আর্গুমেন্ট গ্রহণ করতে পারে। ফাংশনের সিগনেচারে টাইপের আগে `...` বসিয়ে Variadic function ডিফাইন করা হয়, যা একই ধরনের একাধিক আর্গুমেন্ট গ্রহণের সুযোগ দেয়।

---

### উদাহরণ: Variadic Function

```go
package main

func print(numbers ...int) { // Variadic ফাংশন যা সংখ্যাগুলো প্রিন্ট করে
	for _, n := range numbers {
		println(n) // প্রতিটি সংখ্যা প্রিন্ট করা হচ্ছে
	}
}

func main() {
	print(1, 2, 3, 4, 5) // Variadic ফাংশনে একাধিক আর্গুমেন্ট পাঠানো
	print(1, 2)           // কম আর্গুমেন্টসহ কল
	print(1)              // একক আর্গুমেন্ট
	print()               // কোন আর্গুমেন্ট না পাঠিয়ে কল
}
```

---

### ব্যাখ্যা:

* `print(numbers ...int)` একটি Variadic ফাংশন, যা যেকোনো সংখ্যক `int` মান নিতে পারে।
* `main()` এ বিভিন্ন আর্গুমেন্ট সেট দিয়ে `print()` ফাংশন কল করা হয়েছে:

  * ৫টি সংখ্যা
  * ২টি সংখ্যা
  * ১টি সংখ্যা
  * কোন সংখ্যা পাঠানো হয়নি

---

### Variadic Functions সম্পর্কে আরো:

* Variadic ফাংশনগুলো আসলে সব আর্গুমেন্টকে একটি slice হিসেবে ধরে।
* slice-কে সরাসরি `...` ব্যবহার করে unpack করে Variadic ফাংশনে পাঠানো যায়।

---

### উদাহরণ: Slice পাঠানো Variadic Function এ

```go
package main

func print(numbers ...int) {
	for _, n := range numbers {
		println(n)
	}
}

func main() {
	nums := []int{10, 20, 30}
	print(nums...) // slice unpack করে পাঠানো
}
```

---

## Interview প্রশ্ন ও উত্তর

### ১. Go তে Variadic Function ব্যবহারের সুবিধা কী?

**উত্তর:**
Variadic function ফাংশন কলকে ফ্লেক্সিবল করে তোলে। একাধিক ভিন্ন সংখ্যক আর্গুমেন্ট সহজে পাঠানো যায়, আলাদা আলাদা ফাংশন বানানোর প্রয়োজন পড়ে না।

---

### ২. Variadic function কি একাধিক টাইপের আর্গুমেন্ট নিতে পারে?

**উত্তর:**
না, Go তে Variadic function একটাই টাইপ গ্রহণ করে। তবে `interface{}` ব্যবহার করলে বিভিন্ন ধরনের ডেটা নেয়া যায়, কিন্তু টাইপ সেফটি হারায়।

**উদাহরণ:**

```go
package main

import "fmt"

func printAnything(values ...interface{}) {
	for _, v := range values {
		fmt.Println(v)
	}
}

func main() {
	printAnything(1, "Hello", 3.14, true) // একাধিক টাইপ গ্রহণ
}
```

---

### ৩. একটি ফাংশনে একাধিক Variadic প্যারামিটার রাখা যাবে?

**উত্তর:**
না, Go তে একাধিক Variadic প্যারামিটার রাখা যায় না। শুধুমাত্র একটি Variadic প্যারামিটার থাকতে পারে।

---

### ৪. যদি Variadic ফাংশনে কোন আর্গুমেন্ট না দেওয়া হয়, তখন কী হয়?

**উত্তর:**
প্যারামিটার হিসেবে ফাংশনের ভিতরে একটি খালি slice তৈরি হয়। তাই ফাংশন তখনো কাজ করবে।

**উদাহরণ:**

```go
package main

func print(numbers ...int) {
	if len(numbers) == 0 {
		println("No numbers provided")
		return
	}
	for _, n := range numbers {
		println(n)
	}
}

func main() {
	print() // কোন আর্গুমেন্ট নেই
}
```

---

### ৫. Slice কি Variadic ফাংশনে পাঠানো যাবে?

**উত্তর:**
হ্যাঁ, slice কে `...` ব্যবহার করে unpack করে পাঠানো যায়।

**উদাহরণ:**

```go
package main

func print(numbers ...int) {
	for _, n := range numbers {
		println(n)
	}
}

func main() {
	nums := []int{10, 20, 30}
	print(nums...) // slice unpack করে পাঠানো
}
```


