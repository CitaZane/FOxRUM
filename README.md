[//]: # "Do Ctrl + Shift + V for a better viewing experience"

# FOxRUM

Real time forum project.

## Project setup

- Make sure when installing dependencies / starting servers, to be in the right folder!

#### For front-end server

```
cd ./front-end
```

#### For back-end server

```
cd ./server
```

- When both servers are up and running, you can interact with forum here -> [http://localhost:8080](http://localhost:8080/login)

- Check out [Audit questions](https://github.com/01-edu/public/tree/master/subjects/real-time-forum/audit)

## Install

#### Front-end server (Vuejs)

```
npm install
```

## Execute

#### Front-end server (Vuejs)

```
npm run serve
```

#### Back-end server (go)

```
go run .
```

## Resources & Acknowledgments

- [SVG filter calcualtion](https://codepen.io/sosuke/pen/Pjoqqp)
- [Websocket connection](https://www.whichdev.com/go-vuejs-chat/)
- [Presisting state with pinia plugin](https://github.com/prazdevs/pinia-plugin-persistedstate)
- [Icons from fontawesome](https://fontawesome.com/icons)
- Artwork by Me :)

## Author

- Zane Krūmiņa

### What I learned

- First interaction with front end framework - Vue
- Front end form validation with validation library
- Websocket basics
- Real time messaging and notifications
- Creating a database structure
- Using interfaces with go for better compatibility
- Using channels/ gorutines with Go

### To be improved

- UX/UI -accessibility
- Design is not responsive
- Forum safety - more validation for user input
- Error handling could be improved/polished
- CSS cleanup (could be some repetetive styles)
