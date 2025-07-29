[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-22
**Category:** interview-qa/Scope
**Tags:** [go, Scope]
]


# স্কোপ (Scope) 

## পরিচিতি

Go (Golang)-এ **scope** বলতে কোডের সেই অংশকে বোঝানো হয় যেখানে একটি ভেরিয়েবল, ফাংশন বা কনস্ট্যান্ট অ্যাক্সেসযোগ্য থাকে। স্কোপ নির্ধারণ করে কোন ভেরিয়েবল বা ফাংশন কোথা থেকে দেখা যাবে এবং কতক্ষণ সক্রিয় থাকবে। এটি বোঝা গুরুত্বপূর্ণ কারণ স্কোপ Go কম্পাইলার কীভাবে সিম্বল রিজলভ করে তা নির্ধারণ করে।

Go-তে বিভিন্ন ধরণের স্কোপ আছে:

1. **Package Scope**: প্যাকেজ লেভেলে সংজ্ঞায়িত ভেরিয়েবল বা ফাংশন
2. **Function Scope**: ফাংশনের মধ্যে সংজ্ঞায়িত ভেরিয়েবল
3. **Block Scope**: লুপ, if-statement ইত্যাদির মধ্যে সংজ্ঞায়িত ভেরিয়েবল
4. **Global Scope**: পুরো প্রোগ্রামে অ্যাক্সেসযোগ্য ভেরিয়েবল বা ফাংশন

এই ডকুমেন্টে আমরা এই স্কোপগুলো উদাহরণসহ আলোচনা করব এবং স্কোপ সম্পর্কিত সাধারণ ভুল ও সেরা চর্চাগুলো তুলে ধরব।

---

## Go-তে Scope-এর প্রকারভেদ

### 1. Package Scope

Go-তে প্যাকেজ স্কোপ বলতে এমন ভেরিয়েবল বা ফাংশন বোঝানো হয় যেগুলো তাদের ডিক্লেয়ার করা প্যাকেজের মধ্যে যেকোনো ফাংশন থেকে অ্যাক্সেস করা যায়।

#### উদাহরণ: Package Scope

```go
package main

import "fmt"

// Global variable (Package Scope)
var globalVar = 10

func printGlobalVar() {
    // Accessing package-scoped variable
    fmt.Println("Global Variable:", globalVar)
}

func main() {
    printGlobalVar() // Output: Global Variable: 10
}
```

উপরের উদাহরণে, `globalVar` প্যাকেজ লেভেলে ডিক্লেয়ার করা হয়েছে, তাই এটি `main` এবং `printGlobalVar` ফাংশনের মধ্যে অ্যাক্সেসযোগ্য।

---

### 2. Function Scope

যেসব ভেরিয়েবল ফাংশনের ভিতরে ডিক্লেয়ার করা হয়, সেগুলো শুধুমাত্র ঐ ফাংশনের মধ্যেই অ্যাক্সেসযোগ্য। একে বলা হয় **function scope**।

#### উদাহরণ: Function Scope

```go
package main

import "fmt"

func myFunction() {
    // Variable inside the function
    var functionVar = "I am local to the function"
    fmt.Println(functionVar) // Output: I am local to the function
}

func main() {
    // Un-commenting the next line would cause an error
    // fmt.Println(functionVar) // Error: undefined functionVar
    myFunction()
}
```

উদাহরণে, `functionVar` শুধুমাত্র `myFunction()` এর মধ্যে ব্যবহারযোগ্য এবং `main()` এর মধ্যে নয়।

---

### 3. Block Scope

Block scope বোঝায় এমন ভেরিয়েবল যেগুলো কোডের একটি নির্দিষ্ট ব্লকের ভিতরে ডিক্লেয়ার করা হয়, যেমন `if` স্টেটমেন্ট, লুপ ইত্যাদি।

#### উদাহরণ: Block Scope

```go
package main

import "fmt"

func main() {
    if true {
        // Variable inside the block
        var blockVar = "I exist only within this block"
        fmt.Println(blockVar) // Output: I exist only within this block
    }
    // Un-commenting the next line would cause an error
    // fmt.Println(blockVar) // Error: undefined blockVar
}
```

এখানে `blockVar` কেবল `if` ব্লকের মধ্যে অ্যাক্সেসযোগ্য। বাইরে ব্যবহার করলে কম্পাইল-টাইম এরর হবে।

---

### 4. Global Scope (Cross-package Scope)

Go-তে, গ্লোবাল ভেরিয়েবল (প্যাকেজ-লেভেল ভেরিয়েবল) একই প্যাকেজের একাধিক ফাইলে শেয়ার করা যায়। তবে, অন্য প্যাকেজ থেকে অ্যাক্সেস করতে চাইলে তা **exported** হতে হবে।

#### উদাহরণ: Cross-file Scope

```go
// global/shared.go
package global

import "fmt"

var SharedVar = "This is a shared variable"

func DisplaySharedVar() {
    fmt.Println(SharedVar)
}

// main.go
package main

import (
    "example.com/project/global" // adjust based on your module path
)

func main() {
    global.DisplaySharedVar() // Output: This is a shared variable
}

```

