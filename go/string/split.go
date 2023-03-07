package string

import (
	"fmt"
	"strconv"
	"strings"
)

// SplitNum Split the string into numbers according to the rules, where a dash indicates a continuous group of data
// and a comma separates individual data.
// 0-9,25-35,233,244
func SplitNum(value string) (map[string]struct{}, error) {
	sMap := make(map[string]struct{})
	valueArr := strings.Split(value, ",")
	for _, valueItem := range valueArr {
		if valueItem == "" {
			continue
		}

		numArr := strings.Split(valueItem, "-")
		if len(numArr) != 2 && len(numArr) != 1 {
			return nil, fmt.Errorf("%s value is invaildate", valueItem)
		}

		if len(numArr) == 1 {
			_, err := strconv.ParseInt(numArr[0], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("%s value's %s start index is not int64", valueItem, numArr[0])
			}
			sMap[numArr[0]] = struct{}{}
			continue
		}

		if len(numArr) == 2 && numArr[0] == "" || numArr[1] == "" {
			return nil, fmt.Errorf("%s value is invaildate", valueItem)
		}

		startIndex, err := strconv.ParseInt(numArr[0], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%s value's %s start index is not int64", valueItem, numArr[0])
		}

		endIndex, err := strconv.ParseInt(numArr[1], 10, 64)
		if err != nil {
			return nil, fmt.Errorf("%s value's %s end index is not int64", valueItem, numArr[1])
		}

		for i := startIndex; i < endIndex; i++ {
			sMap[fmt.Sprintf("%d", i)] = struct{}{}
		}
	}
	return sMap, nil
}
