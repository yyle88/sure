# Background of Creation

## Design Intent
This tool will allow the code to provide options for panic/warning when errors occur.

## Core Functions
The core of this package provides two functions:
```go
// Must - Hard, panics and crashes directly when there is an error, halting the process
func Must(err error)

// Soft - Soft, prints a warning log when there is an error, continues the process
func Soft(err error)
```

## Simple Usage
You can use it like this in your code logic:
```go
data, err := json.Marshal(example)
sure.Must(err) // Does nothing if there is no error, crashes with panic if there is an error
```
Or use:
```go
sure.Soft(err) // Does nothing if there is no error, prints a warning log if there is an error
```
Of course, it’s clear that this kind of brute force throwing exceptions or directly ignoring errors might not be suitable in formal scenarios and needs careful consideration.

## Code Generation
Use code generation to automatically give certain functions in a class the ability to `must`, or certain functions in a package the ability to `must`.

## Other Background:
During development, I encountered a very confusing situation, which was that I saw others defining the `Must` function like this: `a := Must(defaultValue)`, where if a calculation error occurs, it returns a default value, i.e., even if there is no result, it must return a result.

For example, this function:
```go
func (j *Json) MustInt(args ...int) int {
	var def int

	switch len(args) {
	case 0:
	case 1:
		def = args[0]
	default:
		log.Panicf("MustInt() received too many arguments %d", len(args))
	}

	i, err := j.Int()
	if err == nil {
		return i
	}

	return def
}
```
I looked at it for a long time before understanding the meaning of `must` here. They even made the default value optional, which made it clear that their understanding of `must` was different from mine.

In contrast, when they actually needed to directly warn on errors, they used the `Require` function. Oh my god, it seems that people who aren't proficient in English might not even understand what `must` means, which is pretty unfortunate.

My understanding of `must` is that you must accomplish something to get a result, or you must accept the punishment (i.e., `panic`), while their understanding of `must` is that no matter what difficulty you face, you must give a result, and if the result can't be obtained, you return the default value. They care more about guaranteeing a result.

Since there is ambiguity in understanding `must`, let me clarify: my `must` means you must accomplish something, and if no error occurs, it proceeds; otherwise, it will `panic`. The entire project is based on this context.

For example:
```go
func (T *SimpleMust) Strings(key string) (res []string) {
	res, err1 := T.T.Strings(key)
	sure.Must(err1)
	return res
}
```
Call:
```go
tags := sim.Must().Strings("tags")
```
If it can’t be obtained, it directly `panics` and crashes, effectively asserting when fetching the value.

The entire project is based on this context.

## Similar Packages
Of course, I spent a lot of time choosing between package names like `mustsoft` (soft and hard), `mustgo` (hard), `flexible` (flexible), and `mustdone` (must achieve), and finally decided that it could go well with `github.com/yyle88/done`, so I named it `sure`, which is relatively short.

Of course, the `github.com/yyle88/done` package can also solve the problem, but it still requires an extra `nice` call, and every time you wrap it outside, the code becomes harder to read when the wrapping exceeds two layers.

For example:
```go
defer func() { // This is an operation when closing db *gorm.DB, assume it is in a test case, where a temporary DB is created, and it can be closed like this after it ends
    done.Done(done.VCE(db.DB()).Nice().Close())
}()
```
The readability of this code has significantly decreased.

## Final Effect
In practice, the usage scenarios are rare, and in 90% of scenarios, regenerating code externally to empower methods is not as good as using `github.com/yyle88/done`. In the remaining 10% of cases, simply using `if err != nil { panic(err) }` works fine.

Of course, if you particularly want to use it, you can, as it can save three lines of code in crucial moments.

---
