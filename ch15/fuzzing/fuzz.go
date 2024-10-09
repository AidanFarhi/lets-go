package fuzzing

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"strings"
)

func ParseData(r io.Reader) ([]string, error) {
	s := bufio.NewScanner(r)
	if !s.Scan() {
		return nil, errors.New("empty")
	}
	countStr := s.Text()
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, err
	}
	if count < 0 {
		return nil, errors.New("no negatives")
	}
	if count > 10_000 {
		return nil, errors.New("too many lines")
	}
	out := make([]string, 0, count)
	for i := 0; i < count; i++ {
		if !s.Scan() {
			return nil, errors.New("too few lines")
		}
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			return nil, errors.New("blank line")
		}
		out = append(out, line)
	}
	return out, nil
}

func ToData(s []string) []byte {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(len(s)))
	b.WriteRune('\n')
	for _, v := range s {
		b.WriteString(v)
		b.WriteRune('\n')
	}
	return b.Bytes()
}
