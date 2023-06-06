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

# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### GitHubEventProcessor <a name="GitHubEventProcessor" id="cdk-eventbridge-partner-processors.GitHubEventProcessor"></a>

CDK wrapper for the GitHub Eventbridge processor.

> [https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github)

#### Initializers <a name="Initializers" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.Initializer"></a>

```typescript
import { GitHubEventProcessor } from 'cdk-eventbridge-partner-processors'

new GitHubEventProcessor(scope: Construct, id: string, props: GitHubProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubEventProcessor.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubEventProcessor.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubEventProcessor.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-eventbridge-partner-processors.GitHubProps">GitHubProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-eventbridge-partner-processors.GitHubProps">GitHubProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubEventProcessor.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubEventProcessor.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.isConstruct"></a>

```typescript
import { GitHubEventProcessor } from 'cdk-eventbridge-partner-processors'

GitHubEventProcessor.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubEventProcessor.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubEventProcessor.property.invocationAlarm">invocationAlarm</a></code> | <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm">InvocationAlarm</a></code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubEventProcessor.property.partnerEventsFunction">partnerEventsFunction</a></code> | <code>aws-cdk-lib.aws_lambda.Function</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `invocationAlarm`<sup>Required</sup> <a name="invocationAlarm" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.property.invocationAlarm"></a>

```typescript
public readonly invocationAlarm: InvocationAlarm;
```

- *Type:* <a href="#cdk-eventbridge-partner-processors.InvocationAlarm">InvocationAlarm</a>

---

##### `partnerEventsFunction`<sup>Required</sup> <a name="partnerEventsFunction" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.property.partnerEventsFunction"></a>

```typescript
public readonly partnerEventsFunction: Function;
```

- *Type:* aws-cdk-lib.aws_lambda.Function

---


### InvocationAlarm <a name="InvocationAlarm" id="cdk-eventbridge-partner-processors.InvocationAlarm"></a>

Cloudwatch Alarm used across this construct library.

#### Initializers <a name="Initializers" id="cdk-eventbridge-partner-processors.InvocationAlarm.Initializer"></a>

```typescript
import { InvocationAlarm } from 'cdk-eventbridge-partner-processors'

new InvocationAlarm(scope: Construct, id: string, props: InvocationAlarmProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarmProps">InvocationAlarmProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-eventbridge-partner-processors.InvocationAlarm.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-eventbridge-partner-processors.InvocationAlarm.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-eventbridge-partner-processors.InvocationAlarm.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-eventbridge-partner-processors.InvocationAlarmProps">InvocationAlarmProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk-eventbridge-partner-processors.InvocationAlarm.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-eventbridge-partner-processors.InvocationAlarm.isConstruct"></a>

```typescript
import { InvocationAlarm } from 'cdk-eventbridge-partner-processors'

InvocationAlarm.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-eventbridge-partner-processors.InvocationAlarm.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-eventbridge-partner-processors.InvocationAlarm.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---


### PartnerProcessor <a name="PartnerProcessor" id="cdk-eventbridge-partner-processors.PartnerProcessor"></a>

Abstract class for Lambda-driven Eventbridge integrations.

This only works because the pattern for the S3 Keys lines up to: lambda-templates/<partner>-lambdasrc.zip

> [https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html)

#### Initializers <a name="Initializers" id="cdk-eventbridge-partner-processors.PartnerProcessor.Initializer"></a>

```typescript
import { PartnerProcessor } from 'cdk-eventbridge-partner-processors'

new PartnerProcessor(scope: Construct, id: string, props: PartnerFunctionProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerProcessor.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerProcessor.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerProcessor.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-eventbridge-partner-processors.PartnerFunctionProps">PartnerFunctionProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-eventbridge-partner-processors.PartnerProcessor.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-eventbridge-partner-processors.PartnerProcessor.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-eventbridge-partner-processors.PartnerProcessor.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-eventbridge-partner-processors.PartnerFunctionProps">PartnerFunctionProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerProcessor.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk-eventbridge-partner-processors.PartnerProcessor.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerProcessor.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-eventbridge-partner-processors.PartnerProcessor.isConstruct"></a>

```typescript
import { PartnerProcessor } from 'cdk-eventbridge-partner-processors'

