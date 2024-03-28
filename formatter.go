package go_resource_formatter

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Unit = string

const (
	Unit_Ki Unit = "Ki"
	Unit_Mi Unit = "Mi"
	Unit_Gi Unit = "Gi"
	Unit_Ti Unit = "Ti"
	Unit_Pi Unit = "Pi"
	Unit_Ei Unit = "Ei"
	Unit_K  Unit = "K"
	Unit_M  Unit = "M"
	Unit_G  Unit = "G"
	Unit_T  Unit = "T"
	Unit_P  Unit = "P"
	Unit_E  Unit = "E"
)

var unitMapping = map[Unit]float64{
	Unit_E:  math.Pow(1000, 6),
	Unit_P:  math.Pow(1000, 5),
	Unit_T:  math.Pow(1000, 4),
	Unit_G:  math.Pow(1000, 3),
	Unit_M:  math.Pow(1000, 2),
	Unit_K:  1000,
	Unit_Ei: math.Pow(1024, 6),
	Unit_Pi: math.Pow(1024, 5),
	Unit_Ti: math.Pow(1024, 4),
	Unit_Gi: math.Pow(1024, 3),
	Unit_Mi: math.Pow(1024, 2),
	Unit_Ki: 1024,
}

// TransferResource 函数用于将含有单位的字符串资源转换为以字节为单位的浮点数值。
// 比如:
// 1Gi 会被转换为 1024 * 1024 * 1024 字节。
// 1GiB 会被转换为 1024 * 1024 * 1024 字节。
// 1G 会被转换为 1000 * 1000 * 1000 字节。
// 1GB 会被转换为 1000 * 1000 * 1000 字节。
func TransferResource(s string) (float64, error) {
	// 正则表达式匹配数字和单位
	re := regexp.MustCompile(`(\d+(?:\.\d+)?)([EPTGMK]?i?)`)
	match := re.FindStringSubmatch(s)

	if match == nil {
		return 0, fmt.Errorf("invalid format")
	}

	// 提取数值和单位
	numberStr, unit := match[1], match[2]
	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number value")
	}
	if multiplier, exists := unitMapping[unit]; exists {
		return number * multiplier, nil
	} else {
		return 0, fmt.Errorf("unknown unit")
	}
}

// Float64Format 格式化浮点数
// f 为格式化字符串
// 比如: Float64Format(3.1415926, "%.2f") 返回 "3.14"
func Float64Format(v float64, f string) string {
	if v == 0 {
		return ""
	}
	return fmt.Sprintf(f, v)
}

// ResourceFormat 按1000进制格式化资源数量
// v 为资源数量，f 为格式化字符串
// f 示例："%.2f"，不要有单位，本返回会自动添加单位
// pretty 为 true 时，会自动去掉小数点后多余的零
// 进制为 1000
// 比如: ResourceFormat(1000, "%.2f") 返回 "1.00KB"
func ResourceFormat(v float64, f string, pretty bool) string {
	if v == 0 {
		return ""
	}
	if v < unitMapping[Unit_K] {
		return fmt.Sprintf(f, v)
	}
	var unit Unit
	if v < unitMapping[Unit_M] {
		unit = Unit_K
	} else if v < unitMapping[Unit_G] {
		unit = Unit_M
	} else if v < unitMapping[Unit_T] {
		unit = Unit_G
	} else if v < unitMapping[Unit_P] {
		unit = Unit_T
	} else {
		unit = Unit_P
	}
	s := Float64Format(v/unitMapping[unit], f)
	if pretty {
		s = PrettyFloatStr(s)
	}
	return s + unit + "B"
}

// ResourceFormatTo 将资源数量格式化为指定单位
// v 为资源数量，to 为目标单位，f 为格式化字符串
// 示例:
// ResourceFormatTo(1000000, Unit_K, "%.2f") 返回 1000.00
// ResourceFormatTo(1000000, Unit_M, "%.2f") 返回 1.00
// ResourceFormatTo(1000000, Unit_Mi, "%.2f") 返回 0.95
func ResourceFormatTo(v float64, to Unit, f string) float64 {
	f2 := unitMapping[to]
	ret := v / f2
	str := Float64Format(ret, f)
	ret, _ = strconv.ParseFloat(str, 64)
	return ret
}

// ResourceFormatToStr 将资源数量格式化为指定单位并返回字符串
// v 为资源数量，to 为目标单位，f 为格式化字符串
// 示例:
// ResourceFormatToStr(1000001, Unit_M, lo.ToPtr("%.2f")) 返回 1MB
// ResourceFormatToStr(1000001, Unit_Mi, lo.ToPtr("%.2f")) 返回 0.95MiB
func ResourceFormatToStr(v float64, to Unit, f string, pretty bool) string {
	f2 := unitMapping[to]
	ret := v / f2
	s := Float64Format(ret, f)
	if pretty {
		s = PrettyFloatStr(s)
	}
	return s + to + "B"
}

// ResourceFormat1024 按1024进制格式化资源数量
// v 为资源数量，f 为格式化字符串
// f 示例："%.2f"，不要有单位，本返回会自动添加单位
// pretty 为 true 时，会自动去掉小数点后多余的零
// 进制为 1024
func ResourceFormat1024(v float64, f string, pretty bool) string {
	if v == 0 {
		return ""
	}
	if v < unitMapping[Unit_Ki] {
		return fmt.Sprintf(f, v)
	}
	var unit Unit
	if v < unitMapping[Unit_Mi] {
		unit = Unit_Ki
	} else if v < unitMapping[Unit_Gi] {
		unit = Unit_Mi
	} else if v < unitMapping[Unit_Ti] {
		unit = Unit_Gi
	} else if v < unitMapping[Unit_Pi] {
		unit = Unit_Ti
	} else {
		unit = Unit_Pi
	}
	s := Float64Format(v/unitMapping[unit], f)
	if pretty {
		s = PrettyFloatStr(s)
	}
	return s + unit + "B"
}

// ResourceStringFormat1024 使用1024进制格式化资源字符串为合适的单位
// 示例：
// ResourceStringFormat1024("1024MiB", "%.2f", true) 返回 1GiB
// ResourceStringFormat1024("0.2GB", "%.2f", true) 返回 190.73MiB
func ResourceStringFormat1024(s string, f string, pretty bool) (string, error) {
	v, err := TransferResource(s)
	if err != nil {
		return "", err
	}
	return ResourceFormat1024(v, f, pretty), nil
}

// ResourceStringFormat 使用1000进制格式化资源字符串为合适的单位
// 示例：
// ResourceStringFormat("1024MiB", "%.2f", true)("1000MB", "%.2f", false) 返回 1.00GB
// ResourceStringFormat("0.2GB", "%.2f", true) 返回 200MB
func ResourceStringFormat(s string, f string, pretty bool) (string, error) {
	v, err := TransferResource(s)
	if err != nil {
		return "", err
	}
	return ResourceFormat(v, f, pretty), nil
}

// PercentFormat 将小数转成百分比格式
func PercentFormat(v float64, f string) string {
	if f == "" {
		f = "%.2f"
	}
	return fmt.Sprintf(f, v*100) + "%"
}

// PrettyFloatStr 格式化浮点数，去除末尾的0和.
func PrettyFloatStr(s string) string {
	s = strings.TrimRight(s, "0")
	return strings.TrimRight(s, ".")
}
