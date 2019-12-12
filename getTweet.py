import json
import os
from requests_oauthlib import OAuth1Session
from datetime import datetime, timezone, timedelta
from dateutil import parser
from pytz import timezone
import boto3
from boto3.dynamodb.conditions import Key, Attr


def lambda_handler(event, context):
    
    count = 0
    count = search_tweet(count)
    
    put_item(count)


def search_tweet(count):
    
    CK  = os.environ['CONSUMER_KEY']
    CS  = os.environ['CONSUMER_SECRET_KEY']
    AT  = os.environ['ACCESS_TOKEN']
    ATS = os.environ['ACCESS_TOKEN_SECRET']
    twitter = OAuth1Session(CK, CS, AT, ATS)
    
    url = 'https://api.twitter.com/1.1/search/tweets.json'
    query = 'Java　レガシー'
    
    params = {
        'count' : 100,
        'q' : query
    }
    
    req = twitter.get(url, params = params)
    
    print('currentTime:', datetime.now())
    
    if req.status_code == 200:
        res = json.loads(req.text)
    
        for tweet in res['statuses']:
            print("-" * 20)
            print('timeLag:',datetime.now() - parser.parse(tweet['created_at']).replace(tzinfo=None))
            
    
            if datetime.now() - parser.parse(tweet['created_at']).replace(tzinfo=None) <= timedelta(hours=1):
                count += 1
    
    else:
        print('Faiiled: %d' % req.status_code)
    
    return count
    
    
def put_item(count):
    
    dynamodb = boto3.resource('dynamodb')
    
    get_current_time = datetime.now() + timedelta(hours=9)
        
    table = dynamodb.Table('CountTweets')
    table.put_item (
        Item={
            "Date": get_current_time.strftime('%Y-%m-%d'),
            "Time": get_current_time.strftime('%H:%M:%S'),
            "CountTweets": count
        }
    )