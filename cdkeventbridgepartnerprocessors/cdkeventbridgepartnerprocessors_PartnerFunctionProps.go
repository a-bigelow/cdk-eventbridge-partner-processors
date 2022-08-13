// cdk-eventbridge-partner-processors
package cdkeventbridgepartnerprocessors

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

type PartnerFunctionProps struct {
	// The partner to create an events processor for.
	EventbridgePartner Partner `field:"required" json:"eventbridgePartner" yaml:"eventbridgePartner"`
	// Eventbus to send Partner events to.
	EventBus awsevents.IEventBus `field:"required" json:"eventBus" yaml:"eventBus"`
	// Maximum number of concurrent invocations on the fURL function before triggering the alarm.
	LambdaInvocationAlarmThreshold *float64 `field:"required" json:"lambdaInvocationAlarmThreshold" yaml:"lambdaInvocationAlarmThreshold"`
	// SM Secret containing the secret string used to validate webhook events.
	WebhookSecret awssecretsmanager.ISecret `field:"required" json:"webhookSecret" yaml:"webhookSecret"`
}

