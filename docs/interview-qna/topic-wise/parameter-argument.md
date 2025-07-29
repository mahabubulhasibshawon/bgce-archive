[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-22
**Category:** interview-qa/Argument vs Parameter
**Tags:** [go, Parameter, Argument]
]


## Parameter এবং Argument এর মধ্যে পার্থক্য

### Parameter

**Parameter** হলো একটি ভেরিয়েবল in the declaration of a **function**. এটি নির্ধারণ করে কী ধরনের ইনপুট একটি **function** গ্রহণ করবে যখন এটি কল করা হবে। এটি একটি placeholder যা সেই মানগুলোর জন্য ব্যবহৃত হয় যা **function**-এ পাঠানো হবে।

### Argument

**Argument** হলো বাস্তব value যা পাঠানো হয় একটি **function**‑এ যখন এটি কল করা হয়। **Arguments** সেই parameters-এর সাথে মিলে যায় যা **function signature**‑এ ডিফাইন করা থাকে। এই গুলো হলো সেই আসল মান যা **function**‑কে execution সময় প্রদান করা হয়।

### Example Code

```go
package main

import "fmt"

// Function with parameters 'a' and 'b'
func add(a int, b int) int {
    return a + b
}

func main() {
    // Calling the function with arguments 5 and 3
    result := add(5, 3)
    fmt.Println("Result:", result) // Output: Result: 8
}
```

উপরের example‑এ:

* `a` ও `b` হল **parameters** of the `add` **function**
* `5` ও `3` হল **arguments** যা **add** **function**‑এ pass করা হয় যখন এটি `main` **function**‑এ call করা হয়।

---

### Further Explanation

**Parameters** ডিক্লেয়ার করা হয় **function’s signature**‑এর মধ্যে। এগুলো **function**‑কে input values প্রদান করে যা **function’s body**‑এর মধ্যে ব্যবহার করা হয়।

**Arguments** সরবরাহ করা হয় যখন **function** call করা হয়। এই **arguments**‑এর value গুলো সেই соответствующий **parameters**‑এ অ্যাসাইন করা হয়, ফলে **function** সেই মান অনুযায়ী অপারেশন সম্পাদন করতে পারে।

### Types of Parameters

* **Formal Parameters**: এইগুলো হলো parameters যা **function definition**‑এ উপস্থিত থাকে। এগুলো **expected values**‑কে represent করে **inside the function**।
* **Actual Parameters**: এগুলো হলো **arguments** যা **function call** করার সময় pass করা হয়।

---

### Interview Questions and Answers

#### 1.একটি function **arguments** এ pass‑by‑value এবং pass‑by‑reference এর মধ্যে পার্থক্য কি ?

**Answer:**

* **Pass‑by‑value**: **function** একটি copy of the **argument’s** value পায়। parameter‑এ যেকোনো পরিবর্তন আসল **argument**‑কে প্রভাবিত করে না।
* **Pass‑by‑reference**: **function** একটি reference (memory address) পায় **argument**‑এর, তাই parameter‑এ করা পরিবর্তন সরাসরি আসল **argument**‑কে প্রভাবিত করে।

**Code Example**:

```go
package main

import "fmt"

// Pass‑by‑value
func modifyValue(x int) {
    x = 20
}

// Pass‑by‑reference
func modifyReference(x *int) {
    *x = 20
}

func main() {
    a := 10
    modifyValue(a)
    fmt.Println("After modifyValue:", a) // Output: 10 (no change)

    modifyReference(&a)
    fmt.Println("After modifyReference:", a) // Output: 20 (value changed)
}
```

#### 2. আপনি কিভাবে Go তে default **arguments** handle করেন ?

**Answer:**
Go directly default **arguments** support করে না। পরিবর্তে, আপনি variadic **functions** ব্যবহার করতে বা explicit চেক করতে পারেন যদি কোনো value প্রদান করা হয় কি না।

**Code Example**:

```go
package main

import "fmt"

func greet(name string, message string) {
    if message == "" {
        message = "Hello" // default message
    }
    fmt.Println(message, name)
}

func main() {
    greet("John", "")  // Output: Hello John
    greet("Jane", "Hi") // Output: Hi Jane
}
```

#### 3. **variadic function** কি এবং এটা কিভাবে কাজ করে ?

**Answer:**
A **variadic function** হলো একটি **function** যা variable সংখ্যক **arguments** নিতে পারে। এটি `...` ব্যবহার করে **parameter list**‑এ ডিফাইন করা হয়। এর ফলে আপনি যেকোনো সংখ্যক **arguments** পাঠাতে পারেন।

**Code Example**:

```go
package main

import "fmt"

func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))         // Output: 6
    fmt.Println(sum(10, 20, 30, 40)) // Output: 100
}
```

#### 4. যদি একটি **function** এর **parameters** থেকে বেশি **arguments** দিয়ে invoke করা হয় তবে কি হবে ?

**Answer:**
Go‑তে যদি পাঠানো **arguments**‑এর সংখ্যা **parameters**‑এর সংখ্যার থেকে বেশি হয়, তাহলে code‑এ compile‑time error হবে। Go compile‑time‑এ **arguments** এবং **parameters** এর সংখ্যা যাচাই করে, এবং mismatch থাকলে error।

**Code Example**:

```go
package main

import "fmt"

func greet(name string) {
    fmt.Println("Hello, ", name)
}

func main() {
    greet("John", "Doe") // Compile‑time error: too many arguments
}
```

#### 5. Go তে কি আপনি একটি **function** থেকে multiple values return করতে পারবেন ?

**Answer:**
Yes, Go একটি **function** কে multiple values return করতে দেয়। আপনি বিভিন্ন টাইপের values return করতে পারেন।

**Code Example**:

```go
package main

import "fmt"

// Function returning two values
func divide(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    quotient, remainder := divide(10, 3)
    fmt.Println("Quotient:", quotient)   // Output: Quotient: 3
    fmt.Println("Remainder:", remainder) // Output: Remainder: 1
}
```
