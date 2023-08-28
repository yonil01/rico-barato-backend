package helper

import (
	"backend-ccff/internal/logger"
	"encoding/json"
	"fmt"
	"strings"
)

func SliceToString(s []string) string {
	var elements string

	for i, e := range s {
		if e == "" {
			continue
		}
		if i == 0 {
			elements = fmt.Sprintf("'%s'", strings.ToLower(e))
		} else {
			elements = fmt.Sprintf("%s,'%s'", elements, strings.ToLower(e))
		}

	}
	return elements
}

func SliceInt64ToString(ids []int64) string {
	var elements string
	for i, e := range ids {
		if e == 0 {
			continue
		}
		if i == 0 {
			elements = fmt.Sprintf("'%d'", e)
		} else {
			elements = fmt.Sprintf("%s,'%d'", elements, e)
		}

	}
	return elements
}

func SliceInt64ToStringInteger(ids []int64) string {
	var elements string
	for i, e := range ids {
		if e == 0 {
			continue
		}
		if i == 0 {
			elements = fmt.Sprintf("%d", e)
		} else {
			elements = fmt.Sprintf("%s,%d", elements, e)
		}

	}
	return elements
}

func SliceIntToString(ids []int) string {
	var elements string
	for i, e := range ids {
		if i == 0 {
			elements = fmt.Sprintf("'%d'", e)
		} else {
			elements = fmt.Sprintf("%s,'%d'", elements, e)
		}

	}
	return elements
}

func SlicePointerToString(s []*string) string {
	var elements string
	for i, e := range s {
		if e == nil {
			continue
		}
		if i == 0 {
			elements = fmt.Sprintf("'%s'", strings.ToLower(*e))
		} else {
			elements = fmt.Sprintf("%s,'%s'", elements, strings.ToLower(*e))
		}

	}
	return elements
}

func InterfaceToMapInterface(s interface{}) (map[string]interface{}, error) {
	rs := make(map[string]interface{}, 0)
	data, err := json.Marshal(s)
	if err != nil {
		logger.Error.Println("couldn't convert interface to map[string]interface{}: ", err)
		return rs, err
	}
	err = json.Unmarshal(data, &rs)
	if err != nil {
		logger.Error.Println("couldn't convert interface to map[string]interface{}: ", err)
		return rs, err
	}
	return rs, nil

}
