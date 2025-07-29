[**Author:** @mdimamhosen, @mahabubulhasibshawon
**Date:** 2025-04-19
**Category:** interview-qa/boolean
**Tags:** [go, boolean, data-types]
]

# একটি বুলিয়ান ডেটা-টাইপ হতে পারে "TRUE" অথবা "FALSE"

```go
package main

import "fmt"

func main() {
	isGolangPL := true
	isHtmlPL := false
	fmt.Println(isGolangPL)
	fmt.Println(isHtmlPL)
}
```

## Frequently Asked Questions

### Qপ্রশ্ন ১: কিভাবে শর্তযুক্ত স্টেটমেন্টে বুলিয়ান ভ্যালু ব্যবহার করা যায়?

**উত্তর:** বুলিয়ান ভ্যালু প্রায়ই শর্তযুক্ত স্টেটমেন্টে ব্যবহৃত হয় একটি প্রোগ্রামের প্রবাহ নিয়ন্ত্রণ করতে। 
উদাহরণস্বরূপ:

```go
package main

import "fmt"

func main() {
	isEven := true
	if isEven {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}
}
```

### প্রশ্ন ২: বুলিয়ান ভ্যালু কি সরাসরি তুলনা করা যায়?

**উত্তর:** হ্যাঁ, বুলিয়ান ভ্যালুগুলো সরাসরি তুলনা করা যায় তুলনা অপারেটর ব্যবহার করে। উদাহরণস্বরূপ:

```go
package main

import "fmt"

func main() {
	isTrue := true
	isFalse := false
	fmt.Println(isTrue == isFalse) // Output: false
	fmt.Println(isTrue != isFalse) // Output: true
}
```
