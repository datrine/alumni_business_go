package jobs

import (
	"fmt"

	paystack "github.com/datrine/alumni_business/Utils/Paystack"
	"github.com/robfig/cron/v3"
)

func init() {
	c := cron.New()
	id, err := c.AddFunc("* * * * * ", paystack.VerifyPaymentsJob)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id)
	c.Start()
}
