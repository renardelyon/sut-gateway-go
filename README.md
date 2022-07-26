# sut-gateway-go

Microservice skeleton

## Infrastructure
![infra](https://user-images.githubusercontent.com/79710053/179886093-449a6144-b195-4e3d-bff0-17cc1b3ba111.png)

## How to run in local?
1. Run all required microservices
  <ul>
  <li><a href="https://github.com/renardelyon/sut-notification-go" target="_blank">sut-notification-go</a></li>
  <li><a href="https://github.com/renardelyon/sut-order-go" target="_blank">sut-order-go</a></li>
  <li><a href="https://github.com/renardelyon/sut-product-go" target="_blank">sut-product-go</a></li>
  <li><a href="https://github.com/renardelyon/sut-auth-go" target="_blank">sut-auth-go</a></li>
  <li><a href="https://github.com/renardelyon/sut-storage-go" target="_blank">sut-storage-go</a></li>
  </ul>

2. Create file env and name it `dev.env`. its content can be seen in code block below. 
```
PORT=:50056
AUTH_SVC_URL=localhost:50051
ORDER_SVC_URL=localhost:50052
STORAGE_SVC_URL=localhost:50053
PRODUCT_SVC_URL=localhost:50054
NOTIF_SVC_URL=localhost:50055

```

3. Execute command below
```
make init # initialize go.mod
make tidy # Tidy up go module
```

4. Adding go bin into path env variables
```
export PATH=$PATH:$(go env GOPATH)/bin
```

5. Adding folder with `pb` as name into ther project root directory

6. Generate protobuf by executing command below
```
make proto-gen
```

7. Run the application
```
make run
```

## API Documentation
https://app.swaggerhub.com/apis-docs/renardelyon/sut-gateway_go/0.1.0
