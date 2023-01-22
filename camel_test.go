package strcase

import (
	"fmt"
	"testing"
)

type TT struct {
	str, out string
}

func TestCamel(t *testing.T) {
	samples := []TT{
		{"sample text", "sampleText"},
		{"sample-text", "sampleText"},
		{"sample_text", "sampleText"},
		{"sample___text", "sampleText"},
		{"sampleText", "sampleText"},
		{"inviteYourCustomersAddInvites", "inviteYourCustomersAddInvites"},
		{"sample 2 Text", "sample2Text"},
		{"   sample   2    Text   ", "sample2Text"},
		{"   $#$sample   2    Text   ", "sample2Text"},
		{"SAMPLE 2 TEXT", "sample2Text"},
		{"___$$Base64Encode", "base64Encode"},
		{"FOO:BAR$BAZ", "fooBarBaz"},
		{"FOO#BAR#BAZ", "fooBarBaz"},
		{"something.com", "somethingCom"},
		{"$something%", "something"},
		{"something.com", "somethingCom"},
		{"•¶§ƒ˚foo˙∆˚¬", "foo"},
		{"CStringRef", "cstringRef"},
		{"5test", "5Test"},
		{"test5", "test5"},
		{"THE5r", "the5R"},
		{"5TEst", "5Test"},
		{"_5TEst", "5Test"},
		{"@%#&5TEst", "5Test"},
		{"edf_6N", "edf6N"},
		{"f_pX9", "fPX9"},
		{"p_z9Rg", "pZ9Rg"},
		{"2FA Enabled", "2FaEnabled"},
		{"Enabled 2FA", "enabled2Fa"},
	}

	for _, sample := range samples {
		if out := Camel(sample.str); out != sample.out {
			t.Errorf("got %q from %q, expected %q", out, sample.str, sample.out)
		}
	}
}

func ExampleCamel() {
	fmt.Println(Camel("sample text"))
	fmt.Println(Camel("sample-text"))
	fmt.Println(Camel("sample_text"))
	fmt.Println(Camel("sample___text"))
	fmt.Println(Camel("sampleText"))
	fmt.Println(Camel(">> samPLE text <<"))
	fmt.Println(Camel("sample 2 Text"))
	fmt.Println(Camel("SAMPLE 2 TEXT"))
	fmt.Println(Camel("invite-Your-Customers-Add-Invites"))
	fmt.Println(Camel("2FA Enabled"))
	fmt.Println(Camel("Enabled 2FA"))
	// Output:
	// sampleText
	// sampleText
	// sampleText
	// sampleText
	// sampleText
	// samPleText
	// sample2Text
	// sample2Text
	// inviteYourCustomersAddInvites
	// 2FaEnabled
	// enabled2Fa
}

func ExampleCamel_withSpace() {
	fmt.Println(Camel("   sample   2    Text   "))
	// Output:
	// sample2Text
}

func ExampleCamel_long() {
	fmt.Println(Camel("super long sentence that is not really necessary but we need to try it"))
	// Output:
	// superLongSentenceThatIsNotReallyNecessaryButWeNeedToTryIt
}

func ExampleCamel_multiline() {
	fmt.Println(
		Camel(
			`here
is
a
multiline
string`,
		),
	)
	// Output:
	// hereIsAMultilineString
}

func ExampleCamel_quoted() {
	fmt.Println(Camel("\"hello world\""))
	// Output:
	// helloWorld
}

func ExampleCamel_swedish() {
	// Swedish and non ASCII char are not supported :(
	// open an issue if this is something you need
	//
	// want:
	// när-såg-du-en-kråka-väl-bita-en-man
	fmt.Println(Camel("När såg du en kråka väl bita en man?"))
	// Output:
	// nRSGDuEnKrKaVLBitaEnMan
}

func ExampleCamel_nonASCII() {
	fmt.Println(Camel("   $#$sample   2    Text   "))
	fmt.Println(Camel("___$$Base64Encode"))
	fmt.Println(Camel("FOO#BAR#BAZ"))
	fmt.Println(Camel("FOO:BAR$BAZ"))
	fmt.Println(Camel("FOO#BAR#BAZ"))
	fmt.Println(Camel("something.com"))
	fmt.Println(Camel("$something%"))
	fmt.Println(Camel("something.com"))
	fmt.Println(Camel("•¶§ƒ˚foo˙∆˚¬"))
	fmt.Println(Camel("•¶§ƒ˚foo˙∆˚¬bar"))
	fmt.Println(Camel("•¶§ƒ˚foo˙∆˚¬ bar"))
	fmt.Println(Camel("CStringRef"))

	// Output:
	// sample2Text
	// base64Encode
	// fooBarBaz
	// fooBarBaz
	// fooBarBaz
	// somethingCom
	// something
	// somethingCom
	// foo
	// fooBar
	// fooBar
	// cstringRef
}

func ExampleCamel_cryptic() {
	fmt.Println(Camel("5test"))
	fmt.Println(Camel("test5"))
	fmt.Println(Camel("THE5r"))
	fmt.Println(Camel("5TEst"))
	fmt.Println(Camel("_5TEst"))
	fmt.Println(Camel("@%#&5TEst"))
	fmt.Println(Camel("edf_6N"))
	fmt.Println(Camel("f_pX9"))
	fmt.Println(Camel("p_z9Rg"))
	fmt.Println(Camel("@49L0S145_¬fwHƒ0TSLNVp"))
	fmt.Println(Camel("lk0B@bFmjrLQ_Z6YL"))
	// Output:
	// 5Test
	// test5
	// the5R
	// 5Test
	// 5Test
	// 5Test
	// edf6N
	// fPX9
	// pZ9Rg
	// 49L0S145FwH0Tslnvp
	// lk0BBFmjrLqZ6Yl
}
