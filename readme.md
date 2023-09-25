##### Config
Copy .env.sample to .env and put in your data

##### Running Dev Mode
- Start ngrok `ngrok http 1323`
- Run `npm update:env` to update env file url
- Run `npm update:url` to update url in partner dashboard
 
Then start main.go with Nodemon:
```
nodemon -e go --signal SIGTERM --exec 'go' run .
```

##### Install Nodemon
`npm i -g nodemon`