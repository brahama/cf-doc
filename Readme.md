## permanent fork
This repo is a permanent fork of the original project.

### Installation via brew
```
brew tap daniel-ciaglia/taps
brew tap-pin daniel-ciaglia/taps
brew install daniel-ciaglia/taps/cf-doc
```
See the [Homebrew documentation](https://docs.brew.sh/Taps) for more information on managing taps.


[![Build Status](https://travis-ci.org/brahama/cf-doc.svg?branch=dev)](https://travis-ci.org/brahama/cf-doc)
  `cf-doc(1)` &sdot; a quick utility to generate docs from Cloudformation templates.

  Inspired and "transformed" from [Segmentio Terraform](https://github.com/segmentio/terraform-docs)  


## Features

  - View docs for inputs and outputs
  - Generate docs for inputs and outputs
  - Generate JSON docs (for customizing presentation)
  - Generate markdown tables of inputs and outputs

## Installation (Pending...)

  - `go get github.com/.....`
  - [Binaries](https://github.com/......)
  - `brew install cloudformation-docs` (on macOS)

## Usage

```bash

  Usage:
    cf-doc [json | md | markdown] <file>...
    cf-doc -h | --help

  Examples:

    # View inputs and outputs
    $ cf-doc my-template.yaml

    # Generate a JSON of inputs and outputs
    $ cf-doc my-template.yaml

    # Generate markdown tables of inputs and outputs
    $ cf-doc md my-template.yaml


  Options:
    -h, --help     show help information

```

## Example

Given a simple template at `./_example/asg.yaml`:

```yaml
#
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


Outputs:
  asgid:
    Description: AsgBase Logical ID
    Value: !Ref 'asg'
    Export:
      Name: !Sub "${AWS::StackName}-asgid"


```

To view docs:

```bash
$ cf-doc json _example/asg.yaml
```

To output JSON docs:

```bash
$ cf-doc json _example/asg.yaml
{
    "Usage": "AWS Cloudformation Template for AutoScalingGroups (ASG)\n\n  Template usage:\n \n       This template depends on the LC template and also the VPC stack.\n       This is a NON working template. only for demonstration purposes for cf-doc\n\n",
    "Parameters": [
        {
            "Name": "pLaunchConfigurationName",
            "Description": "Launch configuration name",
            "Default": "",
            "Type": "String",
            "AllowedValues": ""
        },
        {
            "Name": "pVPCZoneIdentifier",
            "Description": "Subnets List of VPC",
            "Default": "",
            "Type": "CommaDelimitedList",
            "AllowedValues": ""
        }
    ],
    "Outputs": [
        {
            "Name": "asgid",
            "Description": "AsgBase Logical ID",
            "Export": "${AWS::StackName}-asgid"
        }
    ]
}
```

To output markdown docs:

```bash
$ cf-doc md _example/asg.yaml

AWS Cloudformation Template for AutoScalingGroups (ASG)

  Template usage:

       This template depends on the LC template and also the VPC stack.
       This is a NON working template. only for demonstration purposes for cf-doc



## Parameters

| Name | Description | Type | Default | Allowed Values |
|------|-------------|:----:|:-------:|:---------------|
| pLaunchConfigurationName | Launch configuration name | String |  |  |
| pVPCZoneIdentifier | Subnets List of VPC | CommaDelimitedList |  |  |

## Outputs

| Name | Description | Export |
|------|-------------|--------|
| asgid | AsgBase Logical ID | ${AWS::StackName}-asgid |

```

## License

MIT License

Copyright (c) 2017 Brahama

Copyright (c) 2017 Segment, Inc

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
