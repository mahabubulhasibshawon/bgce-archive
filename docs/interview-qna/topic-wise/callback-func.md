# Callback Functions

যদি একটি ফাংশনকে আরেকটি ফাংশনের আর্গুমেন্ট হিসেবে পাস করা হয়, তাহলে সেই ধরণের ফাংশনকে Higher-Order function বলা হয়। এইভাবে একটি ফাংশনকে আর্গুমেন্ট হিসেবে পাস করাকে Go তে callback function বা first-class function ও বলা হয়।


```go
package main
import "fmt"

func addName(name string, callback func(string)) {
    callback(name)
}

func main() {
    addName("HuXn", func(nm string) {
        fmt.Printf("Hi, my name is %v\n", nm)
    })
}
```
