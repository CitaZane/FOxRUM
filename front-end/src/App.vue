<template>
  <NavBar :resp="resp"/>
  <router-view :resp="resp"/>
  <div :class="popupState" id="popup">Welcome to the Forest, {{name}}</div>
</template>

<script>
import NavBar from '@/components/navbar/NavBar.vue'
import { mapState, mapWritableState, mapActions } from 'pinia' //helper for store
import {useMainStore} from '@/stores/main.js'

export default {
  components: {NavBar},
  data:function(){
    return{
      serverUrl: "ws://localhost:8081/ws",
      resp:{
        activeChats:{},
        userList:[],
        userStats: {}, 
        categories:[],
        posts:[], 
        post:{},
        comments:[],
      }
    }
  },
  computed:{
    ...mapWritableState(useMainStore,["ws"]),
    ...mapState(useMainStore,["isLoggedIn"]),
    ...mapWritableState(useMainStore,["name"]),
    ...mapWritableState(useMainStore,['chatNotification']),
    ...mapState(useMainStore,["darkMode"]),
    popupState(){
        return (this.darkMode)? "welcome-popup-dark": "welcome-popup"
    }
  },
  methods:{
    ...mapActions(useMainStore, {resetStore : 'reset'}),
    listenToWebsocket() {
        this.ws.addEventListener('message', (event) => { this.onWebsocketMessage(event) });
    },
      onWebsocketMessage(event) {
        let data = event.data;
        data = data.split(/\r?\n/);

        for (let i = 0; i < data.length; i++) {
          let msg = JSON.parse(data[i]);
          switch (msg.action) {
          case "get-all-users":
              this.handleGetAllUsers(msg);
              break;
          case "get-username":
              this.name = (msg.message)
              break;
          case "user-logout":
              this.handleUserLogout(msg);
              break;
          case "user-join":
              this.handleUserJoin(msg);
              break;
          case "user-stats":
              this.handleUserStats(msg);
              break;
          case "join-chat":
              this.handleJoinChat(msg);
              break;
          case "new-chat-message":
              this.handleNewChatMessage(msg);
              break;
          case "chat-message-request":
            this.handleChatMessageRequest(msg);
              break;
          case "get-categories":
            this.handleGetCategories(msg);
              break;
          case "get-posts":
            this.handleGetPosts(msg);
              break;
          case "get-category-posts":
            this.handleGetPosts(msg);
              break;
          case "more-posts":
            this.handleMorePosts(msg);
              break;
          case "get-user-posts":
            this.handleGetUserPosts(msg);
              break;
          case "get-post":
            this.handleGetPost(msg);
              break;
          case "new-comment":
            this.handleNewComment(msg);
              break;
          case "get-comments":
            this.handleGetComments(msg);
              break;
          case "more-comments":
            this.handleMoreComments(msg);
              break;
          case "throw-out":
            this.handleThrowOut(msg);
              break;
          case "type":
            this.handleType(msg);
              break;
          default:
              break;
          }
        }
      },
      handleGetAllUsers(msg) {
        this.resp.userList = msg.users;
      },
      handleUserJoin(msg){
        if (msg.sender.name == this.name)return; // catch if current user
        if (msg.message == "signup"){
          // Add to all user list
          this.resp.userList.push(msg.sender)
        }else{
          // change status 
          let index = this.resp.userList.findIndex((user)=> user.name == msg.sender.name)
          if (index != -1)this.resp.userList[index].online = true
        } 
      },
      handleUserLogout(msg){
        let index = this.resp.userList.findIndex((user)=> user.name == msg.sender.name)
        if (index != -1)this.resp.userList[index].online = false
      },
       handleUserStats(msg){
         this.resp.userStats = msg.stats
       },
       //chat joined with message history from server
      handleJoinChat(msg){
        if (typeof msg.Target == 'undefined') return //target holds chatId
        if (msg.message != "valid"){
          this.$router.push({path:'/not-found'})
          return
        }
        let response = {messages: msg.chatmessages, unread: false}
        this.resp.activeChats[msg.Target] = response;
        this.handleChatNotification()
      },
      // detected incoming message from server
      handleNewChatMessage(msg){
        let isAuthor = msg.sender.name == this.name // holds tru for msg sauthor
        let chatOpen = this.$route.params.id == msg.sender.name //holds true if open
        // add message to the list for author or chat is open
        if(isAuthor || chatOpen){
          // check if chat is not empty
          if (!this.resp.activeChats[msg.Target].messages){
            this.resp.activeChats[msg.Target].messages = []
          }
          this.resp.activeChats[msg.Target].messages.unshift(msg.chatmessages[0]);
          this.resp.activeChats[msg.Target].unread=false;
          this.pushUserUpInList(msg)
          return
        }
        // Chat is not open
        // notify about new message using empty object with connection as key
          let response = {messages: [], unread: true}
          this.resp.activeChats[msg.Target] = response;
          this.chatNotification = true;
          this.pushUserUpInList(msg)
      },
      pushUserUpInList(msg){ // Move user upwards in user list
        if (msg.chatmessages[0].sender == this.name){
          let index = this.resp.userList.findIndex((user)=> user.name == msg.chatmessages[0].reciver)
          if (index != -1)this.resp.userList[index].lastMsgTime =  msg.chatmessages[0].time
        }else{
          let index = this.resp.userList.findIndex((user)=> user.name == msg.chatmessages[0].sender)
          if (index != -1)this.resp.userList[index].lastMsgTime =  msg.chatmessages[0].time
        }
      },
      handleChatMessageRequest(msg){
        if(!msg.chatmessages) return
        this.resp.activeChats[msg.Target].messages = [].concat(this.resp.activeChats[msg.Target].messages, msg.chatmessages)
      },
      // Check in no unread messages and configure global notification state
      handleChatNotification(){
        for (let chat in this.resp.activeChats){
          if(this.resp.activeChats[chat].unread){
            this.chatNotification = true;
            return
          }
        }
        this.chatNotification = false;
      },
      handleGetCategories(msg){
        this.resp.categories = msg.categories
      },
      handleGetPosts(msg){
        if (!msg.posts){
          this.$router.push({path:'/not-found'})
          return
        }
        this.resp.posts = msg.posts
      },
      handleGetUserPosts(msg){
        this.resp.posts = msg.posts
      },
      handleMorePosts(msg){
        if(!msg.posts) return
        this.resp.posts.push(...msg.posts)
      },
      handleGetPost(msg){
        if(!msg.posts) return
        if(!msg.posts[0]) {
          this.$router.push({path:'/not-found'})
          return
        }
        this.resp.post = msg.posts[0]
      },
      handleNewComment(msg){
        if(!msg.comments) return
        this.resp.comments.unshift(...msg.comments)
      },
      handleGetComments(msg){
        if(!msg.comments){
          this.resp.comments = []
        }else{
          this.resp.comments = msg.comments
        }     
      },
      handleMoreComments(msg){
        if(!msg.comments) return;
        this.resp.comments.push(...msg.comments)
      },
      handleThrowOut(){
        this.resetStore()
        this.$router.push({name:'login'})
      },
      handleType(msg){
        if (msg.message == "start"){
            this.resp.activeChats[msg.Target].typingInAction=true;
        }else{
          this.resp.activeChats[msg.Target].typingInAction=false;
        }

      },
      showPopup(){
        let popup = document.querySelector('#popup')
        popup.style.display = "block"
        setTimeout(function() {
            let popup = document.querySelector('#popup')
        popup.style.display = "none"
        }, 3000);
      }
  },
  mounted(){
    if(this.isLoggedIn && Object.keys(this.ws).length == 0){
      this.ws = new WebSocket( this.serverUrl);
      this.listenToWebsocket()
    }
    
  },
  watch:{
    ws(to, from){
      if (from == null && to != null){
        this.listenToWebsocket()
      }
    },
    name(name){
      if (!name) return;
      this.showPopup()
    }
  }
}
</script>


