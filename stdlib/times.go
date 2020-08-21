package stdlib

import (
	"time"

	"github.com/zeaphoo/nanojs/v2"
)

var timesModule = map[string]nanojs.Object{
	"format_ansic":        &nanojs.String{Value: time.ANSIC},
	"format_unix_date":    &nanojs.String{Value: time.UnixDate},
	"format_ruby_date":    &nanojs.String{Value: time.RubyDate},
	"format_rfc822":       &nanojs.String{Value: time.RFC822},
	"format_rfc822z":      &nanojs.String{Value: time.RFC822Z},
	"format_rfc850":       &nanojs.String{Value: time.RFC850},
	"format_rfc1123":      &nanojs.String{Value: time.RFC1123},
	"format_rfc1123z":     &nanojs.String{Value: time.RFC1123Z},
	"format_rfc3339":      &nanojs.String{Value: time.RFC3339},
	"format_rfc3339_nano": &nanojs.String{Value: time.RFC3339Nano},
	"format_kitchen":      &nanojs.String{Value: time.Kitchen},
	"format_stamp":        &nanojs.String{Value: time.Stamp},
	"format_stamp_milli":  &nanojs.String{Value: time.StampMilli},
	"format_stamp_micro":  &nanojs.String{Value: time.StampMicro},
	"format_stamp_nano":   &nanojs.String{Value: time.StampNano},
	"nanosecond":          &nanojs.Int{Value: int64(time.Nanosecond)},
	"microsecond":         &nanojs.Int{Value: int64(time.Microsecond)},
	"millisecond":         &nanojs.Int{Value: int64(time.Millisecond)},
	"second":              &nanojs.Int{Value: int64(time.Second)},
	"minute":              &nanojs.Int{Value: int64(time.Minute)},
	"hour":                &nanojs.Int{Value: int64(time.Hour)},
	"january":             &nanojs.Int{Value: int64(time.January)},
	"february":            &nanojs.Int{Value: int64(time.February)},
	"march":               &nanojs.Int{Value: int64(time.March)},
	"april":               &nanojs.Int{Value: int64(time.April)},
	"may":                 &nanojs.Int{Value: int64(time.May)},
	"june":                &nanojs.Int{Value: int64(time.June)},
	"july":                &nanojs.Int{Value: int64(time.July)},
	"august":              &nanojs.Int{Value: int64(time.August)},
	"september":           &nanojs.Int{Value: int64(time.September)},
	"october":             &nanojs.Int{Value: int64(time.October)},
	"november":            &nanojs.Int{Value: int64(time.November)},
	"december":            &nanojs.Int{Value: int64(time.December)},
	"sleep": &nanojs.UserFunction{
		Name:  "sleep",
		Value: timesSleep,
	}, // sleep(int)
	"parse_duration": &nanojs.UserFunction{
		Name:  "parse_duration",
		Value: timesParseDuration,
	}, // parse_duration(str) => int
	"since": &nanojs.UserFunction{
		Name:  "since",
		Value: timesSince,
	}, // since(time) => int
	"until": &nanojs.UserFunction{
		Name:  "until",
		Value: timesUntil,
	}, // until(time) => int
	"duration_hours": &nanojs.UserFunction{
		Name:  "duration_hours",
		Value: timesDurationHours,
	}, // duration_hours(int) => float
	"duration_minutes": &nanojs.UserFunction{
		Name:  "duration_minutes",
		Value: timesDurationMinutes,
	}, // duration_minutes(int) => float
	"duration_nanoseconds": &nanojs.UserFunction{
		Name:  "duration_nanoseconds",
		Value: timesDurationNanoseconds,
	}, // duration_nanoseconds(int) => int
	"duration_seconds": &nanojs.UserFunction{
		Name:  "duration_seconds",
		Value: timesDurationSeconds,
	}, // duration_seconds(int) => float
	"duration_string": &nanojs.UserFunction{
		Name:  "duration_string",
		Value: timesDurationString,
	}, // duration_string(int) => string
	"month_string": &nanojs.UserFunction{
		Name:  "month_string",
		Value: timesMonthString,
	}, // month_string(int) => string
	"date": &nanojs.UserFunction{
		Name:  "date",
		Value: timesDate,
	}, // date(year, month, day, hour, min, sec, nsec) => time
	"now": &nanojs.UserFunction{
		Name:  "now",
		Value: timesNow,
	}, // now() => time
	"parse": &nanojs.UserFunction{
		Name:  "parse",
		Value: timesParse,
	}, // parse(format, str) => time
	"unix": &nanojs.UserFunction{
		Name:  "unix",
		Value: timesUnix,
	}, // unix(sec, nsec) => time
	"add": &nanojs.UserFunction{
		Name:  "add",
		Value: timesAdd,
	}, // add(time, int) => time
	"add_date": &nanojs.UserFunction{
		Name:  "add_date",
		Value: timesAddDate,
	}, // add_date(time, years, months, days) => time
	"sub": &nanojs.UserFunction{
		Name:  "sub",
		Value: timesSub,
	}, // sub(t time, u time) => int
	"after": &nanojs.UserFunction{
		Name:  "after",
		Value: timesAfter,
	}, // after(t time, u time) => bool
	"before": &nanojs.UserFunction{
		Name:  "before",
		Value: timesBefore,
	}, // before(t time, u time) => bool
	"time_year": &nanojs.UserFunction{
		Name:  "time_year",
		Value: timesTimeYear,
	}, // time_year(time) => int
	"time_month": &nanojs.UserFunction{
		Name:  "time_month",
		Value: timesTimeMonth,
	}, // time_month(time) => int
	"time_day": &nanojs.UserFunction{
		Name:  "time_day",
		Value: timesTimeDay,
	}, // time_day(time) => int
	"time_weekday": &nanojs.UserFunction{
		Name:  "time_weekday",
		Value: timesTimeWeekday,
	}, // time_weekday(time) => int
	"time_hour": &nanojs.UserFunction{
		Name:  "time_hour",
		Value: timesTimeHour,
	}, // time_hour(time) => int
	"time_minute": &nanojs.UserFunction{
		Name:  "time_minute",
		Value: timesTimeMinute,
	}, // time_minute(time) => int
	"time_second": &nanojs.UserFunction{
		Name:  "time_second",
		Value: timesTimeSecond,
	}, // time_second(time) => int
	"time_nanosecond": &nanojs.UserFunction{
		Name:  "time_nanosecond",
		Value: timesTimeNanosecond,
	}, // time_nanosecond(time) => int
	"time_unix": &nanojs.UserFunction{
		Name:  "time_unix",
		Value: timesTimeUnix,
	}, // time_unix(time) => int
	"time_unix_nano": &nanojs.UserFunction{
		Name:  "time_unix_nano",
		Value: timesTimeUnixNano,
	}, // time_unix_nano(time) => int
	"time_format": &nanojs.UserFunction{
		Name:  "time_format",
		Value: timesTimeFormat,
	}, // time_format(time, format) => string
	"time_location": &nanojs.UserFunction{
		Name:  "time_location",
		Value: timesTimeLocation,
	}, // time_location(time) => string
	"time_string": &nanojs.UserFunction{
		Name:  "time_string",
		Value: timesTimeString,
	}, // time_string(time) => string
	"is_zero": &nanojs.UserFunction{
		Name:  "is_zero",
		Value: timesIsZero,
	}, // is_zero(time) => bool
	"to_local": &nanojs.UserFunction{
		Name:  "to_local",
		Value: timesToLocal,
	}, // to_local(time) => time
	"to_utc": &nanojs.UserFunction{
		Name:  "to_utc",
		Value: timesToUTC,
	}, // to_utc(time) => time
}

