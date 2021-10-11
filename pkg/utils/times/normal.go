package times

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type TimeNormal struct { // 内嵌方式（推荐）
	time.Time
}

func (t TimeNormal) MarshalJSON() ([]byte, error) {
	// tune := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	value, err := t.Value()
	if err != nil {
		return nil, err
	}
	if value == nil {
		return json.Marshal(nil)
	}
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

// Value insert timestamp into mysql need this function.
func (t TimeNormal) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *TimeNormal) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeNormal{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
