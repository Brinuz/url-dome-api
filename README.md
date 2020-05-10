# url-at-minimal-api

## Run locally on docker
##### Run redis
`> docker run -p 6379:6379 --name redisdb -d redis`

##### Build using Cloud Build
`> gcloud builds submit --tag gcr.io/url-dome/url-at-minimal-api`

##### Deploy using Cloud Run
`> gcloud run deploy url-at-minimal-api --platform managed --region europe-west1 --allow-unauthenticated --image gcr.io/url-dome/url-at-minimal-api`