এই ক্ষেত্রে, `SharedVar` একই প্যাকেজের একাধিক ফাইলে অ্যাক্সেসযোগ্য।

---

## Scope-সম্পর্কিত নিয়মাবলি

1. **Variable Shadowing**: যখন একটি লোকাল ভেরিয়েবলের নাম বাইরের স্কোপের ভেরিয়েবলের সাথে একই হয়, তখন লোকাল ভেরিয়েবলটি বাইরের ভেরিয়েবলকে "ছায়া" দেয়।
2. **Exported Variables**: বড় হাতের অক্ষরে শুরু হওয়া ভেরিয়েবল অন্য প্যাকেজ থেকেও অ্যাক্সেসযোগ্য হয়। ছোট হাতের অক্ষরে শুরু হলে তা শুধু নিজের প্যাকেজেই সীমাবদ্ধ।
3. **Short Declaration**: `:=` অপারেটর ব্যবহার করে ফাংশন বা ব্লকের মধ্যে সংক্ষিপ্তভাবে ভেরিয়েবল ডিক্লেয়ার ও ইনিশিয়ালাইজ করা যায়।

---

## Scope-এ সাধারণ ভুল

1. **Variable Shadowing**: একই নামের ভেরিয়েবল ভেতরের স্কোপে আবার ডিক্লেয়ার করলে বাইরের ভেরিয়েবল ঢেকে যায়, যা অনিচ্ছাকৃত আচরণ তৈরি করতে পারে।

   ```go
   package main

   import "fmt"

   var x = 10

   func main() {
       var x = 20 // Shadows outer x
       fmt.Println(x) // Output: 20
   }

   fmt.Println(x) // Output: 10
   ```

2. **Unexported Variables অ্যাক্সেস করার চেষ্টা**: ছোট হাতের অক্ষরে ডিক্লেয়ার করা ভেরিয়েবল অন্য প্যাকেজ থেকে অ্যাক্সেস করা যায় না।

   ```go
   package main

   import "fmt"

   var unexportedVar = "I am private"

   func main() {
       fmt.Println(unexportedVar) // Accessible within the same package
   }
   ```

---

## 5টি ইন্টারভিউ প্রশ্ন কোডসহ

### Q1: `for` লুপে ডিক্লেয়ার করা ভেরিয়েবলের স্কোপ কী?

#### কোড:

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        var loopVar = "Loop variable"
        fmt.Println(loopVar)
    }
    // Uncommenting the line below will cause an error because loopVar is out of scope
    // fmt.Println(loopVar) // Error: undefined loopVar
}
```

**উত্তর**: `loopVar` কেবল `for` লুপের ভিতরেই অ্যাক্সেসযোগ্য।

---

### Q2: ভিন্ন স্কোপে একই নামের ভেরিয়েবল ডিক্লেয়ার করলে কী হয়?

#### কোড:

```go
package main

import "fmt"

var x = 10

func main() {
    fmt.Println(x) // Output: 10
    var x = 20 // Shadows the outer x
    fmt.Println(x) // Output: 20
}
```

**উত্তর**: `main()` এর ভিতরের `x` বাইরের `x`-কে ছায়া দেয় (shadow করে) এবং শুধুমাত্র main-এর মধ্যে কার্যকর থাকে।

---

### Q3: `:=` অপারেটর ভেরিয়েবল স্কোপে কীভাবে কাজ করে?

#### কোড:

```go
package main

import "fmt"

var a = 5

func main() {
    fmt.Println(a) // Output: 5
    a := 10 // Short declaration: Creates a new a variable within main's scope
    fmt.Println(a) // Output: 10
}
```

**উত্তর**: `:=` নতুন ভেরিয়েবল তৈরি করে, যা বাইরের `a`-কে shadow করে।

---

### Q4: Go কীভাবে গ্লোবাল স্কোপ ভেরিয়েবল হ্যান্ডল করে একাধিক ফাইলে?

#### কোড:

```go
// global/shared.go
package global

import "fmt"

var SharedVar = "This is a shared variable"

func DisplaySharedVar() {
    fmt.Println(SharedVar)
}

// main.go
package main

import (
    "example.com/project/global" // adjust based on your module path
)

func main() {
    global.DisplaySharedVar() // Output: This is a shared variable
}
```

**উত্তর**: একই প্যাকেজের মধ্যে `SharedVar` একাধিক ফাইলে অ্যাক্সেসযোগ্য।

---

### Q5: Go-তে ভেরিয়েবল নামকরণের নিয়মের গুরুত্ব কী?

#### কোড:

```go
package main

import "fmt"

// Exported variable (visible outside the package)
var ExportedVar = "I am exported"

// Unexported variable (private to the package)
var unexportedVar = "I am unexported"

func main() {
    fmt.Println(ExportedVar) // Accessible
    fmt.Println(unexportedVar) // Accessible within the same package
}
```

**উত্তর**: বড় হাতের অক্ষরে শুরু হওয়া ভেরিয়েবল এক্সপোর্টেড এবং অন্য প্যাকেজ থেকে অ্যাক্সেসযোগ্য, ছোট হাতের হলে তা কেবল নিজের প্যাকেজেই সীমাবদ্ধ।

