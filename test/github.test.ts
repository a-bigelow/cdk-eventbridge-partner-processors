import { App, SecretValue, Stack } from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { EventBus } from 'aws-cdk-lib/aws-events';
import { Secret } from 'aws-cdk-lib/aws-secretsmanager';
import { GitHubEventProcessor } from '../src';

test('That GitHub synths', () => {
  const app = new App();
  const stack = new Stack(app, 'TestGitHubStack', {});

  new GitHubEventProcessor(stack, 'TestGitHubProcessor', {
    eventBus: new EventBus(stack, 'TestEventBus', { eventBusName: 'TheGreatestBus' }),
    gitHubWebhookSecret: new Secret(stack, 'VeryVerySecret', { secretStringValue: SecretValue.unsafePlainText('DontEverDoThis') }),
    lambdaInvocationAlarmThreshold: 2000,
  });

  expect(Template.fromStack(stack));
});