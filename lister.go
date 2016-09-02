package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
	"time"
)

func currentdate() string {
	return time.Now().Format("20060102")
}

type issue struct {
	Id       string
	Severity string
}

type report struct {
	Issues []issue
}

func main() {
	var arn string
	var region string
	var trun string
	// var summary report
	arn = "arn:aws:inspector:us-east-1:133701333029:target/0-M2JoEoLe/template/0-CkJAaP5N"
	region = "us-east-1"
	// datematch := currentdate() + "*"
	datematch := "20160831*"
	counts := make(map[string]int)
	cleaned := make(map[string]string)

	svc := inspector.New(session.New(), &aws.Config{Region: aws.String(region)})
	params := &inspector.ListAssessmentRunsInput{
		AssessmentTemplateArns: []*string{
			aws.String(arn),
		},
		Filter: &inspector.AssessmentRunFilter{
			NamePattern: aws.String(datematch),
			States: []*string{
				aws.String("COMPLETED"),
			},
		},
	}

	resp, _ := svc.ListAssessmentRuns(params)
	for _, run := range resp.AssessmentRunArns {
		trun = *run
	}

	tparams := &inspector.ListFindingsInput{
		AssessmentRunArns: []*string{
			aws.String(trun), // Required
		},
		MaxResults: aws.Int64(50),
	}
	fresp, _ := svc.ListFindings(tparams)

	for _, findingarn := range fresp.FindingArns {
		params := &inspector.DescribeFindingsInput{
			FindingArns: []*string{ // Required
				aws.String(*findingarn), // Required
			},
		}

		resp, _ := svc.DescribeFindings(params)
		for _, finding := range resp.Findings {
			// if *finding.Severity == "High" || *finding.Severity == "Medium" {
			counts[*finding.Severity]++
			cleaned[*finding.Id] = *finding.Severity
			// f := issue{Id: *finding.Id, Severity: *finding.Severity}
			//summary.Issues = append(summary.Issues, f)
			//}
		}
	}
	// fmt.Println(counts)
	// fmt.Println(len(cleaned))
	// fmt.Println(len(summary.Issues))
	for rate, count := range counts {
		fmt.Printf("%s:\t%d\n", rate, count)
	}
	for issue, rate := range cleaned {
		if rate == "High" {
			fmt.Println(rate + "\t" + issue)
		}

	}
	//for _, issue := range summary.Issues {
	//	if issue.Severity == "High" {
	//		fmt.Println(issue)
	//	}
	//}
}
