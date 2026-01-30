import { IEventBus } from 'aws-cdk-lib/aws-events';
import { ISecret } from 'aws-cdk-lib/aws-secretsmanager';
import { Construct } from 'constructs';
import { Partner } from './Partner';
import { PartnerProcessor } from './PartnerFunction';

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
 * @see https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html
 */
export class GitLabEventProcessor extends PartnerProcessor {
  constructor(scope: Construct, id: string, props: GitLabProps) {
    super(scope, id, { ...props, eventbridgePartner: Partner.GITLAB });
  }
}
