package environment

import "os"

func GetEnv(k, v string) string {
	val, ok := os.LookupEnv(k)
	if !ok {
		return v
	}

	return val
}