<style>
*,
*::after,
*::after {
  box-sizing: border-box;
}

* {
  margin: 0;
  padding: 0;
}
html {
  scroll-behavior: smooth;
}
html{
    /* Variables for default theme (light-mode) */
    --main-color: #0f4c5c;
    --main-color-active: #0b3742;
    --accent-color-dark: #e36414;
    --accent-color-light: #fb8b24;
    --accent-color-notification: #fb8c243f;
    --error-color: #e33d14;
    --content-color-dark:#aeb7b3;
    --content-color: #E4E6E5;
    --content-color-active: #d2d6d4; 
    --accent-text-color: #f5f5f5;
    /* Font sizes */
    /* 10 / 14 / 16 / 20 px */
    --font-xs: 0.625rem;
    --font-s: clamp(0.5rem, 1vh, 0.625rem);
    /* --font-s: 0.875rem; */
    --font-s: clamp(0.75rem, 1.3vh, 0.875rem);
    /* --font-m: 1rem; */
    --font-m: clamp(0.875rem, 1.5vh, 1rem);
    /* --font-l: 1.25rem;  */
    --font-l: clamp(1rem, 2vh, 1.25rem);
}
html.dark-mode{
  /* Variables for default theme (light-mode) */
    --main-color: #E4E6E5;
    --main-color-active: #aeb7b3;
    --accent-color-dark: #e36414;
    --accent-color-light: #fb8b24;
    --accent-color-notification: #fb8c243f;
    --error-color: #e33d14;
    --content-color-dark: #0f4c5c;
    --content-color: #0b3742;
    --content-color-active: #146275; 
    --accent-text-color: #0f4c5c;
}
#app {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: var(--main-color);
  display: flex;
  flex-wrap: nowrap;
  height: 100%;
  position: relative;
  
}
/* Some general widely used classes */
.main-container{
  background-color: var(--accent-text-color); 
    width:73vw;
    min-width: 547.5px; 
    height: 90vh;
    min-height: 405px;
    margin-top: 5vh;
    border-radius: 1rem;
    padding: 2vh 2vw 2vh 2vw;
    overflow: hidden;

    display: flex;
    flex-flow: column nowrap;
    justify-content: flex-start;
}
.main-wrapper{
  background-color: var(--content-color);
   width: 77.5vw;
   min-width: 581.25px;
}
.side-wrapper{
    background-color: var(--content-color);
    width: 18vw;
    min-width: 135px;
    height: 100vh;
    min-height: 450px;
}
.side-container{
    height: 90vh;
    min-height: 405px;
    margin-top: 5vh;
    padding: 0vh 2vw 0vh 2vw;
    overflow: hidden;
}
@media (max-width: 900px){
  .main-wrapper{
    width:65.5vw;
    /* min-width: 491.25px; */
  }
  .side-wrapper{
    width:30vw;
    /* min-width: 225px; */
  }
  .main-container{
    width:63.5vw;
    /* min-width:476,25px;  */
  }
}
.main-title{
  font-size: var(--font-l);
  font-weight: bold;
    color: var(--main-color);
}
.input-field{
  background-color: var(--content-color-active);
    border-radius: 0.5rem;
    border-style: none;
    width: 100%;
    padding: 1vh 1vw 1vh 1vw;
    font-size: var(--font-s);
    color:var(--main-color);
}
.input-field:focus{
  outline: 2px solid var(--accent-color-dark);
}
.input-field:autofill {
  background: var(--content-color-active);
}
.input-field::-ms-reveal,
      input::-ms-clear {
        display: none;
      }
