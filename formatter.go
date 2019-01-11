package filtered

import (
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

const defaultFilteredValue = "[FILTERED]"

var alphaNumericReg = regexp.MustCompile("[^a-zA-Z0-9]+")

// Formatter formats logs with filtered fields
type Formatter struct {
	formatter     logrus.Formatter
	fields        map[string]bool
	FilteredValue string
}

// New returns an instance of logrus formatter
func New(fields []string, formatter logrus.Formatter) *Formatter {
	m := make(map[string]bool)
	for _, f := range fields {
		m[normalizeString(f)] = true
	}
	return &Formatter{
		formatter:     formatter,
		fields:        m,
		FilteredValue: defaultFilteredValue,
	}
}

// Format renders a single log entry
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	// for k := range entry.Data {
	// 	for filterKey := range f.fields {
	// 		normalizedField := normalizeString(k)
	// 		if strings.Contains(normalizedField, filterKey) {
	// 			entry.Data[k] = f.FilteredValue
	// 		}
	// 	}
	// }
	f.format(entry.Data)

	data, err := f.formatter.Format(entry)
	return data, err
}

func (f *Formatter) format(m map[string]interface{}) {
	for k, v := range m {
		nestedMap, ok := v.(map[string]interface{})
		if ok {
			f.format(nestedMap)
		}

		for filterKey := range f.fields {
			normalizedField := normalizeString(k)
			if strings.Contains(normalizedField, filterKey) {
				m[k] = f.FilteredValue
			}
		}
	}
}

func normalizeString(str string) string {
	s := strings.ToLower(str)
	s = alphaNumericReg.ReplaceAllString(s, "")

	return s
}
