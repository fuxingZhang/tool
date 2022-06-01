package tool

import (
	"errors"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	successfn := func() error { return nil }
	failedfn := func() error { return errors.New("unexpect error") }

	err := Retry(3, time.Millisecond*100, successfn)
	if err != nil {
		t.Fatal(err)
	}

	err = Retry(3, time.Millisecond*100, failedfn)
	if err == nil {
		t.Fatal(err)
	}
}
