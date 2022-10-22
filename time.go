package timeby77

import (
	"math"
	"strconv"
	"time"
)

func HowLongAgo(timebefore int64, timenow int64) (strres string) {
	var TimeUint = [6]int64{365 * 86400, 30 * 86400, 86400, 3600, 60, 1}
	var StrUnit = [6]string{"年前", "月前", "天前", "小時前", "分鐘前", "秒鐘前"}
	strres = "剛剛"

	if timebefore > timenow {
		return
	}

	diff := timenow - timebefore

	for i := 0; i < len(TimeUint); i++ {
		if diff < TimeUint[i] {
			continue
		}
		var temp = math.Floor(float64(diff / TimeUint[i]))
		diff = diff % TimeUint[i]
		if temp > 0 {
			var tempStr string
			tempStr = strconv.FormatFloat(temp, 'f', -1, 64)
			strres = tempStr + StrUnit[i]
		}
		break
	}
	return strres
}

func Time2DateType1(Input time.Time) (OutPut string) {

	OutPut = Input.Format("01/02 15:04")
	return OutPut
}

func GetAgeNew(BirthDay time.Time, Now time.Time) (age int) {
	YearDiff := Now.Year() - BirthDay.Year()

	if YearDiff < 0 {
		return 0
	}

	if BirthDay.Month() > Now.Month() {
		YearDiff = YearDiff - 1

	} else if Now.Month() == BirthDay.Month() {
		if BirthDay.Day() > Now.Day() {
			YearDiff = YearDiff - 1

		}
	}

	if YearDiff < 18 {
		YearDiff = 0
	}

	return YearDiff

}

func GetConstellation(BirthDay time.Time) (star string) {

	month := BirthDay.Month()
	day := BirthDay.Day()
	star = "未知"
	if month <= 0 || month >= 13 {
		star = "未知"
	}
	if day <= 0 || day >= 32 {
		star = "未知"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "水瓶座"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "雙魚座"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "白羊座"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "金牛座"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "雙子座"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "巨蟹座"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "獅子座"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "處女座"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "天枰座"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "天蠍座"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "射手座"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "摩羯座"
	}
	return star
}

func GetConstellationEn(BirthDay time.Time) (star string) {

	month := BirthDay.Month()
	day := BirthDay.Day()
	star = "Unknown"
	if month <= 0 || month >= 13 {
		star = "Unknown"
	}
	if day <= 0 || day >= 32 {
		star = "Unknown"
	}
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) {
		star = "Aquarius"
	}
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) {
		star = "Pisces"
	}
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) {
		star = "Aries"
	}
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) {
		star = "Taurus"
	}
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
		star = "Gemini"
	}
	if (month == 6 && day >= 22) || (month == 7 && day <= 22) {
		star = "Cancer"
	}
	if (month == 7 && day >= 23) || (month == 8 && day <= 22) {
		star = "Leo"
	}
	if (month == 8 && day >= 23) || (month == 9 && day <= 22) {
		star = "Virgo"
	}
	if (month == 9 && day >= 23) || (month == 10 && day <= 22) {
		star = "Libra"
	}
	if (month == 10 && day >= 23) || (month == 11 && day <= 21) {
		star = "Scorpio"
	}
	if (month == 11 && day >= 22) || (month == 12 && day <= 21) {
		star = "Sagittarius"
	}
	if (month == 12 && day >= 22) || (month == 1 && day <= 19) {
		star = "Capricorn"
	}
	return star
}
