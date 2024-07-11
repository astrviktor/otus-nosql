### 1. Запуск Elasticsearch
```
make compose-up
```

### 2. Проверка Elasticsearch
```
curl --location 'localhost:9200/_cluster/health?pretty'

{
  "cluster_name" : "docker-cluster",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 0,
  "active_shards" : 0,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}

```

### 3. Создание индекса с маппингом

```
curl --location --request PUT 'localhost:9200/test_index' \
--header 'Content-Type: application/json' \
--data '{
    "settings": {
        "analysis": {
            "filter": {
                "ru_stop": {
                    "type": "stop",
                    "stopwords": "_russian_"
                },
                "ru_stemmer": {
                    "type": "stemmer",
                    "language": "russian"
                }
            },
            "analyzer": {
                "my_russian": {
                    "tokenizer": "standard",
                    "filter": [
                        "lowercase",
                        "ru_stop",
                        "ru_stemmer"
                    ]
                }
            }
        }
    },
    "mappings": {
        "properties": {
            "text": {
                "type": "text",
                "analyzer": "my_russian"
            }
        }
    }
}'

{
  "acknowledged": true,
  "shards_acknowledged": true,
  "index": "test_index"
}

```

### 4. Добавление данных

```
curl --location 'localhost:9200/_bulk' \
--header 'Content-Type: application/json' \
--data '{"create":{"_index":"test_index","_id":"1"}}
{"text":"моя мама мыла посуду а кот жевал сосиски"}
{"create":{"_index":"test_index","_id":"2"}}
{"text":"рама была отмыта и вылизана котом"}
{"create":{"_index":"test_index","_id":"3"}}
{"text":"мама мыла раму"}
'

{
  "errors": false,
  "took": 144,
  "items": [
    {
      "create": {
        "_index": "test_index",
        "_id": "1",
        "_version": 1,
        "result": "created",
        "_shards": {
          "total": 2,
          "successful": 1,
          "failed": 0
        },
        "_seq_no": 0,
        "_primary_term": 1,
        "status": 201
      }
    },
    {
      "create": {
        "_index": "test_index",
        "_id": "2",
        "_version": 1,
        "result": "created",
        "_shards": {
          "total": 2,
          "successful": 1,
          "failed": 0
        },
        "_seq_no": 1,
        "_primary_term": 1,
        "status": 201
      }
    },
    {
      "create": {
        "_index": "test_index",
        "_id": "3",
        "_version": 1,
        "result": "created",
        "_shards": {
          "total": 2,
          "successful": 1,
          "failed": 0
        },
        "_seq_no": 2,
        "_primary_term": 1,
        "status": 201
      }
    }
  ]
}

```

### 5. Запрос всех данных

```
curl --location --request GET 'localhost:9200/test_index/_search' \
--header 'Content-Type: application/json' \
--data '{
    "query": {
        "match_all": {}
    }
}'

{
  "took": 54,
  "timed_out": false,
  "_shards": {
    "total": 1,
    "successful": 1,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": {
      "value": 3,
      "relation": "eq"
    },
    "max_score": 1,
    "hits": [
      {
        "_index": "test_index",
        "_id": "1",
        "_score": 1,
        "_source": {
          "text": "моя мама мыла посуду а кот жевал сосиски"
        }
      },
      {
        "_index": "test_index",
        "_id": "2",
        "_score": 1,
        "_source": {
          "text": "рама была отмыта и вылизана котом"
        }
      },
      {
        "_index": "test_index",
        "_id": "3",
        "_score": 1,
        "_source": {
          "text": "мама мыла раму"
        }
      }
    ]
  }
}

```

### 6. Запрос нечеткого поиска

```
curl --location --request GET 'localhost:9200/test_index/_search' \
--header 'Content-Type: application/json' \
--data '{
    "query": {
        "match": {
            "text": "мама ела сосиски"
        }
    }
}'

{
  "took": 20,
  "timed_out": false,
  "_shards": {
    "total": 1,
    "successful": 1,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": {
      "value": 2,
      "relation": "eq"
    },
    "max_score": 1.2535897,
    "hits": [
      {
        "_index": "test_index",
        "_id": "1",
        "_score": 1.2535897,
        "_source": {
          "text": "моя мама мыла посуду а кот жевал сосиски"
        }
      },
      {
        "_index": "test_index",
        "_id": "3",
        "_score": 0.5376842,
        "_source": {
          "text": "мама мыла раму"
        }
      }
    ]
  }
}

```

Приложена postman-коллекция

### 7. Остановка Elasticsearch
```
make compose-down
```
