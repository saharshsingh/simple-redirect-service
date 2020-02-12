# simple-redirect-service
A simple Go app that can be used to redirect to a specific URL.

## Build

        docker build -t saharshsingh/simple-redirect-service .

## Run

Following will start a server on port 8080 that redirects to https://gainznow.com

        docker run -d \
            --name gainznow-redirect \
            -e SRS_REDIRECT_TARGET="https://gainznow.com" \
            -p 8080:8080 \
            saharshsingh/simple-redirect-service