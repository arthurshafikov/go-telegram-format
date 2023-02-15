package tgformat

import "fmt"

type MarkdownFormatter struct{}

func NewMarkdownFormatter() *MarkdownFormatter {
	return &MarkdownFormatter{}
}

func (mf MarkdownFormatter) Bold(str string) string {
	return fmt.Sprintf("*%s*", str)
}

func (mf MarkdownFormatter) Italic(str string) string {
	return fmt.Sprintf("_%s_", str)
}

func (mf MarkdownFormatter) URL(text, url string) string {
	return fmt.Sprintf("[%s](%s)", text, url)
}

func (mf MarkdownFormatter) UserMention(text string, userID int) string {
	return fmt.Sprintf("[%s](tg://user?id=%v)", text, userID)
}

func (mf MarkdownFormatter) InlineCode(code string) string {
	return fmt.Sprintf("`%s`", code)
}

func (mf MarkdownFormatter) Code(code string, language ...string) string {
	if len(language) == 0 {
		language[0] = ""
	}

	return fmt.Sprintf("```%s\n%s```", language[0], code)
}
