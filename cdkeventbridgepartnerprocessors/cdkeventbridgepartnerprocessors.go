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
			_jsii_.MemberProperty{JsiiProperty: "invocationAlarm", GoGetter: "InvocationAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partnerEventsFunction", GoGetter: "PartnerEventsFunction"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GitHubEventProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PartnerProcessor)
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
	_jsii_.RegisterEnum(
		"cdk-eventbridge-partner-processors.Partner",
		reflect.TypeOf((*Partner)(nil)).Elem(),
		map[string]interface{}{
			"GITHUB": Partner_GITHUB,
			"STRIPE": Partner_STRIPE,
			"TWILIO": Partner_TWILIO,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-eventbridge-partner-processors.PartnerFunctionProps",
		reflect.TypeOf((*PartnerFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-eventbridge-partner-processors.PartnerProcessor",
		reflect.TypeOf((*PartnerProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "invocationAlarm", GoGetter: "InvocationAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partnerEventsFunction", GoGetter: "PartnerEventsFunction"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PartnerProcessor{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-eventbridge-partner-processors.StripeEventProcessor",
		reflect.TypeOf((*StripeEventProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "invocationAlarm", GoGetter: "InvocationAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partnerEventsFunction", GoGetter: "PartnerEventsFunction"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_StripeEventProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PartnerProcessor)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-eventbridge-partner-processors.StripeProps",
		reflect.TypeOf((*StripeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-eventbridge-partner-processors.TwilioEventProcessor",
		reflect.TypeOf((*TwilioEventProcessor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "invocationAlarm", GoGetter: "InvocationAlarm"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "partnerEventsFunction", GoGetter: "PartnerEventsFunction"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TwilioEventProcessor{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_PartnerProcessor)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-eventbridge-partner-processors.TwilioProps",
		reflect.TypeOf((*TwilioProps)(nil)).Elem(),
	)
}
