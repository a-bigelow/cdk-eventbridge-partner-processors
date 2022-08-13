import { awscdk } from 'projen';
const project = new awscdk.AwsCdkConstructLibrary({
  author: 'a-bigelow',
  authorAddress: 'adam@adambigelow.com',
  cdkVersion: '2.1.0',
  defaultReleaseBranch: 'main',
  name: 'cdk-eventbridge-partner-processors',
  projenrcTs: true,
  release: false,
  repositoryUrl: 'https://github.com/a-bigelow/cdk-eventbridge-partner-processors.git',

  // deps: [],                /* Runtime dependencies of this module. */
  // description: undefined,  /* The description is just a string that helps people understand the purpose of the package. */
  // devDeps: [],             /* Build dependencies for this module. */
  // packageName: undefined,  /* The "name" in package.json. */
});
project.synth();