::placeholder { /* Chrome, Firefox, Opera, Safari 10.1+ */
  color: var(--main-color);
  opacity: 1; /* Firefox */
}
.btn{
    background-color: var(--accent-color-dark);
    color:var(--accent-text-color);
    font-size: var(--font-s); 
    font-weight: bold; 
    text-align: center;
    padding: 1vh 3vw 1vh 3vw;
    border-radius: 0.5rem;
    cursor: pointer;
}
.btn:hover{
  background-color: var(--accent-color-light);
}
.divider{
  width:100%;
    height: 0.5px;
    border: 0 none;
    background-color: var(--content-color-dark);
    margin-top:2vh;
    margin-bottom: 4vh;
    margin-left: auto;
    margin-right: auto;
}
.side-divider{
  margin-top:4vh;
}
.welcome-popup{
  background: url("@/assets/popup.png");
}
.welcome-popup-dark{
  background: url("@/assets/popup_dark.png");
}
.welcome-popup, .welcome-popup-dark{
  display: none;
  position:absolute;
  top: 2vh;
  left: 4.5vw;
  background-size:contain;
  background-repeat: no-repeat;
  z-index: 10 !important;
  font-size: var(--font-m);
  font-weight: bold;
  white-space: wrap;
  width: 18vw;
  min-width: 215px;
  aspect-ratio: 2.3/1;
  padding-top: 2vw;
  padding-left: 4vw;
  padding-right: 2vw;
  color:var(--main-color);
}
@media (max-width: 900px){
  .welcome-popup, .welcome-popup-dark{
    left:  33.75px;
    padding-top: 25px;
    padding-left: 40px;
    padding-right: 25px;
  }
}


/* scrollbar */
* {
  scrollbar-width: thin;
  scrollbar-color: var(--accent-color-dark) var(--content-color);
}
::-webkit-scrollbar {
    width: 6px;
    height: 6px;
}
/* Track */
::-webkit-scrollbar-track {
    background: var(--content-color-dark);
    -webkit-border-radius: 10px;
    border-radius: 6px;
}
 
/* Handle */
::-webkit-scrollbar-thumb {
    -webkit-border-radius: 6px;
    border-radius: 6px;
    /* background:var(--content-color);  */
    background: var(--accent-color-dark);
}
::-webkit-scrollbar-thumb:hover {
    background: var(--accent-color-light); 
}

</style>
