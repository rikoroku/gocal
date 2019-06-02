# Goedule
## Status
WIP

## Setup
This step is first time only.

### Set environment variable.

You need to turn on the Google Calendar API and get client_secret.json.

About turn on the Google Calendar API and get client_secret.json:

1. Use this [wizard](https://console.developers.google.com/start/api?id=calendar) to create or select a project in the Google Developers Console and automatically turn on the API. Click *Continue*, then *Go to credentials*.
2. On the *Add credentials to your project* page, click the *Cancel* button.
3. At the top of the page, select the *OAuth consent screen* tab. Select an *Email address*, enter a Product name if not already set, and click the *Save* button.
4. Select the *Credentials* tab, click the *Create credentials* button and select *OAuth client ID*.
5. Select the application type *Other*, enter the name "Goedule", and click the Create button.
6. Click *OK* to dismiss the resulting dialog.
7. Click the (Download JSON) button to the right of the client ID.
8. Move this file to your working directory and rename it `client_secret.json`.

```
$ export CLIENT_SECRET_FILE="Your client_secret.json path"
$ source ~/.bash_profile
```

### Authorization
The authorization procedure is as follows.

```
$ go run Goedule setup

Go to the following link in your browser then type the authorization code:
https://xxxx/

Got resulting code? Please input your resulting code
```
