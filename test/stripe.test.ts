import { App, SecretValue, Stack } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { EventBus } from 'aws-cdk-lib/aws-events';
import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
import { StripeEventProcessor } from '../src';

test('That Stripe synths', () => {
  const app = new App();
  const stack = new Stack(app, 'TestStripeStack', {});

  new StripeEventProcessor(stack, 'TestStripeProcessor', {
    eventBus: new EventBus(stack, 'TestEventBus', { eventBusName: 'TheGreatestBus' }),
    webhookSecret: new Secret(stack, 'VeryVerySecret', { secretStringValue: SecretValue.unsafePlainText('DontEverDoThis') }),
    lambdaInvocationAlarmThreshold: 2000,
  });

  expect(Template.fromStack(stack));
});