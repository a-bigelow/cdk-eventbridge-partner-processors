// cdk-eventbridge-partner-processors
package cdkeventbridgepartnerprocessors

import (
	_init_ "github.com/a-bigelow/cdk-eventbridge-partner-processors/cdkeventbridgepartnerprocessors/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/a-bigelow/cdk-eventbridge-partner-processors/cdkeventbridgepartnerprocessors/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Cloudwatch Alarm used across this construct library.
type InvocationAlarm interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for InvocationAlarm
type jsiiProxy_InvocationAlarm struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_InvocationAlarm) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewInvocationAlarm(scope constructs.Construct, id *string, props *InvocationAlarmProps) InvocationAlarm {
	_init_.Initialize()

	j := jsiiProxy_InvocationAlarm{}

	_jsii_.Create(
		"cdk-eventbridge-partner-processors.InvocationAlarm",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewInvocationAlarm_Override(i InvocationAlarm, scope constructs.Construct, id *string, props *InvocationAlarmProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-eventbridge-partner-processors.InvocationAlarm",
		[]interface{}{scope, id, props},
		i,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func InvocationAlarm_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-eventbridge-partner-processors.InvocationAlarm",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvocationAlarm) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