PartnerProcessor.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-eventbridge-partner-processors.PartnerProcessor.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerProcessor.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerProcessor.property.invocationAlarm">invocationAlarm</a></code> | <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm">InvocationAlarm</a></code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerProcessor.property.partnerEventsFunction">partnerEventsFunction</a></code> | <code>aws-cdk-lib.aws_lambda.Function</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-eventbridge-partner-processors.PartnerProcessor.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `invocationAlarm`<sup>Required</sup> <a name="invocationAlarm" id="cdk-eventbridge-partner-processors.PartnerProcessor.property.invocationAlarm"></a>

```typescript
public readonly invocationAlarm: InvocationAlarm;
```

- *Type:* <a href="#cdk-eventbridge-partner-processors.InvocationAlarm">InvocationAlarm</a>

---

##### `partnerEventsFunction`<sup>Required</sup> <a name="partnerEventsFunction" id="cdk-eventbridge-partner-processors.PartnerProcessor.property.partnerEventsFunction"></a>

```typescript
public readonly partnerEventsFunction: Function;
```

- *Type:* aws-cdk-lib.aws_lambda.Function

---


### StripeEventProcessor <a name="StripeEventProcessor" id="cdk-eventbridge-partner-processors.StripeEventProcessor"></a>

CDK wrapper for the Stripe Eventbridge processor.

> [https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github)

#### Initializers <a name="Initializers" id="cdk-eventbridge-partner-processors.StripeEventProcessor.Initializer"></a>

```typescript
import { StripeEventProcessor } from 'cdk-eventbridge-partner-processors'

new StripeEventProcessor(scope: Construct, id: string, props: StripeProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.StripeEventProcessor.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.StripeEventProcessor.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.StripeEventProcessor.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-eventbridge-partner-processors.StripeProps">StripeProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-eventbridge-partner-processors.StripeEventProcessor.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-eventbridge-partner-processors.StripeEventProcessor.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-eventbridge-partner-processors.StripeEventProcessor.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-eventbridge-partner-processors.StripeProps">StripeProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.StripeEventProcessor.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk-eventbridge-partner-processors.StripeEventProcessor.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.StripeEventProcessor.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-eventbridge-partner-processors.StripeEventProcessor.isConstruct"></a>

```typescript
import { StripeEventProcessor } from 'cdk-eventbridge-partner-processors'

StripeEventProcessor.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-eventbridge-partner-processors.StripeEventProcessor.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.StripeEventProcessor.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-eventbridge-partner-processors.StripeEventProcessor.property.invocationAlarm">invocationAlarm</a></code> | <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm">InvocationAlarm</a></code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.StripeEventProcessor.property.partnerEventsFunction">partnerEventsFunction</a></code> | <code>aws-cdk-lib.aws_lambda.Function</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-eventbridge-partner-processors.StripeEventProcessor.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `invocationAlarm`<sup>Required</sup> <a name="invocationAlarm" id="cdk-eventbridge-partner-processors.StripeEventProcessor.property.invocationAlarm"></a>

```typescript
public readonly invocationAlarm: InvocationAlarm;
```

- *Type:* <a href="#cdk-eventbridge-partner-processors.InvocationAlarm">InvocationAlarm</a>

---

##### `partnerEventsFunction`<sup>Required</sup> <a name="partnerEventsFunction" id="cdk-eventbridge-partner-processors.StripeEventProcessor.property.partnerEventsFunction"></a>

```typescript
public readonly partnerEventsFunction: Function;
```

- *Type:* aws-cdk-lib.aws_lambda.Function

---


### TwilioEventProcessor <a name="TwilioEventProcessor" id="cdk-eventbridge-partner-processors.TwilioEventProcessor"></a>

CDK wrapper for the Twilio Eventbridge processor.

> [https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github)

#### Initializers <a name="Initializers" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.Initializer"></a>

```typescript
import { TwilioEventProcessor } from 'cdk-eventbridge-partner-processors'

new TwilioEventProcessor(scope: Construct, id: string, props: TwilioProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioEventProcessor.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioEventProcessor.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioEventProcessor.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-eventbridge-partner-processors.TwilioProps">TwilioProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-eventbridge-partner-processors.TwilioProps">TwilioProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioEventProcessor.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioEventProcessor.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.isConstruct"></a>

```typescript
import { TwilioEventProcessor } from 'cdk-eventbridge-partner-processors'

