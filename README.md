# Go-DataSearch

##### Run `./start.sh` to download the dependencies and run the the application

To run the application, you have to define the environment variables, default values of the variables are defined inside `start.sh`

- SERVER_ADDRESS    `[IP Address of the machine]`
- SERVER_PORT       `[Port of the machine]`
- DB_USER           `[Database username]`
- DB_PASSWD         `[Database password]`
- DB_ADDR           `[IP address of the database]`
- DB_PORT           `[Port of the database]`
- DB_NAME           `[Name of the database]`

# create .gitignore file
```sh
git init
# create a github repository  github.com/new

git remote add origin https://github.com/Super-Phantoms/Go-DataSearch.git
# add all files inside this project
git add .
git commit -m "extracted the error and logger from api"
#adding tag
git tag v1.0.0
git push origin master --tags

# search key words: using go moudles
# change git repository
git status
git add .gitignore ReadMe.md
git commit -m "added zap dependency"
git tag v1.0.1
git push origin master --tags
git push -f origin master