package types

import (
	"strings"

	"go/ast"
)

const (
	suppressPrefix                 = "ginkgo-linter:"
	suppressLengthAssertionWarning = suppressPrefix + "ignore-len-assert-warning"
	suppressNilAssertionWarning    = suppressPrefix + "ignore-nil-assert-warning"
	suppressErrAssertionWarning    = suppressPrefix + "ignore-err-assert-warning"
)

type Config struct {
	SuppressLen        Boolean
	SuppressNil        Boolean
	SuppressErr        Boolean
	AllowHaveLen0      Boolean
	AllowSprintfInDesc Boolean
}

func (s *Config) AllTrue() bool {
	return bool(s.SuppressLen && s.SuppressNil && s.SuppressErr && s.AllowSprintfInDesc)
}

func (s *Config) Clone() Config {
	return Config{
		SuppressLen:        s.SuppressLen,
		SuppressNil:        s.SuppressNil,
		SuppressErr:        s.SuppressErr,
		AllowHaveLen0:      s.AllowHaveLen0,
		AllowSprintfInDesc: s.AllowSprintfInDesc,
	}
}

func (s *Config) UpdateFromComment(commentGroup []*ast.CommentGroup) {
	for _, cmntList := range commentGroup {
		if s.AllTrue() {
			break
		}

		for _, cmnt := range cmntList.List {
			commentLines := strings.Split(cmnt.Text, "\n")
			for _, comment := range commentLines {
				comment = strings.TrimPrefix(comment, "//")
				comment = strings.TrimPrefix(comment, "/*")
				comment = strings.TrimSuffix(comment, "*/")
				comment = strings.TrimSpace(comment)

				s.SuppressLen = s.SuppressLen || (comment == suppressLengthAssertionWarning)
				s.SuppressNil = s.SuppressNil || (comment == suppressNilAssertionWarning)
				s.SuppressErr = s.SuppressErr || (comment == suppressErrAssertionWarning)
			}
		}
	}
}

func (s *Config) UpdateFromFile(cm ast.CommentMap) {

	for key, commentGroup := range cm {
		if s.AllTrue() {
			break
		}

		if _, ok := key.(*ast.GenDecl); ok {
			s.UpdateFromComment(commentGroup)
		}
	}
}
