# githubmosaic

[![Build Status](https://travis-ci.org/jizuscreed/githubmosaic.svg?branch=master)](https://travis-ci.org/jizuscreed/githubmosaic)

This app generates git repo with commits from a custom image which paint the image at the github dashboard. Like this

![](http://jizuscreed.ru/images/mosaic.png)

#### Installation
1. Clone repository from github
2. Build it by ``go build mosaic.go``

#### Using
1. Run compiled application with image file and new repo dir parameters (both paths are relative). For example: ``mosaic face.jpg test``. New repo directory path are relative from directory parent to current. Process may take 5-40 minutes depending from image and pc performance.
2. Create repository at github
3. Add origin to just created repo and push
4. PROFIT!!!

#### Forking and modifying
Feel free, but all comments in code are in russian. Ke-ke-ke

Site: [http://jizuscreed.ru/githubmosaic/](http://jizuscreed.ru/githubmosaic/)

