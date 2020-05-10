# url-dome-api

## Run locally on docker
##### Run postgres
`> docker run -p 2345:5432 --name url-dome-postgres -e POSTGRES_PASSWORD=secretpassword -e POSTGRES_DB=url-dome -d postgres`

##### Run redis
`> docker run -p 6379:6379 --name redisdb -d redis`

## Build
##### Build using Cloud Build
`> gcloud builds submit --tag gcr.io/url-dome/url-at-minimal-api`

## Deploy
##### Deploy using Cloud Run
`> gcloud run deploy url-at-minimal-api --platform managed --region europe-west1 --allow-unauthenticated --image gcr.io/url-dome/url-at-minimal-api`