TwilioEventProcessor.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioEventProcessor.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioEventProcessor.property.invocationAlarm">invocationAlarm</a></code> | <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarm">InvocationAlarm</a></code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioEventProcessor.property.partnerEventsFunction">partnerEventsFunction</a></code> | <code>aws-cdk-lib.aws_lambda.Function</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `invocationAlarm`<sup>Required</sup> <a name="invocationAlarm" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.property.invocationAlarm"></a>

```typescript
public readonly invocationAlarm: InvocationAlarm;
```

- *Type:* <a href="#cdk-eventbridge-partner-processors.InvocationAlarm">InvocationAlarm</a>

---

##### `partnerEventsFunction`<sup>Required</sup> <a name="partnerEventsFunction" id="cdk-eventbridge-partner-processors.TwilioEventProcessor.property.partnerEventsFunction"></a>

```typescript
public readonly partnerEventsFunction: Function;
```

- *Type:* aws-cdk-lib.aws_lambda.Function

---


## Structs <a name="Structs" id="Structs"></a>

### GitHubProps <a name="GitHubProps" id="cdk-eventbridge-partner-processors.GitHubProps"></a>

#### Initializer <a name="Initializer" id="cdk-eventbridge-partner-processors.GitHubProps.Initializer"></a>

```typescript
import { GitHubProps } from 'cdk-eventbridge-partner-processors'

const gitHubProps: GitHubProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubProps.property.eventBus">eventBus</a></code> | <code>aws-cdk-lib.aws_events.IEventBus</code> | Eventbus to send GitHub events to. |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubProps.property.lambdaInvocationAlarmThreshold">lambdaInvocationAlarmThreshold</a></code> | <code>number</code> | Maximum number of concurrent invocations on the fURL function before triggering the alarm. |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubProps.property.webhookSecret">webhookSecret</a></code> | <code>aws-cdk-lib.aws_secretsmanager.ISecret</code> | SM Secret containing the secret string used to validate webhook events. |

---

##### `eventBus`<sup>Required</sup> <a name="eventBus" id="cdk-eventbridge-partner-processors.GitHubProps.property.eventBus"></a>

```typescript
public readonly eventBus: IEventBus;
```

- *Type:* aws-cdk-lib.aws_events.IEventBus

Eventbus to send GitHub events to.

---

##### `lambdaInvocationAlarmThreshold`<sup>Required</sup> <a name="lambdaInvocationAlarmThreshold" id="cdk-eventbridge-partner-processors.GitHubProps.property.lambdaInvocationAlarmThreshold"></a>

```typescript
public readonly lambdaInvocationAlarmThreshold: number;
```

- *Type:* number

Maximum number of concurrent invocations on the fURL function before triggering the alarm.

---

##### `webhookSecret`<sup>Required</sup> <a name="webhookSecret" id="cdk-eventbridge-partner-processors.GitHubProps.property.webhookSecret"></a>

```typescript
public readonly webhookSecret: ISecret;
```

- *Type:* aws-cdk-lib.aws_secretsmanager.ISecret

SM Secret containing the secret string used to validate webhook events.

---

### InvocationAlarmProps <a name="InvocationAlarmProps" id="cdk-eventbridge-partner-processors.InvocationAlarmProps"></a>

#### Initializer <a name="Initializer" id="cdk-eventbridge-partner-processors.InvocationAlarmProps.Initializer"></a>

```typescript
import { InvocationAlarmProps } from 'cdk-eventbridge-partner-processors'

const invocationAlarmProps: InvocationAlarmProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarmProps.property.eventFunction">eventFunction</a></code> | <code>aws-cdk-lib.aws_lambda.IFunction</code> | The function to monitor. |
| <code><a href="#cdk-eventbridge-partner-processors.InvocationAlarmProps.property.threshold">threshold</a></code> | <code>number</code> | Lambda Invocation threshold. |

---

##### `eventFunction`<sup>Required</sup> <a name="eventFunction" id="cdk-eventbridge-partner-processors.InvocationAlarmProps.property.eventFunction"></a>

