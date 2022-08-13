import { IEventBus } from 'aws-cdk-lib/aws-events';
import { ISecret } from 'aws-cdk-lib/aws-secretsmanager';
import { Construct } from 'constructs';
import { Partner } from './Partner';
import { PartnerProcessor } from './PartnerFunction';

export interface StripeProps {

  /**
   * SM Secret containing the secret string used to validate webhook events.
   */
  readonly webhookSecret: ISecret;

  /**
   * Eventbus to send GitHub events to.
   */
  readonly eventBus: IEventBus;

  /**
   * Maximum number of concurrent invocations on the fURL function before triggering the alarm.
   */
  readonly lambdaInvocationAlarmThreshold: number;

}

/**
 * CDK wrapper for the GitHub Eventbridge processor.
 * @see https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-saas-furls.html#furls-connection-github
 */
export class StripeEventProcessor extends PartnerProcessor {
  constructor(scope: Construct, id: string, props: StripeProps) {
    super(scope, id, { ...props, eventbridgePartner: Partner.STRIPE });
  }
}