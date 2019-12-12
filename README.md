# java-legacy-bot

「Javaはレガシー」と一日で何回ツイートされたか集計して自動ツイートするbot

## 概要

Twitter APIを利用して検索クエリを投げて、指定したワードが含まれているツイート件数を取得し、データベースに保存して、一日の終わりに集計結果を自動ツイートします。

ツイートの取得、データベースへのアクセス、集計ツイートはAWSのLambdaを使って自動化し、データベースはDynamoDBを使用しています。

<img src="https://user-images.githubusercontent.com/43922475/70678984-1723f400-1cd7-11ea-87bb-57ab4f41cd7d.png">

## ツイートの取得
 Twitter APIへクエリを投げるのと、ツイート件数の取得は`search_tweet()`内で行なっています。`query`部分をいじることで好きな検索条件で検索できます。

```
query = 'Java　レガシー'
    
    params = {
        'count' : 100,
        'q' : query
    }
```

 以下の部分で取得したツイートの件数をカウントしています。条件式`if datetime.now() - parser.parse(tweet['created_at']).replace(tzinfo=None) <= timedelta(hours=1)`は現在時刻とツイートされた時刻を比較して、1時間以内のものだけをカウントするようにするものですが、これはTwitter APIの仕様で一度に取得できるツイートが１００件までなので、1時間ごとに区切って取得するためです。今回はあまりツイートされないキーワードなので1時間にしていますが、もっとメジャーなキーワードを検索するあ場合、もっと短く区切るべきかもしれません。`<= timedelta(hours=1)`に好きなdatetime.timedelta型を指定すれば変更できます。

```
if req.status_code == 200:
        res = json.loads(req.text)
    
        for tweet in res['statuses']:
            print("-" * 20)
            print('timeLag:',datetime.now() - parser.parse(tweet['created_at']).replace(tzinfo=None))
            
    
            if datetime.now() - parser.parse(tweet['created_at']).replace(tzinfo=None) <= timedelta(hours=1):
                count += 1
```
## データベースへの登録
　データベースへの登録は`put_item(count)`で行なっています。日本時間に合わせるため、現在時刻に9時間を足しています。
```
get_current_time = datetime.now() + timedelta(hours=9)
        
    table = dynamodb.Table('CountTweets')
    table.put_item (
        Item={
            "Date": get_current_time.strftime('%Y-%m-%d'),
            "Time": get_current_time.strftime('%H:%M:%S'),
            "CountTweets": count
        }
    )
```
## 集計
　
## 集計結果のツイート
