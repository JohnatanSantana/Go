import pandas as pd
import csv
import json
from pymongo import MongoClient
from elasticsearch import Elasticsearch
from elasticsearch import helpers
import redis
from pyhive import presto



# MONGO DB
username = 'root'
password = 'root'
server = '3.82.6.150'
client = MongoClient('mongodb://{}:27017/'.format(server))

banco = client['john']
collection = banco['nba']
collection = banco.album

# ELASTIC
es = Elasticsearch(hosts= server)

# REDIS

r = redis.Redis(
    host=server,
    port=6379)


#CONNECT PRESTO
connect = presto.connect(host=server,port=8080)

def read_presto():
    df = pd.read_sql_query('select * from elastic.default.john',connect)
    print(df.head())

def write_mongo(data_dict):
    try:
        data_id = collection.insert_many(data_dict)
        print(data_id)
    except NameError:
        print(NameError)


def read_csv(path):
    data = pd.read_csv(path)
    data_json = json.loads(data.to_json(orient='records'))
    return data_json


def read_mongo():
    mongo_data = collection.find()
    print('oi')
    print(list(mongo_data))


def write_elastic(data):
    try:
        # //data_id = es.index(index='john', doc_type='nba', body=data)
        data_id = helpers.bulk(es,index='john', doc_type='nba',actions=data)
        print(data_id)
    except NameError as e:
        print(e)


# def write_redis(data):


path = '/home/johnatan/Documents/go/Go/ApiMongo/mongo/files/NBAAllStars2000-2016.csv'
df = read_csv(path)

read_presto()
# write_elastic(df)
# write_mongo(df)
# read_mongo()
