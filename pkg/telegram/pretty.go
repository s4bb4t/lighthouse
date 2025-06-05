package telegram

import (
	"fmt"
	"github.com/s4bb4t/lighthouse/pkg/core/levels"
	"github.com/s4bb4t/lighthouse/pkg/core/sp"
	"strings"
)

func prettify(err error) string {
	var b strings.Builder
	b.Grow(600)
	e := sp.Ensure(err).Spin(levels.LevelDebug)

	const divider = "\u2800" // non-breaking space
	section := func(title string) string {
		return "┌─ *" + title + "*\n"
	}

	b.WriteString(divider + "\n\n\n")

	b.WriteString(section("Level"))
	b.WriteString("🚨 *" + escape(e.ReadLevel().String()) + "*\n\n")

	b.WriteString(section("Message"))
	b.WriteString("📝 `" + escape(e.ReadMsg(sp.En)) + "`\n\n")

	if e.ReadCaused() != nil {
		b.WriteString(section("Cause"))
		b.WriteString("💥 `" + escape(e.ReadCaused().Error()) + "`\n\n")
	}

	b.WriteString(section("Description"))
	b.WriteString("📖 " + escape(e.ReadDesc()) + "\n\n")

	b.WriteString(section("Hint"))
	b.WriteString("💡 _" + escape(e.ReadHint()) + "_\n\n")

	b.WriteString(section("Source"))
	b.WriteString("🧭 || `" + escape(e.ReadSource()) + "` ||\n\n")

	b.WriteString(section("Time"))
	b.WriteString("⏰ " + escape(e.ReadTime().Format("2006-01-02 15:04:05")) + "\n\n")

	meta := e.AllMeta()
	if len(meta) > 0 {
		b.WriteString(section("Meta"))
		for key, val := range meta {
			b.WriteString("  • *" + escape(key) + "* → `" + escape(fmt.Sprintf("%v", val)) + "`\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("\n\n\n" + divider)
	return b.String()
}

func escape(s string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
		"||", "\\||",
	)
	return replacer.Replace(s)
}
