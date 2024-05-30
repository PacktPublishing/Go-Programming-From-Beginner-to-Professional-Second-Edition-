# Go Programming - From Beginner to Professional Second Edition

<a href="https://www.packtpub.com/product/go-programming-from-beginner-to-professional-second-edition/9781803243054"> <img src="https://content.packt.com/B18621/cover_image_small.jpg" alt="Go Programming - From Beginner to Professional - Second Edition" itemprop="url" height="256px" align="right">

This is the code repository for [Go Programing - From Beginner to Professional Second Edition](https://www.packtpub.com/product/go-programming-from-beginner-to-professional-second-edition/9781803243054), published by Packt.

**Learn everything you need to build modern software using Go**

## What is this book about?
Master real-world development in “Go Programming – From Beginner to Professional”. From the basics to advanced topics, gain proficiency by mastering applications, database interactions, RESTful APIs, concurrency and more using Go idiomatic practices.
 
This book covers the following exciting features:
* Understand the Go syntax and apply it proficiently to handle data and write functions
* Debug your Go code to troubleshoot development problems
* Safely handle errors and recover from panics
* Implement polymorphism using interfaces and gain insight into generics
* Work with files and connect to popular external databases
* Create an HTTP client and server and work with a RESTful web API
* Use concurrency to design efficient software
* Use Go tools to simplify development and improve your code

If you feel this book is for you, get your [copy](https://a.co/d/b0BTbEk) today!

## Instructions and Navigations

All of the code is organized into folders. For example, Chapter02.

The code will look like the following:

```
package main
import "fmt"
func main() {
 fmt.Println(10 > 5)
 fmt.Println(10 == 5)
}
```

## OS Specific Instructions
* Sometimes some of the executable generated while running "go build" may not directly work for Windows, try running"go run ." or "go run main.go" 
* For Exercise 13.03, "cat" command is OS specific, "type" command would work for Windows.
* Running Exercise 13.04 may show some OS specific errors, replacing "echo" in code with "print" works for Windows.
* Exercise 14.01 applies for unix-like systems (linux and macOS). Windows does not work with signals the same way so this exercise will not work for windows. 

**Following is what you need for this book:**
Designed for both complete beginners in Go as well as professionals transitioning from another programming language, this book equips developers with skills needed to build real-world projects and launch their career in Go. With a step-by-step approach, beginners can grasp Go fundamentals even without prior programming experience, and gradually advance to idiomatic Go best practices, exploring the latest features of the language.

With the following software and hardware list you can run all code files present in the book (Chapter 1-21).

## Software and Hardware List
| Software/ Hardware required | OS required/ Other requirements |
| ------------------------------------ | ----------------------------------- |
| Go Compiler | Windows, macOS, Linux (https://packt.live/375DQDA) |
| Git | Windows, macOS, Linux (https://packt.live/35ByRug) |
| Visual Studio Code | Windows, macOS, Linux (https://packt.live/35KD2Ek) |
| Docker | Windows (https://packt.live/2EKGDG6), macOS (https://packt.live/34VJLJD), Linux |
| PostgreSQL | Windows, macOS, Linux (https://www.postgresql.org/download/) |

## Errata
* Page 10 (Exercise 1.01, Step 5): **rand.Seed(time.Now().UnixNano())** _should be_ **source := rand.NewSource(time.Now().UnixNano())**
                                                                                    **rng := rand.New(source)**
* Page 10 (Exercise 1.01, Step 6): **r := rand.Intn(5) + 1** _should be_ **r := rng.Intn(5) + 1**
* Page 10 (Exercise 1.01, Step 6): **between 0 and then** _should be_ **between o and 4, then**                                 
* Page 38 (Paragraph 1 , line 4): **it's a pointer because its output starts with &.** _should be_ **it's a pointer because we declared the variable with &.**
* Page 41 (Exercise 1.15, Step 4): **pointer** should be neglected
* Page 164 (Paragraph 2, line 3): **Go can assign a variable to a function** _should be_ **Go can assign a function to a variable**
* Page 194 (Exercise 5.07, Step 4): **func main() {import "fmt"** _should be_ **package main**
                                    **counter :==4**                          **import "fmt"**
                                                                              **func main() {**
                                                                              **counter := 4**                              
* Page 194 (Exercise 5.07, Step 7): **I--** _should be_ **i--**
* Page 248 (Paragraph 3, line 1): **implement an interface explicitly** _should be_ **are implemented explicitly**
* Page 267 (3rd block of code, line 2): **s interface{}** _should be_ **i interface{}**
* Page 297 (Introduction, Paragraph 1): It erroneously talks about the next chapter instead of the previous. This _should be_ **In the previous chapter, we delved into the realm of Go generics, an innovative advancement in the language offering type parameters, constraints, and type inference.**

## Related products
* gRPC Go for Professionals [[Packt]](https://www.packtpub.com/product/grpc-go-for-professionals/9781837638840) [[Amazon]](https://a.co/d/e8CWmQd)
* Building Modern CLI Applications in Go [[Packt]](https://www.packtpub.com/product/building-modern-cli-applications-in-go/9781804611654) [[Amazon]](https://a.co/d/7RhPoLY)

## Get to Know the Author
**Samantha Coyle**
a Software Engineer at Diagrid, specializes in Go for cloud-native developer tooling, adept at abstracting application development challenges. Committed to Open Source, she actively contributes to projects like Dapr and Testcontainers. With a diverse background in retail computer vision solutions and successful stabilization of industrial edge use cases, particularly within biopharma data pipelines, Samantha brings invaluable expertise to her projects. Her expertise extends to being CKAD-certifi ed and reviewing Go textbooks. She is passionate about empowering early-career, diverse professionals.

Samantha is part of a family of gophers, and enjoys GopherCon with her brother and identical twin sister. A seasoned speaker, she has presented at various conferences, including GopherCon
