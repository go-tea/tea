package tea

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type compileSet map[string]*regexp.Regexp

var compileVars = compileSet{}

//func (d compileSet) Set(name, regex string) error {
func (d compileSet) Set(name string, regex interface{}) error {

	if !strings.HasPrefix(name, "@") {
		return errors.New("compile name of regexp must start with a colon '@'")
	}

	switch regex.(type) {
	case string:
		// fix regex
		rxstr := regex.(string)
		if !strings.HasPrefix(rxstr, "^") {
			rxstr = fmt.Sprintf("^%s$", rxstr)
		}

		r := regexp.MustCompile(rxstr)
		d[name] = r

		return nil
	case *regexp.Regexp:

		rxstr := regex.(*regexp.Regexp).String()
		if !strings.HasPrefix(rxstr, "^") {
			rxstr = fmt.Sprintf("^%s$", rxstr)
			r := regexp.MustCompile(rxstr)
			d[name] = r

		} else {
			d[name] = regex.(*regexp.Regexp)
		}
		return nil
	default:
		return errors.New("regex is not string or regexp.Regexp")
	}

}
