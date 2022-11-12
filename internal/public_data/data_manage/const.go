package data_manage

type Period int

const (
	Minutely Period = iota
	Five_Minutely
	Quarterly
	Hourly
	Daily
	Weekly
	Monthly
)

func BuildPeriod(period Period) string {
	switch period {
	case Minutely:
		return "1m"
	case Five_Minutely:
		return "5m"
	case Quarterly:
		return "15m"
	case Hourly:
		return "1h"
	case Daily:
		return "1d"
	case Weekly:
		return "1w"
	case Monthly:
		return "1m"
	}
	return ""
}
