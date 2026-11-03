package errpref

import (
	"errors"
	"fmt"
	"testing"
)

// ============================================================================
// Simulated Function Chain (Deep Call Stack)
// ============================================================================

// Series04Func01 calls Series04Func02 and adds context if it fails.
func Series04Func01() error {
	factory := (&ErrorGenericFactory{}).NewErrorGenericConfig("series04.go", "Service", "Series04Func01")

	err := Series04Func02()
	if err != nil {
		return factory.GetErrorFromFunctionCall(
			"",
			"Series04Func02",
			"orchestrating downstream sequence",
			"top-level operation failed",
			err,
		)
	}
	return nil
}

// Series04Func02 calls Series03Func03 and adds context if it fails.
func Series04Func02() error {
	factory := (&ErrorGenericFactory{}).NewErrorGenericConfig("series04.go", "Service", "Series04Func02")

	err := Series03Func03()
	if err != nil {
		return factory.GetErrorFromFunctionCall(
			"",
			"Series03Func03",
			"validating intermediate system state",
			"mid-level verification failed",
			err,
		)
	}
	return nil
}

// Series03Func03 calls Series03Func04 and adds context if it fails.
func Series03Func03() error {
	factory := (&ErrorGenericFactory{}).NewErrorGenericConfig("series03.go", "Repository", "Series03Func03")

	err := Series03Func04()
	if err != nil {
		return factory.GetErrorFromFunctionCall(
			"",
			"Series03Func04",
			"executing data processing operation",
			"low-level repository failed",
			err,
		)
	}
	return nil
}

// Series03Func04 generates the root ErrorNativeToFunction error.
func Series03Func04() error {
	factory := (&ErrorGenericFactory{}).NewErrorGenericConfig("series03.go", "Repository", "Series03Func04")

	// Simulating a calculation or operational bounds error native to this function
	return factory.GetErrorNativeToFunction(
		"userID",
		"/var/log/app",
		"data.bin",
		"https://example.com",
		"internal buffer alignment",
		"encountered fatal calculation error: value out of bounds",
	)
}

// ============================================================================
// Test Execution (TestFunc04)
// ============================================================================

func TestFunc04(t *testing.T) {
	// Execute the full call chain
	err := Series04Func01()

	if err == nil {
		t.Fatal("Expected an error from Series04Func01, but got nil")
	}

	// 1. Print the entire combined error string to visualize the stack trace
	fmt.Println("--- zzzt_ErrorGenericFactory_0010_test.go TestFunc04 ---")
	fmt.Println("--- Testing ErrorGenericFactory ---")
	fmt.Println("--- FULL COMBINED ERROR OUTPUT ---")
	fmt.Println(err.Error())
	fmt.Println("---------------------------------")

	// 2. Inspect the chain to confirm the original ErrorNativeToFunction is fully visible
	var nativeErr *ErrorNativeToFunction

	if errors.As(err, &nativeErr) {
		t.Log("✅ Success: Target ErrorNativeToFunction was found in the error chain!")

		// You can directly access the promoted inner data fields on the cast type
		fmt.Printf("Root Cause parameters identified at top:\n")
		// (Assuming fields were exported, or using accessors if they remain lowercase)
		// For private struct fields inside the same package, you can read them.
		// If testing externally, your struct fields would need to be Capitalized (Exported).
	} else {
		t.Error("❌ Failure: Could not find ErrorNativeToFunction in the returned error chain.")
	}

	// 3. Confirm we can also extract intermediate layer contexts if needed

	/*

	   if errors.As(err, &callErr) {
	     t.Logf("✅ Success: Intercepted intermediate wrapper showing called function: %s", err)
	   }

	*/

	if _, ok := errors.AsType[*ErrorFromFunctionCall](err); ok {
		t.Logf("✅ Success: Intercepted intermediate wrapper showing called function: %s", err)
	}
}
