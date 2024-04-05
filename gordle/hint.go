package gordle

import "strings"

type hint byte

type feedback []hint

const (
	absentCharacter hint = iota
	wrongPosition
	correctPosition
)

func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "🔲"
	case correctPosition:
		return "💚"
	case wrongPosition:
		return "🟡"
	default:
		return "💔"
	}
}

func (fb feedback) StringConcat() string {
	sb := strings.Builder{}
	for _, f := range fb {
		sb.WriteString(f.String())
	}
	return sb.String()
}

// Helper function to test for equality
func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}
	for i, value := range fb {
		if value != other[i] {
			return false
		}
	}
	return true
}
