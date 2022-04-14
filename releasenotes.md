# *errpref* (Error Prefix) Release Notes Version 1.7.1

This version of ***errpref*** was compiled and tested using ***Go Version 1.18.1***.

This version supports ***Go*** modules.

## Version 1.7.1

#### Compiled with Go Version 1.18.1

This version of ***errpref*** was compiled with ***Go Version 1.18.1***. Currently ***errpref*** does not employ ***generics***.



#### Added New Method *ErrPrefixDto.XCpy()*

The new new method ***ErrPrefixDto.XCpy()*** is designed for use in calling subsidiary methods with error context information. 

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
    // is submitted with the new error context info.
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
    // is submitted with the new error context info.
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

## Version 1.7

##### Documentation Updates
Modified Best Practices documentation for implementing ErrPrefixDto in calls to **Internal or Private Methods**. See the [REAEDME File](README.md).



##### ErrPrefixDto

1. Added new method ***ErrPrefixDto{}.NewFromErrPrefDto()***
   This method can be used to reduce the lines of code required to implement ***ErrPrefixDto*** objects in **Internal or Private Methods**. See the [REAEDME File](README.md).
   
   ```go
   func (ePrefDto ErrPrefixDto) NewFromErrPrefDto(
      dto *ErrPrefixDto,
      newErrPrefix string,
      newErrContext string) (
      newErrPrefDto *ErrPrefixDto,
      err error)
   ```
   
2. Added new method ***ErrPrefixDto.DeleteLastErrPrefix()***
   This method deletes the last Error Prefix Information object in the current ***ErrPrefixDto*** collection.
   
3. Added new method ***ErrPrefixDto.GetLastErrPrefix()***
   This method returns a deep copy of the last Error Prefix Information object in the current ***ErrPrefixDto*** collection.
   
4. Added new method ***ErrPrefixDto.ReplaceLastErrPrefix()***
   This method deletes and replaces the Last Error Prefix Information object in the current ***ErrPrefixDto*** collection with new error prefix and context information.



##### New Leading and Trailing Text Feature

***ErrPrefixDto*** has added a new feature which allows the user to configure leading and trailing strings. These strings will be added to the beginning and ending of error prefix text displays. The leading and trailing text strings may be comprised of any string of characters including new lines ('\n'), tabs ('\t') and line separators.

For more information see the documentation on the following methods:
1. **ErrPrefixDto.SetLeadingTextStr()**
2. ***ErrPrefixDto.SetTrailingTextStr(***)
3. ***ErrPrefixDto.GetLeadingTextStr()***
4. ***ErrPrefixDto.GetTrailingTextStr()***
5. ***ErrPrefixDto.ClearLeadingTextStr()***
6. ***ErrPrefixDto.ClearTrailingTextStr()***



##### Enforcing Minimum Error Prefix Line Length

Users have always had the ability to set the maximum line length for error prefix text displays. As of this release, the minimum line length will be enforced by methods in the ***ErrPrefixDto*** and ***ErrPref*** types. The Minimum Error Prefix Line Length is '10' characters. Depending on the function called, any attempt to set a value less than the Minimum Error Prefix Line Length will cause a reset to the default Line Length value (40-characters), or generate an error. For more information see the documentation on the following methods:
1. ***ErrPref.SetMaxErrPrefTextLineLength()***
2. ***ErrPrefixDto.SetMaxTextLineLen()***
3. ***ErrPrefixDto.GetMaxTextLineLen()***
4. ***ErrPrefixDto. SetMaxTextLineLenToDefault()***



##### Added More Tests

Code coverage for tests is now at 96%.




## Version 1.6.1

#### Changes
Documentation updates.

## Version 1.6.0

#### Changes

##### Directory Structure: Development Environment and Package Distribution

This version marks a paradigm change in organization of the ***errpref*** project. 

Moving forward, all development and testing will be conducted in the development environment ***errprefops*** located in software repository https://github.com/MikeAustin71/errprefops .  

Storage and distribution of the ***errpref*** software package will be processed through software repository https://github.com/MikeAustin71/errpref

##### ErrPrefixDto

