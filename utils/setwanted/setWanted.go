package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

func main() {

	s := bufio.NewScanner(os.Stdin)
	//	str := `/home/nunnatsa/GIT/ginkgolinter/testdata/src/a/len/c.go:21:3: ginkgo-linter: found issues: (1) wrong length assertion; consider using ` + "`g.Expect(s).Should(g.HaveLen(4))`" + ` instead
	///home/nunnatsa/GIT/ginkgolinter/testdata/src/a/len/c.go:22:3: ginkgo-linter: found issues: (1) wrong length assertion; consider using ` + "`g.Expect(s).WithOffset(1).Should(g.HaveLen(4))`" + ` instead
	///home/nunnatsa/GIT/ginkgolinter/testdata/src/a/len/c.go:23:3: ginkgo-linter: found issues: (1) wrong length assertion; consider using ` + "`g.Expect(s).WithOffset(1).Should(g.HaveLen(4))`" + ` instead
	///home/nunnatsa/GIT/ginkgolinter/testdata/src/a/len/c.go:24:3: ginkgo-linter: found issues: (1) wrong length assertion; consider using ` + "`g.Expect(s).WithOffset(1).ShouldNot(g.HaveLen(5))`" + ` instead
	//`
	//	s := bufio.NewScanner(bytes.NewBuffer([]byte(str)))
	m := parse(s)

	errgrp := errgroup.Group{}

	for f, mp := range m {
		func(file string, msgs map[int][]string) {
			errgrp.Go(func() error {
				return fixFile(file, msgs)
			})
		}(f, mp)
	}

	if err := errgrp.Wait(); err != nil {
		panic(err)
	}
}

func parse(scanner *bufio.Scanner) map[string]map[int][]string {
	m := make(map[string]map[int][]string)
	for scanner.Scan() {
		fileName, lineNum, msg, err := parseLine(scanner.Text())
		if err == nil {
			if _, exists := m[fileName]; !exists {
				m[fileName] = make(map[int][]string)
			}
			m[fileName][lineNum] = append(m[fileName][lineNum], msg)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m
}

func parseLine(line string) (string, int, string, error) {
	idx := strings.Index(line, ":")
	fileName := line[:idx]
	line = line[idx+1:]
	idx = strings.Index(line, ":")
	lineNum, err := strconv.Atoi(line[:idx])
	if err != nil {
		return "", 0, "", err
	}

	idx = strings.Index(line, " ")
	msg := line[idx+1:]

	msg = strings.ReplaceAll(msg, "(", "\\(")
	msg = strings.ReplaceAll(msg, ")", "\\)")
	msg = strings.ReplaceAll(msg, "[", "\\[")
	msg = strings.ReplaceAll(msg, "]", "\\]")
	msg = strings.ReplaceAll(msg, "*", "\\*")
	msg = strings.ReplaceAll(msg, ".", "\\.")
	msg = strings.ReplaceAll(msg, "`", ".")

	return fileName, lineNum, msg, nil
}

func fixFile(fileName string, msgs map[int][]string) error {
	fmt.Println("fixing file", fileName)
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for n, msgList := range msgs {
		if n > len(lines) {
			return fmt.Errorf("wrong line numbe %d for file %s; msgs = %v", n, fileName, msgList)
		}
		line := lines[n-1]
		idx := strings.Index(line, "//")

		msg := strings.Join(msgList, "` `")
		if idx == -1 {
			line = line + " // want `" + msg + "`"
		} else {
			line = line[:idx] + "// want `" + msg + "`"
		}
		lines[n-1] = line
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}
