import { awscdk } from 'projen';
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'a-bigelow',
  authorAddress: 'adam@adambigelow.com',
  cdkVersion: '2.37.1',
  defaultReleaseBranch: 'main',
  depsUpgrade: true,
  autoApproveUpgrades: true,
  autoApproveOptions: {
    allowedUsernames: ['a-bigelow'],
  },
  name: 'cdk-eventbridge-partner-processors',
  keywords: ['cdk', 'eventbridge', 'github', 'gitlab', 'stripe', 'twilio'],
  projenrcTs: true,
  release: true,
  repositoryUrl: 'https://github.com/a-bigelow/cdk-eventbridge-partner-processors.git',
  gitignore: ['.idea/', 'package-lock.json'],
  publishToPypi: {
    distName: 'a-bigelow.cdk-eventbridge-partner-processors',
    module: 'a-bigelow.cdk-eventbridge-partner-processors',
  },
  publishToGo: {
    moduleName: 'github.com/a-bigelow/cdk-eventbridge-partner-processors-go',
  },
  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});

// Copy GitLab Lambda function code to lib directory after compilation
project.postCompileTask.exec('mkdir -p lib/gitlab-lambda && cp -r src/gitlab-lambda/* lib/gitlab-lambda/');

project.synth();