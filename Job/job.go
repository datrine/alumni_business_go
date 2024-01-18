package job

import (
	"fmt"

	paystack "github.com/datrine/alumni_business/Utils/Paystack"
	"github.com/robfig/cron/v3"
)

func init() {
	c := cron.New()
	_, err := c.AddFunc("* * * * * ", paystack.VerifyPaymentsJob)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Start()
}
