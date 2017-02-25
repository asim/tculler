# TCuller

TCuller is a tweet culler. It let's you delete all your tweets and favorites.

```shell
go get github.com/asim/tculler
```

```shell
tculler \
-consumer_key $CONSUMER_KEY -consumer_secret $CONSUMER_SECRET \
-access_token $ACCESS_TOKEN -access_secret $ACCESS_SECRET \
-username $USERNAME -csv_file=$CSV_FILE -older_than=$OLDER_THAN \
tweets
```

## Usage

```shell
Usage of tculler:
  -access_secret string
    	Twitter access secret
  -access_token string
    	Twitter access token
  -consumer_key string
    	Twitter consumer key
  -consumer_secret string
    	Twitter consumer secret
  -csv_file string
    	Path to tweets csv file (default "tweets.csv")
  -older_than string
    	Delete tweets older than (1m|1h|1d) default deletes all
  -username string
    	Username for media timeline deletion

tculler <tweets|search|faves|media|csv>
```
