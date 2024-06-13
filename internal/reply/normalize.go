package reply

import "strings"

func normalizeReply(reply string) []string {
	parts := strings.Split(strings.ReplaceAll(reply, "\n", ""), " ")

	for i, part := range parts {
		parts[i] = strings.ReplaceAll(part, "\n", "")
		parts[i] = strings.ReplaceAll(part, "\x00", "")
	}

	return parts
}
