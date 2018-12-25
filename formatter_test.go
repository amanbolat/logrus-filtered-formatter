package filtered_test

import (
	"encoding/json"
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
	}

	formatter := filtered.New(fields, &logrus.JSONFormatter{})

	for k, v := range mustFilteredFields {
		b, err := formatter.Format(logrus.WithField(k, v))
		assert.NoError(t, err)

		entry := make(map[string]interface{})
		err = json.Unmarshal(b, &entry)
		assert.NoError(t, err)

		assert.Equal(t, "[FILTERED]", entry[k])
	}
}