```typescript
public readonly eventFunction: IFunction;
```

- *Type:* aws-cdk-lib.aws_lambda.IFunction

The function to monitor.

---

##### `threshold`<sup>Required</sup> <a name="threshold" id="cdk-eventbridge-partner-processors.InvocationAlarmProps.property.threshold"></a>

```typescript
public readonly threshold: number;
```

- *Type:* number

Lambda Invocation threshold.

---

### PartnerFunctionProps <a name="PartnerFunctionProps" id="cdk-eventbridge-partner-processors.PartnerFunctionProps"></a>

#### Initializer <a name="Initializer" id="cdk-eventbridge-partner-processors.PartnerFunctionProps.Initializer"></a>

```typescript
import { PartnerFunctionProps } from 'cdk-eventbridge-partner-processors'

const partnerFunctionProps: PartnerFunctionProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerFunctionProps.property.eventbridgePartner">eventbridgePartner</a></code> | <code><a href="#cdk-eventbridge-partner-processors.Partner">Partner</a></code> | The partner to create an events processor for. |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerFunctionProps.property.eventBus">eventBus</a></code> | <code>aws-cdk-lib.aws_events.IEventBus</code> | Eventbus to send Partner events to. |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerFunctionProps.property.lambdaInvocationAlarmThreshold">lambdaInvocationAlarmThreshold</a></code> | <code>number</code> | Maximum number of concurrent invocations on the fURL function before triggering the alarm. |
| <code><a href="#cdk-eventbridge-partner-processors.PartnerFunctionProps.property.webhookSecret">webhookSecret</a></code> | <code>aws-cdk-lib.aws_secretsmanager.ISecret</code> | SM Secret containing the secret string used to validate webhook events. |

---

##### `eventbridgePartner`<sup>Required</sup> <a name="eventbridgePartner" id="cdk-eventbridge-partner-processors.PartnerFunctionProps.property.eventbridgePartner"></a>

```typescript
public readonly eventbridgePartner: Partner;
```

- *Type:* <a href="#cdk-eventbridge-partner-processors.Partner">Partner</a>

The partner to create an events processor for.

---

##### `eventBus`<sup>Required</sup> <a name="eventBus" id="cdk-eventbridge-partner-processors.PartnerFunctionProps.property.eventBus"></a>

```typescript
public readonly eventBus: IEventBus;
```

- *Type:* aws-cdk-lib.aws_events.IEventBus

Eventbus to send Partner events to.

---

##### `lambdaInvocationAlarmThreshold`<sup>Required</sup> <a name="lambdaInvocationAlarmThreshold" id="cdk-eventbridge-partner-processors.PartnerFunctionProps.property.lambdaInvocationAlarmThreshold"></a>

```typescript
public readonly lambdaInvocationAlarmThreshold: number;
```

- *Type:* number

Maximum number of concurrent invocations on the fURL function before triggering the alarm.

---

##### `webhookSecret`<sup>Required</sup> <a name="webhookSecret" id="cdk-eventbridge-partner-processors.PartnerFunctionProps.property.webhookSecret"></a>

```typescript
public readonly webhookSecret: ISecret;
```

- *Type:* aws-cdk-lib.aws_secretsmanager.ISecret

SM Secret containing the secret string used to validate webhook events.

---

### StripeProps <a name="StripeProps" id="cdk-eventbridge-partner-processors.StripeProps"></a>

#### Initializer <a name="Initializer" id="cdk-eventbridge-partner-processors.StripeProps.Initializer"></a>

```typescript
import { StripeProps } from 'cdk-eventbridge-partner-processors'

const stripeProps: StripeProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.StripeProps.property.eventBus">eventBus</a></code> | <code>aws-cdk-lib.aws_events.IEventBus</code> | Eventbus to send GitHub events to. |
| <code><a href="#cdk-eventbridge-partner-processors.StripeProps.property.lambdaInvocationAlarmThreshold">lambdaInvocationAlarmThreshold</a></code> | <code>number</code> | Maximum number of concurrent invocations on the fURL function before triggering the alarm. |
| <code><a href="#cdk-eventbridge-partner-processors.StripeProps.property.webhookSecret">webhookSecret</a></code> | <code>aws-cdk-lib.aws_secretsmanager.ISecret</code> | SM Secret containing the secret string used to validate webhook events. |

