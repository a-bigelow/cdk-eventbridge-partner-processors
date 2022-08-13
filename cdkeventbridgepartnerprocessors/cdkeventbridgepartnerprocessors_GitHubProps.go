// cdk-eventbridge-partner-processors
package cdkeventbridgepartnerprocessors

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

type GitHubProps struct {
	// Eventbus to send GitHub events to.
	EventBus awsevents.IEventBus `field:"required" json:"eventBus" yaml:"eventBus"`
	// SM Secret containing the secret string used to validate webhook events.
	GitHubWebhookSecret awssecretsmanager.ISecret `field:"required" json:"gitHubWebhookSecret" yaml:"gitHubWebhookSecret"`
	// Maximum number of concurrent invocations on the fURL function before triggering the alarm.
	LambdaInvocationAlarmThreshold *float64 `field:"required" json:"lambdaInvocationAlarmThreshold" yaml:"lambdaInvocationAlarmThreshold"`
}

