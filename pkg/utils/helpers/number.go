package helpers

import (
	"fmt"
	"math"
	"strconv"
)

// 保留位数四舍五入

func Decimal(value float64, num int) float64 {
	// 保留位数
	keepNum := fmt.Sprintf("%v", num)
	value, _ = strconv.ParseFloat(fmt.Sprintf("%."+keepNum+"f", value), 64)
	return value
}

// 保留两位小数，舍弃尾数，无进位运算
// 主要逻辑就是先乘，trunc之后再除回去，就达到了保留N位小数的效果

func FormatFloat(num float64, decimal int) (float64, error) {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	res := strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
	return strconv.ParseFloat(res, 64)
}
