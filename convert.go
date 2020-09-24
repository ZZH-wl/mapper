package mapper

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
	"math/big"
	"reflect"
	"strconv"
	"time"
)

// Convert is the target string
type Convert string

// Set string
func (f *Convert) Set(v string) {
	if v != "" {
		*f = Convert(v)
	} else {
		f.Clear()
	}
}

// Clear string
func (f *Convert) Clear() {
	*f = Convert(0x1E)
}

// Exist check string exist
func (f Convert) Exist() bool {
	return string(f) != string(0x1E)
}

// Bool string to bool
func (f Convert) Bool() (bool, error) {
	return strconv.ParseBool(f.String())
}

// Float32 string to float32
func (f Convert) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

// Float64 string to float64
func (f Convert) Float64() (float64, error) {
	return strconv.ParseFloat(f.String(), 64)
}

// Int string to int
func (f Convert) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int(v), err
}

// Int8 string to int8
func (f Convert) Int8() (int8, error) {
	v, err := strconv.ParseInt(f.String(), 10, 8)
	return int8(v), err
}

// Int16 string to int16
func (f Convert) Int16() (int16, error) {
	v, err := strconv.ParseInt(f.String(), 10, 16)
	return int16(v), err
}

// Int32 string to int32
func (f Convert) Int32() (int32, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int32(v), err
}

// Int64 string to int64
func (f Convert) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(f.String(), 10) // octal
		if !ok {
			return v, err
		}
		return ni.Int64(), nil
	}
	return v, err
}

// Uint string to uint
func (f Convert) Uint() (uint, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint(v), err
}

// Uint8 string to uint8
func (f Convert) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

// Uint16 string to uint16
func (f Convert) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(f.String(), 10, 16)
	return uint16(v), err
}

// Uint32 string to uint32
func (f Convert) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint32(v), err
}

// Uint64 string to uint64
func (f Convert) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(f.String(), 10)
		if !ok {
			return v, err
		}
		return ni.Uint64(), nil
	}
	return v, err
}

// String string to string
func (f Convert) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}

// ToString interface to string
func ToString(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

// ToInt64 interface to int64
func ToInt64(value interface{}) (d int64) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		panic(fmt.Errorf("ToInt64 need numeric not `%T`", value))
	}
	return
}

type argInt []int

// get int by index from int slice
func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	}
	if len(args) > 0 {
		r = args[0]
	}
	return
}

// TimeToUnix transform time to Unix time, the number of seconds elapsed
func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

// UnixToTime transform Unix time to local Time
func UnixToTime(tt int64) time.Time {
	return time.Unix(tt, 0)
}

func TimestampToTime(timestampObj *tspb.Timestamp) time.Time {
	t, _ := ptypes.Timestamp(timestampObj)
	return t
}

func TimeToTimestamp(timeObj time.Time) *tspb.Timestamp {
	t, _ := ptypes.TimestampProto(timeObj)
	return t
}

// TimeToUnixLocation transform time to Unix time with time location
// location like "Asia/Shanghai"
func TimeToUnixLocation(t time.Time, location string) (int64, error) {
	timeStr := t.Format("2006-01-02 15:04:05")
	loc, err := time.LoadLocation(location)
	if err != nil {
		return 0, err
	}
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	if err != nil {
		return 0, err
	}
	return tt.Unix(), err
}

// UnixToTimeLocation transform Unix time to local Time with time location
// location like "Asia/Shanghai"
func UnixToTimeLocation(tt int64, location string) (time.Time, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return time.Now(), err
	}
	time.Local = loc
	return time.Unix(tt, 0), nil
}

func TimeStrAutoToTime(timeStr string) time.Time {
	format, timeStr, timeZone := CheckTimeStrFormat(timeStr)
	loc, _ := time.LoadLocation(timeZone)
	formatTime, _ := time.ParseInLocation(format, timeStr, loc)
	return formatTime
}

func CheckTimeStrFormat(timeStr string) (string, string, string) {
	var format, timezone string
	switch len(timeStr) {
	case 8:
		format = "20060102"
	case 10:
		format = "2006-01-02"
	case 14:
		format = "20060102150405"
	case 19:
		format = "2006-01-02 15:04:05"
	case 24, 25, 26:
		timezone = checkTimezone(timeStr[19:])
		format = "2006-01-02 15:04:05"
		timeStr = timeStr[:10] + " " + timeStr[11:19]
	}
	if checkFormat(timeStr) {
		format = "2006/01/02" + format[10:]
	}
	return format, timeStr, timezone
}

func checkFormat(str string) bool {
	if len(str) < 10 || str[4:5] != "/" || str[7:8] != "/" {
		return false
	}
	return true
}

func checkTimezone(timezone string) string {
	switch timezone {
	case "+0000", "+00:00", "Z00:00", "Z0000":
		timezone = "UTC"
	case "+0100", "+01:00", "Z01:00", "Z0100":
		timezone = "Europe/London"
	case "+0200", "+02:00", "Z02:00", "Z0200":
		timezone = "Africa/Cairo"
	case "+0300", "+03:00", "Z03:00", "Z0300":
		timezone = "Europe/Moscow"
	case "+0400", "+04:00", "Z04:00", "Z0400":
		timezone = "Asia/Dubai"
	case "+0500", "+05:00", "Z05:00", "Z0500":
		timezone = "Asia/Yekaterinburg"
	case "+0600", "+06:00", "Z06:00", "Z0600":
		timezone = "Asia/Urumqi"
	case "+0700", "+07:00", "Z07:00", "Z0700":
		timezone = "Asia/Jakarta"
	case "+0800", "+08:00", "Z08:00", "Z0800":
		timezone = "Asia/Shanghai"
	case "+0900", "+09:00", "Z09:00", "Z0900":
		timezone = "Asia/Tokyo"
	case "+1000", "+010:00", "Z010:00", "Z1000":
		timezone = "Australia/Brisbane"
	case "+1100", "+011:00", "Z011:00", "Z1100":
		timezone = "Pacific/Guadalcanal"
	case "+1200", "+012:00", "Z012:00", "Z1200":
		timezone = "Pacific/Nauru"
	case "-0100", "-01:00", "Z013:00", "Z1300":
		timezone = "Atlantic/Cape_Verde"
	case "-0200", "-02:00", "Z014:00", "Z1400":
		timezone = "Atlantic/South_Georgia"
	case "-0300", "-03:00", "Z015:00", "Z1500":
		timezone = "America/Sao_Paulo"
	case "-0400", "-04:00", "Z016:00", "Z1600":
		timezone = "America/Toronto"
	case "-0500", "-05:00", "Z017:00", "Z1700":
		timezone = "America/Cayman"
	case "-0600", "-06:00", "Z018:00", "Z1800":
		timezone = "America/Costa_Rica"
	case "-0700", "-07:00", "Z019:00", "Z1900":
		timezone = "America/Phoenix"
	case "-0800", "-08:00", "Z020:00", "Z2000":
		timezone = "America/Anchorage"
	case "-0900", "-09:00", "Z021:00", "Z2100":
		timezone = "Pacific/Gambier"
	case "-1000", "-010:00", "Z022:00", "Z2200":
		timezone = "Pacific/Honolulu"
	case "-1100", "-011:00", "Z023:00", "Z2300":
		timezone = "Pacific/Midway"
	default:
		timezone = "UTC"
	}
	return timezone
}
