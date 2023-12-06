by [andrewsaputra](https://github.com/andrewsaputra)

# Message Queue System

## Summary

Design and implement a simple  message queueing system using Go programming language and RabbitMQ or Kafka.

The system includes the following parts :

#### API
Design an API where it should receive a product data and store it in the database, below are the parameters that should be passed in the API

- user_id
- product_name
- product_description (text)
- product_images (array of image urls)
- product_price (number)

#### Producer
After storing the product details in the database, `product_id` should be passed on the message queue.


#### Consumer
Based on the `product_id`, images should be download and compressed and stored in local. After storing, a local location path should be added as an array value in the `products` table in the `compressed_images` column.

## Database Schema

#### Users

| Column | Type | Notes |
| --- | --- | --- |
| id | INT | PRIMARY KEY |
| name | VARCHAR(255) | |
| email | VARCHAR(255) | |
| created_at | BIGINT | |
| updated_at | BIGINT | |


#### Products
(data should be added from API only)

| Column | Type | Notes |
| --- | --- | --- |
| id | INT | PRIMARY KEY |
| user_id | INT | FOREIGN KEY |
| name | VARCHAR(255) | |
| description | TEXT | |
| images | TEXT[] | |
| compressed_images | TEXT[] | |
| price | NUMERIC(10,2) | |
| created_at | BIGINT | |
| updated_at | BIGINT | |


## Testing [WIP]
Include unit tests, integration and benchmark tests.

---
---

## How to Use

### Setting Up Applications

1. Install and Run Rabbit MQ
```
docker pull rabbitmq
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management
```
Alternatively, follow the official [installation guide](https://www.rabbitmq.com/download.html).

2. Deploy Database to PostgreSQL
```
cd $REPO_ROOT/database
psql -U postgres -f db.sql
```

Sample Result :
```
demo_message_queue=# \dt
          List of relations
 Schema |   Name   | Type  |  Owner   
--------+----------+-------+----------
 public | products | table | postgres
 public | users    | table | postgres
(2 rows)

demo_message_queue=# select * from users;
 id |    name    |        email         |  created_at   | updated_at 
----+------------+----------------------+---------------+------------
  1 | John Smith | john.smith@email.com | 1701848883054 |          0
  2 | Laura Jane | laura.jane@email.com | 1701848883054 |          0
(2 rows)

demo_message_queue=# select * from products;
 id | user_id | name | description | images | compressed_images | price | created_at | updated_at 
----+---------+------+-------------+--------+-------------------+-------+------------+------------
(0 rows)
```

3. Copy product image files to `images` folder

4. Run API Service
```
cd $REPO_ROOT/apiserver
go run .
```

| Path | Method | Payload | Description |
| --- | --- | --- | --- |
| `/status` | GET | - | Application health check |
| `/users` | POST | [AddUserDTO](https://github.com/andrewsaputra/go-message-queue-exercise/blob/main/apiserver/api/data-types.go) | Create new user record |
| `/users/{id}` | GET | - | Get user record |
| `/users/{id}` | DELETE | - | Delete user record |
| `/products` | POST | [AddProductDTO](https://github.com/andrewsaputra/go-message-queue-exercise/blob/main/apiserver/api/data-types.go) | Create new product record and trigger publishing to RabbitMQ |
| `/products/{id}` | GET | - | Get product record |
| `/products/{id}` | DELETE | - | Delete product record |


5. Run Consumer Service
```
cd $REPO_ROOT/consumer
go run .
```

### Running the Scenario
Call `API Service` to add new product record.
Example :
```
curl -w "\n" -X POST localhost:3000/products -d '{"user_id":2, "product_name":"product 1", "product_description":"description qwerty", "product_price" : 9.99, "product_images":["pizza.jpg", "spaghetti.jpg", "ramen.jpg"]}'
```

Sample Result :
```
{"Data":{"Id":2,"UserId":2,"Name":"product 1","Description":"desc 1","Images":["pizza.jpg","spaghetti.jpg","ramen.jpg"],"Price":9.99,"compressed_images":[],"CreatedAt":1701847644879,"UpdatedAt":0},"Message":""}
```

Upon successful consumption you should see similar result on `Consumer Service`'s logs :
```
ProductConsumerService.OnConsumed : 1
ProductConsumerService.OnConsumed : 2
...
```

New compressed images will be saved to `images/compressed` folder.