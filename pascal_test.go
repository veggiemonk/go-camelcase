package strcase

import (
	"fmt"
	"testing"
)

func TestPascal(t *testing.T) {
	samples := []TT{
		{"sample text", "SampleText"},
		{"sample-text", "SampleText"},
		{"sample_text", "SampleText"},
		{"sample___text", "SampleText"},
		{"sampleText", "SampleText"},
		{"inviteYourCustomersAddInvites", "InviteYourCustomersAddInvites"},
		{"sample 2 Text", "Sample2Text"},
		{"   sample   2    Text   ", "Sample2Text"},
		{"   $#$sample   2    Text   ", "Sample2Text"},
		{"SAMPLE 2 TEXT", "Sample2Text"},
		{"___$$Base64Encode", "Base64Encode"},
		{"FOO:BAR$BAZ", "FooBarBaz"},
		{"FOO#BAR#BAZ", "FooBarBaz"},
		{"something.com", "SomethingCom"},
		{"$something%", "Something"},
		{"something.com", "SomethingCom"},
		{"•¶§ƒ˚foo˙∆˚¬", "Foo"},
		{"CStringRef", "CstringRef"},
		{"5test", "5Test"},
		{"test5", "Test5"},
		{"THE5r", "The5R"},
		{"5TEst", "5Test"},
		{"_5TEst", "5Test"},
		{"@%#&5TEst", "5Test"},
		{"edf_6N", "Edf6N"},
		{"f_pX9", "FPX9"},
		{"p_z9Rg", "PZ9Rg"},
		{"2FA Enabled", "2FaEnabled"},
		{"Enabled 2FA", "Enabled2Fa"},
	}

	for _, sample := range samples {
		if out := Pascal(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func ExamplePascal() {
	fmt.Println(Pascal("sample text"))
	fmt.Println(Pascal("sample-text"))
	fmt.Println(Pascal("sample_text"))
	fmt.Println(Pascal("sample___text"))
	fmt.Println(Pascal("sampleText"))
	fmt.Println(Pascal(">> samPLE text <<"))
	fmt.Println(Pascal("sample 2 Text"))
	fmt.Println(Pascal("SAMPLE 2 TEXT"))
	fmt.Println(Pascal("invite-Your-Customers-Add-Invites"))
	fmt.Println(Pascal("2FA Enabled"))
	fmt.Println(Pascal("Enabled 2FA"))
	// Output:
	// SampleText
	// SampleText
	// SampleText
	// SampleText
	// SampleText
	// SamPleText
	// Sample2Text
	// Sample2Text
	// InviteYourCustomersAddInvites
	// 2FaEnabled
	// Enabled2Fa
}

func ExamplePascal_withSpace() {
	fmt.Println(Pascal("   sample   2    Text   "))
	// Output:
	// Sample2Text
}

func ExamplePascal_long() {
	fmt.Println(Pascal("super long sentence that is not really necessary but we need to try it"))
	// Output:
	// SuperLongSentenceThatIsNotReallyNecessaryButWeNeedToTryIt
}

func ExamplePascal_multiline() {
	fmt.Println(
		Pascal(
			`here
is
a
multiline
string`,
		),
	)
	// Output:
	// HereIsAMultilineString
}

func ExamplePascal_quoted() {
	fmt.Println(Pascal("\"hello world\""))
	// Output:
	// HelloWorld
}

func ExamplePascal_swedish() {
	// Swedish and non ASCII char are not supported :(
	// open an issue if this is something you need
	//
	// want:
	// när-såg-du-en-kråka-väl-bita-en-man
	fmt.Println(Pascal("När såg du en kråka väl bita en man?"))
	// Output:
	// NRSGDuEnKrKaVLBitaEnMan
}

func ExamplePascal_nonASCII() {
	fmt.Println(Pascal("   $#$sample   2    Text   "))
	fmt.Println(Pascal("___$$Base64Encode"))
	fmt.Println(Pascal("FOO#BAR#BAZ"))
	fmt.Println(Pascal("FOO:BAR$BAZ"))
	fmt.Println(Pascal("FOO#BAR#BAZ"))
	fmt.Println(Pascal("something.com"))
	fmt.Println(Pascal("$something%"))
	fmt.Println(Pascal("something.com"))
	fmt.Println(Pascal("•¶§ƒ˚foo˙∆˚¬"))
	fmt.Println(Pascal("•¶§ƒ˚foo˙∆˚¬bar"))
	fmt.Println(Pascal("•¶§ƒ˚foo˙∆˚¬ bar"))
	fmt.Println(Pascal("CStringRef"))

	// Output:
	// Sample2Text
	// Base64Encode
	// FooBarBaz
	// FooBarBaz
	// FooBarBaz
	// SomethingCom
	// Something
	// SomethingCom
	// Foo
	// FooBar
	// FooBar
	// CstringRef
}

func ExamplePascal_cryptic() {
	fmt.Println(Pascal("5test"))
	fmt.Println(Pascal("test5"))
	fmt.Println(Pascal("THE5r"))
	fmt.Println(Pascal("5TEst"))
	fmt.Println(Pascal("_5TEst"))
	fmt.Println(Pascal("@%#&5TEst"))
	fmt.Println(Pascal("edf_6N"))
	fmt.Println(Pascal("f_pX9"))
	fmt.Println(Pascal("p_z9Rg"))
	fmt.Println(Pascal("@49L0S145_¬fwHƒ0TSLNVp"))
	fmt.Println(Pascal("lk0B@bFmjrLQ_Z6YL"))
	// Output:
	// 5Test
	// Test5
	// The5R
	// 5Test
	// 5Test
	// 5Test
	// Edf6N
	// FPX9
	// PZ9Rg
	// 49L0S145FwH0Tslnvp
	// Lk0BBFmjrLqZ6Yl
}
