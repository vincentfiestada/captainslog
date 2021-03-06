package caller_test

import (
	"testing"

	"github.com/vincentfiestada/captainslog/v2/caller"
	"github.com/vincentfiestada/captainslog/v2/preflight"
)

func TestGetName(test *testing.T) {
	t := preflight.Unit(test)

	// GetName(1) should return the name of the calling function
	this := "github.com/vincentfiestada/captainslog/v2/caller_test.TestGetName"
	t.Expect(caller.GetName(1)).Equals(this)

	func() {
		t.Expect(caller.GetName(2)).Equals(this)
	}()
}

func TestShorten(test *testing.T) {
	t := preflight.Unit(test)

	path := "github.com/vincentfiestada/captainslog/caller_test.TestShorten"

	// should return the most specific path possible
	t.Expect(caller.Shorten(path, 15)).Equals("TestShorten")
	t.Expect(caller.Shorten(path, 11)).Equals("TestShort..")
	t.Expect(caller.Shorten(path, 30)).Equals("caller_test.TestShorten")
}
