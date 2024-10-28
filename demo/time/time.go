package time

import (
	"strings"
	"time"
)

const (
	TimeFormatDate      = "2006-01-02 15:04:05"
	TimeFormatDay       = "2006-01-02"
	TimeFormatYear      = "2006"
	TimeFormatDateAdmin = "2006/1/2 15:04:05"
)

// DefaultLocation UTC+8（错误在开发阶段就能发现）
func DefaultLocation() (loc *time.Location) {
	loc, _ = time.LoadLocation("Asia/Shanghai")
	return
}

// GetZeroTimeInLoc 获取0点0时0分的时间（指定时区的时间）
func GetZeroTimeInLoc(now time.Time, loc *time.Location) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
}

func GetLocationUTC13() *time.Location {
	return time.FixedZone("UTC+13", 13*60*60)
}

func GetLocationUTC12() *time.Location {
	return time.FixedZone("UTC+12", 12*60*60)
}

func GetLocationUTC11() *time.Location {
	return time.FixedZone("UTC+11", 11*60*60)
}

func GetLocationUTC10() *time.Location {
	return time.FixedZone("UTC+10", 10*60*60)
}

func GetLocationUTC9() *time.Location {
	return time.FixedZone("UTC+9", 9*60*60)
}

func GetLocationUTC8() *time.Location {
	return time.FixedZone("UTC+8", 8*60*60)
}

func GetLocationUTC7() *time.Location {
	return time.FixedZone("UTC+7", 7*60*60)
}

func GetLocationUTC6Dot30() *time.Location {
	return time.FixedZone("UTC+6:30", 6*60*60+30*60)
}

func GetLocationUTC6() *time.Location {
	return time.FixedZone("UTC+6", 6*60*60)
}

func GetLocationUTC5Dot45() *time.Location {
	return time.FixedZone("UTC+5:45", 5*60*60+45*60)
}

func GetLocationUTC5Dot30() *time.Location {
	return time.FixedZone("UTC+5:30", 5*60*60+30*60)
}

func GetLocationUTC5() *time.Location {
	return time.FixedZone("UTC+5", 5*60*60)
}

func GetLocationUTC4Dot30() *time.Location {
	return time.FixedZone("UTC+4:30", 4*60*60+30*60)
}

func GetLocationUTC4() *time.Location {
	return time.FixedZone("UTC+4", 4*60*60)
}

func GetLocationUTC3Dot30() *time.Location {
	return time.FixedZone("UTC+3:30", 3*60*60+30*60)
}

func GetLocationUTC3() *time.Location {
	return time.FixedZone("UTC+3", 3*60*60)
}

func GetLocationUTC2() *time.Location {
	return time.FixedZone("UTC+2", 2*60*60)
}

func GetLocationUTC1() *time.Location {
	return time.FixedZone("UTC+1", 60*60)
}

func GetLocationUTC0() *time.Location {
	return time.FixedZone("UTC+0", 0)
}

func GetLocationUTCReduce1() *time.Location {
	return time.FixedZone("UTC-1", -60*60)
}

func GetLocationUTCReduce2() *time.Location {
	return time.FixedZone("UTC-2", -2*60*60)
}

func GetLocationUTCReduce3() *time.Location {
	return time.FixedZone("UTC-3", -3*60*60)
}

func GetLocationUTCReduce4() *time.Location {
	return time.FixedZone("UTC-4", -4*60*60)
}

func GetLocationUTCReduce4Dot30() *time.Location {
	return time.FixedZone("UTC-4:30", -4*60*60-30*60)
}

func GetLocationUTCReduce5() *time.Location {
	return time.FixedZone("UTC-5", -5*60*60)
}

func GetLocationUTCReduce6() *time.Location {
	return time.FixedZone("UTC-6", -6*60*60)
}

func GetLocationUTCReduce8() *time.Location {
	return time.FixedZone("UTC-8", -8*60*60)
}

func GetLocationUTCReduce10() *time.Location {
	return time.FixedZone("UTC-10", -10*60*60)
}

func GetLocationUTCReduce11() *time.Location {
	return time.FixedZone("UTC-11", -11*60*60)
}

