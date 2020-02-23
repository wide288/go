package datetime

import "time"

const (
	dateFormat string = "2006-01-02 03:04:05"
)

/**
 * 时间戳的值是否为今天
 * @return true 是, false 不是今天
 */
func IsTody(unix int64) bool {
	tm := time.Unix(unix, 0)
	tmToday := time.Now()

	if tmToday.Day() == tm.Day() {
		return true
	}
	return false
}

/**
 * 二个时间戳是否同一天
 * @return true 是, false 不是今天
 */
func IsSameDay(one, two int64) bool {
	tm := time.Unix(one, 0)
	tmAnother := time.Unix(two, 0)

	if tmAnother.Day() == tm.Day() {
		return true
	}
	return false
}

/**
 * 取得最近几个月的时间戳
 */
func GetLastMonth(num int) (int64, int64) {
	y, m, _ := time.Now().Date()
	dateFrom := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)

	return dateFrom.Unix(), dateFrom.AddDate(0, num, 0).Unix()
}

func TimeUnixToStr(unix int64) string {
	tm := time.Unix(unix, 0)
	return tm.Format(dateFormat)
}