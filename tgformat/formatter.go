package tgformat

const (
	MarkdownParseMode = "Markdown"
)

type Formatter interface {
	Bold(str string) string
	Italic(str string) string
	URL(text, url string) string
	UserMention(text string, userID int) string
	InlineCode(code string) string
	Code(code string, language ...string) string
}

func NewFormatter(style string) (Formatter, error) {
	var f Formatter

	switch style {
	case MarkdownParseMode:
		f = NewMarkdownFormatter()
	default:
		return nil, ErrUnsupportedStyle
	}

	return f, nil
}
