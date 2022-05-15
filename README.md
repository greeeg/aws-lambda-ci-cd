# AWS Lambda CI/CD

> This repository is the source code for an article on how to set up a CI/CD pipeline for a serverless application using AWS, Terraform, GitHub Actions & Golang.

## Getting started

In order to use this project, make sure to set up your AWS account, GitHub repository & local machine accordingly.

### Setting up your AWS account

- Create an S3 bucket that will be used to store Terraform’s state
- Create a IAM user with programmatic access & `AdministatorAccess` policy

### Setting up your local machine

- Install Terraform with a version compatible with the one used in this project

### Setting up your GitHub repository

- Define two repository secrets for `AWS_ACCESS_KEY_ID` & `AWS_SECRET_ACCESS_KEY` to be used in the CD GitHub Actions workflow