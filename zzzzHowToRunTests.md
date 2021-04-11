# Running Tests

## Run Basic Tests
Open a command prompt in this directory:
 *("./MikeAustin71/errprefops/errpref")*

and run the following command:

***go test -v > zzzzz_tests.txt***

This will generate a '.txt' file in the current directory which 
contains all the test results. 

The **-v** option signals a request for verbose output.

## Resources
***http://codesamplez.com/development/golang-unit-testing***

## Dependencies
None so far.

## Running Tests with code coverage

First pull down and install the ***cover*** package.

***go get golang.org/x/tools/cmd/cover***
  
Next, follow the test execution protocol.  
  
### Test Execution with Code Coverage

***go test -cover -v > zzzzz_tests.txt***  
     

### Cover Profile

Generate the code coverage detail:

***go test -coverprofile=zzzzz_coverage.out***

The following provides for code coverage display in your
browser. Run this on the command line:

  ***go tool cover -html=zzzzz_coverage.out***

## FileNaming Conventions
The use of *zzzzz_* as a file name prefix for output files ensures
that the output file will be listed last in the directory.