# ErrPref - Adding Function Chains to Error Messages

Error Prefix - A lightweight GoLang type for attaching function chain lists and error context
strings to error messages.

This software package was written in the [Go](https://golang.org/) programming language, a.k.a. ***Golang***.

## Table of Contents

- [ErrPref - Adding Function Chains to Error Messages](#errpref---adding-function-chains-to-error-messages)
  - [Table of Contents](#table-of-contents)
  - [The Problem](#the-problem)
  - [The Solution](#the-solution)
  - [Definition of Terms](#definition-of-terms)
    - [Error Prefix](#error-prefix)
    - [Error Context](#error-context)
  - [String Formatting Conventions](#string-formatting-conventions)
    - [Error Prefix Delimiters](#error-prefix-delimiters)
    - [Error Context Delimiters](#error-context-delimiters)
  - [ErrPref](#errpref)
    - [Example Usage Summary](#example-usage-summary)
    - [Conclusion](#conclusion)
  - [ErrPrefixDto - Error Prefix Data Transfer Object](#errprefixdto---error-prefix-data-transfer-object)
  - [Source Code Documentation](#source-code-documentation)
  - [Version](#version)
  - [License](#license)
  - [Comments And Questions](#comments-and-questions)

## The Problem

As my Go programs, types and methods have grown in sophistication and complexity, the need for improved error tracking, detailed error messages, and a record of code execution has become more important. The terms 'Function Chains' or 'Method Chains' as used here describes a list of the functions called prior to encountering a particular error. Adding function chain documentation to returned error messages seems to make error tracking and management a much easier proposition.

Basically, when I see an error message, especially during the development phase, the first question that comes to mind is, **How the Hell did I get here?** Answering that question quickly and with certainty usually requires a list of functions ordered by their execution sequence.  Adding that information to error messages is the focus of this project.

## The Solution

The ***ErrPref*** or error prefix project is intended to provide better function or method documentation in error messages returned by Go functions. Two types of objects have been defined for this purpose: ***ErrPref*** and ***ErrPrefixDto***. Both of these types are designed to receive function chain information, format it and return the formatted strings for inclusion in error messages.

The idea is not to go overboard. I tried my hand at an error handler for Go Programs earlier and quickly realized that complexity was growing exponentially. At least at the outset of this project, the idea is to keep ***ErrPref*** simple, as a fast and efficient mechanism for adding error prefix text and function chain lists to error messages. That said, ***ErrPref*** is a work in progress.
The current code is version 1.0.0. We shall see where this leads. For examples and documentation, refer to the source code file, ***errpref.go***.

## Definition of Terms

For our purposes error prefix consists of two elements, a function name, and an optional error context string.

### Error Prefix

**Error Prefix** text is designed to be configured at the beginning of error messages and is most often used to document the thread of code execution by listing the calling sequence for a specific list of functions and/or methods.

### Error Context

**Error Context** strings are designed to provide additional information about the function or method identified by the associated **Error Prefix** text. Typical context information might include variable names, variable values and additional details on function execution.

Error Context strings are optional and are always associated with an Error Prefix string.

## String Formatting Conventions

When using the methods provided by types ***ErrPref*** and ***ErrPrefixDto*** it's best to use the following formatting conventions:

### Error Prefix Delimiters

Error Prefix Element are delimited by one of two delimiters: new line character (**\n**) or in-line delimiter string (**[SPACE]-[SPACE]**).

### Error Context Delimiters

Error Context Sub-Elements likewise have two delimiters: new line (**\n[SPACE]:[SPACE]\[SPACE]**) or in-line delimiter **([SPACE]:[SPACE]**)

## ErrPref

Type, ***ErrPref*** is designed to format text strings for use in error messages. ***ErrPref*** is simple, lightweight, easy to use and seems to work in a variety of situations. The concept is straight forward, "put raw text strings in - get formatted error information out".

The ***ErrPref*** methods do not use pointer receivers. All receivers are value receivers. 

### Example Usage Summary

***ErrPref*** passes error prefix information by string. An example calling sequence is shown as follows:

``` go

func(tx1 *Tx1) Something() {

 tx2 := Tx2{}

 err := Tx2.SomethingElse("Tx1.Something()")

 if err !=nil {
   fmt.Printf("%v\n",
   err.Error())
   return
 }

}

func(tx2 *Tx2) SomethingElse(errPrefix string) error {

  // errorPrefix now equals "Tx1.Something()"
  errorPrefix := ErrPref{}.EPrefOld(errPrefix)

  // errorPrefix now equals
  // "Tx1.Something() - Tx2.SomethingElse()"
  errorPrefix = ErrPref{}.EPref(
    errorPrefix,
    "Tx2.SomethingElse()")

  tx3 := Tx3{}

  err := Tx3.DoSomething(errorPrefix)

  if err !=nil {
   return err
  }
}

func(tx3 *Tx3) DoSomething(errorPrefix string) error {

    // Adding Error Context
    errorPrefix = ErrPref{}.EPrefCtx(
                   "Tx3.DoSomething()",
                   "A=B/C C== 0.0")

    .... bad code ....

    if isDivideByZero {
      return fmt.Errorf("%v\n" +
      "I divided by zero and got this error.\n",
      errorPrefix)
    }
}
```

When this error is returned up the function chain and finally printed out, the text will look like this:

``` go
 Tx1.Something() - Tx2.SomethingElse()
 Tx3.DoSomething() : A=B/C C== 0.0
 I divided by zero and got this error.

```

### Conclusion

By now you're saying, **"Whoopee - all you've done is concatenated a series of strings!"**

That's what I thought, at first. However, there is the problem of delimiting error prefixes and error context information plus the problem of determining how much information to show on each line. For types ***ErrPref*** and ***ErrPrefixDto***, the user may configure the line length which will determine whether a new line character (**\n**) or in-line delimiter string (**[SPACE]-[SPACE]**) is used to separate error prefix information.

## ErrPrefixDto - Error Prefix Data Transfer Object

The type, ***ErrPrefixDto***, offers the same services as type, ***ErrPref***, but is packaged in a different architecture. While ***ErrPref*** methods receive a string and instantly return a formatted string, ***ErrPrefixDto*** encapsulates error prefix information in an internal array of Error Prefix Information objects. Strings are only created when the type's ***String()*** method is called. Instances of ***ErrPrefixDto*** are designed to be passed as input parameters to subsidiary methods.

With the sole exception of the ***String()*** method, All ***ErrPrefixDto*** methods have pointer receivers.

Consider what the example function calling sequence would look like if ***ErrPrefixDto*** was employed:

``` go
func(tx1 *Tx1) Something() {

  ePrefDto := ErrPrefixDto{}.New()

  ePrefDto.SetEPref("Tx1.Something()")

  tx2 := Tx2{}

  err := Tx2.SomethingElse(ePrefDto)

  if err !=nil {
    fmt.Printf("%v\n",
    err.Error())
    return
  }

}

func(tx2 *Tx2) SomethingElse(ePrefDto ErrPrefixDto) error {

  ePrefDto.SetEPref("Tx2.SomethingElse()")

  tx3 := Tx3{}

  err := Tx3.DoSomething(ePrefDto)

  if err !=nil {
    return err
  }
}

func(tx3 *Tx3) DoSomething(ePrefDto ErrPrefixDto) error {

    // Add error prefix and error context
    // information.
    ePrefDto.SetEPrefCtx(
       "Tx3.DoSomething()",
       "A=B/C C== 0.0")

    .... bad code ....

    if isDivideByZero {
      return fmt.Errorf("%v\n" +
      "I divided by zero and got this error.\n",
      ePrefDto.String())
    }
}


```

When this error is returned up the function chain and finally printed out, the text will look like this:

``` text
 Tx1.Something() - Tx2.SomethingElse()
 Tx3.DoSomething() : A=B/C C== 0.0
 I divided by zero and got this error.

```

## Source Code Documentation

[Source Documentation](https://pkg.go.dev/github.com/MikeAustin71/errpref)

## Version

The latest version is Version 1.2.0.

Original Release Date: 2021-02-11 16:46:00 USACST

This Version supports *Go* modules.

[Version 1.2.0 Release Notes](./releasenotes.md)

## License

Use of this source code is governed by the (open-source) MIT-style license which can be found in the LICENSE file located in this directory.

[MIT License](./LICENSE)

## Comments And Questions

Send questions or comments to:

``` text
    mike.go@paladinacs.net
```
