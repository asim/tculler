CONSUMER_KEY=
CONSUMER_SECRET=
ACCESS_TOKEN=
ACCESS_SECRET=
OLDER_THAN=
USERNAME=
CSV_FILE=

tculler \
-consumer_key $CONSUMER_KEY -consumer_secret $CONSUMER_SECRET \
-access_token $ACCESS_TOKEN -access_secret $ACCESS_SECRET \
-username $USERNAME -csv_file=$CSV_FILE -older_than=$OLDER_THAN $1
