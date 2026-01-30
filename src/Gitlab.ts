import * as path from 'path';
import { CfnOutput, Duration } from 'aws-cdk-lib';
import { IEventBus } from 'aws-cdk-lib/aws-events';
import { Code, Function, FunctionUrlAuthType, Runtime } from 'aws-cdk-lib/aws-lambda';
import { ISecret } from 'aws-cdk-lib/aws-secretsmanager';
import { Construct } from 'constructs';
import { InvocationAlarm } from './Util';

export interface GitLabProps {

  /**
   * SM Secret containing the secret string used to validate webhook events.
   */
  readonly webhookSecret: ISecret;

  /**
   * Eventbus to send GitLab events to.
   */
  readonly eventBus: IEventBus;

  /**
   * Maximum number of concurrent invocations on the fURL function before triggering the alarm.
   */
  readonly lambdaInvocationAlarmThreshold: number;

}

/**
 * CDK wrapper for the GitLab Eventbridge processor.
 * Unlike other partner processors, GitLab uses a custom Lambda implementation
 * as AWS does not provide a pre-built Lambda for GitLab webhooks.
 * @see https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html
 */
export class GitLabEventProcessor extends Construct {
  public partnerEventsFunction: Function;
  public invocationAlarm: InvocationAlarm;

  constructor(scope: Construct, id: string, props: GitLabProps) {
    super(scope, id);

    this.partnerEventsFunction = new Function(this, 'gitlabEventsFunction', {
      code: Code.fromAsset(path.join(__dirname, 'gitlab-lambda')),
      handler: 'app.lambda_handler',
      runtime: Runtime.PYTHON_3_9,
      memorySize: 128,
      timeout: Duration.seconds(100),
      reservedConcurrentExecutions: 10,
      environment: {
        GITLAB_WEBHOOK_SECRET_ARN: props.webhookSecret.secretArn,
        EVENT_BUS_NAME: props.eventBus.eventBusName,
      },
    });

    this.invocationAlarm = new InvocationAlarm(this, 'gitlabInvocationAlarm', {
      threshold: props.lambdaInvocationAlarmThreshold,
      eventFunction: this.partnerEventsFunction,
    });

    const fURL = this.partnerEventsFunction.addFunctionUrl({ authType: FunctionUrlAuthType.NONE });

    props.webhookSecret.grantRead(this.partnerEventsFunction);
    props.eventBus.grantPutEventsTo(this.partnerEventsFunction);

    new CfnOutput(this, 'gitlabFunctionUrl', { value: fURL.url });
  }
}
