[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** interview-qa/strings
**Tags:** [go, strings, operations]]

# GoLang Strings

এই ডকুমেন্টে GoLang-এ string অপারেশনগুলো নিয়ে উদাহরণসহ আলোচনা করা হয়েছে।

## প্রায় জিজ্ঞাসিত প্রশ্নোত্তর

### ১. Go-তে কীভাবে স্ট্রিং যুক্ত (concatenate) করবেন?

আপনি `+` অপারেটর ব্যবহার করে স্ট্রিং যুক্ত করতে পারেন।

```go
fmt.Println("Hello, " + "World!")
```

### ২. Go-তে স্ট্রিং-এর দৈর্ঘ্য কীভাবে জানবেন?

`len` ফাংশন ব্যবহার করে string-এর দৈর্ঘ্য পেতে পারেন।

```go
message := "Hello World!"
fmt.Println("Length:", len(message))
```

### ৩. Go-তে কীভাবে একটি substring বের করবেন?

slice notation ব্যবহার করে substring extract করা যায়।

```go
message := "Hello World!"
fmt.Println("Substring:", message[0:5]) // "Hello"
```

### ৪. Go-তে দুটি স্ট্রিং কীভাবে তুলনা করবেন?

`==` ও `!=` অপারেটর বা `strings.Compare` ফাংশন ব্যবহার করতে পারেন।

```go
msg1 := "one"
msg2 := "two"
fmt.Println("Equal?", msg1 == msg2)
fmt.Println("Not Equal?", msg1 != msg2)
fmt.Println(strings.Compare(msg1, msg2))
```

### ৫. Go-তে কিভাবে চেক করবেন কোনো স্ট্রিং-এর মধ্যে নির্দিষ্ট substring আছে কিনা?

`strings.Contains` ফাংশন ব্যবহার করুন।

```go
message := "Hello World!"
fmt.Println(strings.Contains(message, "World")) // true
```

### ৬. Go-তে স্ট্রিং কীভাবে uppercase বা lowercase করবেন?

`strings.ToUpper` ও `strings.ToLower` ফাংশন ব্যবহার করুন।

```go
message := "Hello World!"
fmt.Println(strings.ToUpper(message)) // "HELLO WORLD!"
fmt.Println(strings.ToLower(message)) // "hello world!"
```

### ৭. Go-তে কীভাবে একটি স্ট্রিংকে ভাগ করবেন (split) multiple substring-এ?

`strings.Split` ফাংশন ব্যবহার করুন।

```go
message := "Hello World!"
fmt.Println(strings.Split(message, " ")) // ["Hello", "World!"]
```

### ৮. Go-তে কীভাবে স্ট্রিং-এর substring replace করবেন?

`strings.Replace` ফাংশন ব্যবহার করুন।

```go
message := "Hello World!"
fmt.Println(strings.Replace(message, "World", "Go", 1)) // "Hello Go!"
```

### ৯. Go-তে কীভাবে স্ট্রিং থেকে leading ও trailing spaces সরাবেন?

`strings.TrimSpace` ফাংশন ব্যবহার করুন।

```go
message := "  Hello World!  "
fmt.Println(strings.TrimSpace(message)) // "Hello World!"
```

### ১০. Go-তে একটি স্ট্রিং এর প্রথম ক্যারেক্টার কীভাবে পাবেন?

index ব্যবহার করে স্ট্রিং-এর প্রথম ক্যারেক্টার অ্যাক্সেস করুন।

```go
message := "Hello World!"
fmt.Printf("First character: %c\n", message[0])
```

---

## মৌলিক String অপারেশন

### Concatenation

আপনি `+` অপারেটর দিয়ে স্ট্রিং যোগ করতে পারেন।

```go
fmt.Println("Hello, " + "World!")
```

### Formatting Strings

Go-তে বিভিন্নভাবে string format করা যায় `fmt.Printf` ব্যবহার করে।

```go
message := "Hello World!"
fmt.Printf("Data: %v\n", message)
fmt.Printf("Data: %+v\n", message)
fmt.Printf("Data: %#v\n", message)
fmt.Printf("Data: %T\n", message)
fmt.Printf("Data: %q\n", message)
fmt.Printf("First character: %c\n", message[0])
```

### String Length

`len` ফাংশন ব্যবহার করে string-এর দৈর্ঘ্য জানতে পারেন।

```go
fmt.Println("Length:", len(message))
```

### Substrings

slice notation ব্যবহার করে substring বের করুন।

```go
fmt.Println("Substring:", message[0:5]) // এটি "Hello" প্রিন্ট করবে
```

### String Comparison

`==`, `!=` অপারেটর বা `strings.Compare` ফাংশন ব্যবহার করে স্ট্রিং তুলনা করুন।

```go
msg1 := "one"
msg2 := "two"

fmt.Println("Equal?", msg1 == msg2)
fmt.Println("Not Equal?", msg1 != msg2)
fmt.Println(strings.Compare(msg1, msg2))
```

---

## অতিরিক্ত String ফাংশন

### Contains

স্ট্রিং-এর মধ্যে নির্দিষ্ট substring আছে কিনা চেক করুন।

```go
fmt.Println(strings.Contains(message, "World")) // true
```

### ToUpper ও ToLower

স্ট্রিং uppercase বা lowercase-এ রূপান্তর করুন।

```go
fmt.Println(strings.ToUpper(message)) // "HELLO WORLD!"
fmt.Println(strings.ToLower(message)) // "hello world!"
```

### Split

স্ট্রিংকে substring-এর slice-এ ভাগ করুন।

```go
fmt.Println(strings.Split(message, " ")) // ["Hello", "World!"]
```

### Replace

Substring প্রতিস্থাপন করুন।

```go
fmt.Println(strings.Replace(message, "World", "Go", 1)) // "Hello Go!"
```

### Trim

leading ও trailing spaces সরান।

```go
fmt.Println(strings.TrimSpace("  Hello World!  ")) // "Hello World!"
```

---

যদি আরো string ফাংশন বা বিস্তারিত দরকার হয়, তবে আপনি [Go documentation](https://golang.org/pkg/strings/) দেখতে পারেন
