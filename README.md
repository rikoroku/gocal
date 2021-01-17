# gocal

gocal is a golang application that allows you to access your Google Calendar from a command line. It's easy to manage your events.

## Installation

```bash
$ go get -u github.com/rikoroku/gocal
```

The steps below are the first time only.

#### 1. Set environment variable.

You need to turn on the Google Calendar API and get client_secret.json.

About turn on the Google Calendar API and get client_secret.json:

1. Use this [wizard](https://console.developers.google.com/start/api?id=calendar) to create or select a project in the Google Developers Console and automatically turn on the API. Click *Continue*, then *Go to credentials*.
1. On the *Add credentials to your project* page, click the *Cancel* button.
1. At the top of the page, select the *OAuth consent screen* tab. Select an *Email address*, enter a Product name if not already set, and click the *Save* button.
1. Select the *Credentials* tab, click the *Create credentials* button and select *OAuth client ID*.
1. Select the application type *Other*, enter the name "gocal", and click the Create button.
1. Click *OK* to dismiss the resulting dialog.
1. Click the (Download JSON) button to the right of the client ID.
1. Move this file to your working directory and rename it `client_secret.json`.

```
$ export GOCAL_CLIENT_SECRET_FILE_PATH="Your client_secret.json path"
```

#### 2. Authorization
The authorization procedure is as follows.

```
$ gocal setup

Go to the following link in your browser then type the authorization code:
https://xxxx/

Got resulting code? Please input your resulting code
```

## Contributing

1. Fork it ( https://github.com/rikoroku/gocal )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

Bug reports and pull requests are welcome.

## License

This tool is available as open-source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
