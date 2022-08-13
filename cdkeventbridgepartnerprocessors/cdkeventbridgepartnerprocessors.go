package cdkeventbridgepartnerprocessors

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-eventbridge-partner-processors.GitHubEventProcessor",
		reflect.TypeOf((*GitHubEventProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "githubEventsFunction", GoGetter: "GithubEventsFunction"},
			_jsii_.MemberProperty{JsiiProperty: "invocationAlarm", GoGetter: "InvocationAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GitHubEventProcessor{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-eventbridge-partner-processors.GitHubProps",
		reflect.TypeOf((*GitHubProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-eventbridge-partner-processors.InvocationAlarm",
		reflect.TypeOf((*InvocationAlarm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_InvocationAlarm{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-eventbridge-partner-processors.InvocationAlarmProps",
		reflect.TypeOf((*InvocationAlarmProps)(nil)).Elem(),
	)
}
