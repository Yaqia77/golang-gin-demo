package jobs

import (
	"fmt"
	"time"
)

// MyCronJob is an example cron job
func MyCronJob() {
    fmt.Println("Cron job executed at", time.Now().Format(time.RFC1123))
}


// AnotherCronJob is another example cron job
func AnotherCronJob() {
	fmt.Println("Another cron job executed at", time.Now().Format(time.RFC1123))
}
