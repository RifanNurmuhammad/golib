package golib

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	length := 10
	rs := RandomString(length)

	// set invalid string that
	// should not be contained in random string
	invalidString := `!@#$%^&*()_+`

	if len(rs) != length {
		t.Errorf("length of random string is not equal %d", length)
	}

	if strings.Contains(rs, invalidString) {
		t.Fatal("random string contains symbols")
	}
}

func TestValidateEmail(t *testing.T) {
	var (
		email string
		err   error
	)

	// test valid email should be not error/valid
	email = "julius.bernhard@bhinneka.com"
	if err = ValidateEmail(email); err != nil {
		t.Fatal("testing valid email is not valid")
	}

	// test invalid email should be error
	email = "julius.@bernhard@bhinneka.com"
	if err = ValidateEmail(email); err == nil {
		t.Fatal("testing invalid email is not valid")
	}
}

func TestValidateURL(t *testing.T) {
	var (
		url string
		err error
	)

	url = "http://www.bhinneka.com"
	if err = ValidateURL(url); err != nil {
		t.Fatal("testing 1st valid URL is not valid")
	}

	url = "www.bhinneka.com"
	if err = ValidateURL(url); err != nil {
		t.Fatal("testing 2nd valid URL is not valid")
	}

	url = "ftp://www.bhinneka.com"
	if err = ValidateURL(url); err != nil {
		t.Fatal("testing 3rd valid URL is not valid")
	}

	url = "https:///www.bhinneka.com"
	if err = ValidateURL(url); err == nil {
		t.Fatal("testing invalid URL is not valid")
	}
}

func TestValidatePhoneNumber(t *testing.T) {
	var (
		tel string
		err error
	)

	tel = "08119889788"
	if err = ValidatePhoneNumber(tel); err != nil {
		t.Fatal("testing valid phone number is not valid")
	}

	tel = "081-1988-9788"
	if err = ValidatePhoneNumber(tel); err == nil {
		t.Fatal("testing 1st invalid phone number is not valid")
	}

	tel = "0811"
	if err = ValidatePhoneNumber(tel); err == nil {
		t.Fatal("testing 2nd invalid phone number - not greater than 5 chars is not valid")
	}
}

func TestValidatePhoneAreaNumber(t *testing.T) {
	var (
		area string
		err  error
	)

	area = "+62"
	if err = ValidatePhoneAreaNumber(area); err != nil {
		t.Fatal("testing valid area number is not valid")
	}

	area = "+6 2"
	if err = ValidatePhoneAreaNumber(area); err == nil {
		t.Fatal("testing 1st invalid area number is not valid")
	}

	area = "+"
	if err = ValidatePhoneAreaNumber(area); err == nil {
		t.Fatal("testing 2nd invalid area number is not valid")
	}
}

func TestValidateAlphaNumeric(t *testing.T) {
	var (
		alpha string
	)

	alpha = "Some days are beautiful."
	if !ValidateAlphaNumeric(alpha) {
		t.Fatal("testing valid alpha numeric is not valid")
	}

	alpha = "Some days are beautiful. :) :*"
	if ValidateAlphaNumeric(alpha) {
		t.Fatal("testing 1st invalid alpha numeric is not valid")
	}

	alpha = `<img src="http://example.com/image.jpg" />`
	if ValidateAlphaNumeric(alpha) {
		t.Fatal("testing 2nd invalid alpha numeric is not valid")
	}
}

func TestValidateNumeric(t *testing.T) {
	t.Run("Test Validate Numeric", func(t *testing.T) {
		boolFalse := ValidateNumeric("1.0.1")
		assert.False(t, boolFalse)

		boolTrue := ValidateNumeric("0123456789")
		assert.True(t, boolTrue)
	})
}

func TestValidateAlphabet(t *testing.T) {
	t.Run("Test Validate Alphabet", func(t *testing.T) {
		boolTrue := ValidateAlphabet("huFtBanGeT")
		assert.True(t, boolTrue)

		boolFalse := ValidateAlphabet("1FgH^*")
		assert.False(t, boolFalse)
	})
}

func TestValidateAlphabetWithSpace(t *testing.T) {
	t.Run("Test Validate Alphabet With Space", func(t *testing.T) {
		boolFalse := ValidateAlphabetWithSpace("huFtBanGeT*")
		assert.False(t, boolFalse)

		boolTrue := ValidateAlphabetWithSpace("huFt BanGeT")
		assert.True(t, boolTrue)
	})
}

func TestValidateAlphanumeric(t *testing.T) {
	t.Run("Test Validate Alphabet Numeric", func(t *testing.T) {
		boolTrue := ValidateAlphanumeric("okesip12", true)
		assert.True(t, boolTrue)

		boolTrue = ValidateAlphanumeric("okesip", false)
		assert.True(t, boolTrue)

		boolFalse := ValidateAlphanumeric("1FgH^*", false)
		assert.False(t, boolFalse)
	})
}

func TestValidateAlphanumericWithSpace(t *testing.T) {
	t.Run("Test Validate Alphabet Numeric With Space", func(t *testing.T) {
		boolTrue := ValidateAlphanumericWithSpace("oke sip1", false)
		assert.True(t, boolTrue)

		boolTrue = ValidateAlphanumericWithSpace("OKE sip1", false)
		assert.True(t, boolTrue)

		boolFalse := ValidateAlphanumericWithSpace("okesip1", true)
		assert.False(t, boolFalse)

		boolFalse = ValidateAlphanumericWithSpace("okesip1@", true)
		assert.False(t, boolFalse)
	})
}

func TestGenerateRandomID(t *testing.T) {
	t.Run("Test Generate Random ID", func(t *testing.T) {
		var res string
		randomID := GenerateRandomID(5)
		assert.IsType(t, res, randomID)

		randomID = GenerateRandomID(5, "00")
		assert.IsType(t, res, randomID)
	})
}

func TestRandomNumber(t *testing.T) {
	t.Run("Test Generate Random Number", func(t *testing.T) {
		var res string
		randomNumber := RandomNumber(5)

		assert.IsType(t, res, randomNumber)
	})
}
