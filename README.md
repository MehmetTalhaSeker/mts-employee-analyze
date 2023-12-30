# Employee Analyze App
Analyze your employees feedbacks using kafka, elastic search, golang, java and kafka-streams. 

## Requirements
- Docker
- Maven

## Installation

1. Install dependencies of kafka-streams java app:

```
cd kafka-streams && mvn clean install                                                                                                                            
```


2. Build and run docker containers:

```
docker-compose up --build -d
```

## Examples

```
curl -X POST http://localhost:8080 \
-H "Content-Type: application/json" \
-d '{
  "employee_id": "1",
  "keywords": "vending-machine"
}'
```

```
curl -X GET http://localhost:9200/systemindex/_search \
-H "Content-Type: application/json"
```