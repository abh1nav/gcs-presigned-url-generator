# GCS Pre-signed URL Generator

This script allows you to generate a pre-signed URL for a GCS object with support for resumable uploads.

## IAM Config

Make sure you have a service account that has Storage Admin permissions. Allow the service account to create tokens:

```shell
gcloud iam service-accounts add-iam-policy-binding SERVICE_ACCOUNT_EMAIL \
  --member="serviceAccount:SERVICE_ACCOUNT_EMAIL" \
  --role="roles/iam.serviceAccountTokenCreator" \
  --project="PROJECT_NAME"
```

## How To

To run this, once the IAM config above has been completed, generate a token for the service account and place it in the same directory as `main.go`.  

Run the following to generate a token:

```
go run main.go \
  --bucket="BUCKET_NAME" \
  --key="object/key.mp4" \
  --credentials="credentials.json"
```

(assuming the crenditial file is called `credentials.json`).
