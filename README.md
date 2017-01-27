# TCuller

TCuller is a tweet culler. It let's you delete all your tweets and favorites.

```shell
tculler -consumer_key $CONSUMER_KEY \
	-consumer_secret $CONSUMER_SECRET \
	-access_token $ACCESS_TOKEN \
	-access_secret $ACCESS_SECRET \
	-username $USERNAME \
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
  -username string
    	Username for media timeline deletion

tculler <tweets|faves|media>
```
