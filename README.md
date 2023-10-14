# go-aws-cognito

### (local) usage of AWS Cognito via Cobra and/or AWS CLI.

- create user pool via AWS CLI
```
aws cognito-idp create-user-pool --pool-name "My pool" --endpoint-url=http://localhost:9229 --profile localstack
```
- create user pool via CLI
```
go run main.go --config=config.yaml create-pool --name="My pool"
```
---

- describe user pool via AWS CLI
```
aws cognito-idp describe-user-pool --user-pool-id ID --endpoint-url=http://localhost:9229 --profile localstack
```
- describe user pool via CLI
```
go run main.go --config=config.yaml describe-pool --poolID=ID
```
---

- create user pool client via AWS CLI
```
aws cognito-idp create-user-pool-client --user-pool-id ID --client-name "My client" \
    --endpoint-url=http://localhost:9229 --profile localstack
```
- create user pool client via CLI
```
go run main.go --config=config.yaml create-pool-client --poolID=ID --name="My client"
```
---

---
※ References:
- [cognito-idp](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/cognito-idp/index.html)
