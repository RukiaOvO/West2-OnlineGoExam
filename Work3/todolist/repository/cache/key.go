package cache

import "fmt"

func TaskViewKey(id int64) string {
	return fmt.Sprintf("view:task:%d", id)
}
