package doc

import (
	"testing"
)

var content = []byte(yamlInputContent)

func TestCfDocUsage(t *testing.T) {
	expected := `AWS Cloudformation Template for AutoScalingGroups (ASG)

  Template usage:

       This template depends on the LC template and also the VPC stack.
       This is a NON working template. only for demonstration purposes for cf-doc

`
	actual := Create(content).Usage
	if actual != expected {
		t.Errorf("Test failed, got %s, expected %s", actual, expected)
	}
}

func TestCfDocParams(t *testing.T) {
	expected := `pLaunchConfigurationNameStringLaunch configuration namepVPCZoneIdentifierCommaDelimitedListSubnets List of VPC`
	params := Create(content).Parameters
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
	expected := `asgidAsgBase Logical ID${AWS::StackName}-asgid`
	outs := Create(content).Outputs
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

const yamlInputContent = `#
#  Template usage:
#
#       This template depends on the LC template and also the VPC stack.
#       This is a NON working template. only for demonstration purposes for cf-doc
#

AWSTemplateFormatVersion: '2010-09-09'
Description: 'AWS Cloudformation Template for AutoScalingGroups (ASG)'
# Testing Parameters
Parameters:
  pLaunchConfigurationName:
    Type: String
    Description: Launch configuration name
  pVPCZoneIdentifier:
    Type: CommaDelimitedList
    Description: Subnets List of VPC

Conditions:
  LoadBalancerNamesNull: !Equals
    - !Join
      - ''
      - !Ref pLoadBalancerNames
    - ''

Resources:
  asg:
    Type: AWS::AutoScaling::AutoScalingGroup
    Properties:
      LaunchConfigurationName: !Ref 'pLaunchConfigurationName'
      VPCZoneIdentifier: !Ref 'pVPCZoneIdentifier'
      Cooldown: !Ref 'pCooldown'
      MinSize: !Ref 'pMinSize'
      MaxSize: !Ref 'pMaxSize'
      DesiredCapacity: !Ref 'pDesiredCapacity'
      HealthCheckGracePeriod: !Ref 'pHealthCheckGracePeriod'
      HealthCheckType: !Ref 'pHealthCheckType'
      Tags:
      - Key: Name
        Value: ''
        PropagateAtLaunch: true

Outputs:
  asgid:
    Description: AsgBase Logical ID
    Value: !Ref 'asg'
    Export:
      Name: !Sub "${AWS::StackName}-asgid"`
