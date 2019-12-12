# java-legacy-bot

「Javaはレガシー」と一日で何回ツイートされたか集計して自動ツイートするbot

## 概要

Twitter APIを利用して検索クエリを投げて、指定したワードが含まれているツイート件数を取得し、データベースに保存して、一日の終わりに集計結果を自動ツイートします。

ツイートの取得、データベースへのアクセス、集計ツイートはAWSのLambdaを使って自動化し、データベースはDynamoDBを使用しています。

<img src="https://user-images.githubusercontent.com/43922475/70678984-1723f400-1cd7-11ea-87bb-57ab4f41cd7d.png">

## ツイートの取得
Twitter APIへクエリを投げるのと、ツイート件数の取得は'search_tweet()'内で行なっています。'query'部分をいじることで好きな検索条件で検索できます。

'''
query = 'Java　レガシー'
    
    params = {
        'count' : 100,
        'q' : query
    }
'''

## 集計結果のツイート
