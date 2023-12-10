
# RSS Aggregator using Go and Postgresql DB

    This is a Go server which can automatically extract RSS from multiple feeds.
    These feeds are posted by users and other users can subscribe to them. 



    ## Following Operations can be performed:

    * `create new user`
    * `get user details`

    * `post new feeds`
    * `get all feeds`

    * `subscribe to feeds`
    * `unsubscribe to feeds`
    * `know which feeds a user followed`

    



    ## Run Project

    1. create a Postgres database and set its url in .env
    1. go build
    2. ./rss_aggregator