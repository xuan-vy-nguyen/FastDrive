# CNPM-Project01

## Hello teacher
This project has 2 branch:
- Master: Future-Project.
- IntroToSE: is used for submitting Software-Engineering Final-Project
- Mobile: is used for submitting Mobile-Development Final-Project <br/>
Current Branch is Master!

## Intro
First, you should go to this link and learn more about Git if you are beginner.
```bash
https://github.github.com/training-kit/downloads/github-git-cheat-sheet.pdf
```

## Clone 
Use this command to clone.

```bash
git clone https://github.com/xuan-vy-nguyen/CNPM-Project01.git
```

## Requirements
In this project, we should learn and use MongoDB, Golang, Docker, Heroku, JWT, IFTTT, Redis in backend-server. Why? Because they are so clearly and friendly.
In FrontEnd, there are many options: App, WebApp, Web,... you can discuss and vote about this topic on teamChat.

## How to setup server
0. Clone this repo

1. first install go-mongo-driver
- install golang & mongodb
- install dep
- go to golang_workspace and create PROJECT_WORKSPACE
- type "dep init"
- How to use Mongo-Driver? follow this link :
```
https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial 
https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj
```

2.  use jwt
- https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/

3. upload large file to mongoDB
- https://stackoverflow.com/questions/39039560/upload-image-from-android-to-golang-server-and-save-it-in-mongodb
- https://www.mongodb.com/blog/post/quick-start-golang--mongodb--a-quick-look-at-gridfs

4.  Setup on Heroku
- heroku.yml
- https://devcenter.heroku.com/articles/build-docker-images-heroku-yml
- view heroku.logs
- https://devcenter.heroku.com/articles/logging

5.  setup on VM Google 
- use UFW https://www.digitalocean.com/community/tutorials/how-to-set-up-a-firewall-with-ufw-on-ubuntu-16-04
- Allow port 80
- redirect port 80 to port server

## API 
Import APIServer.json to Postman and follow it. It's so ez.

## About content
- I include Example/Pytorch/fast_neural_style in my repo. U can clone it in here: https://github.com/pytorch/examples.git.
- I am going to develop this system. All I want is that this API can be embedded with lifelog retrieval or image enhancement, become a LARGE SYSTEM in AI-Engineering.