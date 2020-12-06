package timeby77

import (
	"math"
	"strconv"
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