---

##### `eventBus`<sup>Required</sup> <a name="eventBus" id="cdk-eventbridge-partner-processors.StripeProps.property.eventBus"></a>

```typescript
public readonly eventBus: IEventBus;
```

- *Type:* aws-cdk-lib.aws_events.IEventBus

Eventbus to send GitHub events to.

---

##### `lambdaInvocationAlarmThreshold`<sup>Required</sup> <a name="lambdaInvocationAlarmThreshold" id="cdk-eventbridge-partner-processors.StripeProps.property.lambdaInvocationAlarmThreshold"></a>

```typescript
public readonly lambdaInvocationAlarmThreshold: number;
```

- *Type:* number

Maximum number of concurrent invocations on the fURL function before triggering the alarm.

---

##### `webhookSecret`<sup>Required</sup> <a name="webhookSecret" id="cdk-eventbridge-partner-processors.StripeProps.property.webhookSecret"></a>

```typescript
public readonly webhookSecret: ISecret;
```

- *Type:* aws-cdk-lib.aws_secretsmanager.ISecret

SM Secret containing the secret string used to validate webhook events.

---

### TwilioProps <a name="TwilioProps" id="cdk-eventbridge-partner-processors.TwilioProps"></a>

#### Initializer <a name="Initializer" id="cdk-eventbridge-partner-processors.TwilioProps.Initializer"></a>

```typescript
import { TwilioProps } from 'cdk-eventbridge-partner-processors'

const twilioProps: TwilioProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioProps.property.eventBus">eventBus</a></code> | <code>aws-cdk-lib.aws_events.IEventBus</code> | Eventbus to send GitHub events to. |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioProps.property.lambdaInvocationAlarmThreshold">lambdaInvocationAlarmThreshold</a></code> | <code>number</code> | Maximum number of concurrent invocations on the fURL function before triggering the alarm. |
| <code><a href="#cdk-eventbridge-partner-processors.TwilioProps.property.webhookSecret">webhookSecret</a></code> | <code>aws-cdk-lib.aws_secretsmanager.ISecret</code> | SM Secret containing the secret string used to validate webhook events. |

---

##### `eventBus`<sup>Required</sup> <a name="eventBus" id="cdk-eventbridge-partner-processors.TwilioProps.property.eventBus"></a>

```typescript
public readonly eventBus: IEventBus;
```

- *Type:* aws-cdk-lib.aws_events.IEventBus

Eventbus to send GitHub events to.

---

##### `lambdaInvocationAlarmThreshold`<sup>Required</sup> <a name="lambdaInvocationAlarmThreshold" id="cdk-eventbridge-partner-processors.TwilioProps.property.lambdaInvocationAlarmThreshold"></a>

```typescript
public readonly lambdaInvocationAlarmThreshold: number;
```

- *Type:* number

Maximum number of concurrent invocations on the fURL function before triggering the alarm.

---

##### `webhookSecret`<sup>Required</sup> <a name="webhookSecret" id="cdk-eventbridge-partner-processors.TwilioProps.property.webhookSecret"></a>

```typescript
public readonly webhookSecret: ISecret;
```

- *Type:* aws-cdk-lib.aws_secretsmanager.ISecret

SM Secret containing the secret string used to validate webhook events.

---



## Enums <a name="Enums" id="Enums"></a>

### Partner <a name="Partner" id="cdk-eventbridge-partner-processors.Partner"></a>

Supported partners with fURL integrations.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-eventbridge-partner-processors.Partner.GITHUB">GITHUB</a></code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.Partner.STRIPE">STRIPE</a></code> | *No description.* |
| <code><a href="#cdk-eventbridge-partner-processors.Partner.TWILIO">TWILIO</a></code> | *No description.* |

---

##### `GITHUB` <a name="GITHUB" id="cdk-eventbridge-partner-processors.Partner.GITHUB"></a>

---


##### `STRIPE` <a name="STRIPE" id="cdk-eventbridge-partner-processors.Partner.STRIPE"></a>

---


##### `TWILIO` <a name="TWILIO" id="cdk-eventbridge-partner-processors.Partner.TWILIO"></a>

---

