// cdk-eventbridge-partner-processors
package cdkeventbridgepartnerprocessors

import (
	_init_ "github.com/a-bigelow/cdk-eventbridge-partner-processors/cdkeventbridgepartnerprocessors/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// CDK wrapper for the GitHub Eventbridge processor.
// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github
//
type StripeEventProcessor interface {
	PartnerProcessor
	InvocationAlarm() InvocationAlarm
	SetInvocationAlarm(val InvocationAlarm)
	// The tree node.
	Node() constructs.Node
	PartnerEventsFunction() awslambda.Function
	SetPartnerEventsFunction(val awslambda.Function)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for StripeEventProcessor
type jsiiProxy_StripeEventProcessor struct {
	jsiiProxy_PartnerProcessor
}

func (j *jsiiProxy_StripeEventProcessor) InvocationAlarm() InvocationAlarm {
	var returns InvocationAlarm
	_jsii_.Get(
		j,
		"invocationAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StripeEventProcessor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StripeEventProcessor) PartnerEventsFunction() awslambda.Function {
	var returns awslambda.Function
	_jsii_.Get(
		j,
		"partnerEventsFunction",
		&returns,
	)
	return returns
}


func NewStripeEventProcessor(scope constructs.Construct, id *string, props *StripeProps) StripeEventProcessor {
	_init_.Initialize()

	j := jsiiProxy_StripeEventProcessor{}

	_jsii_.Create(
		"cdk-eventbridge-partner-processors.StripeEventProcessor",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewStripeEventProcessor_Override(s StripeEventProcessor, scope constructs.Construct, id *string, props *StripeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-eventbridge-partner-processors.StripeEventProcessor",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_StripeEventProcessor) SetInvocationAlarm(val InvocationAlarm) {
	_jsii_.Set(
		j,
		"invocationAlarm",
		val,
	)
}

func (j *jsiiProxy_StripeEventProcessor) SetPartnerEventsFunction(val awslambda.Function) {
	_jsii_.Set(
		j,
		"partnerEventsFunction",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func StripeEventProcessor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-eventbridge-partner-processors.StripeEventProcessor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StripeEventProcessor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

