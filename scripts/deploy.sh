#!/bin/bash

cd ../dist
zip -r -D ./dist/Alexa-Guess-Anything-API.zip *
aws lambda update-function-code --function-name  Alexa-Guess-Anything-API --zip-file fileb://dist/Alexa-Guess-Anything-API.zip