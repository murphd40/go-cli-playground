package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Scanner struct {
	src *bufio.Reader
	curr byte
	err error

	tokenBuffer *bytes.Buffer
}

type ScanCondition func(*Scanner, byte) bool

func (s *Scanner) Init(src io.Reader) {
	s.src = bufio.NewReader(src)
	s.tokenBuffer = new(bytes.Buffer)
	// read first byte
	s.read()
}

func (s *Scanner) read() {
	s.curr, s.err = s.src.ReadByte()
}

func (s *Scanner) isErr() bool {
	return s.err != nil
}

func (s *Scanner) ScanUntil(condition ScanCondition) int {
	s.tokenBuffer.Reset()
	for !s.isErr() {
		if condition(s, s.curr) {
			break
		}
		s.tokenBuffer.WriteByte(s.curr)

		s.read()
	}
	return 0
}

func (s *Scanner) Word() string {
	s.ScanUntil(IsNotWhitespace)
	s.ScanUntil(IsWhitespace)
	return s.Token()
}

func (s *Scanner) Line() string {
	s.ScanUntil(NewLine)
	return s.Token()
}

func IsWhitespace(s *Scanner, b byte) bool {
	return isSpace(b)
}

func IsNotWhitespace(s *Scanner, b byte) bool {
	return !isSpace(b)
}

func EOF(s *Scanner, b byte) bool {
	return true
}

func NewLine(s *Scanner, b byte) bool {
	return b == '\n'
}

func (s *Scanner) Token() string {
	return s.tokenBuffer.String()
}

func isSpace(c byte) bool {
	return c <= ' ' && (c == ' ' || c == '\t' || c == '\r' || c == '\n')
}

func main() {
	text := `# HELP kube_customresource_status Cluster Service Version install status`
	
	var s Scanner

	s.Init(strings.NewReader(text))

	fmt.Println(s.Word())
	fmt.Println(s.Word())
	fmt.Println(s.Word())
	s.ScanUntil(IsNotWhitespace)
	fmt.Println(s.Line())
}