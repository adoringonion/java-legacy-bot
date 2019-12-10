import json
import config
from requests_oauthlib import OAuth1Session
from datetime import datetime, timezone, timedelta
from dateutil import parser
from pytz import timezone

CK = config.API_KEY
CS = config.API_SECRET_KEY
AT = config.ACCESS_TOKEN
ATS = config.ACCESS_TOKEN_SECRET
twitter = OAuth1Session(CK, CS, AT, ATS)

url = 'https://api.twitter.com/1.1/search/tweets.json'
keyword = 'Java　レガシー'
count = 0

params = {
    'count' : 100,
    'q' : keyword
}

req = twitter.get(url, params = params)

if req.status_code == 200:
    res = json.loads(req.text)
    print(len(res['statuses']))

    for tweet in res['statuses']:
        print(datetime.now() - parser.parse(tweet['created_at']).replace(tzinfo=None))
        print("-" * 10)

        if datetime.now() - parser.parse(tweet['created_at']).replace(tzinfo=None) <= timedelta(hours=1):
            count += 1
        
    
    print(count)

else:
    print('Faiiled: %d' % req.status_code)