
# RSS Aggregator
### _Build using Go and Postgresql DB_


 This is a Go server which can automatically extract RSS from multiple feeds.
 These feeds are posted by users and other users can subscribe to them. 


## Following Operations can be performed:

- `create new user`
 `get user details`

- `post new feeds`
 `get all feeds`

- `subscribe to feeds`
 `unsubscribe to feeds`
 `know which feeds a user followed`

    



## Run Project

```sh
    1. create a Postgres database and set its url in .env
    2. go build
    3.  ./rss_aggregator
```
