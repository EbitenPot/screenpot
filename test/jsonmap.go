// Created by EldersJavas(EldersJavas&gmail.com)

package test

import "strings"

func SerJSON(map1 map[string]string) string {
	result := ""
	for k, v := range map1 {
		if result == "" {
			if strings.HasPrefix(v, "{") == true || strings.HasPrefix(v, "[{") == true {
				result = strings.Join([]string{`"`, k, `":`, v}, "")
			} else {
				result = strings.Join([]string{`"`, k, `":`, `"`, v, `"`}, "")
			}

		} else {
			if strings.HasPrefix(v, "{") == true || strings.HasPrefix(v, "[{") == true {
				result += strings.Join([]string{`,"`, k, `":`, v}, "")
			} else {
				result += strings.Join([]string{`,"`, k, `":`, `"`, v, `"`}, "")

			}
		}
	}

	if result != "" {
		result = strings.Join([]string{"{", result, "}"}, "")
	}
	return result

}
