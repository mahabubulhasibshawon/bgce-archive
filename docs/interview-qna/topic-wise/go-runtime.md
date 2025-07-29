[**Author:** @mahabubulhasibshawon
**Date:** 2025-07-29
**Category:** e.g., interview-qa/topic-wise
**Tags:** [go, concurrency, channels, go-runtime]
]



# Go Runtime Explained — The Mini Operating System of Go

## Go Runtime কী?

* Go প্রোগ্রামের ভিতরে থাকা **একটি ছোট্ট অপারেটিং সিস্টেমের মতো সফটওয়্যার কম্পোনেন্ট**
* Go Runtime **goroutine scheduling**, **memory management**, **garbage collection**, **syscalls** ইত্যাদি কাজ করে
* Go প্রোগ্রাম চালু হওয়ার সাথে সাথেই Go Runtime ইনিশিয়ালাইজ হয় এবং কাজ শুরু করে
* এটি OS এর kernel এর ওপর তৈরি, কিন্তু অনেক কাজ নিজেই করে

---

## Go Runtime এর প্রধান দায়িত্বসমূহ

| দায়িত্ব                  | বর্ণনা                                                                                          |
| ------------------------ | ----------------------------------------------------------------------------------------------- |
| **Goroutine Scheduler**  | Goroutine গুলোকে OS thread এর উপরে ম্যানেজ করে, কখন কোন goroutine চলবে ঠিক করে দেয়              |
| **Memory Manager**       | Heap allocation, stack management করে, এবং প্রয়োজনে stack গুলোর সাইজ ডায়নামিক বাড়ায়             |
| **Garbage Collector**    | অপ্রয়োজনীয় মেমোরি পরিষ্কার করে, মেমোরি লিক প্রতিরোধ করে                                         |
| **Syscall Interface**    | OS এর সাথে যোগাযোগ করে, system calls পরিচালনা করে যেমন ফাইল পড়া/লেখা, নেটওয়ার্ক অপারেশন ইত্যাদি |
| **Scheduler Management** | Logical Processor (P), OS Thread (M), এবং Goroutine (G) এর মধ্যকার সমন্বয় করে                   |

---

## Go Runtime কিভাবে কাজ করে?

### ১. Initialization

* প্রোগ্রাম শুরু হলে Go Runtime নিজেকে initialize করে
* সিস্টেমের CPU core সংখ্যা দেখে logical processor (P) তৈরি করে
* প্রতিটি P এর জন্য একটি OS thread (M) তৈরি হয়
* Main goroutine (G) তৈরি হয়, যেটা `main()` ফাংশন execute করে

### ২. Goroutine Scheduling

* Thousands বা লাখ লাখ goroutine create করা যায়
* Go Runtime scheduler ছোট ছোট টুকরো (time slice) দিয়ে goroutine গুলোকে OS thread এ ভাগ করে execute করায়
* OS thread (M) CPU core এ বসে goroutine (G) execute করে
* Scheduler balanced রাখে workload, blocking গুলো handle করে

### ৩. Memory Management

* প্রতিটি goroutine এর stack শুরু হয় ছোট (2KB) এবং প্রয়োজন মতো বড় হয়
* Heap allocator মেমোরি allocate ও free করে
* Garbage collector চালু থাকে, মেমোরি স্বয়ংক্রিয়ভাবে পরিষ্কার করে

### ৪. System Call & Runtime Calls

* Go Runtime OS এর system calls abstract করে দেয়
* Runtime calls যেমন `runtime.Gosched()`, `runtime.GOMAXPROCS()` ইত্যাদি control execution

---

## M\:P\:G মডেল সংক্ষেপে

| Component         | Meaning            | কাজের ধরন                             |
| ----------------- | ------------------ | ------------------------------------- |
| **M** (Machine)   | OS Thread          | OS thread, যা CPU core এ run করে      |
| **P** (Processor) | Logical Processor  | Scheduler এর unit, M কে G দান করে     |
| **G** (Goroutine) | Lightweight Thread | কাজের unit, যা scheduler schedule করে |

Go Runtime এই তিনের মধ্যে সমন্বয় করে concurrency পরিচালনা করে।

---

## Go Runtime এর কাজের Flow

```plaintext
Program Start
    ↓
Go Runtime Initialize
    ↓
CPU core অনুযায়ী P তৈরি
    ↓
M (OS Thread) তৈরি, P-র সাথে map হয়
    ↓
Main Goroutine তৈরি (G)
    ↓
Scheduler G গুলোকে M এর উপর run করে
    ↓
Goroutine execute হয় CPU core এ
    ↓
Memory manager মেমোরি allocate/free করে
    ↓
Garbage collector মেমোরি পরিষ্কার করে
```

---

## কিছু গুরুত্বপূর্ণ Runtime ফাংশন

* `runtime.Gosched()` — বর্তমানে চলমান goroutine কে voluntary CPU থেকে ছাড়িয়ে দেয়, scheduler অন্য goroutine রান করতে পারে
* `runtime.GOMAXPROCS(n)` — Go Runtime-কে কয়টি logical processor ব্যবহার করতে হবে নির্দেশ দেয়
* `runtime.Goexit()` — goroutine নিজেই execution শেষ করে

---

## সংক্ষেপে

* Go Runtime হলো Go প্রোগ্রামের **concurrency এবং execution control করার মূল স্তর**
* OS thread এর উপরে lightweight goroutine scheduling করে
* স্মার্ট মেমোরি ম্যানেজমেন্ট ও garbage collection পরিচালনা করে
* Programmers এর concurrency management অনেক সহজ ও efficient করে দেয়

---
