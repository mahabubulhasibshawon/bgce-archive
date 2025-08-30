[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** interview-qa/constants
**Tags:** [go, constants, beginner]
]


### 📌 **Constant কখন ব্যবহৃত হয়?**

* যখন আপনার এমন একটি মান প্রয়োজন যা প্রোগ্রাম চলাকালীন সময়ে কখনোই পরিবর্তন হবে না, তখন constant ব্যবহার করা উচিত।
* এটি সাধারণত magic numbers বা fixed messages/states/states নির্ধারণ করার জন্য ব্যবহৃত হয়।

---

### ⚙️ **Go build করার সময় Constant কোথায় যায়?**

Go প্রোগ্রাম build করার সময়, compiler `const` দিয়ে declared মানগুলোকে **code segment** বা **text segment**-এ রাখে।

* `const` variable গুলো **compile-time**-এ evaluate হয়, অর্থাৎ run-time এ গিয়ে তা আর হিসাব করে না।
* এদের জন্য **memory allocation হয় না heap বা stack-এ**, বরং এরা **program binary-র read-only অংশে** store থাকে।
* এই কারণে constant গুলো **immutable (অপরিবর্তনীয়)** হয় এবং runtime-এ পরিবর্তন করা যায় না।

### 🧠 উদাহরণস্বরূপ:

```go
const PI = 3.14
```

* এই মানটি runtime-এ কোনো memory assign করে রাখবে না।
* বরং যেখানেই `PI` ব্যবহার করা হবে, সেখানে সরাসরি `3.14` inject করে দেবে Go compiler।

---

### 🚀 **Performance এর দিক থেকে Constant এর উপকারিতা:**

1. **No Memory Overhead:**
   Constant মানের জন্য আলাদা memory allocation হয় না (heap/stack এ না), ফলে এটি memory-efficient।

2. **Fast Execution:**
   যেহেতু constant-এর মান compile time-এ জানা থাকে, Go compiler এগুলোকে সরাসরি machine instruction এ transform করে — এতে execution দ্রুত হয়।

3. **Safe & Predictable:**
   Constant মান runtime-এ হঠাৎ করে পরিবর্তন হয়ে যাবার ভয় থাকে না, যা bug প্রতিরোধে সাহায্য করে।

---

### 📍 Stack, Heap, এবং Code Segment – Constant কোথায়?

| Variable Type | Stored In    | Mutable? | Allocated at |
| ------------- | ------------ | -------- | ------------ |
| `var`         | Stack / Heap | ✅ হ্যাঁ  | Runtime      |
| `const`       | Code Segment | ❌ না     | Compile Time |

---

### 🎯 যখন constant ব্যবহার করা উচিত:

* যখন আপনি এমন কিছু ডেটা ব্যবহার করছেন যা **একই থাকে** পুরো প্রোগ্রাম জুড়ে (যেমন `PI`, `MAX_RETRIES`, `VERSION` ইত্যাদি)
* যখন আপনি এমন মান ব্যবহার করছেন যা **পরিবর্তন করার প্রয়োজন নেই বা করা উচিত নয়**
* যখন আপনি **better performance ও memory efficiency** চান

---

### 🛑 ভুল ধারণা:

অনেকে মনে করে constant মান মানে শুধু number বা string। কিন্তু Go তে constant হিসেবে `bool`, `rune`, এবং untyped value-ও ব্যবহার করা যায় — যদিও complex type যেমন slice, map, array constant হিসেবে ব্যবহার করা যায় না।


---

# Constant Rules (নিয়মাবলী)

### 1. Constant-এর নাম সাধারণ variable এর মতোই rule follow করে

### 2. Constant-এর নাম সাধারণত uppercase অক্ষরে লেখা হয়

### 3. Constants ফাংশনের ভিতর এবং বাহির — দুই জায়গাতেই ঘোষণা করা যায়

---

## Frequently Asked Questions

### 1. **Go-তে constant বলতে কী বোঝায়?**

**উত্তর:**
একটি constant হলো এমন একটি variable যার মান একবার সেট করার পর পরিবর্তন করা যায় না। Constant ঘোষণা করতে `const` কিওয়ার্ড ব্যবহার করা হয়।

**Code Example:**

```go
package main
import "fmt"

const PI = 3.14

func main() {
    fmt.Println("The value of PI is:", PI)
}
```

---

### 2. **Constant কি ফাংশনের ভিতর ঘোষণা করা যায়?**

**উত্তর:**
হ্যাঁ, constant ফাংশনের ভিতরেও এবং বাইরেও ঘোষণা করা যায়।

**Code Example:**

```go
package main
import "fmt"

func main() {
    const GREETING = "Hello, World!"
    fmt.Println(GREETING)
}
```

---

### 3. **Constant কি শুধু সংখ্যা রাখে, না কি অন্য টাইপও রাখা যায়?**

**উত্তর:**
Constant শুধু সংখ্যা না, string, boolean, এমনকি character টাইপের মানও রাখতে পারে।

**Code Example:**

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

### 4. **Constant কি runtime-এ হিসাব করে মান রাখতে পারে?**

**উত্তর:**
না, constant-কে এমন মান দিতে হবে যা compile time-এ নির্ধারণ করা যায়।

**Code Example:**

```go
package main
import "fmt"

const VALUE = 10 * 2 // Valid

func main() {
    fmt.Println("Value:", VALUE)
}
```

---

### 5. **যদি constant-এর মান পরিবর্তন করার চেষ্টা করি, তাহলে কী হয়?**

**উত্তর:**
যদি আপনি constant-এর মান পরিবর্তনের চেষ্টা করেন, তাহলে compile time-এ error হবে।

**Code Example:**

```go
package main
import "fmt"

const NAME = "John"

func main() {
    // NAME = "Doe" // Uncommenting this line will cause a compilation error
    fmt.Println(NAME)
}
```

---

### 6. **Constant কি expression-এ ব্যবহার করা যায়?**

**উত্তর:**
হ্যাঁ, constant ব্যবহার করে অন্য constant-এর মান হিসাব করা যায়।

**Code Example:**

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

### 7. **`const` এবং `var` এর মধ্যে পার্থক্য কী?**

**উত্তর:**
`const` ব্যবহার করা হয় এমন মানের জন্য যেটা কখনোই পরিবর্তন হবে না, আর `var` ব্যবহার করা হয় পরিবর্তনশীল মান রাখার জন্য।

**Code Example:**

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

### 8. **Constant কি array বা slice টাইপের হতে পারে?**

**উত্তর:**
না, constant কখনো array, slice বা map টাইপের হতে পারে না।

**Code Example:**

```go
package main
import "fmt"

func main() {
    // const ARR = [3]int{1, 2, 3} // This will cause a compilation error
    fmt.Println("Constants cannot be arrays or slices.")
}
```

---

### 9. **Go-তে constant কি export করা যায়?**

**উত্তর:**
হ্যাঁ, যদি constant-এর নাম uppercase letter দিয়ে শুরু হয়, তাহলে তা exported হয় এবং অন্য প্যাকেজ থেকেও access করা যায়।

**Code Example:**

```go
package main
import "fmt"

const ExportedConstant = "I am exported!"

func main() {
    fmt.Println(ExportedConstant)
}
```

---

### 10. **Go-তে untyped constant কী?**

**উত্তর:**
Untyped constant এমন constant যেটার কোনো নির্দিষ্ট টাইপ নেই যতক্ষণ না তা কোনো variable এ assign করা হয়।

**Code Example:**

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
