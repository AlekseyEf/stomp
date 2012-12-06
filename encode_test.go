package stomp

import (
	. "launchpad.net/gocheck"
	"testing"
)

func TestEncode(t *testing.T) { TestingT(t) }

type EncodeSuite struct{}

var _ = Suite(&EncodeSuite{})

func (s *EncodeSuite) TestEncodeValue(c *C) {
	c.Check(encodeValue("Contains\r\nNewLine and : colon and \\ backslash"),
		Equals,	"Contains\\r\\nNewLine and \\c colon and \\\\ backslash")
}

func (s *EncodeSuite) TestUnencodeValue(c *C) {
	c.Check(unencodeValue("Contains\\r\\nNewLine and \\c colon and \\\\ backslash"),
		Equals,	"Contains\r\nNewLine and : colon and \\ backslash")
}
