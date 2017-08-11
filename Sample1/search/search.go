package search

import(
	"log"
	"sync"
)

//小写字母开头的变量不公开，大写字母开头的公开
var matchers = make(map[string]Matcher)

