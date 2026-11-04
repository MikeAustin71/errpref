## Table of Contents
- [***errpref*** - Adding function execution chains and error context to Error Messages In Go](#errpref---adding-function-execution-chains-and-error-context-to-error-messages-in-go)
- [The Problem](#the-problem)
- [A New Solution: ***ErrorGenericFactory***](#a-new-solution-errorgenericfactory)
		- [Specialized Error Types](#specialized-error-types)
		- [Custom Error Creation and Usage](#custom-error-creation-and-usage)
			- [1. Create an Instance of the Factory](#1-create-an-instance-of-the-factory)
			- [2. Select and Generate a Specialized Custom Error Instance](#2-select-and-generate-a-specialized-custom-error-instance)
			- [3. Apply the Specialized Error Instance As Needed](#3-apply-the-specialized-error-instance-as-needed)
		- [***ErrorGenericFactory*** Example Usage](#errorgenericfactory-example-usage)
			- [Call Tree for Test](#call-tree-for-test)
			- [Test Source Code](#test-source-code)
		- [***ErrorGenericFactory*** Source Code](#errorgenericfactory-source-code)
- [An Earlier Solution](#an-earlier-solution)
		- [Wait a Minute - Let's Think About This](#wait-a-minute---lets-think-about-this)
	- [What *errpref* Does and Doesn't Do](#what-errpref-does-and-doesnt-do)
	- [Definition of Terms](#definition-of-terms)
		- [Error Prefix](#error-prefix)
		- [Error Context](#error-context)
	- [Getting Started](#getting-started)
		- [Two Types To Choose From](#two-types-to-choose-from)
		- [ErrPrefixDto - A Full Featured Solution](#errprefixdto---a-full-featured-solution)
			- [Public Facing Methods](#public-facing-methods)
			- [Internal or Private Methods](#internal-or-private-methods)
			- [Error Context Examples](#error-context-examples)
				- [Example-1 Error Flow](#example-1-error-flow)
				- [Example-2 Calling Subsidiary Functions](#example-2-calling-subsidiary-functions)
			- [Customizing Input and Output String Delimiters](#customizing-input-and-output-string-delimiters)
			- [Maximum Text Line Length](#maximum-text-line-length)
			- [Left Margin Feature](#left-margin-feature)
			- [Leading And Trailing Text Strings](#leading-and-trailing-text-strings)
			- [Text Display On/Off Switch](#text-display-onoff-switch)
			- [***Is Last Text Line Terminated With New Line***](#is-last-text-line-terminated-with-new-line)
		- [ErrPref - Quick And Simple Solution](#errpref---quick-and-simple-solution)
			- [Example Usage Summary](#example-usage-summary)
			- [String Formatting Conventions](#string-formatting-conventions)
				- [Error Prefix Delimiters](#error-prefix-delimiters)
				- [Error Context Delimiters](#error-context-delimiters)
	- [Exchanging Error Prefix Information with User Defined Types](#exchanging-error-prefix-information-with-user-defined-types)
		- [*IBuilderErrorPrefix*](#ibuildererrorprefix)
		- [IBasicErrorPrefix](#ibasicerrorprefix)
		- [Custom Input String Delimiters](#custom-input-string-delimiters)
	- [Usage Examples](#usage-examples)
		- [Test Code](#test-code)
		- [Example Application](#example-application)
	- [User Package Import and Configuration](#user-package-import-and-configuration)
		- [Required ***Golang*** Source Code File Import Statement](#required-golang-source-code-file-import-statement)
		- [***go.mod*** File Configuration](#gomod-file-configuration)
			- [Option 1: Go Get Command](#option-1-go-get-command)
				- [How a User Employs `go get` (The Standard Workflow)](#how-a-user-employs-go-get-the-standard-workflow)
				- [Step 1: Initialize Their Own Project (If they haven't already)](#step-1-initialize-their-own-project-if-they-havent-already)
				- [Step 2: Run the Modern `go get` Command](#step-2-run-the-modern-go-get-command)
				- [Step 3: What Go Autonomously Does to `go.mod`](#step-3-what-go-autonomously-does-to-gomod)
			- [Option 2: ***go.mod*** Manual Configuration](#option-2-gomod-manual-configuration)
		- [](#)
	- [External Dependencies](#external-dependencies)
	- [Tests](#tests)
	- [OS Support](#os-support)
	- [Version](#version)
	- [***errPref*** Source Code Repository](#errpref-source-code-repository)
	- [Release Notes](#release-notes)
	- [License](#license)
	- [Comments And Questions](#comments-and-questions)



# ***errpref*** - Adding function execution chains and error context to Error Messages In Go

[![Go Reference](https://pkg.go.dev/badge/github.com/MikeAustin71/errpref.svg)](https://pkg.go.dev/github.com/MikeAustin71/errpref)

**The *errpref* software package contains GoLang custom types designed to facilitate and improve the process of identifying, tracing and understanding errors generated in code.**

The ***errpref*** software package was written in the [Go](https://golang.org/) programming language, a.k.a. ***Golang***.

The current version of ***errpref*** is ***Version 2.0.1*** which features:

- Support for ***Go Version 1.27.3***
-  A new custom error type, ***ErrorGenericFactory*** 

For more details, see the [Release Notes](./releasenotes.md).



# The Problem


As my **Go** programs, types and methods have grown in sophistication and complexity, the need for improved error tracking, detailed error messages, and a bread crumb trail of code execution has become more important. The terms 'Function Chains' or 'Method Chains', as used here, describes a list of the functions called prior to encountering a particular error. Adding function chain documentation to returned error messages makes error tracking and management a much easier proposition.

Basically, when I see an error message, especially during the development phase, the first question that comes to mind is, **How the Hell did I get here?** Answering that question quickly, and with certainty, usually requires a list of functions ordered by their execution sequence.  Adding that information to error messages is the focus of this project.



# A New Solution: ***ErrorGenericFactory***

Based on testing conducted to date, the ***ErrorGenericFactory*** custom error type was found to be most effective in providing detailed, flexible, and informative error messages for various error scenarios. These features allow developers to create custom error messages with specific error codes and descriptions, making it easier to identify and resolve issues in their Go code.

***ErrorGenericFactory*** utilizes a bubble up approach. As errors bubble up through the function call tree, the features offered by ***ErrorGenericFactory*** allow developers to add context and trace identification information at every level of the call tree. 

Based on testing conducted to date, the ***ErrorGenericFactory*** custom error type has been found to be most effective in providing detailed and informative error messages for various error scenarios. It allows developers to create custom error messages with specific error codes and descriptions, making it easier to identify and resolve issues in their Go code.

***ErrorGenericFactory*** utilizes a bubble up approach. As errors bubble up through the function call tree, the features offered by ***ErrorGenericFactory*** allow developers to add context and trace identification information at every level of the call tree. 

The ***ErrorGenericFactory*** type is located in source code file: **`errpref/ErrorGenericFactory_0010.go`** .

### Specialized Error Types

***ErrorGenericFactory*** incorporates the following specialized error types:

1. **ErrorInputParameterOutOfRange**
2. **ErrorInputParameterNilPointer**
3. **ErrorNativeToFunction**
4. **ErrorFromFunctionCall**

These specialized error types will be identifiable when inspecting a series of **wrapped** errors. Reference the tests included in **`errpref/zzzt_ErrorGenericFactory_0010_test.go`** .

### Custom Error Creation and Usage

#### 1. Create an Instance of the Factory

Call the Factory Constructor - ***ErrorGenericFactory.NewErrorGenericConfig()***

#### 2. Select and Generate a Specialized Custom Error Instance

Call one of the methods returning a specialized custom error instance. Example: ***ErrorGenericFactory. GetErrorFromFunctionCall()***

#### 3. Apply the Specialized Error Instance As Needed

See the Example Usage below [Test Source Code](#Test Source Code)



### ***ErrorGenericFactory*** Example Usage

#### Call Tree for Test

```go
TestFunc04() // Starting point.TestFunc04() calls Series04Func01() 
            |
            |----> Series04Func01() // Receives and Returns Error
                 |
                 ----> Series04Func02() // Receives and Returns Error
                     |
                     ----> Series04Func03()  // Receives and Returns Error
                         |
                          ----> Series04Func04() // Base error generated here
```



#### Test Source Code

```go
import (
	"fmt"

	ePref "github.com/MikeAustin71/errpref"
)

type TestSeries04Functions struct {
}

func (tf *TestSeries04Functions) TestFunc04() {

	fmt.Printf("\nExecuting TestFunc04() ...\n")

    // NewErrorGenericConfig()
    // Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"TestFunc02")

	err := tf.Series04Func01() // First function call in the chain

	if err != nil {

        // Error returned.
        // Add context information for called function.
		combinedWrappedErr := errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series04Func01()",
			"",
			"",
			err)

		fmt.Printf("%s\n", combinedWrappedErr)

		innerMostErr := errGenericCfg.GetInnerMostError(combinedWrappedErr)

		fmt.Println("Inner Most Error:")

		fmt.Printf("%s\n", innerMostErr)

		return
	}

	fmt.Println("TestFunc02() - Series04Func01() completed without error!")

	return
}

func (tf *TestSeries04Functions) Series04Func01() error {

    // NewErrorGenericConfig()
    // Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"Series04Func01")

	err := tf.Series04Func02()

	if err != nil {

        // Error returned.
        // Add context information for called function.
		return errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series04Func02()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries04Functions) Series04Func02() error {

    // NewErrorGenericConfig()
    // Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"Series04Func02")

	err := tf.Series04Func03()

	if err != nil {

        // Error returned.
        // Add context information for called function.
		return errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series04Func03()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries04Functions) Series04Func03() error {

    // NewErrorGenericConfig()
    // Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"Series04Func03")

	err := tf.Series04Func04()

	if err != nil {

        // Error returned.
        // Add context information for called function.
		return errGenericCfg.GetErrorFromFunctionCall(
			"",
			"Series04Func04()",
			"",
			"",
			err)

	}

	return nil
}

func (tf *TestSeries04Functions) Series04Func04() error {

    // NewErrorGenericConfig()
    // Add identification information for this function
	errGenericCfg := new(ePref.ErrorGenericFactory).NewErrorGenericConfig(
		"TestFuncs_04.go",
		"TestSeries04Functions",
		"Series04Func04")

        // Error originated in this function
        // Add context information and error message
        // for this error. 
	err := errGenericCfg.GetErrorNativeToFunction(
		"",
		"",
		"",
		"",
		"Serious Error Occurred!!",
		"I'm feeling down - and sent this error.")

	return err
}
```



### ***ErrorGenericFactory*** Source Code

```go
// Source Code File: ErrorGenericFactory_0010.go

package errpref

import (
	"errors"
	"fmt"
)

// ============================================================================
// 1. Core Configuration Types
// ============================================================================

type ErrorGenericFactory struct {
	sourceFile         string // Source file containing 'thisFunction'.
	thisFunctionType   string // The 'type' in which 'thisFunction' (method) resides.
	thisFunction       string // The current function processing this error.
	calledFunctionType string // The 'type' of the method identified in
	// 'calledFunction', below.
	calledFunction string // The method or function called by 'thisFunction'.
	parameter      string // Parameter associated with this error.
	path           string // File path associated with this error.
	fileName       string // File name associated with this error.
	url            string // URL associated with this error.
	errorContext   string // Error context associated with this error.
	errorMessage   string // The error message associated with this error.
}

// ============================================================================
// 2. Specialized Struct-Based Error Types
// ============================================================================

// ErrorInputParameterOutOfRange represents a parameter value out of bounds.
type ErrorInputParameterOutOfRange struct {
	ErrorGenericFactory // Value embedding ensures an isolated, safe snapshot
}

func (e *ErrorInputParameterOutOfRange) Error() string {
	return e.buildErrorStr()
}

// ErrorInputParameterNilPointer represents a required pointer passing a nil value.
type ErrorInputParameterNilPointer struct {
	ErrorGenericFactory
}

func (e *ErrorInputParameterNilPointer) Error() string {
	return e.buildErrorStr()
}

// ErrorNativeToFunction represents an internal operational failure within a function.
type ErrorNativeToFunction struct {
	ErrorGenericFactory
}

func (e *ErrorNativeToFunction) Error() string {
	return e.buildErrorStr()
}

// ErrorFromFunctionCall represents an error context layer wrapping a downstream failure.
type ErrorFromFunctionCall struct {
	ErrorGenericFactory
	unwrapped error // Keeps track of the original wrapped error
}

func (e *ErrorFromFunctionCall) Error() string {
	errorStr := e.buildErrorStr()
	if e.unwrapped == nil {
		return fmt.Sprintf("%s\nReturned error from called Function is nil!", errorStr)
	}
	return fmt.Sprintf("%s: %v", errorStr, e.unwrapped)
}

// Unwrap enables standard library tools (errors.Is, errors.As) to traverse past this type.
func (e *ErrorFromFunctionCall) Unwrap() error {
	return e.unwrapped
}

// ============================================================================
// 3. Factory Constructors and Builders
// ============================================================================

// NewErrorGenericConfig is designed to be called at the beginning of a function when
// only three elements are known: sourceFile, parentType, and thisFunction.
func (egc *ErrorGenericFactory) NewErrorGenericConfig(
	sourceFile string,
	parentType string,
	thisFunction string) *ErrorGenericFactory {

	return &ErrorGenericFactory{
		sourceFile:       sourceFile,
		thisFunctionType: parentType,
		thisFunction:     thisFunction,
	}
}

// GetErrorInputParameterOutOfRange yields a unique ErrorInputParameterOutOfRange error.
func (egc *ErrorGenericFactory) GetErrorInputParameterOutOfRange(
	parameter string,
	path string,
	fileName string,
	url string,
	errorContext string,
	errorMessage string) error {

	egc.parameter = parameter
	egc.path = path
	egc.fileName = fileName
	egc.url = url
	egc.errorContext = errorContext
	egc.errorMessage = errorMessage

	// Dereferencing (*egc) makes a clean, isolated value copy for this instance
	return &ErrorInputParameterOutOfRange{ErrorGenericFactory: *egc}
}

// GetErrorInputParameterNilPointer yields a unique ErrorInputParameterNilPointer error.
func (egc *ErrorGenericFactory) GetErrorInputParameterNilPointer(
	parameter string,
	path string,
	fileName string,
	url string,
	errorContext string,
	errorMessage string) error {

	egc.parameter = parameter
	egc.path = path
	egc.fileName = fileName
	egc.url = url
	egc.errorContext = errorContext
	egc.errorMessage = errorMessage

	return &ErrorInputParameterNilPointer{ErrorGenericFactory: *egc}
}

// GetErrorNativeToFunction yields a unique ErrorNativeToFunction error.
func (egc *ErrorGenericFactory) GetErrorNativeToFunction(
	parameter string,
	path string,
	fileName string,
	url string,
	errorContext string,
	errorMessage string) error {

	egc.parameter = parameter
	egc.path = path
	egc.fileName = fileName
	egc.url = url
	egc.errorContext = errorContext
	egc.errorMessage = errorMessage

	return &ErrorNativeToFunction{ErrorGenericFactory: *egc}
}

// GetErrorFromFunctionCall returns an inspectable, wrapped FunctionCallError.
func (egc *ErrorGenericFactory) GetErrorFromFunctionCall(
	calledFunctionType string,
	calledFunction string,
	errorContext string,
	errorMessage string,
	returnedCalledFunctionError error) error {

	egc.calledFunctionType = calledFunctionType
	egc.calledFunction = calledFunction
	egc.errorContext = errorContext
	egc.errorMessage = errorMessage

	return &ErrorFromFunctionCall{
		ErrorGenericFactory: *egc,
		unwrapped:           returnedCalledFunctionError,
	}
}

// GetInnerMostError
//
// Returns the innermost (or last) error from a series of
// wrapped errors.
func (egc *ErrorGenericFactory) GetInnerMostError(
	wrappedErrors error) error {

	if wrappedErrors == nil {
		return errors.New("ErrorGenericConfig.GetInnerMostError()\n" +
			"Input parameter 'wrappedErrors' is nil")
	}

	innerMostError := wrappedErrors

	for {

		unwrappedErr := errors.Unwrap(innerMostError)

		if unwrappedErr == nil {
			break
		}

		innerMostError = unwrappedErr
	}

	return innerMostError
}

// ============================================================================
// 4. Utility Context Setters
// ============================================================================

// SetDetailErrorContext updates the detail variables if non-empty arguments are provided.
func (egc *ErrorGenericFactory) SetDetailErrorContext(
	parameter string,
	path string,
	fileName string,
	url string) {

	if parameter != "" {
		egc.parameter = parameter
	}
	if path != "" {
		egc.path = path
	}
	if fileName != "" {
		egc.fileName = fileName
	}
	if url != "" {
		egc.url = url
	}
}

// SetErrorMessageElements updates context and messaging details if non-empty arguments are provided.
func (egc *ErrorGenericFactory) SetErrorMessageElements(
	errorContext string,
	errorMessage string) {

	if errorContext != "" {
		egc.errorContext = errorContext
	}
	if errorMessage != "" {
		egc.errorMessage = errorMessage
	}
}

// buildErrorStr constructs a string representation of the populated elements.
func (egc *ErrorGenericFactory) buildErrorStr() string {
	errStr := ""
	foundParametersCount := 0

	if egc.sourceFile != "" {
		errStr = "\n"
		errStr += fmt.Sprintf("Source File: %s\n", egc.sourceFile)
		foundParametersCount++
	}

	if egc.thisFunctionType != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("This Function Type: %s\n", egc.thisFunctionType)
		foundParametersCount++
	}

	if egc.thisFunction != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("This Function: %s\n", egc.thisFunction)
		foundParametersCount++
	}

	if egc.calledFunctionType != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("Called Function Type: %s\n", egc.calledFunctionType)
		foundParametersCount++
	}

	if egc.calledFunction != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("Called Function: %s\n", egc.calledFunction)
		foundParametersCount++
	}

	if egc.parameter != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("Parameter: %s\n", egc.parameter)
		foundParametersCount++
	}

	if egc.path != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("Path: %s\n", egc.path)
		foundParametersCount++
	}

	if egc.fileName != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("FileName: %s\n", egc.fileName)
		foundParametersCount++
	}

	if egc.url != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("URL: %s\n", egc.url)
		foundParametersCount++
	}

	if egc.errorContext != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("Error Context: %s\n", egc.errorContext)
		foundParametersCount++
	}

	if egc.errorMessage != "" {
		if foundParametersCount == 0 {
			errStr = "\n"
		}

		errStr += fmt.Sprintf("Error Message: %s\n", egc.errorMessage)
		foundParametersCount++
	}

	if foundParametersCount == 0 {
		errStr += "\nNo parameters found for this ErrorFunctionCall instance\n"
	}

	return errStr
}

```



# An Earlier Solution

This earlier version of ***errpref*** or Error Prefix project endeavored to provide better function or method documentation in error messages returned by **Go** functions. Two types have been defined for this purpose: ***ErrPref*** and ***ErrPrefixDto***. Both of these types are designed to receive function chain information, format it and return the formatted strings for inclusion in error messages.

From a usage standpoint, ***ErrPref*** and ***ErrPrefixDto*** collect information in string form at each method or function in the execution chain. If an error occurs, the user is responsible for extracting the collected information from ***ErrPref*** or ***ErrPrefixDto*** and integrating that information into the returned error message. While these types are designed to collect error prefix and error context information, they can in fact be used to collect any required informational or error related messaging in string format. This could be error numbers, execution or debugging information or virtually any text specified by the developer. In addition, despite the name 'error prefix', there is no requirement that this information be attached to the beginning of an error message. As a practical matter, the information could be injected as a string anywhere in the returned error message or informational message display.

### Wait a Minute - Let's Think About This

So, if you only take a quick look at the ***errpref*** package you might be tempted to say, **"Whoopee - all you've done is concatenate a series of strings!"**

That's what I thought, at first. However, there is the problem of delimiting error prefixes and error context information plus the problem of determining how much information to show on each line, plus the problem of receiving error prefix information from external sources in various formats - and the list goes on. 

With types ***ErrPref*** and  ***ErrPrefixDto***, the user may configure the line length which will determine whether a new line character (**\n**) or in-line delimiter string (**[SPACE]-[SPACE]**) is used to separate error prefix information.

***ErrPrefixDto*** utilizes a series of interfaces to receive and exchange data with external sources. In addition, ***ErrPrefixDto*** implements a variable or custom string delimiter feature which allows the user to process error prefix strings received from external sources in a variety of formats.

To summarize: If you are serious about error management and tracking, the ***errpref*** package might be worth a second look.



## What *errpref* Does and Doesn't Do

The ***errpref*** package generates text strings which can be populated with error prefix and error context information. It does NOT create error messages. Error messages are still created by the user with the standard tools provided by the **Go Programming Language**.  

***errpref*** is used to generate the error prefix, function chain and error context information. After that, it is up to the user to add this information to the error message. 

```go
	if inComingFormatterCurrency == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'inComingFormatterCurrency' is"+
			" a 'nil' pointer!\n",
			errPrefixDto.String())
		return err
	}

```

In the example above, ***errPrefixDto*** is an instance of the ***errpref*** type ***ErrPrefixDto*** previously configured with a series of method names documenting the execution path which led to this point in the code. As the example shows, the user is still responsible for configuring the error message in its entirety. 

While the ***errpref*** package was specifically designed to document function or method chains of execution, the example above shows that, as a practical matter, the user is free to position this text anywhere in the error message. Also, this text is not reserved exclusively for error messages. It could be used in informational messages or any other type of text display.

The typical error message decorated with error prefix and context information will typically look something like this:

```tex
main()-mainTest004()
TestFuncsDto.TestAlphaDto001()
 :  A->B
TestFuncDtoAlpha01.Tx1DoSomething()
testFuncDtoAlpha02.tx2DoSomethingElse()
testFuncDtoAlpha03.tx3DoAnything()
testFuncDtoAlpha04.tx4DoNothing()
 :  A/B==4
testFuncDtoAlpha05.tx5DoSomethingBig()
 :  A->B
testFuncDtoAlpha06.tx6TryGivingUp()
 :  A/B = C B==0
Example Error: An Error Occurred! Something bad.
Like Divide by Zero
```

This example error message was taken from the ***errpref*** example application, [errorPrefixExamples](https://github.com/MikeAustin71/errorPrefixExamples).



## Definition of Terms

For our purposes, error prefix information consists of two elements: 

1. One or more function or method names listed in sequence of code execution. 
2. An optional error context string associated with a specific function or method name.

### Error Prefix

**Error Prefix** text is designed to be configured at the beginning of error messages and is most often used to document the thread of code execution by listing the calling sequence for a specific list of functions and/or methods. As a previously stated, there is no hard and fast requirement that it be attached to the beginning of an error message. The user is free to inject this text string where ever it is most useful.

### Error Context

**Error Context** strings are designed to provide additional information about the function or method identified by the
associated **Error Prefix** text. Typical context information might include variable names, variable values, error numbers or details about function execution.

An Error Context cannot be created as a stand-alone object. It is always paired with an Error Prefix. Error Context strings are optional, but when created they are always associated with an Error Prefix string.



## Getting Started

### Two Types To Choose From

Currently, the ***errpref*** package provides two types offering error prefix formatting services: ***ErrPrefixDto*** and ***ErrPref***. Type ***ErrPref*** offers basic error prefix formatting while type ***ErrPrefixDto*** provides a much wider range of formatting and data transmission capabilities.

### ErrPrefixDto - A Full Featured Solution

The Error Prefix Data Transfer Object, ***ErrPrefixDto***, offers the same services as type, ***ErrPref***, but is packaged in a different architecture. While ***ErrPref*** methods receive a string and instantly return a formatted string, ***ErrPrefixDto*** encapsulates error prefix information in an internal array of Error Prefix Information objects (***ErrorPrefixInfo***). Strings are only created when the type's ***String()*** or ***StrMaxLineLen()*** methods are called. Instances of ***ErrPrefixDto*** are designed to be passed as input parameters to subsidiary methods. The act of passing ***ErrPrefixDto*** objects to successive functions in the execution sequence effectively builds documentation of the execution function chain. 

The ***String()*** method has a value receiver.  Virtually all other ***ErrPrefixDto*** methods have pointer receivers.

#### Public Facing Methods
A public method may receive error prefix information in a variety of formats from sources both internal and external to your application. To date, the best use scenario for a Public Facing Function or Method follows this pattern:

``` go
func (numStrBasic *NumStrBasic) GetCountryFormatters(
  fmtCollection *FormatterCollection,
  countryCulture CountryCultureId,
  errorPrefix interface{}) error {

  var ePrefix *ErrPrefixDto
  var err error

  ePrefix,
    err = ErrPrefixDto{}.NewIEmpty(
    errorPrefix,
    "NumStrBasic.GetCountryFormatters()",
    "")

  if err != nil {
    return err
  }

  return numStrBasicQuark{}.ptr().
    getCountryFormatters(
      fmtCollection,
      countryCulture,
      ePrefix)
}
```

By calling ***ErrPrefixDto{}.NewIEmpty()*** with an empty interface, functions can pass error prefix information in any of 10-different formats:

  1. **nil** - A nil value is valid and generates an empty collection of error prefix and error context information.
  2. **Stringer** - The Stringer interface from the '**fmt**' package. This interface has only one method:

``` go
    	type Stringer interface {
    		String() string
    	}
```

  3. **string** - A string containing error prefix information. For maximum effectiveness, this should be delimited with the standard [String Delimiters](#string-formatting-conventions).
  5. **\[\]string** - A one-dimensional slice of strings containing error prefix information. Typically, the one-dimensional slice contains only function or method names.
  1. **\[\]\[2\]string** - A two-dimensional slice of strings containing error prefix and error context information. slice\[x\]\[0\] holds function or method names. slice\[x\]\[1\] holds associated error context information.
  2. ***strings.Builder*** - An instance of strings.Builder. Error prefix information extracted from ***strings.Builder*** will be imported into the new returned instance of ***\*ErrPrefixDto***.
  3. ***\*strings.Builder*** - A pointer to an instance of strings.Builder.  Extracted error prefix information will be imported into the new returned instance of ***\*ErrPrefixDto***.
  4. ***ErrPrefixDto*** - An instance of ***ErrPrefixDto***. The ***ErrorPrefixInfo*** collection extracted from this object will be copied to the new returned instance of ***\*ErrPrefixDto***.
  6.  **\*ErrPrefixDto** - A pointer to an instance of **ErrPrefixDto**. The ***ErrorPrefixInfo*** collection extracted from this object will be copied to the new returned instance of  ***\*ErrPrefixDto***.
  7.   **IBasicErrorPrefix** - An interface to a method generating a two-dimensional slice of strings containing error prefix and error context information.

``` go
             type IBasicErrorPrefix interface {
               GetEPrefStrings() [][2]string
             }

```

The Method Signature for ***ErrPrefixDto{}.NewIEmpty()*** is shown below:

``` go
func (ePrefDto ErrPrefixDto) NewIEmpty(
	iEPref interface{},
	newErrPrefix string,
	newErrContext string) (
	*ErrPrefixDto,
	error)
```

This pattern allows for use of the **nil** value. This means that parameter ***iEPref*** will accept a **nil** value. If no error prefix information is present or required, just pass a **nil** value for the ***iEPref*** parameter.

Finally, notice how the currently executing method name (***NumStrBasic.GetCountryFormatters()***) is added to the error prefix method chain:

``` go
	ePrefix,
		err = ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrBasic.GetCountryFormatters()",
		"")
```

The final empty string parameter is optional and could be used to add error context information.



#### Internal or Private Methods

In the **Public Facing Methods** example, above, method ***GetCountryFormatters()*** makes a call to the  internal or private method, ***getCountryFormatters()***. Calls like this one to a supporting or subsidiary method usually pass a pointer to an instance of ***ErrPrefixDto***.  To date, the best use scenario for calls to subsidiary methods follows this pattern:

``` go

func (nStrBasicQuark numStrBasicQuark) getCountryFormatters(
	fmtCollection *FormatterCollection,
	countryCulture CountryCultureId,
	errPrefix *ErrPrefixDto) (
	err error) {
	
    ePrefix,
	err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"numStrBasicQuark.getCountryFormatters()",
		"")

	if err != nil {
		return err
	}

	if fmtCollection == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fmtCollection' is "+
			"a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

    ...
```

Notice how the function name is added to the error prefix chain:

``` go
	ePrefix,
	err := ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefix,
		"numStrBasicQuark.getCountryFormatters()",
		"")
```

This is the method signature for ***NewFromErrPrefDto()***:

```
func (ePrefDto ErrPrefixDto) NewFromErrPrefDto(
	dto *ErrPrefixDto,
	newErrPrefix string,
	newErrContext string) (
	newErrPrefDto *ErrPrefixDto,
	err error)
```

This method makes provision for the inclusion of an optional error context string.

Also, notice that this pattern allows for use of the **nil** value for parameter, ***dto***. If no error prefix information is present or required, just pass a **nil** parameter value.

This pattern provides a separate function chain string for each method. This architecture allows for multiple calls from parent methods without adding unnecessary and irrelevant text to the function chain. If an error occurs, only the relevant error prefix and error context information will be returned in the formatted error message.



#### Error Context Examples

##### Example-1 Error Flow

Recall that **Error Context** strings are designed to provide additional information about the function or method identified by the associated **Error Prefix** text. Typical context information might include variable names, variable values and additional details on function execution.

In this example, a function chain is built by calls to multiple levels of the code hierarchy.  The final call to method **Tx3.DoSomething()** triggers an error thereby returning the names of all methods in the call chain plus error context information associated with those methods.

``` go
func(tx1 *Tx1) Something() {

  ePrefDto := ErrPrefixDto{}.New()

  ePrefDto.SetEPref("Tx1.Something()")

  tx2 := Tx2{}

  err := tx2.SomethingElse(&ePrefDto)

  if err !=nil {
    fmt.Printf("%v\n",
    err.Error())
    return
  }

}

func(tx2 *Tx2) SomethingElse(ePrefDto *ErrPrefixDto) error {

	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"Tx2.SomethingElse()",
		"")

	if err != nil {
		return err
	}

  	tx3 := Tx3{}

  	err := tx3.DoSomething(ePrefix)

  	if err !=nil {
    	return err
  	}
}

func(tx3 *Tx3) DoSomething(ePrefDto *ErrPrefixDto) error {

   // Add error prefix and error context
    // information.
	ePrefix,
		err := ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"Tx3.DoSomething()",
		"A=B/C C= 0")

	if err != nil {
		return err
	}

    .... bad code ....

    if isDivideByZero {
      return fmt.Errorf("%v\n" +
      "I divided by zero and got this error.\n",
      ePrefix.String())
    }
}


```

When this error is returned up the function chain and finally printed out, the text will look like this:

``` text
 Tx1.Something() - Tx2.SomethingElse()
 Tx3.DoSomething() : A=B/C C= 0
 I divided by zero and got this error.

```

##### Example-2 Calling Subsidiary Functions

See the two examples of ***ErrPrefixDto.XCpy()*** included in the code shown below. **ErrPrefixDto.XCpy()** sends a deep copy of the original  ***ePrefix*** modified with error prefix information. The original ***ePrefix*** instance is **NOT** modified.

```go
func (stdLine *TextLineSpecStandardLine) AddTextField(
	iTextField ITextFieldSpecification,
	errorPrefix interface{}) (
	lastIndexId int,
	err error) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lastIndexId = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.AddTextField()",
		"")

	if err != nil {
		return lastIndexId, err
	}

	if iTextField == nil {
		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'iTextField' is 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

    // In this example, 'XCpy()' is used to 
    // send error context info to a subsidiary
    // method. The original instance of 'ePrefix'
    // is not unchanged. A deep copy of 'ePrefix'
    // is submitted with the error context info.
	err = iTextField.IsValidInstanceError(
		ePrefix.XCpy("iTextField"))

	if err != nil {
		return lastIndexId, err
	}

	var newTextField ITextFieldSpecification

    // In this example, 'XCpy()' is used to 
    // send error context info to a subsidiary
    // method. The original instance of 'ePrefix'
    // is not unchanged. A deep copy of 'ePrefix'
    // is submitted with the error context info.
	newTextField,
		err = iTextField.CopyOutITextField(
		ePrefix.XCpy("iTextField->newTextField"))

	if err != nil {
		return lastIndexId, err
	}

	stdLine.textFields = append(stdLine.textFields,
		newTextField)

	lastIndexId = len(stdLine.textFields) - 1

	return lastIndexId, err
}

```



#### Customizing Input and Output String Delimiters

Unlike type ***ErrPref***, ***ErrPrefixDto*** allows users to configure the string delimiters used to parse input strings containing error prefix information received from external sources. This facilitates the exchange of error prefix information received from outside sources by ensuring that error prefix string components are properly separated and identified during the parsing operation. In addition, users may also specify the string delimiters used to combine and join error prefix components for generation and formatting of output text incorporated in error messages.

Internally, ***ErrPrefixDto*** maintains two types of string delimiters, **Input String Delimiters** and **Output String Delimiters**.

**Input String Delimiters** are used to parse raw string values containing error prefix and error context information received as input parameters from external sources. Methods performing this type of operation include:

```go
ErrPrefixDto.NewEPrefOld
ErrPrefixDto.NewFromStrings()
ErrPrefixDto.NewFromStrings()
ErrPrefixDto.NewIEmptyWithDelimiters()
ErrPrefixDto.SetEPrefOld()
ErrPrefixDto.XEPrefOld()
ErrPrefixDto.ZEPrefOld()
```

 Output String Delimiters are used by ***ErrPrefixDto*** instances to join or concatenate individual error prefix and error context components to form presentation text for output and use in preparation of error message strings. Methods performing this type of operation are: 

```go
ErrPrefixDto.String()
ErrPrefixDto.StrMaxLineLen()
```

 Initially, both the Input and Output String Delimiters for any given ***ErrPrefixDto*** instance are set to the system default values. This means that if the Input and Output String Delimiters were not directly configured by the user, the system default string delimiters are applied.

The system defaults for both Input and Output String Delimiters are listed as follows:  

```go
 New Line Error Prefix Delimiter  = "\n"
 In-Line Error Prefix Delimiter   = " - "
 New Line Error Context Delimiter = "\n :  "
 In-Line Error Context Delimiter  = " : "
```

if required, the Input and Output String Delimiters can always be reset to system default values by calling method:

```go
func (ePrefDto *ErrPrefixDto) SetStrDelimitersToDefault()
```



However, users do have the option of setting custom Input and Output String Delimiters using methods:

```go
func (ePrefDto *ErrPrefixDto) SetInputStringDelimiters(
	inputStrDelimiters ErrPrefixDelimiters,
	ePrefix string) error

func (ePrefDto *ErrPrefixDto) SetOutputStringDelimiters(
	outputStrDelimiters ErrPrefixDelimiters,
	ePrefix string) error

```

In terms of output text, a variety of interesting and creative display effects can be achieved through manipulating the Output String Delimiters. However, when configuring Output String Delimiters it may be necessary to coordinate the Maximum Text Line Length which is initially set to a default of 40-characters.  Configuring Output String Delimiters should be coupled with managing the Maximum Text Line Length discussed below.

The current settings for String Delimiters can be monitored using methods:

```go
 ErrPrefixDto.GetStrDelimiters()
 ErrPrefixDto.GetInputStringDelimiters()
 ErrPrefixDto.GetOutputStringDelimiters()
```



#### Maximum Text Line Length

Users have the option of setting the Maximum Text Line Length for error prefix strings. The default Maximum Text Line Length is 40-characters. The Minimum Text Line Length is 10-characters.

The following error prefix text display demonstrates the default 40-character maximum line length.

```tex
main()-mainTest004()
TestFuncsDto.TestAlphaDto001()
 :  A->B
TestFuncDtoAlpha01.Tx1DoSomething()
testFuncDtoAlpha02.tx2DoSomethingElse()
testFuncDtoAlpha03.tx3DoAnything()
testFuncDtoAlpha04.tx4DoNothing()
 :  A/B==4
testFuncDtoAlpha05.tx5DoSomethingBig()
 :  A->B
testFuncDtoAlpha06.tx6TryGivingUp()
 :  A/B = C B==0
Example Error: An Err
```



The next example shows the same error prefix information displayed with a custom Maximum Text Line Length of 70-characters.

```tex
main()-mainTest005() - TestFuncsDto.TestAlphaDto002()
TestFuncDtoAlpha01.Tx1DoSomething()
testFuncDtoAlpha02.tx2DoSomethingElse()
testFuncDtoAlpha03.tx3DoAnything()
testFuncDtoAlpha04.tx4DoNothing() : A/B==4
testFuncDtoAlpha05.tx5DoSomethingBig() : A->B
testFuncDtoAlpha06.tx6TryGivingUp() : A/B = C B==0
Example Error: An Error Occurred! Something bad.
Like Divide by Zero!
```

The Maximum Text Line Length is monitored and controlled using methods:

-  ***ErrPrefixDto.SetMaxTextLineLen()***
-  ***ErrPrefixDto.GetMaxTextLineLen()***
-  ***ErrPrefixDto.SetMaxTextLineLenToDefault()*** 

Be advised that the minimum text line Length is 10-characters. Any attempt to set the Maximum Text Line Length to a value less than 10-characters will result in reinstatement of the default Maximum Text Line Length.

This example error messages shown above were taken from the ***errpref*** example application, [errorPrefixExamples](https://github.com/MikeAustin71/errorPrefixExamples).



#### Left Margin Feature

The default left margin length for text returned by the method ***ErrPrefixDto.String()*** is zero (0). In other words, by default, there is no left margin.

Example error prefix information text display with left margin length of zero:

```tex
Tx1.Something() - Tx2.SomethingElse()
Tx3.DoSomething() - Tx4() - Tx5()
Tx6.DoSomethingElse()
Tx7.TrySomethingNew()
 :  something->newSomething
Tx8.TryAnyCombination()
Tx9.TryAHammer() : x->y - Tx10.X()
Tx11.TryAnything() - Tx12.TryASalad()
Tx13.SomeFabulousAndComplexStuff()
Tx14.MoreAwesomeGoodness()
 :  A=7 B=8 C=9
```

The following ***ErrPrefixDto*** methods allow the user to control both the left margin length, and the character used to populate the left margin:

1. ***ErrPrefixDto.SetLeftMarginLength()*** 
2. ***ErrPrefixDto.SetLeftMarginChar()***
3. ***ErrPrefixDto.GetLeftMarginLength()***
4. ***ErrPrefixDto.GetLeftMarginChar()***

In the next example, the left margin length is set to three (3), and the left margin character is set to the empty space character (' ') or ASCII 0x20:

```tex
   Tx1.Something() - Tx2.SomethingElse()
   Tx3.DoSomething() - Tx4() - Tx5()
   Tx6.DoSomethingElse()
   Tx7.TrySomethingNew()
    :  something->newSomething
   Tx8.TryAnyCombination()
   Tx9.TryAHammer() : x->y - Tx10.X()
   Tx11.TryAnything() - Tx12.TryASalad()
   Tx13.SomeFabulousAndComplexStuff()
   Tx14.MoreAwesomeGoodness()
    :  A=7 B=8 C=9
```



In the final example, the left margin length is set to three (3), and the left margin character is set to the asterisk character ('*'): 

```tex
***Tx1.Something() - Tx2.SomethingElse()
***Tx3.DoSomething() - Tx4() - Tx5()
***Tx6.DoSomethingElse()
***Tx7.TrySomethingNew()
*** :  something->newSomething
***Tx8.TryAnyCombination()
***Tx9.TryAHammer() : x->y - Tx10.X()
***Tx11.TryAnything() - Tx12.TryASalad()
***Tx13.SomeFabulousAndComplexStuff()
***Tx14.MoreAwesomeGoodness()
*** :  A=7 B=8 C=9
```



#### Leading And Trailing Text Strings

Users have the option of adding leading or trailing text strings to the error prefix display. Leading Text Strings will be added to the beginning of the text display while Trailing Text Strings will be added to the end of the text display. 

Leading and Trailing Text Strings may be comprised of any valid string of characters including new lines ('\n'), tabs ('\t') and line separators.

The addition of Leading and Trailing Text Strings is controlled through the following methods:

- ***ErrPrefixDto.SetLeadingTextStr(***)
- ***ErrPrefixDto.SetTrailingTextStr()***
- ***ErrPrefixDto.GetLeadingTextStr()***
- ***ErrPrefixDto.GetTrailingTextStr()***
- ***ErrPrefixDto.ClearLeadingTextStr()***
- ***ErrPrefixDto.ClearTrailingTextStr()***



Consider the following example:

```go
	ePDto := errpref.ErrPrefixDto{}.New()

	maxLineLen := 50

	ePDto.SetMaxTextLineLen(maxLineLen)

	var twoDSlice [][2]string

	twoDSlice = make([][2]string, 14)

	twoDSlice[0][0] = "Tx1.Something()"
	twoDSlice[0][1] = ""

	twoDSlice[1][0] = "Tx2.SomethingElse()"
	twoDSlice[1][1] = ""

	twoDSlice[2][0] = "Tx3.DoSomething()"
	twoDSlice[2][1] = ""

	twoDSlice[3][0] = "Tx4()"
	twoDSlice[3][1] = ""

	twoDSlice[4][0] = "Tx5()"
	twoDSlice[4][1] = ""

	twoDSlice[5][0] = "Tx6.DoSomethingElse()"
	twoDSlice[5][1] = ""

	twoDSlice[6][0] = "Tx7.TrySomethingNew()"
	twoDSlice[6][1] = "something->newSomething"

	twoDSlice[7][0] = "Tx8.TryAnyCombination()"
	twoDSlice[7][1] = ""

	twoDSlice[8][0] = "Tx9.TryAHammer()"
	twoDSlice[8][1] = "x->y"

	twoDSlice[9][0] = "Tx10.X()"
	twoDSlice[9][1] = ""

	twoDSlice[10][0] = "Tx11.TryAnything()"
	twoDSlice[10][1] = ""

	twoDSlice[11][0] = "Tx12.TryASalad()"
	twoDSlice[11][1] = ""

	twoDSlice[12][0] = "Tx13.SomeFabulousAndComplexStuff()"
	twoDSlice[12][1] = ""

	twoDSlice[13][0] = "Tx14.MoreAwesomeGoodness()"
	twoDSlice[13][1] = "A=7 B=8 C=9"

	ePDto.SetEPrefStrings(twoDSlice)

	leadingText := "\n" +
		strings.Repeat("-", maxLineLen)

	trailingText := strings.Repeat("-", maxLineLen) +
		"\n"

	ePDto.SetLeadingTextStr(leadingText)
	ePDto.SetTrailingTextStr(trailingText)
	ePDto.SetIsLastLineTermWithNewLine(true)

	outputStr := fmt.Sprintf("%v"+
		"Error: Divide by Zero!",
		ePDto.String())

	fmt.Printf(outputStr)

```



The the example above the final ***outputStr*** will be formatted as:

```tex
 --------------------------------------------------
 Tx1.Something() - Tx2.SomethingElse()
 Tx3.DoSomething() - Tx4() - Tx5()
 Tx6.DoSomethingElse()
 Tx7.TrySomethingNew() : something->newSomething
 Tx8.TryAnyCombination() - Tx9.TryAHammer() : x->y
 Tx10.X() - Tx11.TryAnything() - Tx12.TryASalad()
 Tx13.SomeFabulousAndComplexStuff()
 Tx14.MoreAwesomeGoodness() : A=7 B=8 C=9
 --------------------------------------------------
 Error: Divide by Zero!

```



#### Text Display On/Off Switch

***ErrPrefixDto*** implements a feature which allows the user to turn off the generation of output strings containing error prefix information.  The typical means of injecting output text into an error message follows this example:

```go
    if isDivideByZero {
      return fmt.Errorf("%v\n" +
      "I divided by zero and got this error.\n",
      ePrefDto.String())
    }

```

In this example. the call to ***ePrefDto.String()*** will generate a string containing error prefix information which is then inserted into the returned error message. However, ***ErrPrefixDto*** now provides a mechanism for "Turning Off" this error prefix string generation. Users control this "On/Off" switch through calls to method:

```go
func (ePrefDto *ErrPrefixDto) SetTurnOffTextDisplay(
	turnOffTextDisplay bool)
```

Calling this method with parameter ***turnOffTextDisplay*** set to **true** will ensure the future calls to ***ErrPrefixDto.String()*** and ***ErrPrefixDto.StrMaxLineLen()*** will return an empty string. Assuming Text Display was turned-off, the call to ***ePrefDto.String()*** in the "Divide By Zero" example above, would return an empty string. The net result is that no error prefix information would be injected in to the error message in this example.



#### ***Is Last Text Line Terminated With New Line***

By default, the last line of error prefix strings returned by the methods ***ErrPrefixDto.String()***  and ***ErrPrefixDto.StrMaxLineLen()*** **ARE NOT terminated with a new line character ('\n')**. Consider the following example:

```go
    if isDivideByZero {
      return fmt.Errorf("%v\n" + // <------ User manually inserted new character after error prefix string.
      "I divided by zero and got this error.\n",
      ePrefDto.String())
    }

```

Here, the user is manually inserting a new line character after the error prefix string because the error prefix does NOT include a new line character.  Therefore, when this error is returned up the function chain and finally printed out, the text will look like this:

``` text
 Tx1.Something() - Tx2.SomethingElse()
 Tx3.DoSomething() : A=B/C C== 0.0    <------ User added a new line character
 I divided by zero and got this error.

```

If the last line of error prefix strings should be automatically terminated with a new line character ('\n'),  call the following method:

```go
func (ePrefDto *ErrPrefixDto) SetIsLastLineTermWithNewLine(
	isLastLineTerminatedWithNewLine bool)
```

Setting input parameter ***isLastLineTerminatedWithNewLine*** to **true** will effectively add a new line character ('\n') to all error prefix strings returned by methods:

```go
func (ePrefDto ErrPrefixDto) String() string

func (ePrefDto *ErrPrefixDto) StrMaxLineLen(
	maxLineLen int) string
```

In the 'Divide by Zero' example, the user no longer needs to manually add the new line character:

```
 if isDivideByZero {
      return fmt.Errorf("%v" + <------ // User did NOT insert new character after error prefix string.
      "I divided by zero and got this error.\n",
      ePrefDto.String())
    }


```

When this error is returned up the function chain and finally printed out, the text will look like this:

``` tex
 Tx1.Something() - Tx2.SomethingElse()
 Tx3.DoSomething() : A=B/C C== 0.0 <--- New line character automatically added to end of error prefix string.
 I divided by zero and got this error.

```





### ErrPref - Quick And Simple Solution

Type ***ErrPref*** is designed to format error prefix text strings for use in error messages. ***ErrPref*** is simple, lightweight, easy to use and seems to work in a variety of situations. The concept is straight forward, "**Put raw text strings in - Get formatted error prefix information out**".

The ***ErrPref*** methods do not use pointer receivers. All receivers are value receivers.

The ***ErrPref*** type does not store error prefix information like the ***ErrPrefixDto*** type. Again, the concept is, "**Put raw strings in - Get formatted error prefix strings out**".

#### Example Usage Summary

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

#### String Formatting Conventions

When using the methods provided by type ***ErrPref***, incoming error prefix strings are parsed for function names, method names and error context information based on specific string delimiters. For type ***ErrPref***, these string delimiters are standardized, fixed and not subject to customization by the user. The system default input and output string delimiters are defined as follows:

##### Error Prefix Delimiters

Error Prefix Elements are delimited by one of two delimiters: 

- **New Line Delimiter** = "**\n**" Delimits function/method names on a single line.
- **In-Line Delimiter** = " **-** " (**[SPACE]-[SPACE]**) Delimits multiple function/method names on the same line.

##### Error Context Delimiters

Associated Error Context Sub-Elements likewise have two delimiters: 

- **New Line Delimiter String** = "**\n[SPACE]:[SPACE]\[SPACE]**" Delimits error context information on a single line.
- **In-Line Delimiter String** = "**[SPACE]:[SPACE]**" Delimits error context information displayed with function/method name on same line.

If variable or custom input and output string delimiters are required, use type ***ErrPrefixDto***.





## Exchanging Error Prefix Information with User Defined Types



### *IBuilderErrorPrefix*

Custom user defined types supporting the ***IBuilderErrorPrefix*** interface will be able to receive data from and insert data into instances of ***ErrPrefixDto***. The following ***ErrPrefixDto*** methods interoperate with the ***IBuilderErrorPrefix*** interface:

1. ***ErrPrefixDto.CopyInFromIBuilder()***
2. ***ErrPrefixDto.CopyOutToIBuilder()***
3. ***ErrPrefixDto.NewIEmpty()***
4. ***ErrPrefixDto.NewIEmptyWithDelimiters()***
5. ***ErrPrefixDto.SetIBuilder()***

Taken together, these methods facilitate the import and export of error prefix and context information between ***ErrPrefixDto*** and user defined types implementing the ***IBuilderErrorPrefix*** interface.

The ***IBuilderErrorPrefix*** interface is defined as follows:

```go
type IBuilderErrorPrefix interface {
	GetEPrefStrings() [][2]string

	SetEPrefStrings(twoDStrArray [][2]string)

	String() string
}
```



### IBasicErrorPrefix

Custom user defined types supporting the ***IBasicErrorPrefix*** will be able to transmit error prefix information to an instance of ***ErrPrefixDto***.

```
type IBasicErrorPrefix interface {
	GetEPrefStrings() [][2]string
}
```

The two-dimensional string array returned by GetEPrefStrings() is intended to be populated with error prefix and error context string pairs.

***ErrPrefixDto*** methods specifically configured to handle and process ***IBasicErrorPrefix*** objects are listed as follows:

1. ***ErrPrefixDto.NewIEmpty()***
2. ***ErrPrefixDto.NewIEmptyWithDelimiters()***
3. ***ErrPrefixDto.SetIBasic()***



### Custom Input String Delimiters

Type ***ErrPrefixDto*** allows users to configure input string delimiters used to parse input strings containing error prefix information received from external sources. This facilitates the exchange of error prefix information received from outside sources by ensuring that error prefix string components are properly separated and identified during the parsing operation. If the string delimiters for incoming strings is known, ***ErrPrefixDto*** input string delimiters can be coordinated and configured to correctly parse the incoming error prefix and error context information.

**Input String Delimiters** are used to parse raw string values containing error prefix and error context information. These parsing operations are routinely handled through the following methods:

1. ***ErrPrefixDto.NewEPrefOld()***
2. ***ErrPrefixDto.NewFromStrings()***
3. ***ErrPrefixDto.NewIEmptyWithDelimiters()***
5. ***ErrPrefixDto.SetEPrefOld()***
6. ***ErrPrefixDto.XEPrefOld()***
7. ***ErrPrefixDto.ZEPrefOld()***



## Usage Examples

### Test Code

Test Code is located in the **errpref** directory of the source code repository. All files beginning with the letters "**zzzt_**" and ending with "**_test.go**" contain test code. The **errpref** directory is located here: [Test Code](https://github.com/MikeAustin71/errpref)

Test Code provides many usage examples for types ***ErrPref*** and ***ErrPrefixDto***.

### Example Application

Additional code examples can be found in the Error Prefix Examples Application located at [https://github.com/MikeAustin71/errorPrefixExamples](https://github.com/MikeAustin71/errorPrefixExamples).  This application also contains a "Concurrency" example.



## User Package Import and Configuration

There are two basic requirements:

1. **The User must add `import` statements to those source code files incorporating `errPref` functions, types, methods, constants etc.**
2. **The User MUST configure the local `go.mod` file to reference the `errPref` package.**

### Required ***Golang*** Source Code File Import Statement

```go
// For the later errpref versions (Version 2)

import (

	ePref "github.com/MikeAustin71/errpref/v2"
)

// For earlier 'errpref' versions (Version 1)
import (

	ePref "github.com/MikeAustin71/errpref"
)


```

The import example above shows an alias of '**ePref**' which is optional.

### ***go.mod*** File Configuration

**Users** must also configure their local ***go.mod*** file. There are two options to accomplish this: 

#### Option 1: Go Get Command

```go
go get github.com/MikeAustin71/errpref/v2/@v2.0.1
```

​	-- or --

```go
go get github.com/MikeAustin71/errpref/@latest
```

##### How a User Employs `go get` (The Standard Workflow)

When a developer wants to use your `errpref` library in their own source code, they will follow these exact steps from their terminal:

##### Step 1: Initialize Their Own Project (If they haven't already)

The **User** creates their project folder and initializes their own ***`gomod`*** module:

```bash
go mod init mytestproject
```

At this point, their `go.mod` file is completely empty of dependencies.

##### Step 2: Run the Modern `go get` Command

Instead of opening `go.mod` and guessing the syntax, the user simply runs the command in their terminal:

```
go get github.com/MikeAustin71/errpref/v2@v2.0.1
```

##### Step 3: What Go Autonomously Does to `go.mod`

The moment that command executes, the Go toolchain does three things simultaneously:

1. **Downloads the source code** for `v2.0.1` into the local machine's global module cache.

2. **Modifies the user's `go.mod` file automatically**, adding the exact block you expected:

   ```
   require github.com/MikeAustin71/errpref/v2 v2.0.1
   ```

   

3. **Creates or updates a `go.sum` file** next to it, which stores the cryptographic security hashes of your code to ensure it's never altered downstream.

------

#### Option 2: ***go.mod*** Manual Configuration

If the ***go.mod*** file was previously created, simply add the following ***require*** statement to the existing ***go.mod*** file.

```bash
require github.com/MikeAustin71/errpref/v2 v2.0.1
```

###  

## External Dependencies

None. This software package is not dependent on any external module or package.



## Tests

Source code tests are located in the **errpref** directory of the source code repository. All files beginning with the letters "**zzzt_**" and ending with "**_test.go**" contain test code. The **errpref** directory is located here: [Test Code](https://github.com/MikeAustin71/errpref)

To run the test code, first review the command syntax in **`~/errprefops/technotes/HowToRunTests.md`** .

Test results are stored in the text file, **`~/errpref/zzzzz_tests.txt`**

All 500+ tests are currently passing including new tests for newly added type ***ErrorGenericFactory***.  



## OS Support

***errpref*** versions prior to **Version 2.0.1** were tested successfully on Windows 10 and Ubuntu 20.04 LTS.

***errpref*** **Version 2.0.1** has been tested on Windows 11. It should have no problems running on **Linux**, but as yet, <u>no tests have been run on **Linux** operating systems</u>.



## Version

The latest version is ***Version 2.0.1***.

This version was compiled and tested using using ***Go*** ***Version 1.27.3***.

For more details, see the [Release Notes](./releasenotes.md).



## ***errPref*** Source Code Repository

The source code repository for ***errPref*** is **github.com/MikeAustin71/errpref/** . This repository is used to distribute release versions.

The development repository is ***errprefops*** located at ***github.com/MikeAustin71/errprefops***



## Release Notes

 [Release Notes](./releasenotes.md)



## License

Use of this source code is governed by the (open-source) MIT-style license which can be found in the LICENSE file
located in this directory.

[MIT License](./LICENSE)



## Comments And Questions

Send questions or comments to:

``` text
    mike.go@paladinacs.net
```

