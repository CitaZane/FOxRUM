<template>
    <main class="main-wrapper">
        <div class="main-container">

          <div class="main-header">
            <h2 class="main-title">{{ chatRoomName}}</h2>
            <hr class="divider">
          </div>

          <div  class="main-body">

            <div  v-if="chatRoom" class="chat-container">
              <div v-if="chatListEmpty" class="emptyList">Hmm... There are no messages yet.</div>
                <typing-in-progress-box v-if="typingInAction" :resp="resp"></typing-in-progress-box>
                <chat-box v-for="message in chatRoomMessages()" :key="message.time" :message="message"></chat-box>
            </div>

            <div v-else :class="foxClass" class="chat-container" >
            </div>

            <div :class="inputClass" class="chat-input">
              <input v-model="message.content" name="" class="input-field" placeholder="Don't you got nothing to say?" @keyup.enter.exact="sendMessage()" >

              <div class="send-group">
                <span class="btn " @click="sendMessage">Send</span>
              </div>
            </div>

          </div>

        <!-- <div v-else class="main-body">
          <div :class="foxClass" class="chat-container" >
          </div>
        </div> -->
        
        </div>
    </main>
</template>

<script>
import _ from 'lodash';
import ChatBox from "@/components/main/chat/ChatBox.vue"
import { mapState} from 'pinia' //helper for store
import {useMainStore} from '@/stores/main.js'
import TypingInProgressBox from './chat/TypingInProgressBox.vue';

export default {
  components:{ChatBox , TypingInProgressBox},
  props:["resp", "chatRoom"],
  data:function(){
    return{
      fetchInAction: false,
      message:{
        sender:"",
        content:"",
        reciver:""
      },
      typingInProgress: false,
      typeCallCounter : 0,
    }
  },
  computed:{
    ...mapState(useMainStore,["ws"]),
    ...mapState(useMainStore,["name"]),
    ...mapState(useMainStore,["darkMode"]),
    inputClass(){
      return (!this.chatRoom)? "hidden-input": "visible-input"
    }, 
    chatRoomName(){
      return this.chatRoom || "Don't be quite, find a fox to talk to" // returns first thruthy
    },
    foxClass(){
            return (this.darkMode)? "fox-drawing-dark": "fox-drawing"
    },
    chatListEmpty(){
      return (!this.chatRoomMessages())? true: false;
    },
    typingInAction(){
      let connection = this.findChatRoomName()
      return this.resp.activeChats[connection]?.typingInAction
      // return true
    }
  },
  methods:{
    /* ------------- return messages from all chat msg collection ------------- */
    chatRoomMessages(){
      let connection = this.findChatRoomName()
      return this.resp.activeChats[connection]?.messages
    },
    /* --------------------- send new message to server side -------------------- */
    sendMessage(){
      if(this.message.content === "") return
      this.typingInProgress = true
        this.stopTyping()
        this.message.sender = this.name
        this.message.reciver = this.chatRoom
            this.ws.send(JSON.stringify({
            action: 'new-chat-message',
            message: JSON.stringify(this.message),
        }));
        this.message.content = "";

    },
    /* -------- construct chat room name based on sender and reciver name ------- */
    findChatRoomName(){
      let connection = ""
      if (this.chatRoom<this.name){
        connection+=this.chatRoom+":"+this.name
      }else{
        connection+=this.name+":"+this.chatRoom
      }
      return connection
    },
    handleScroll(){
        let chatBox = document.querySelector(".chat-container");
        if(chatBox.childElementCount ==0)return;
        let position = chatBox.scrollTop* -1 + chatBox.clientHeight
        // give 100 px offset as fetching border
        if(position > chatBox.scrollHeight-100 && !this.fetchInAction){
          this.fetchMessages()
        }
    },
    fetchMessages(){
      let targetChat = this.$route.params.id || "";
      if(targetChat == "")return;
      this.fetchInAction = true
      let offset = this.chatRoomMessages()?.length || 0;
      if(offset == 0){
        this.ws.send(JSON.stringify({
          action: 'join-chat',
          target: targetChat,
      }))
      }else{
        this.ws.send(JSON.stringify({
          action: 'chat-message-request',
          target: targetChat, 
          message: offset.toString(),
      }))
      }  
    },
    handleInput(){
      if (!this.typingInProgress){
        this.typingInProgress = true
        this.startTyping()      
      }else{
        this.typingInProgress = true
      }
    },
    handleKeyUp(){
      console.log("Stop")
      this.typingInProgress = false
      this.stopTyping()
    },
    startTyping(){
       this.ws.send(JSON.stringify({
          action: 'type',
          message: "start", 
          target: this.chatRoom
      }))
    },
    stopTyping(){
        this.ws.send(JSON.stringify({
          action: 'type',
          message: "stop", 
          target: this.chatRoom
      }))
    }
  },
  updated(){
    this.fetchInAction = false
    let chatBox = document.querySelector(".chat-container");
    chatBox.addEventListener("scroll", _.debounce(this.handleScroll, 150))
  },
  beforeMount(){
    this.fetchMessages()
  },
  mounted(){
    let inputBox = document.querySelector(".input-field");
    inputBox.addEventListener("keydown",  _.debounce(this.handleInput, 10, {'leading': true, 'trailing': false}));
     inputBox.addEventListener("keyup",  _.debounce(this.handleKeyUp, 500))
    // inputBox.addEventListener("keyup", _.debounce(this.handleInput, 200, {'leading': true, 'trailing': true}))
  },
}
</script>

<style scoped>

/* 1100 */
.main-body{
  height: 78vh;
  min-height: 350px;
}
.main-header{
    height: 8vh;
    min-height: 36px;
}
@media (max-width: 1100px){
  .main-title{
    margin-top: 1vh;
    margin-bottom: 4vh;
}
  .main-body{
    margin-top: 2vh;
    height: 76vh;
    min-height: 342px;
  }
}
@media (max-width: 900px){
  .main-title{
    margin-top:0;
    margin-bottom:0;
  }
}
@media (max-width: 728px){
  .main-title{
    margin-top: 1vh;
    margin-bottom: 4vh;
  }
}
.chat-container{
  height: 90%;
  overflow-y: auto;
  display:flex;
  flex-direction: column-reverse;
}
.chat-input{
  display: flex;
  align-items: baseline;

}
.hidden-input{
  visibility: hidden;
}
.visible-input{
  visibility: visible;
}
.send-group{
  margin-left:1vw;
   height: 10%;
   margin-top:1vw;
}
.fox-drawing{
   background-image: url("@/assets/pointing-fox.png");
}
.fox-drawing-dark{
   background-image: url("@/assets/pointing-fox_dark.png");
}
.fox-drawing, .fox-drawing-dark{
  height: 100%;
    background-repeat: no-repeat;
    background-size: contain;
    
}
@media (max-width: 900px){
  .fox-drawing, .fox-drawing-dark{
    background-size: cover;
}
}

</style>
