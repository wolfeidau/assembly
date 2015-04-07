package assembly

import "time"

// AssemblyDateFormat is the format of the dates used throughout the
// api, note this odd string is used to parse/format dates in go
const AssemblyDateFormat = "2006-01-02 15:04:05 MST"

// Timestamp custom timestamp to support assembly api timestamps
type Timestamp struct {
	time.Time
}

// NewTimestamp make a new timestamp using the time suplied.
func NewTimestamp(t time.Time) *Timestamp {
	return &Timestamp{t}
}

func (ts Timestamp) String() string {
	return ts.Time.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (ts Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(ts.Format(`"` + AssemblyDateFormat + `"`)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ts *Timestamp) UnmarshalJSON(data []byte) (err error) {
	(*ts).Time, err = time.Parse(`"`+AssemblyDateFormat+`"`, string(data))
	return
}

// Equal reports whether t and u are equal based on time.Equal
func (ts Timestamp) Equal(u Timestamp) bool {
	return ts.Time.Equal(u.Time)
}
