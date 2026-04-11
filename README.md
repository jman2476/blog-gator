```                                                                                  
     ,o888888o.          .8.    8888888 8888888888 ,o888888o.     8 888888888o.   
    8888     `88.       .888.         8 8888    . 8888     `88.   8 8888    `88.  
 ,8 8888       `8.     :88888.        8 8888   ,8 8888       `8b  8 8888     `88  
 88 8888              . `88888.       8 8888   88 8888        `8b 8 8888     ,88  
 88 8888             .8. `88888.      8 8888   88 8888         88 8 8888.   ,88'  
 88 8888            .8`8. `88888.     8 8888   88 8888         88 8 888888888P'   
 88 8888   8888888 .8' `8. `88888.    8 8888   88 8888        ,8P 8 8888`8b       
 `8 8888       .8'.8'   `8. `88888.   8 8888   `8 8888       ,8P  8 8888 `8b.     
    8888     ,88'.888888888. `88888.  8 8888    ` 8888     ,88'   8 8888   `8b.   
     `8888888P' .8'       `8. `88888. 8 8888       `8888888P'     8 8888     `88. 

    ____  __               ___                    _             __            
   / __ )/ /___  ____ _   /   | ____ _____ ______(_)___ _____ _/ /_____  _____
  / __  / / __ \/ __ `/  / /| |/ __ `/ __ `/ ___/ / __ `/ __ `/ __/ __ \/ ___/
 / /_/ / / /_/ / /_/ /  / ___ / /_/ / /_/ / /  / / /_/ / /_/ / /_/ /_/ / /    
/_____/_/\____/\__, /  /_/  |_\__, /\__, /_/  /_/\__, /\__,_/\__/\____/_/     
              /____/         /____//____/       /____/                        
```
# Gator
Gator: Blog Aggrigator is a locally hosted command line tool for fetching, aggrigating and browsing your favorite RSS feeds. Simply register a username, add a feed you would like to follow, and browse the articles from those blogs. 

To install and use Gator, please read through and follow carefully the instructions bellow.

## Requirements: PostgreSQL and Go
Gator is built in [Go](https://go.dev/) and uses [PostgreSQL](https://www.postgresql.org/) for the database. You will need to install Go in order to download Gator, and you will need to set up PostgreSQL so the it will work properly.

### Install Go
*You may skip this step if you already have Go version 1.26 or newer installed*

Go, or Golang, is the language that Gator is built on, and will be necessary for installing Gator later in this guide. The easiest way to download Go will be through `webi`, and the official Go installation documentation will be linked underneath that. 

If you already have Go but do not know what version you have installed, run `go version` in your terminal to check. You want to have version 1.26.0 or newer.

**Webi install:**
Go to [https://webinstall.dev/golang/](https://webinstall.dev/golang/), find your operating system and copy-paste the instructions into your terminal. If you are on Windows and using WSL, choose the Linux tab and use that to install. After pasting in your terminal, press enter and follow any instructions that may come up.

Once the installation is complete, run `go version` to verify that it worked. If this did not work on your machine, try the official installation instructions below.

**Official installation:** Use the [official Golang download instructions](https://go.dev/doc/install) for your operating system. Again, if you are using WSL on Windows, make sure to follow the Linux instructions.

Once the installation is complete, run `go version` to verify that it worked.

### Install PostgreSQL
*You may skip this step if you already have PostgreSQL version 15 or newer installed*

## Installation


### Config Setup









## Credits
Lovingly handtyped by [Jeremy McKeegan](https://github.com/jman2476)

Idea and guided project from [Boot.dev](https://www.boot.dev/courses/build-blog-aggregator-golang)

Title font made with tool from [Pator JK](https://patorjk.com/software/taag/)