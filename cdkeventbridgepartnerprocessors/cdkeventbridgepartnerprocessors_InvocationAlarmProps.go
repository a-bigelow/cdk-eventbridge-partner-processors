// cdk-eventbridge-partner-processors
package cdkeventbridgepartnerprocessors

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type InvocationAlarmProps struct {
	// The function to monitor.
	EventFunction awslambda.IFunction `field:"required" json:"eventFunction" yaml:"eventFunction"`
	// Lambda Invocation threshold.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

