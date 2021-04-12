# ErrPref Release Notes Version 1.5.1

This version of ***ErrPref*** was compiled and tested using Go 1.16.3.

This version supports ***Go*** modules.

## Version 1.5.1

Original Release Date: 2021-04-07 21:39:00 USA CST

Compiled and Tested: Go 1.16.3

1. Fixed Go Mod File to assign correct version.

## Version 1.5.0

Original Release Date: 2021-04-07 17:34:00 USA CST

Compiled and Tested: Go 1.16.3

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

1. Added IErrorPrefix interface. 

2. Added ErrPrefixDto method XCtxEmpty.

3. Added ErrPrefixDto method XSetEmpty.

4. Changed behavior of ErrPrefixDto method SetCtx. Empty input strings will now delete last error context.

5. Changed behavior of ErrPrefixDto method XCtx. Empty input strings will now delete last error context.

6. Added ErrPrefixDto method NewFromIErrorPrefix.

## Version 1.2.0

Original Release Date: 2021-02-11 16:46:00 USA CST

Compiled and Tested: Go 1.15.8

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