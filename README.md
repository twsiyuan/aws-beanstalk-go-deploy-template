# AWS Beanstalk Go Deploy Template

A simple golang project template for deploy to AWS beanstalk (using source code)

## How to deploy

- Clone ```simple-template```
- Put webapp go source code into ```src/app``` folder
- Make sure webapp that listen to port ```os.Getenv("PORT")```
- Collect all dependency package source to into ```src/vendor``` folder using management tool, such as [govendor](https://github.com/kardianos/govendor), [dep](https://github.com/golang/dep), [godep](https://github.com/tools/godep), [glide](https://github.com/bumptech/glide), ...etc
- Modify config files if necessary
- Compress all files under ```simple-template``` into a zip file (as a source bundle)
- Upload source bundle to AWS beanstalk and deploy it
- Done

For more information, see AWS document: [Deploying Go Applications to Elastic Beanstalk Applications](http://docs.aws.amazon.com/elasticbeanstalk/latest/dg/create_deploy_go.html)

## File structure
- .ebextensions (environment configs)
  - applogs.config (rotate app logs to S3 config)
- src
  - app
    - main.go (webapp source code, that contains the ```main``` package )
- build.sh (build commands)
- Buildfile (custom build config)
- Procfile (process config)