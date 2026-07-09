package reverse_words_in_a_string

import "strings"

func reverseWords(s string) string {
    for i := 0; i < len(s); i += 1 {
        if s[i] != ' ' {
            s = s[i:]
            break
        }
    }

    for i := len(s) - 1; i >= 0; i -= 1 {
        if s[i] != ' ' {
            s = s[:i + 1]
            break
        }
    }

    var builder strings.Builder
    builder.Grow(len(s))
    j := len(s)

    for i := len(s) - 1; i >= 0; i -= 1 {
        if s[i] != ' ' {
            continue
        }

        builder.WriteString(s[i + 1:j])
        builder.WriteByte(' ')
        
        for ; i >= 0; i -= 1 {
            if s[i] != ' ' {
                j = i + 1
                break
            }
        }
    }

    builder.WriteString(s[:j])
    return builder.String()
}