func timesSleep(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	i1, ok := nanojs.ToInt64(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	time.Sleep(time.Duration(i1))
	ret = nanojs.UndefinedValue

	return
}

func timesParseDuration(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	s1, ok := nanojs.ToString(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	dur, err := time.ParseDuration(s1)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &nanojs.Int{Value: int64(dur)}

	return
}

func timesSince(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(time.Since(t1))}

	return
}

func timesUntil(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(time.Until(t1))}

	return
}

func timesDurationHours(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	i1, ok := nanojs.ToInt64(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Float{Value: time.Duration(i1).Hours()}

	return
}

func timesDurationMinutes(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	i1, ok := nanojs.ToInt64(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Float{Value: time.Duration(i1).Minutes()}

	return
}

func timesDurationNanoseconds(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	i1, ok := nanojs.ToInt64(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: time.Duration(i1).Nanoseconds()}

	return
}

func timesDurationSeconds(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	i1, ok := nanojs.ToInt64(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Float{Value: time.Duration(i1).Seconds()}

	return
}

func timesDurationString(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	i1, ok := nanojs.ToInt64(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.String{Value: time.Duration(i1).String()}

	return
}

func timesMonthString(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	i1, ok := nanojs.ToInt64(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.String{Value: time.Month(i1).String()}

	return
}

func timesDate(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 7 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	i1, ok := nanojs.ToInt(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}
	i2, ok := nanojs.ToInt(args[1])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}
	i3, ok := nanojs.ToInt(args[2])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}
	i4, ok := nanojs.ToInt(args[3])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}
	i5, ok := nanojs.ToInt(args[4])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "fifth",
			Expected: "int(compatible)",
			Found:    args[4].TypeName(),
		}
		return
	}
	i6, ok := nanojs.ToInt(args[5])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "sixth",
			Expected: "int(compatible)",
			Found:    args[5].TypeName(),
		}
		return
	}
	i7, ok := nanojs.ToInt(args[6])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "seventh",
			Expected: "int(compatible)",
			Found:    args[6].TypeName(),
		}
		return
	}

	ret = &nanojs.Time{
		Value: time.Date(i1,
			time.Month(i2), i3, i4, i5, i6, i7, time.Now().Location()),
	}

	return
}

