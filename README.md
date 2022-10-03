# Eventbridge SaaS Partner fURLs

[![View on Construct Hub](https://constructs.dev/badge?package=cdk-eventbridge-partner-processors)](https://constructs.dev/packages/cdk-eventbridge-partner-processors)

[![npm version](https://badge.fury.io/js/cdk-eventbridge-partner-processors.svg)](https://badge.fury.io/js/cdk-eventbridge-partner-processors)
[![PyPI version](https://badge.fury.io/py/a-bigelow.cdk-eventbridge-partner-processors.svg)](https://badge.fury.io/py/a-bigelow.cdk-eventbridge-partner-processors)
[![Go project version](https://badge.fury.io/go/github.com%2Fa-bigelow%2Fcdk-eventbridge-partner-processors-go.svg)](https://badge.fury.io/go/github.com%2Fa-bigelow%2Fcdk-eventbridge-partner-processors-go)

This CDK Construct library provides CDK constructs for the 1st-party (i.e. developed by AWS) lambda fURL webhook receivers for:

- GitHub
- Stripe
- Twilio

## Usage Examples (Simplified)

These examples are consistent for all 3 primary exported constructs of this library:
- `GitHubEventProcessor`
- `TwilioEventProcessor`
- `StripeEventProcessor`

>Note: Click on the above `View on Construct Hub` button to view auto-generated examples in Python/Go

### Typescript

```ts
import { GitHubEventProcessor } from 'cdk-eventbridge-partner-processors';
import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { EventBus } from 'aws-cdk-lib/aws-events';
import { Secret } from 'aws-cdk-lib/aws-secretsmanager';

export class MyStackWithABetterName extends Stack {
    constructor(scope: Construct, id: string, props: StackProps) {
        super(scope, id, props);

        // This library has no opinion on how you reference your EventBus,
        // It just needs to fulfill the IEventBus protocol
        const myEventBus = new EventBus(this, 'TheBestBusEver', {
            eventBusName: 'TheGreatestBus'
        });

        // This library has no opinion on how you reference your secret,
        // It just needs to fulfill the ISecret protocol
        const mySecret = Secret.fromSecretNameV2(this, 'MyNuclearCodeSecret', '/home/recipes/icbm')
        
        // Function will automatically receive permission to:
        // 1. Post events to the given bus
        // 2. Read the given secret
        const githubEventProcessor = new GitHubEventProcessor(this, 'GitHubProcessor', {
            eventBus: myEventBus,
            webhookSecret: mySecret,
            lambdaInvocationAlarmThreshold: 2000,
        })
        
    }
}
```

### Golang
```go
package main

import (
	partner "github.com/a-bigelow/cdk-eventbridge-partner-processors-go"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ClusterStackProps struct {
	awscdk.StackProps
}

func NewClusterStack(scope constructs.Construct, id string, props *ClusterStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	eventProps := awsevents.EventBusProps{EventBusName: jsii.String("name")}

	eventBus := awsevents.NewEventBus(stack, jsii.String("id"), &eventProps)

	secret := secretsmanager.secret.fromSecretNameV2(scope, jsii.String("secret"), jsii.String("secretName"))
	partnerProcessor := partner.GithubEventProcessor{
		EventBus:                       eventBus,
		WebhookSecret:                  secret,
		LambdaInvocationAlarmThreshold: 2000,
	}

	_ = partner.NewGitHubEventProcessor(stack, jsii.String("id"), partnerProcessor)

	return stack
}
```

### Disclaimer
> :warning: The Lambda Functions that handle the actual event processing in this Library are owned and maintained by Amazon Web Services. This CDK Construct library provides a thin deployment wrapper for these functions. Changes made to the S3 location of the functions will break this library. Until I have a way to deterministically track where these things are, please raise an **issue** if you have reason to believe that the functions have moved.

### AWS Documentation
See [here](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html) for additional information.
