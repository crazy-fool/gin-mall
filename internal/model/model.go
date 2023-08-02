package model

//gorm中重新格式化json时间数据格式返回给前端
import (
	"strings"
	"time"
)

import (
	"database/sql/driver"
	"fmt"
)

const timeFormat = "2006-01-02 15:04:05"
const timezone = "Asia/Shanghai"

// 全局定义
type LocalTime time.Time

func (t LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" || str == "" {
		return nil
	}
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime(t1)
	return err
}

func (t LocalTime) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t LocalTime) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Time(t).In(loc)
}

func (t LocalTime) Value() (driver.Value, error) {
	tTime := time.Time(t)
	return tTime.Format(timeFormat), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
