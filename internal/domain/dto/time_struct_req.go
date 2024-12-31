package dto

import (
	"time"
)

type DatetimeReq time.Time

func (d *DatetimeReq) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}
	*d = DatetimeReq(t)
	return nil
}

func (d DatetimeReq) GetTime() time.Time {
	return time.Time(d)
}

func (d DatetimeReq) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(d).Format(`"2006-01-02T15:04:05"`)), nil
}
