# Eventbridge SaaS Partner fURLs

[![npm version](https://badge.fury.io/js/cdk-eventbridge-partner-processors.svg)](https://badge.fury.io/js/cdk-eventbridge-partner-processors)
[![PyPI version](https://badge.fury.io/py/a-bigelow.cdk-eventbridge-partner-processors.svg)](https://badge.fury.io/py/a-bigelow.cdk-eventbridge-partner-processors)
[![Go project version](https://badge.fury.io/go/github.com%2Fa-bigelow%2Fcdk-eventbridge-partner-processors.svg)](https://badge.fury.io/go/github.com%2Fa-bigelow%2Fcdk-eventbridge-partner-processors)

This CDK Construct library provides CDK constructs for the 1st-party (i.e. developed by AWS) lambda fURL webhook receivers for:

* GitHub
* Stripe
* Twilio

## Usage Examples (Simplified)

These examples are consistent for all 3 primary exported constructs of this library:

* `GitHubEventProcessor`
* `TwilioEventProcessor`
* `StripeEventProcessor`

### Typescript

```go
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

### Disclaimer

> :warning: The Lambda Functions that handle the actual event processing in this Library are owned and maintained by Amazon Web Services. This CDK Construct library provides a thin deployment wrapper for these functions. Changes made to the S3 location of the functions will break this library. Until I have a way to deterministically track where these things are, please raise an **issue** if you have reason to believe that the functions have moved.

### AWS Documentation

See [here](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html) for additional information.