func GetUTCLocationByRegion(region string) *time.Location {
	region = strings.ToUpper(region)
	switch region {
	case "WS":
		return GetLocationUTC13()
	case "FJ", "KI", "MH", "NR", "NZ", "TV", "WF":
		return GetLocationUTC12()
	case "NC", "NF", "SB", "VU":
		return GetLocationUTC11()
	case "AU", "GU", "FM", "MP", "PG":
		return GetLocationUTC10()
	case "JP", "KP", "KR", "PW", "TL":
		return GetLocationUTC9()
	case "BN", "CN", "HK", "MO", "MY", "MN", "PH", "TW", "SG":
		return GetLocationUTC8()
	case "KH", "CX", "ID", "LA", "TH", "VN":
		return GetLocationUTC7()
	case "CC", "MM":
		return GetLocationUTC6Dot30()
	case "BD", "BT", "IO", "KG":
		return GetLocationUTC6()
	case "NP":
		return GetLocationUTC5Dot45()
	case "IN", "LK":
		return GetLocationUTC5Dot30()
	case "KZ", "MV", "PK", "TJ", "TM", "UZ":
		return GetLocationUTC5()
	case "AF":
		return GetLocationUTC4Dot30()
	case "AM", "AZ", "GE", "MU", "OM", "RE", "SC":
		return GetLocationUTC4()
	case "IR":
		return GetLocationUTC3Dot30()
	case "BH", "KM", "DJ", "ER", "ET", "IQ", "KE", "KW", "MG", "YT", "QA", "RU", "SA", "SO", "SD", "TZ", "UG", "AE", "YE":
		return GetLocationUTC3()
	case "BY", "BW", "BG", "BI", "CY", "EG", "EE", "FI", "PS", "GI", "GR", "IL", "JO", "LV", "LB", "LS", "LY", "LT", "MW", "MD", "MZ", "RO", "RW", "ZA", "SS", "SZ", "SY", "TR", "UA", "ZM", "ZW":
		return GetLocationUTC2()
	case "AL", "DZ", "AD", "AO", "AT", "BE", "BJ", "BA", "BV", "CM", "CF", "TD", "CD", "CG", "HR", "CZ", "DK", "GQ", "FR", "FX", "GA", "DE", "VA", "HU", "IT", "XK", "LI", "LU", "MK", "MT", "MC", "ME", "NA", "NL", "NE", "NG", "NO", "PL", "SM", "RS", "SK", "SI", "ES", "SE", "CH", "TN", "TC":
		return GetLocationUTC1()
	case "BF", "CI", "FO", "GM", "GH", "GL", "GG", "GN", "GW", "HM", "IS", "IE", "IM", "JE", "LR", "ML", "MR", "MA", "PT", "SH", "ST", "SN", "SL", "TG", "GB", "EH":
		return GetLocationUTC0()
	case "CV":
		return GetLocationUTCReduce1()
	case "GS":
		return GetLocationUTCReduce2()
	case "AR", "BR", "GF", "SR", "PM", "UY":
		return GetLocationUTCReduce3()
	case "AI", "AG", "AW", "BB", "BM", "BO", "VG", "CL", "CW", "DM", "DO", "FK", "GD", "GP", "GY", "MQ", "MS", "AN", "PY", "PR", "BL", "KN", "LC", "MF", "VC", "SX", "TT", "US", "VI":
		return GetLocationUTCReduce4()
	case "VE":
		return GetLocationUTCReduce4Dot30()
	case "BS", "CA", "KY", "CO", "CU", "EC", "HT", "JM", "PA", "PE":
		return GetLocationUTCReduce5()
	case "BZ", "CR", "SV", "GT", "HN", "MX", "NI", "UM":
		return GetLocationUTCReduce6()
	case "PN":
		return GetLocationUTCReduce8()
	case "CK", "PF":
		return GetLocationUTCReduce10()
	case "AS", "NU", "TK", "TO":
		return GetLocationUTCReduce11()
	default:
		return GetLocationUTCReduce4()
	}
}

// StrToTimeInLoc 指定时区：解析时间字符串（相当于PHP的strtotime）
func StrToTimeInLoc(str, format string, loc *time.Location) (val int64, err error) {
	// 根据指定格式解析时间字符串
	o, err := time.ParseInLocation(format, str, loc)
	if err != nil {
		return
	}
	val = o.Unix()

	return
}

// StrToTimeLocal 节约参数：解析时间字符串（默认时区）
func StrToTimeLocal(str, format string) (val int64, err error) {
	loc := DefaultLocation()

	return StrToTimeInLoc(str, format, loc)
}

// DateStr 根据时间戳返回时间字符串（东八区）
func DateStr(format string, timestamp int64) string {
	loc := DefaultLocation()
	return DateStrInLoc(format, timestamp, loc)
}

// DateStrInLoc 指定时区：根据时间戳返回时间字符串
func DateStrInLoc(format string, timestamp int64, loc *time.Location) string {
	return time.Unix(timestamp, 0).In(loc).Format(format)
}

// TimestampToTime 东八区：把「时间戳」转为Time.time对象
// 如果要指定时区，使用 TimestampToTimeInLoc
func TimestampToTime(timestamp int64) (time.Time, error) {
	loc := DefaultLocation()

	return time.Unix(timestamp, 0).In(loc), nil
}

// MondayTime 获取某个时间的周一(不带格式)
// 时区由now自身的Location决定
func MondayTime(now time.Time) time.Time {
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	addDate := now.AddDate(0, 0, offset)

	return GetZeroTimeInLoc(addDate, now.Location())
}

// GetWeekCount 计算自1970年1月1日以来的星期数
func GetWeekCount(now time.Time) int64 {
	// 获取自1970年1月1日以来的Unix时间戳（秒数）
	unixTime := now.Unix()

	// 一周的秒数
	const weekSeconds = 7 * 24 * 60 * 60

	// 计算自1970年1月1日以来的星期数
	return unixTime / weekSeconds
}
