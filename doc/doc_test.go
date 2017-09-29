package doc

import (
	"io/ioutil"
	"log"
	"testing"
)

func contentHelper(file string) *[]byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return &content
}
func TestCfDocUsage(t *testing.T) {
	content := contentHelper("../_example/asg.yaml")
	expected := `AWS Cloudformation Template for AutoScalingGroups (ASG)

  Template usage:

       This template depends on the LC template and also the VPC stack.
       This is a NON working template. only for demonstration purposes for cf-doc

`
	actual := Create(*content).Usage
	if actual != expected {
		t.Errorf("Test failed, got %s, expected %s", actual, expected)
	}
}

func TestCfDocParams(t *testing.T) {
	content := contentHelper("../_example/asg.yaml")
	expected := `pLaunchConfigurationNameStringLaunch configuration namepVPCZoneIdentifierCommaDelimitedListSubnets List of VPC`
	params := Create(*content).Parameters
	var actual string
	for k := range params {
		actual += params[k].Name
		actual += params[k].Type
		actual += params[k].Description
		actual += params[k].Default
		actual += params[k].AllowedValues
	}
	if actual != expected {
		t.Errorf("Test failed, got %s, expected %s", actual, expected)
	}
}

func TestCfDocOut(t *testing.T) {
	content := contentHelper("../_example/asg.yaml")
	expected := `LogGroupLog group of ECS cluster.${AWS::StackName}-LogGroupasgidAsgBase Logical ID${AWS::StackName}-asgid`
	outs := Create(*content).Outputs
	var actual string
	for k := range outs {
		actual += outs[k].Name
		actual += outs[k].Description
		actual += outs[k].Export
	}
	if actual != expected {
		t.Errorf("Test failed, got %s, expected %s", actual, expected)
	}
}
