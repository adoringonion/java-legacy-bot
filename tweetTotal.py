import boto3
from boto3.dynamodb.conditions import Key, Attr
from requests_oauthlib import OAuth1Session
import os
from datetime import datetime, timezone, timedelta

dynamodb = boto3.resource('dynamodb')

def lambda_handler(event, context):
    
    count = 0
    
    print('today records:', len(get_today_total()))
    
    for item in get_today_total():
        count += item['CountTweets']
        
    print('today total:', count)

    exec_tweet(count)


def get_today_total():
    
    get_current_time_Tokyo = datetime.now() + timedelta(hours=8)
    today = get_current_time_Tokyo.strftime('%Y-%m-%d')
    
    table = dynamodb.Table('CountTweets')
    
    res = table.scan(
            FilterExpression=Attr('Date').eq(today)
        )
    
    if res :
        return res["Items"]
        
    else :
        return("No data")


def exec_tweet(count):

    CK  = os.environ['CONSUMER_KEY']
    CS  = os.environ['CONSUMER_SECRET_KEY']
    AT  = os.environ['ACCESS_TOKEN']
    ATS = os.environ['ACCESS_TOKEN_SECRET']
    twitter = OAuth1Session(CK, CS, AT, ATS)    
    url = 'https://api.twitter.com/1.1/statuses/update.json'
    
    if count == 0:
        tweet = '昨日、Javaはレガシーになりませんでした　\n #java_legacy'
    
    else:
        tweet = '昨日、Javaは' + str(count) + '回レガシーになりました \n #java_legacy'


    param = {"status" : tweet}
    res = twitter.post(url, params = param)

    if res.status_code == 200:
        print("Posted")

    else:
        print("Failed : %d"% res.status_code)