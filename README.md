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

- sign up via AWS CLI
```
aws cognito-idp sign-up --username EMAIL --password PASSWORD --client-id ID \
    --endpoint-url=http://localhost:9229 --profile localstack
```
- sign up via CLI
```
go run main.go sign-up --config=config.yaml -u=EMAIL -p=PASSWORD --clientID=ID
```
---

- confirm sign up via AWS CLI
```
aws cognito-idp confirm-sign-up --client-id ID \
    --username EMAIL --confirmation-code CODE \
    --endpoint-url=http://localhost:9229 --profile localstack
```
- confirm sign up via CLI
```
go run main.go confirm-sign-up --config=config.yaml -u=EMAIL --code=CODE --clientID=ID
```
---

- sign in via AWS CLI
```
aws cognito-idp initiate-auth --client-id ID \
    --auth-parameters USERNAME=EMAIL,PASSWORD=PASSWORD --auth-flow USER_PASSWORD_AUTH \
    --endpoint-url=http://localhost:9229 --profile localstack
```
- sign in via CLI
```
go run main.go sign-in --config=config.yaml -u=EMAIL -p=PASSWORD --clientID=ID
```
---

- delete user via AWS CLI
```
aws cognito-idp delete-user --access-token TOKEN --endpoint-url=http://localhost:9229 --profile localstack
```
- delete user via CLI
```
token=$(cat access-token.txt 2>/dev/null)
go run main.go delete-user --config=config.yaml --token=$token
rm access-token.txt
```
---

- delete user pool client via AWS CLI
```
aws cognito-idp delete-user-pool-client --user-pool-id ID --client-id ID \
    --endpoint-url=http://localhost:9229 --profile localstack
```
- delete user pool client via CLI
```
go run main.go delete-pool-client --config=config.yaml --poolID=ID --clientID=ID
```
---

- delete user pool via AWS CLI
```
aws cognito-idp delete-user-pool --user-pool-id ID --endpoint-url=http://localhost:9229 --profile localstack
```
- delete user pool via CLI
```
go run main.go delete-pool --config=config.yaml --poolID=ID
```

---
※ References:
- [cognito-idp](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/cognito-idp/index.html)
