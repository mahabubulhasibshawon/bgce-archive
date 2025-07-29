[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** interview-qa/maps
**Tags:** [go, maps, data-structure]
]

# Maps 

**Map** হলো একটি ডেটা স্ট্রাকচার যা `key:value` জোড়ায় ডেটা সংরক্ষণ করতে দেয়।
Map একটি **unordered** এবং **changeable** collection, যেখানে **duplicate key** রাখা যায় না।
একটি map-এর default মান হলো `nil`।

```go
userInfo := map[string]int{
    "huxn": 17,
    "alex": 18,
    "john": 27,
}

fmt.Println(userInfo)
userInfo["jordan"] = 15
fmt.Println(userInfo["huxn"])
fmt.Println(userInfo["alex"])
fmt.Println(userInfo["john"])
```

## Frequently Asked Questions

---

### 1. Go-তে কীভাবে map তৈরি করা যায়?

**উত্তর:**
`make` ফাংশন অথবা map literal ব্যবহার করে map তৈরি করা যায়।

```go
// make ফাংশন দিয়ে
userInfo := make(map[string]int)
userInfo["huxn"] = 17
userInfo["alex"] = 18
fmt.Println(userInfo)

// map literal দিয়ে
userInfo := map[string]int{
    "huxn": 17,
    "alex": 18,
}
fmt.Println(userInfo)
```

---

### 2. কীভাবে map-এ কোনো key আছে কিনা চেক করব?

**উত্তর:**
map lookup করার সময় দ্বিতীয় রিটার্ন ভ্যালু দিয়ে চেক করা যায় key আছে কিনা।

```go
userInfo := map[string]int{
    "huxn": 17,
    "alex": 18,
}

value, exists := userInfo["huxn"]
if exists {
    fmt.Println("Key exists with value:", value)
} else {
    fmt.Println("Key does not exist")
}
```

---

### 3. Go-তে map থেকে কোনো key কীভাবে মুছব?

**উত্তর:**
`delete` ফাংশন ব্যবহার করে map থেকে key মুছে ফেলা যায়।

```go
userInfo := map[string]int{
    "huxn": 17,
    "alex": 18,
}
delete(userInfo, "huxn")
fmt.Println(userInfo)
```

---

### 4. সব ধরনের টাইপ কি map-এর key হতে পারে?

**উত্তর:**
না, key অবশ্যই এমন টাইপ হতে হবে যেটা **comparable**, যেমন: `string`, `int` ইত্যাদি।

```go
// বৈধ key
userInfo := map[string]int{
    "huxn": 17,
    "alex": 18,
}

// অবৈধ key (যেমন slice)
// userInfo := map[[]int]int{} // এটি error দেবে
```

---

### 5. Go-তে map কীভাবে iterate করব?

**উত্তর:**
`for` লুপ এবং `range` কীওয়ার্ড ব্যবহার করে map-এ loop চালানো যায়।

```go
userInfo := map[string]int{
    "huxn": 17,
    "alex": 18,
}

for key, value := range userInfo {
    fmt.Printf("Key: %s, Value: %d\n", key, value)
}
```

---

### 6. map-এর zero value কী?

**উত্তর:**
map-এর zero value হলো `nil`।

```go
var userInfo map[string]int
fmt.Println(userInfo == nil) // true
```

---

### 7. Go-তে দুটি map কী compare করা যায়?

**উত্তর:**
না, map সরাসরি compare করা যায় না। ম্যানুয়ালি key-value গুলা মিলিয়ে দেখতে হয়।

```go
map1 := map[string]int{"huxn": 17, "alex": 18}
map2 := map[string]int{"huxn": 17, "alex": 18}

// সরাসরি তুলনা যাবে না
// fmt.Println(map1 == map2) // এটি error দেবে

// ম্যানুয়ালি তুলনা
areEqual := true
for key, value := range map1 {
    if map2[key] != value {
        areEqual = false
        break
    }
}
fmt.Println("Maps are equal:", areEqual)
```

---

### 8. map-এর দৈর্ঘ্য (length) কীভাবে বের করব?

**উত্তর:**
`len` ফাংশন ব্যবহার করে map-এর key-value count জানা যায়।

```go
userInfo := map[string]int{
    "huxn": 17,
    "alex": 18,
}
fmt.Println("Length of map:", len(userInfo))
```

---

### 9. Go-তে কী nested map তৈরি করা যায়?

**উত্তর:**
হ্যাঁ, map-এর value হিসেবে আরেকটি map রাখা যায়।

```go
nestedMap := map[string]map[string]int{
    "group1": {
        "huxn": 17,
        "alex": 18,
    },
    "group2": {
        "john": 27,
    },
}
fmt.Println(nestedMap)
```

---

### 10. Go-তে map কপি কীভাবে করব?

**উত্তর:**
একটি map থেকে আরেকটিতে key-value গুলো ম্যানুয়ালি কপি করতে হয়।

```go
originalMap := map[string]int{
    "huxn": 17,
    "alex": 18,
}
copiedMap := make(map[string]int)

for key, value := range originalMap {
    copiedMap[key] = value
}
fmt.Println("Original Map:", originalMap)
fmt.Println("Copied Map:", copiedMap)
```
