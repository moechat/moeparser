package token

// Options for Tokens - return these bits in GetOptions() to implement this behavior
const (
	// TODO: Interpret as single if there is no closing tag
	PossibleSingle int = 1 << iota
	// TODO: This makes MoeParser ignore any tags inside this tags body. It will be ignored if the Single bit is set.
	NoParseInner
	// TODO: This makes the text inside of tag and passes it as an arg for the output. The text inside will not be parsed.
	TokenBodyAsArg
	// TODO: This makes the tag body become the first arg if there is no first argument (makes [name]arg0[/name] the same as [name=arg0][/name])
	AllowTokenBodyAsFirstArg
	// TODO: remove Converts a number to the number + "px" (ie 12 -> 12px)
	NumberArgToPx
	// TODO: This makes MoeParser match the tags that don't either start with whitespace or the beginning of a line. This is only useful if the token is of the type OpenCloseClass.
	AllowMidWord
	// TODO: This makes MoeParser stop matching:
	// - OpenClass type tokens without leading whitespace/beginning of body
	// - CloseClass type tokens without trailing whitespace/end of body
	// - Single type tokens without leading and trailing whitespace
	DisallowMidWord
	// For compatibility with the bbcode module until it's folded into the moeparsing routines
	HtmlSingle
)

// Token class types - return these in GetType()
const (
	// A single token with no opening or closing token
	SingleToken int = iota
	// A token class that starts a section
	OpenToken
	// A token class that ends a section
	CloseToken
	// A token class that can both begin and end a section
	SymmetricToken
)

// The arguments for a token - these are the regexp capture groups
type TokenArgs struct {
	args     []string
	idByName map[string]int

	size int
}

func NewTokenArgs(args []string, idByName map[string]int) *TokenArgs {
	return &TokenArgs{args, idByName, len(args)}
}

func (ta *TokenArgs) ById(id int) string {
	if id < ta.size {
		return ta.args[id]
	}
	return ""
}

func (ta *TokenArgs) ByName(name string) string {
	if id, ok := ta.idByName[name]; ok {
		return ta.args[id]
	}
	return ""
}

func (ta *TokenArgs) Size() int {
	return ta.size
}

type TokenBuilder interface {
	Build(args *TokenArgs) Token
}

// The Token interface represents an instance of a token class
type Token interface {
	Type() string
}

// A special case of Token used to represent text that isn't matched by any other tokens
// i.e. "hi" in <p>hi</p>
type TextToken struct {
	body string
}

func NewTextToken(body string) *TextToken {
	return &TextToken{body: body}
}

func (tt *TextToken) SetArgs(args *TokenArgs) {
	tt.body = args.ById(0)
}

// Returns the TextToken's body
func (tt *TextToken) Output() (string, error) {
	return tt.body, nil
}

func (tt *TextToken) Type() string {
	return "TEXT"
}
