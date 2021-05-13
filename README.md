# Bookstore REST API using Gin and Gorm - Running in your Kubernetes Cluster

Read the [article](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/).

Read the [article](https://www.digitalocean.com/community/tutorials/how-to-deploy-resilient-go-app-digitalocean-kubernetes).

## STEPS to deploy in Kubernetes cluster

1. Login to your cloud provider account and download kubeconfig of your cluster.

2. Create a namespace  - `kubectl create namespace app`

3. Create a deployment - `kubectl apply -f templates/deployment.yaml`

4. Create a service of type LoadBalancer - `kubectl apply -f templates/service.yaml`

5. Get the LB service's external IP address using the command - `kubectl get services -n app`

6. Interact with the bookstore application running in your Kubernetes cluster using REST APIs. Examples given below.

## REST APIs

```
Create a Book : $ curl -vvv -X POST http://{{External-IP-Of-Service}}/books -H 'Content-Type: application/json' -d '{"title": "Start with Why","author": "Simon Sinek"}'

Get book details for book ID : $ curl -v http://{{External-IP-Of-Service}}/books/2

List all books : $ curl -vvv http://{{External-IP-Of-Service}}/books

Update a book's details by book ID : $ curl -vvv -X PATCH http://{{External-IP-Of-Service}}/books/2 -d '{"title":"Infinite Game"}'

Delete a book by book ID : $ curl -vvv -X DELETE http://{{External-IP-Of-Service}}/books/4
```


