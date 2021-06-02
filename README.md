# apigeeids

## Intrusion Detection in Apigee hybrid

This tool can be used to generate Intrusion Detection System (IDS) rules for Snort.
Indeed, Snort is an IDS that can be used in conjunction with Apigee hybrid.

## Why?

You have installed Apigee hybrid on a kubernetes cluster.

Here are some examples explaining **why** a joint use of Apigee hybrid and an
IDS is beneficial:

- You want to be alerted when an API developer configures a target endpoint
or a service callout that is not allowed
- You want to be alerted on any egress traffic reaching botnets or other 
security sensitive destinations
- You want to be alerted about ongoing calls to certain geographies
- You want to be alerted on obsolete TLS versions used on the ingress of your
hybrid runtime
...

## How?

Using the ```apigeeids``` tool, you can generate list of Snort compliant rules.
This rule set can be used to detect and alert about network connections to some
dedicated IP addresses and TCP ports.
In this case, network connections are originating from an Apigee target
endpoint or a service callout.
The [FeodoTracker](https://feodotracker.abuse.ch/blocklist) blocklist is used in this case.  

## Usage

You have got different options for using ```apigeeids```,
which are the following:

- clone the current git repo and execute (*equivalent of go run...*)

        make run
- clone the current git repo and execute (*equivalent of go build...*)

        make build
        ./bin/apigeeids

## Output

Output is a a set of *.rules files that you can use with your Snort IDS instance.
