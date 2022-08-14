import { awscdk } from 'projen';
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'a-bigelow',
  authorAddress: 'adam@adambigelow.com',
  cdkVersion: '2.37.1',
  defaultReleaseBranch: 'main',
  name: 'cdk-eventbridge-partner-processors',
  keywords: ['cdk', 'eventbridge', 'github', 'stripe', 'twilio'],
  projenrcTs: true,
  release: true,
  repositoryUrl: 'https://github.com/a-bigelow/cdk-eventbridge-partner-processors.git',
  gitignore: ['.idea/'],
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
project.synth();