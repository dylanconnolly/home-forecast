to get a new token for Google API:
```
curl -L -X POST 'https://www.googleapis.com/oauth2/v4/token?client_id=<OAUTH_client_id&client_secret=<OAUTH_client_secret>&refresh_token=<refresh_token>&grant_type=refresh_token'
```