1. Added new method ErrPrefixDto.SetIEmpty() - Sets the data values for the current ErrPrefixDto instance based on any one of 10-valid types passed through an empty interface. See [source code documentation](https://pkg.go.dev/github.com/MikeAustin71/errpref#ErrPrefixDto).

   ```go
   func (ePrefDto *ErrPrefixDto) SetIEmpty(
   	iEPref interface{},
   	callingMethodName string) error
   ```

2. Added new method ErrPrefixDto.Empty() - Reinitializes all member variables for the current ErrPrefixDto instance to their native zero values.

3. Added Left Margin Feature for error prefix strings. Users can now set both the left margin length, and the left margin character used to generate the left margin in error prefix strings returned by method ErrPrefixDto.String().

   - ErrPrefixDto.GetLeftMarginChar() rune
   - ErrPrefixDto.GetLeftMarginLength() int
   - ErrPrefixDto.SetLeftMarginChar(leftMarginCharacter rune)
   - ErrPrefixDto.SetLeftMarginLength(leftMarginLength int)

4. Added new method ErrPrefixDto.GetDelimiters(). This method returns the input and output string delimiters used to delimit error prefix and error context strings.

5. Added new method ErrPrefixDto.CopyOutToIBuilder(). This method takes an object implementing the IBuilderErrorPrefix interface and populates than object with error prefix and context information contained in the current instance of ErrPrefixDto. This method is used to facilitate the exchange of error prefix information with custom user developed types.

6. Added new method ErrPrefixDto.CopyInFromIBuilder(). This method takes an object implementing the IBuilderErrorPrefix interface and copies its error prefix and context information into the current instance of ErrPrefixDto.  This method is used to facilitate the exchange of error prefix information with custom user developed types. This method is used to facilitate the exchange of error prefix information with custom user developed types.

7. Added new method ErrPrefixDto.SetIBuilder(). Like ErrPrefixDto.CopyInFromIBuilder(), this method takes an object implementing the IBuilderErrorPrefix interface and copies its error prefix and context information into the current instance of ErrPrefixDto.  This method is used to facilitate the exchange of error prefix information with custom user developed types. This method is used to facilitate the exchange of error prefix information with custom user developed types.

8. Added new method ErrPrefixDto.IsValidInstanceError(). Returns an error if the current ErrPrefixDto instance is invalid.

9. Added new method ErrPrefixDto.IsValidInstance(). Returns a Boolean flag indicating whether the current ErrPrefixDto instance is valid.

10. Added new method ErrPrefixDto.NewFromStrings(). This method returns a new instance of ErrPrefixDto while allowing the user to set the input and output string delimiters:

    ```go
    func (ePrefDto ErrPrefixDto) NewFromStrings(
       oldErrPrefix string,
       newErrPrefix string,
       newErrContext string,
       inputStrDelimiters ErrPrefixDelimiters,
       outputStrDelimiters ErrPrefixDelimiters,
       ePrefix string) (ErrPrefixDto, error)
    ```

11. **Variable String Delimiter Feature** - Added methods to control input and output string delimiters. These methods effectively implement the Variable String Delimiter Feature. Users are now able to control the string delimiters used to parse input strings containing error prefix information as well as the string delimiters used to format output error prefix text for presentations and display.

    ```
    func (ePrefDto *ErrPrefixDto) SetInputStringDelimiters(
    	inputStrDelimiters ErrPrefixDelimiters,
    	ePrefix string) error
    	
    func (ePrefDto *ErrPrefixDto) SetOutputStringDelimiters(
    	outputStrDelimiters ErrPrefixDelimiters,
    	ePrefix string) error
    
    func (ePrefDto *ErrPrefixDto) GetStrDelimiters() (
    	inputStrDelimiters ErrPrefixDelimiters,
    	outputStrDelimiters ErrPrefixDelimiters)
    	
    func (ePrefDto *ErrPrefixDto) GetInputStringDelimiters() ErrPrefixDelimiters
    
    func (ePrefDto *ErrPrefixDto) GetOutputStringDelimiters() ErrPrefixDelimiters
    
    func (ePrefDto ErrPrefixDto) NewFromStrings(
    	oldErrPrefix string,
    	newErrPrefix string,
    	newErrContext string,
    	inputStrDelimiters ErrPrefixDelimiters,
    	outputStrDelimiters ErrPrefixDelimiters,
    	ePrefix string) (ErrPrefixDto, error)
    
    func (ePrefDto ErrPrefixDto) NewIEmptyWithDelimiters(
    	iEPref interface{},
    	newErrPrefix string,
    	newErrContext string,
    	inputStrDelimiters ErrPrefixDelimiters,
    	outputStrDelimiters ErrPrefixDelimiters,
    	ePrefix string) (
    	*ErrPrefixDto,
    	error)
    ```
    
    
    
12. Additional ***IBuilderErrorPrefix*** Interface Support

    ```go
    func (ePrefDto *ErrPrefixDto) CopyInFromIBuilder(
    	inComingIBuilder IBuilderErrorPrefix,
    	eMsg string) error
    	
    func (ePrefDto *ErrPrefixDto) CopyOutToIBuilder(
    	inComingIBuilder IBuilderErrorPrefix)
    	
    func (ePrefDto *ErrPrefixDto) SetIBuilder(
    	inComingIBuilder IBuilderErrorPrefix,
    	callingMethodName string) error
    ```
    
    
    
    

##### ErrPref

Added new method ErrPref.GetDelimiters(). This method returns the string delimiters used to delimit error prefix and error context strings.



## Version 1.5.2

Original Release Date: 2021-04-09 02:05:00 USA CST

Compiled and Tested: Go 1.16.3

This version supports ***Go*** modules.

#### Changes

1. Fixed link to source code documentation.
2. Added documentation providing more usage examples.
3. Added new method ErrPrefixDto.NewEPrefCtx()

## Version 1.5.1

Original Release Date: 2021-04-07 21:39:00 USA CST

Compiled and Tested: Go 1.16.3

This version supports ***Go*** modules.

#### Changes

1. Fixed Go Mod File to assign correct version.

## Version 1.5.0

Original Release Date: 2021-04-07 17:34:00 USA CST

Compiled and Tested: Go 1.16.3

This version supports ***Go*** modules.

#### Changes

1. Modified interface IErrorPrefix. Added Methods:
   - GetEPrefStrings() [][2]string
   - SetEPrefStrings(twoDStrArray [][2]string)

2. Added interface IBuilderErrorPrefix

3. Added interface IBasicErrorPrefix

4. Added methods on Type ErrPrefixDto
   - AddEPrefStrings()
   - GetEPrefStrings()
   - NewIBasicErrorPrefix()
   - NewIEmpty()
   - SetEPrefStrings()

## Version 1.4.0

Original Release Date: 2021-02-16 16:59:00 USA CST

Compiled and Tested: Go 1.15.8

This version supports ***Go*** modules.

#### Changes

1. Modified the interface IErrorPrefix. Removed references to ErrorPrefixInfo.

2. Added 'Z' methods which return an ErrPrefixDto by value.
   - ZCtx()
   - ZCtxEmpty()
   - ZEPref()
   - ZEPrefCtx()
   - ZEPrefOld()

3. Added tests

## Version 1.3.0

Original Release Date: 2021-02-14 23:38:00 USA CST

Compiled and Tested: Go 1.15.8

This version supports ***Go*** modules.

#### Changes

1. Added IErrorPrefix interface. 

2. Added ErrPrefixDto method XCtxEmpty.

3. Added ErrPrefixDto method XSetEmpty.

4. Changed behavior of ErrPrefixDto method SetCtx. Empty input strings will now delete last error context.

5. Changed behavior of ErrPrefixDto method XCtx. Empty input strings will now delete last error context.

6. Added ErrPrefixDto method NewFromIErrorPrefix.

## Version 1.2.0

Original Release Date: 2021-02-11 16:46:00 USA CST

Compiled and Tested: Go 1.15.8

This version supports ***Go*** modules.

#### Changes

1. Removed Maximum Line Length operation from ErrPrefixDto.String()
2. Added Ptr() method to ErrPrefixDto.
3. Added tests for ErrPrefixDto methods CopyIn() and CopyOut()
4. Added new method ErrPrefixDto.StrMaxLineLen().
5. Changed method name and signature for ErrPrefixDto.SetMaxTextLineLen()
6. Changed method name for ErrPrefixDto.SetMaxTextLineLenToDefault()
7. Added new method ErrPrefixDto.Copy()

## Version 1.1.0

Original Release Date:  2021-02-11 02:11:00 USA CST

Compiled and Tested: Go 1.15.8

This version supports ***Go*** modules.

### ErrPrefixDto

1. Converted method ErrPrefixDto.String() from a pointer receiver to a value receiver.

2. Added 'X' Methods:

   - XCtx() - Returns pointer to current instance.

   - XEPref() - Returns pointer to current instance.

   - XEPrefCtx() - Returns pointer to current instance.

   - XEPrefOld() - Returns pointer to current instance.

     



## Version 1.0.0 

Original Release Date:  2021-02-10 01:17:00 USA CST

Compiled and Tested: Go 1.15.8

This version supports ***Go*** modules.

Initial release of ***ErrPref***.