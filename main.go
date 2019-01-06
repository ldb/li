package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	id := flag.BoolP("id", "i", false, "show instance-ID")
	status := flag.BoolP("status", "s", false, "show instance status")
	instanceType := flag.BoolP("type", "t", false, "show instance type")
	zone := flag.BoolP("zone", "z", false, "show availability-zone")
	help := flag.BoolP("help", "h", false, "show help")
	detailed := flag.BoolP("details", "d", false, "show all information")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ec2Svc := ec2.New(sess)

	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		log.Fatalf("error describing instances: %v", err)
	}

	query := strings.Join(flag.Args(), " ")
	b := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			for _, t := range i.Tags {
				if *t.Key == "Name" && matches(*t.Value, query) {
					s := *t.Value

					if i.PrivateIpAddress != nil {
						s = s + "\t" + *i.PrivateIpAddress
					}

					if (*id || *detailed) && i.InstanceId != nil {
						s = s + "\t" + *i.InstanceId
					}

					if (*status || *detailed) && i.State.Name != nil {
						s = s + "\t" + *i.State.Name
					}

					if (*instanceType || *detailed) && i.InstanceType != nil {
						s = s + "\t" + *i.InstanceType
					}

					if (*zone || *detailed) && i.Placement.AvailabilityZone != nil {
						s = s + "\t" + *i.Placement.AvailabilityZone
					}

					fmt.Fprintln(b, s)
				}
			}
		}
	}
	b.Flush()
}

func matches(value string, query string) bool {
	if query == "" {
		return true
	}

	l, q := strings.ToLower(value), strings.ToLower(query)
	if strings.Contains(l, q) {
		return true
	}

	return false
}
