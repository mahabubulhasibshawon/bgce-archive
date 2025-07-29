[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** e.g., interview-qa/topic_name
**Tags:** [go, concurrency, channels]
]

# Structs (Structures)

একটি struct ব্যবহার করে বিভিন্ন ডেটা টাইপের একাধিক সদস্যকে একটি একক ভেরিয়েবলের মধ্যে সংরক্ষণ করা যায়।

```go
package main
import ("fmt")

type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {
  var userOne Person

  userOne.name = "HuXn"
  userOne.age = 18
  userOne.job = "Programmer"
  userOne.salary = 40000

  fmt.Println(userOne)
  fmt.Println("My name is", userOne.name, "I'm", userOne.age, "Years old", "My Profession is", userOne.job, "My salary is", userOne.salary)
}
```

## Frequently Asked Questions (FAQ)

### ১. Go-তে struct কী?

**উত্তর:** struct হলো একটি composite data type যা বিভিন্ন টাইপের ভেরিয়েবলকে একটি নামের অধীনে গ্রুপ করে।

**উদাহরণ:**

```go
type Person struct {
  name string
  age int
}

func main() {
  p := Person{name: "Alice", age: 30}
  fmt.Println(p)
}
```

---

### ২. Go-তে struct কীভাবে define ও initialize করবেন?

**উত্তর:** `type` কিওয়ার্ড ব্যবহার করে struct define করা হয় এবং struct literal দিয়ে initialize করা যায়।

**উদাহরণ:**

```go
type Car struct {
  brand string
  year  int
}

func main() {
  c := Car{brand: "Toyota", year: 2020}
  fmt.Println(c)
}
```

---

### ৩. Go-তে কি anonymous struct তৈরি করা যায়?

**উত্তর:** হ্যাঁ, আপনি একটি unnamed struct তৈরি করতে পারেন।

**উদাহরণ:**

```go
func main() {
  person := struct {
    name string
    age  int
  }{
    name: "John",
    age: 25,
  }
  fmt.Println(person)
}
```

---

### ৪. Go-তে struct এর fields কীভাবে access ও modify করবেন?

**উত্তর:** dot (`.`) অপারেটর ব্যবহার করে struct এর field access ও পরিবর্তন করা যায়।

**উদাহরণ:**

```go
type Book struct {
  title  string
  author string
}

func main() {
  b := Book{title: "Go Programming", author: "John Doe"}
  b.title = "Advanced Go Programming"
  fmt.Println(b)
}
```

---

### ৫. Go-তে struct এর জন্য কি method তৈরি করা যায়?

**উত্তর:** হ্যাঁ, struct এর জন্য আপনি method তৈরি করতে পারেন।

**উদাহরণ:**

```go
type Rectangle struct {
  width, height float64
}

func (r Rectangle) Area() float64 {
  return r.width * r.height
}

func main() {
  rect := Rectangle{width: 10, height: 5}
  fmt.Println("Area:", rect.Area())
}
```

---

### ৬. struct method-এ value ও pointer receiver-এর মধ্যে পার্থক্য কী?

**উত্তর:** value receiver struct-এর কপি নিয়ে কাজ করে, আর pointer receiver মূল struct-এ পরিবর্তন করে।

**উদাহরণ:**

```go
type Counter struct {
  count int
}

func (c *Counter) Increment() {
  c.count++
}

func main() {
  c := Counter{}
  c.Increment()
  fmt.Println(c.count)
}
```

---

### ৭. Go-তে struct এর ভিতরে অন্য struct রাখা যায় কি?

**উত্তর:** হ্যাঁ, struct embedding এর মাধ্যমে আপনি একটি struct-এর ভিতরে আরেকটি struct রাখতে পারেন।

**উদাহরণ:**

```go
type Address struct {
  city, state string
}

type Person struct {
  name    string
  address Address
}

func main() {
  p := Person{name: "Alice", address: Address{city: "New York", state: "NY"}}
  fmt.Println(p)
}
```

---

### ৮. Go-তে দুটি struct কীভাবে তুলনা করবেন?

**উত্তর:** যদি সব field comparable হয়, তাহলে `==` অপারেটর ব্যবহার করে তুলনা করা যায়।

**উদাহরণ:**

```go
type Point struct {
  x, y int
}

func main() {
  p1 := Point{x: 1, y: 2}
  p2 := Point{x: 1, y: 2}
  fmt.Println(p1 == p2) // true
}
```

---

### ৯. Go-তে struct কি map-এর key হিসেবে ব্যবহার করা যায়?

**উত্তর:** হ্যাঁ, যদি struct-এর সব field comparable হয়, তাহলে struct কে map-এর key হিসেবে ব্যবহার করা যায়।

**উদাহরণ:**

```go
type Point struct {
  x, y int
}

func main() {
  m := make(map[Point]string)
  m[Point{x: 1, y: 2}] = "A Point"
  fmt.Println(m)
}
```

---

### ১০. Go-তে struct-এর field গুলোতে কীভাবে iteration করবেন?

**উত্তর:** `reflect` প্যাকেজ ব্যবহার করে struct-এর ফিল্ড গুলোর উপর iteration করা যায়।

**উদাহরণ:**

```go
import (
  "fmt"
  "reflect"
)

type Person struct {
  Name string
  Age  int
}

func main() {
  p := Person{Name: "Alice", Age: 30}
  v := reflect.ValueOf(p)
  for i := 0; i < v.NumField(); i++ {
    fmt.Println(v.Type().Field(i).Name, v.Field(i))
  }
}
```
