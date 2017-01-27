# TCuller

TCuller is a tweet culler. It let's you delete all your tweets and favorites.

```shell
go get github.com/asim/tculler
```

```shell
tculler -consumer_key $CONSUMER_KEY \
	-consumer_secret $CONSUMER_SECRET \
	-access_token $ACCESS_TOKEN \
	-access_secret $ACCESS_SECRET \
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
  -username string
    	Username for media timeline deletion

tculler <tweets|search|faves|media|csv>
```
