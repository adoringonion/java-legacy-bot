import json
import config
from requests_oauthlib import OAuth1Session
from datetime import datetime, timezone


CK = config.API_KEY
CS = config.API_SECRET_KEY
AT = config.ACCESS_TOKEN
ATS = config.ACCESS_TOKEN_SECRET
twitter = OAuth1Session(CK, CS, AT, ATS)
from pytz import timezone
from dateutil import parser
url = 'https://api.twitter.com/1.1/search/tweets.json'
keyword = 'Java　レガシー'

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

else:
    print('Faiiled: %d' % req.status_code)