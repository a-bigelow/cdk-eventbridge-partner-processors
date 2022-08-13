import { CfnOutput, Duration, Stack } from 'aws-cdk-lib';
import { IEventBus } from 'aws-cdk-lib/aws-events';
import { Code, Function, FunctionUrlAuthType, Runtime } from 'aws-cdk-lib/aws-lambda';
import { Bucket } from 'aws-cdk-lib/aws-s3';
import { ISecret } from 'aws-cdk-lib/aws-secretsmanager';
import { Construct } from 'constructs';
import { Partner } from './Partner';
import { InvocationAlarm } from './Util';

export interface PartnerFunctionProps {

  /**
     * SM Secret containing the secret string used to validate webhook events.
     */
  readonly webhookSecret: ISecret;

  /**
     * Eventbus to send Partner events to.
     */
  readonly eventBus: IEventBus;

  /**
     * Maximum number of concurrent invocations on the fURL function before triggering the alarm.
     */
  readonly lambdaInvocationAlarmThreshold: number;

  /**
     * The partner to create an events processor for.
     */
  readonly eventbridgePartner: Partner;

}

/**
 * CDK wrapper for the GitHub Eventbridge processor.
 * @see https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github
 */
export abstract class PartnerProcessor extends Construct {
  public partnerEventsFunction: Function;
  public invocationAlarm: InvocationAlarm;
  constructor(scope: Construct, id: string, props: PartnerFunctionProps) {
    super(scope, id);

    this.partnerEventsFunction = new Function(this, `${props.eventbridgePartner}EventsFunction`, {
      code: Code.fromBucket(Bucket.fromBucketName(this, 'AWSFunctionBucket', `eventbridge-inbound-webhook-templates-prod-${Stack.of(this).region}`), `lambda-templates/${props.eventbridgePartner}-lambdasrc.zip`),
      handler: 'app.handler',
      runtime: Runtime.PYTHON_3_7,
      memorySize: 128,
      timeout: Duration.seconds(100),
      reservedConcurrentExecutions: 10,
      environment: {
        GITHUB_WEBHOOK_SECRET_ARN: props.webhookSecret.secretArn,
        EVENT_BUS_NAME: props.eventBus.eventBusName,
      },
    });

    this.invocationAlarm = new InvocationAlarm(this, `${props.eventbridgePartner}InvocationAlarm`, {
      threshold: props.lambdaInvocationAlarmThreshold,
      eventFunction: this.partnerEventsFunction,
    });

    const fURL = this.partnerEventsFunction.addFunctionUrl({ authType: FunctionUrlAuthType.NONE });

    props.webhookSecret.grantRead(this.partnerEventsFunction);
    props.eventBus.grantPutEventsTo(this.partnerEventsFunction);

    new CfnOutput(this, `${props.eventbridgePartner}FunctionUrl`, { value: fURL.url });
  }
}