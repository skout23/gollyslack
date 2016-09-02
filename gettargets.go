package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/inspector"
)

func gettargetnames(arn string, region string) ([]string, error) {
	var targets []string
	svc := inspector.New(session.New(), &aws.Config{Region: aws.String(region)})
	params := &inspector.DescribeAssessmentTemplatesInput{
		AssessmentTemplateArns: []*string{ // Required
			aws.String(arn), // Required
		},
	}
	resp, _ := svc.DescribeAssessmentTemplates(params)
	for _, template := range resp.AssessmentTemplates {
		fmt.Println("Assessment Target Arn: " + *template.AssessmentTargetArn)
		params := &inspector.DescribeAssessmentTargetsInput{
			AssessmentTargetArns: []*string{ // Required
				aws.String(*template.AssessmentTargetArn), // Required
			},
		}
		tarnresp, _ := svc.DescribeAssessmentTargets(params)
		for _, target := range tarnresp.AssessmentTargets {
			targets = append(targets, *target.Name)
		}
	}
	return targets, nil
}

func main() {
	var arn string
	var region string
	arn = "arn:aws:inspector:us-east-1:133701333029:target/0-M2JoEoLe/template/0-CkJAaP5N"
	region = "us-east-1"
	fmt.Println("Assement Arn: " + arn)
	var targets []string
	targets, _ = gettargetnames(arn, region)
	fmt.Println(targets[0])
}
