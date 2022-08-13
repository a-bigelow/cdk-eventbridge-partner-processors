import { Duration, Stack } from 'aws-cdk-lib';
import { Alarm, ComparisonOperator, Metric } from 'aws-cdk-lib/aws-cloudwatch';
import { IFunction } from 'aws-cdk-lib/aws-lambda';
import { Construct } from 'constructs';

export interface InvocationAlarmProps {
  /**
     * Lambda Invocation threshold.
     */
  readonly threshold: number;

  /**
     * The function to monitor.
     */
  readonly eventFunction: IFunction;
}

/**
 * Cloudwatch Alarm used across this construct library.
 */
export class InvocationAlarm extends Construct {
  constructor(scope: Construct, id: string, props: InvocationAlarmProps) {
    super(scope, id);

    new Alarm(this, 'InvocationAlarm', {
      alarmDescription: `Alarm for ${Stack.of(this).stackName} - InboundWebhook Lambda for traffic spikes`,
      alarmName: `InboundWebhook-Lambda-Invocation-Alarm-${Stack.of(this).stackName}`,
      evaluationPeriods: 2,
      metric: new Metric({
        metricName: 'Invocations',
        namespace: 'AWS/Lambda',
        period: Duration.seconds(300),
        statistic: 'Sum',
        dimensionsMap: {
          Name: 'FunctionName',
          Value: props.eventFunction.functionName,
        },
      }),
      threshold: props.threshold,
      comparisonOperator: ComparisonOperator.GREATER_THAN_THRESHOLD,
    });
  }
}