func timesNow(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 0 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	ret = &nanojs.Time{Value: time.Now()}

	return
}

func timesParse(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 2 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	s1, ok := nanojs.ToString(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := nanojs.ToString(args[1])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	parsed, err := time.Parse(s1, s2)
	if err != nil {
		ret = wrapError(err)
		return
	}

	ret = &nanojs.Time{Value: parsed}

	return
}

func timesUnix(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 2 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	i1, ok := nanojs.ToInt64(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := nanojs.ToInt64(args[1])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &nanojs.Time{Value: time.Unix(i1, i2)}

	return
}

func timesAdd(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 2 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := nanojs.ToInt64(args[1])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &nanojs.Time{Value: t1.Add(time.Duration(i2))}

	return
}

func timesSub(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 2 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := nanojs.ToTime(args[1])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(t1.Sub(t2))}

	return
}

func timesAddDate(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 4 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	i2, ok := nanojs.ToInt(args[1])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	i3, ok := nanojs.ToInt(args[2])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
		return
	}

	i4, ok := nanojs.ToInt(args[3])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "int(compatible)",
			Found:    args[3].TypeName(),
		}
		return
	}

	ret = &nanojs.Time{Value: t1.AddDate(i2, i3, i4)}

	return
}

func timesAfter(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 2 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := nanojs.ToTime(args[1])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	if t1.After(t2) {
		ret = nanojs.TrueValue
	} else {
		ret = nanojs.FalseValue
	}

	return
}

func timesBefore(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 2 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	t2, ok := nanojs.ToTime(args[1])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.Before(t2) {
		ret = nanojs.TrueValue
	} else {
		ret = nanojs.FalseValue
	}

	return
}

func timesTimeYear(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(t1.Year())}

	return
}

func timesTimeMonth(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(t1.Month())}

	return
}

func timesTimeDay(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(t1.Day())}

	return
}

func timesTimeWeekday(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(t1.Weekday())}

	return
}

func timesTimeHour(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(t1.Hour())}

	return
}

func timesTimeMinute(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(t1.Minute())}

	return
}

func timesTimeSecond(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(t1.Second())}

	return
}

func timesTimeNanosecond(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: int64(t1.Nanosecond())}

	return
}

func timesTimeUnix(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: t1.Unix()}

	return
}

func timesTimeUnixNano(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Int{Value: t1.UnixNano()}

	return
}

func timesTimeFormat(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 2 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	s2, ok := nanojs.ToString(args[1])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
		return
	}

	s := t1.Format(s2)
	if len(s) > nanojs.MaxStringLen {

		return nil, nanojs.ErrStringLimit
	}

	ret = &nanojs.String{Value: s}

	return
}

func timesIsZero(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	if t1.IsZero() {
		ret = nanojs.TrueValue
	} else {
		ret = nanojs.FalseValue
	}

	return
}

func timesToLocal(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Time{Value: t1.Local()}

	return
}

func timesToUTC(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.Time{Value: t1.UTC()}

	return
}

func timesTimeLocation(args ...nanojs.Object) (
	ret nanojs.Object,
	err error,
) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.String{Value: t1.Location().String()}

	return
}

func timesTimeString(args ...nanojs.Object) (ret nanojs.Object, err error) {
	if len(args) != 1 {
		err = nanojs.ErrWrongNumArguments
		return
	}

	t1, ok := nanojs.ToTime(args[0])
	if !ok {
		err = nanojs.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "time(compatible)",
			Found:    args[0].TypeName(),
		}
		return
	}

	ret = &nanojs.String{Value: t1.String()}

	return
}
