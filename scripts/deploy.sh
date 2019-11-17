#!/bin/bash

cd ./dist
[ -e Alexa-Guess-Anything-API.zip ] && rm Alexa-Guess-Anything-API.zip
zip -r -D Alexa-Guess-Anything-API.zip *
aws lambda update-function-code --function-name  Alexa-Guess-Anything-API --zip-file fileb://Alexa-Guess-Anything-API.zip