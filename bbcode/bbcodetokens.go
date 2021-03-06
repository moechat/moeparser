package bbcode

import (
	"github.com/moechat/parser/token"
)

// A type that determines what the parser will replace tags it finds with. The Attributes and CssProps are maps that assign a regexp parser group
type HtmlTags struct {
	Options      int                   // Compatibility options for BBCode until token parsing is complete
	Tags         []string              // HTML tags
	Classes      [][]string            // Classes to give to the HTML elements
	Attributes   []map[int8]string     // HTML tag attributes
	CssProps     []map[int8]string     // CSS Properties
	OutputFunc   func([]string) string // A custom output function; this returns the string to emplace into the HTML.
	InputModFunc func(*[]string)       // A function that takes input and returns input modified (an example use case would be converting a username to a user ID in @tagging)
}

var bbCodeTags = map[string]HtmlTags{
	"b": {Tags: []string{"b"}},
	"i": {Tags: []string{"i"}},
	"u": {
		Tags:    []string{"span"},
		Classes: [][]string{{"underline"}},
	},
	"pre":  {Options: token.NoParseInner, Tags: []string{"pre"}},
	"code": {Options: token.NoParseInner, Tags: []string{"pre", "code"}},
	"color": {
		Tags:     []string{"span"},
		CssProps: []map[int8]string{{0: "color"}},
	},
	"colour": {
		Tags:     []string{"span"},
		CssProps: []map[int8]string{{0: "color"}},
	},
	"size": {
		Options:  token.NumberArgToPx,
		Tags:     []string{"span"},
		CssProps: []map[int8]string{{0: "font-size"}},
	},
	"noparse": {Options: token.NoParseInner},
	"url": {
		Options:    (token.AllowTokenBodyAsFirstArg | token.PossibleSingle),
		Tags:       []string{"a"},
		Attributes: []map[int8]string{{0: "href"}},
	},
	"img": {
		Options: (token.AllowTokenBodyAsFirstArg |
			token.TokenBodyAsArg |
			token.PossibleSingle |
			token.HtmlSingle),
		Tags:       []string{"img"},
		Attributes: []map[int8]string{{0: "src", 1: "title"}},
	},
	"s":    {Tags: []string{"s"}},
	"samp": {Tags: []string{"samp"}},
	"q":    {Tags: []string{"q"}},
}

// One can insert use-case specific BBCode tags by using this function.
//
// IMPORTANT: This is ignored by parser.Parse - you should use AddTokenClass instead!
// Only use this function if you plan on using parser.BbCodeParse()!
//
// This will be deprecated in the future after BBCode functionality is added to AddMatcher.
func AddBbToken(name string, htmlTags HtmlTags) {
	bbCodeTags[name] = htmlTags
}
