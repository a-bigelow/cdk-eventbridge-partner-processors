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

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-eventbridge-partner-processors.GitHubEventProcessor.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

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
| <code><a href="#cdk-eventbridge-partner-processors.GitHubProps.property.gitHubWebhookSecret">gitHubWebhookSecret</a></code> | <code>aws-cdk-lib.aws_secretsmanager.ISecret</code> | SM Secret containing the secret string used to validate webhook events. |
| <code><a href="#cdk-eventbridge-partner-processors.GitHubProps.property.lambdaInvocationAlarmThreshold">lambdaInvocationAlarmThreshold</a></code> | <code>number</code> | Maximum number of concurrent invocations on the fURL function before triggering the alarm. |

---

##### `eventBus`<sup>Required</sup> <a name="eventBus" id="cdk-eventbridge-partner-processors.GitHubProps.property.eventBus"></a>

```typescript
public readonly eventBus: IEventBus;
```

- *Type:* aws-cdk-lib.aws_events.IEventBus

Eventbus to send GitHub events to.

---

##### `gitHubWebhookSecret`<sup>Required</sup> <a name="gitHubWebhookSecret" id="cdk-eventbridge-partner-processors.GitHubProps.property.gitHubWebhookSecret"></a>

```typescript
public readonly gitHubWebhookSecret: ISecret;
```

- *Type:* aws-cdk-lib.aws_secretsmanager.ISecret

SM Secret containing the secret string used to validate webhook events.

---

##### `lambdaInvocationAlarmThreshold`<sup>Required</sup> <a name="lambdaInvocationAlarmThreshold" id="cdk-eventbridge-partner-processors.GitHubProps.property.lambdaInvocationAlarmThreshold"></a>

```typescript
public readonly lambdaInvocationAlarmThreshold: number;
```

- *Type:* number

Maximum number of concurrent invocations on the fURL function before triggering the alarm.

---



