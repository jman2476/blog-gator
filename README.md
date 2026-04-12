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
Go to the [Webi Golang installation page](https://webinstall.dev/golang/), find your operating system and copy-paste the instructions into your terminal. If you are on Windows and using WSL, choose the Linux tab and use that to install. After pasting in your terminal, press enter and follow any instructions that may come up.

Once the installation is complete, run `go version` to verify that it worked. If this did not work on your machine, try the official installation instructions below.

**Official installation:** Use the [official Golang download instructions](https://go.dev/doc/install) for your operating system. Again, if you are using WSL on Windows, make sure to follow the Linux instructions.

Once the installation is complete, run `go version` to verify that it worked. 

**Important for Linux/Windows:** For Windows, you will be asked to create a user and password for admin access to Postgres; make sure you not this for the next step. It may also indicate the port it's running on, note this. For Linux, after installing Postgres, run the command `sudo passwd postgres` and enter a password. Again, save this for later.

### Install PostgreSQL
*You may skip this step if you already have PostgreSQL version 15 or newer installed*
Postgres, or PostgreSQL, is the database used by Gator, and is vital for its operation. If you already have Postgres installed, run `psql --version` in your terminal to verify that you have version 15 or newer. Otherwise, follow the instructions below for your operating system:

**MacOS/Linux/WSL Only: Webi Install:** Go to the [Webi Postgres installation page](https://webinstall.dev/postgres/), copy paste the commands into your terminal, and press enter. Follow any prompts that come up, and then run `psql --version` to verify your installation.

**Windows/All operating systems: EDB Download:** Get the link to the [installer for your operating system](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads) from EnterpriseDB, the official distributor for Postgres, and then follow the instructions in the installer and/or in the [official installation instructions](https://www.enterprisedb.com/docs/supported-open-source/postgresql/installing/) for your operating system and architecture. Once done, run `psql --version` in your terminal to verify your installation.


## Installation
Once Go and Postgres are installed, you can run `go install github.com/jman2476/gator` to install Gator. 

### Config Setup
Before running Gator, you will need to setup a config file in your home directory. Create a file `~/.gatorconfig.json`, and paste the following code into it: 

```
{
    "db_url": "database connection string"
}
```

The database connection string you use will depend on your operating system:
| MacOS | postgres://<your username>:@localhost:5432/gator?sslmode=disable |
| Linux | postgres://postgres:<your postgres password>@localhost:5432/gator?sslmode=disable |
| Windows | postgres://<your postgres username>:<your postgres password>@localhost:5432/gator?sslmode=disable  |


## Usage
To use gater, enter `gator [command]` in your terminal. Below is a table of commands, their arguments, and usage.

```
gator [command] [arg1] [arg2]
```

| Command | Description | Arguments |
| - | - | - |
| register | Register a new user | username |
| login | Login existing user | username |
| users | List stored users | - |
| reset | Delete all users, feeds | - |
| addfeed | Add RSS feed | title, url |
| follow | Adds an existing feed to user's follows by URL | url |
| following | Shows user's followed feeds | - |
| unfollow | Remove feed from user's follows by URL | url |
| feeds | Shows a list of feeds in database | - |
| agg | Gathers posts from all followed feeds | loop_time [Format: 1m, 10s, 4d, etc] |
| browse | Shows a list of arguments | limit [int] |

## Credits
Lovingly handtyped by [Jeremy McKeegan](https://github.com/jman2476)

Idea and guided project from [Boot.dev](https://www.boot.dev/courses/build-blog-aggregator-golang)

Title font made with tool from [Pator JK](https://patorjk.com/software/taag/)