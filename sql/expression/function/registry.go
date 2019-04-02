package function

import (
	"math"

	"gopkg.in/src-d/go-mysql-server.v0/sql"
	"gopkg.in/src-d/go-mysql-server.v0/sql/expression/function/aggregation"
)

// Defaults is the function map with all the default functions.
var Defaults = []sql.Function{
	sql.Function1{
		Name: "count",
		Fn:   func(e sql.Expression) sql.Expression { return aggregation.NewCount(e) },
	},
	sql.Function1{
		Name: "min",
		Fn:   func(e sql.Expression) sql.Expression { return aggregation.NewMin(e) },
	},
	sql.Function1{
		Name: "max",
		Fn:   func(e sql.Expression) sql.Expression { return aggregation.NewMax(e) },
	},
	sql.Function1{
		Name: "avg",
		Fn:   func(e sql.Expression) sql.Expression { return aggregation.NewAvg(e) },
	},
	sql.Function1{
		Name: "sum",
		Fn:   func(e sql.Expression) sql.Expression { return aggregation.NewSum(e) },
	},
	sql.Function1{Name: "is_binary", Fn: NewIsBinary},
	sql.FunctionN{Name: "substring", Fn: NewSubstring},
	sql.FunctionN{Name: "mid", Fn: NewSubstring},
	sql.FunctionN{Name: "substr", Fn: NewSubstring},
	sql.Function1{Name: "year", Fn: NewYear},
	sql.Function1{Name: "month", Fn: NewMonth},
	sql.Function1{Name: "day", Fn: NewDay},
	sql.Function1{Name: "weekday", Fn: NewWeekday},
	sql.Function1{Name: "hour", Fn: NewHour},
	sql.Function1{Name: "minute", Fn: NewMinute},
	sql.Function1{Name: "second", Fn: NewSecond},
	sql.Function1{Name: "dayofweek", Fn: NewDayOfWeek},
	sql.Function1{Name: "dayofyear", Fn: NewDayOfYear},
	sql.Function1{Name: "array_length", Fn: NewArrayLength},
	sql.Function2{Name: "split", Fn: NewSplit},
	sql.FunctionN{Name: "concat", Fn: NewConcat},
	sql.FunctionN{Name: "concat_ws", Fn: NewConcatWithSeparator},
	sql.FunctionN{Name: "coalesce", Fn: NewCoalesce},
	sql.Function1{Name: "lower", Fn: NewLower},
	sql.Function1{Name: "upper", Fn: NewUpper},
	sql.Function1{Name: "ceiling", Fn: NewCeil},
	sql.Function1{Name: "ceil", Fn: NewCeil},
	sql.Function1{Name: "floor", Fn: NewFloor},
	sql.FunctionN{Name: "round", Fn: NewRound},
	sql.Function0{Name: "connection_id", Fn: NewConnectionID},
	sql.Function1{Name: "soundex", Fn: NewSoundex},
	sql.FunctionN{Name: "json_extract", Fn: NewJSONExtract},
	sql.Function1{Name: "ln", Fn: NewLogBaseFunc(float64(math.E))},
	sql.Function1{Name: "log2", Fn: NewLogBaseFunc(float64(2))},
	sql.Function1{Name: "log10", Fn: NewLogBaseFunc(float64(10))},
	sql.FunctionN{Name: "log", Fn: NewLog},
	sql.FunctionN{Name: "rpad", Fn: NewPadFunc(rPadType)},
	sql.FunctionN{Name: "lpad", Fn: NewPadFunc(lPadType)},
	sql.Function1{Name: "sqrt", Fn: NewSqrt},
	sql.Function2{Name: "pow", Fn: NewPower},
	sql.Function2{Name: "power", Fn: NewPower},
	sql.Function1{Name: "ltrim", Fn: NewTrimFunc(lTrimType)},
	sql.Function1{Name: "rtrim", Fn: NewTrimFunc(rTrimType)},
	sql.Function1{Name: "trim", Fn: NewTrimFunc(bTrimType)},
	sql.Function1{Name: "reverse", Fn: NewReverse},
	sql.Function2{Name: "repeat", Fn: NewRepeat},
	sql.Function3{Name: "replace", Fn: NewReplace},
	sql.Function2{Name: "ifnull", Fn: NewIfNull},
	sql.Function2{Name: "nullif", Fn: NewNullIf},
	sql.Function0{Name: "now", Fn: NewNow},
	sql.Function1(Name: "sleep", Fn: NewSleep},
}
