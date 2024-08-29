# Challenge Full Cycle 3.0

## Scope

Create a Docker image that print "Full Cycle Rocks!!" in Go. This image needs to have the minimum size.

## Solution

* Create a layer called "builder" with golang in the lastest version
* Build the app
* Create another layer called "prod" using alpine linux in the latest version
* Copy the app's binary from "builder" layer
* Run the app

## Docker Hub Solution Link

Link: https://hub.docker.com/layers/vinicius91carvalho/fullcycle-challenge-golang/0.0.1/images/sha256-2483ea5a1c2f346e0bcba25d89c9e3ca778a788f8b16308f9af2301e172d9c1a?context=repo