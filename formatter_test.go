package filtered_test

import (
	"github.com/amanbolat/logrus-filtered-formatter"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestFormatter_Format(t *testing.T) {
	fields := []string{
		"password",
		"credit_card",
		"lastName",
		"first_name",
	}

	mustFilteredFields := map[string]interface{}{
		"password": "12345",
		"PassWORD": 123555,
		"super_password_field": "1235555",
		"password ": "pas123",
		"credit card": "6611 3334 6611 3333",
		"last name": "White",
		"FirstName": "John",
		"first_name_old": "Marina",
		"old credit card": "111 222 3333 4444",
		"req": map[string] interface{}{
			"password": "some_nested_password",
			"nested": map[string]interface{}{
				"first_name": "John Doe",
			},
		},
	}

	expectedFilteredFields := map[string]interface{}{
		"password": "[FILTERED]",
		"PassWORD": "[FILTERED]",
		"super_password_field": "[FILTERED]",
		"password ": "[FILTERED]",
		"credit card": "[FILTERED]",
		"last name": "[FILTERED]",
		"FirstName": "[FILTERED]",
		"first_name_old": "[FILTERED]",
		"old credit card": "[FILTERED]",
		"req": map[string] interface{}{
			"password": "[FILTERED]",
			"nested": map[string]interface{}{
				"first_name": "[FILTERED]",
			},
		},
	}

	formatter := filtered.New(fields, &logrus.JSONFormatter{})
	actual, err := formatter.Format(logrus.WithFields(mustFilteredFields))

	jsonFormatter := &logrus.JSONFormatter{}
	expected, err := jsonFormatter.Format(logrus.WithFields(expectedFilteredFields))

	t.Log(string(actual))
	t.Log(string(expected))
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
