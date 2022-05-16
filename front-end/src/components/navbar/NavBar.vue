<template>
    <div class="nav-bar">
        <div id="nav-indicator">
          <div id="helper-rec"></div>
          <div id="helper-circle-one"></div>
          <div id="helper-circle-two"></div>
        </div>
        <link-list v-if="isLoggedIn" :links="linksAuth"></link-list>
        <link-list v-else :links="linksNotAuth"></link-list>
        <input type="checkbox" name="modeSwitch" id="mode-switch">  
    </div>
</template>

<script>
import LinkList from '@/components/navbar/snippets/LinkList.vue'
import { mapState , mapWritableState} from 'pinia' //helper for store
import {useMainStore} from '@/stores/main.js'
let resizeTimer

export default {
  components: { LinkList },
  name: 'NavBar',
  data: function(){
    return {
      linksNotAuth:[{name: "login", id: 1}],
      linksAuth:[
        {name: "profile", id: 2}, 
        {name:"feed", id:3, notificationActive:false}, 
        {name:"chat", id:4 , notificationActive:false}]
    }
  },
    computed:{
      ...mapState(useMainStore,["name"]),
    ...mapState(useMainStore,["isLoggedIn"]),
    ...mapWritableState(useMainStore,["darkMode"]),
    ...mapState(useMainStore,["chatNotification"])
  },
  methods:{
    moveIndicator(path){
      if(!this.isLoggedIn || typeof(this.isLoggedIn) == 'undefined') path= "login"
      if(this.isLoggedIn && path == "login") path="feed"
      let target = document.querySelector(`#nav-link-${path}`)
      let yOffset = target.getBoundingClientRect().top
      let indicator = document.querySelector("#nav-indicator");
      let indicatorHeight = indicator.clientHeight;
      let iconHeight = target.clientHeight;
      let position = yOffset - (indicatorHeight - iconHeight) / 2;
      indicator.style.top = position + "px";
    },
    configurePath(path){
      if(path == "category" || path == "post"|| path == "newpost" || path == ""|| path=="feed") return "feed";
      if(path == "login" || path == "signup") return "login";
      if(path == "chat") return "chat";
      if(path == "profile") return "profile";
      if(path == "logout") return "logout";
      return "feed";
    },
    moveIndicatorAfterResize(){
      let pathName = this.configurePath(this.$route.name)
      this.moveIndicator(pathName)
    },
    detectResize(){
        clearTimeout(resizeTimer);
        resizeTimer = setTimeout(this.moveIndicatorAfterResize, 200)
    },
    chatNotificationOn(){
      this.linksAuth[2].notificationActive = true
    },
    chatNotificationOff(){
      this.linksAuth[2].notificationActive = false
    },
    switchMode(target){
        let element = document.querySelector('html');
      if ( target.checked){
        element.classList.add('dark-mode')
        this.darkMode = true
      }else{
        element.classList.remove('dark-mode')
        this.darkMode = false
      }
    }
  },
  mounted(){
    window.addEventListener("resize", this.detectResize)
    document.querySelector('#mode-switch').addEventListener('input',({target})=> this.switchMode(target))
    if (this.darkMode){
        document.querySelector('#mode-switch').checked = true
        document.querySelector('html').classList.add('dark-mode');
    }
  },
  
  watch:{
    $route (to){
      let goingTo = to.path.split("/")[1];
      let toPath  = this.configurePath(goingTo);
      this.moveIndicator(toPath)
    },
    //detect changes in global chat notification
    chatNotification(unreadMessage){
      if(unreadMessage){
        this.chatNotificationOn();
      }else if (!unreadMessage){
        this.chatNotificationOff();
      }
    }
  }
}
</script>

<style scoped>
    .nav-bar{
        background-color: var(--main-color);
        width: 4.5vw;
        min-width: 33.75px;
        height: 100vh;
        min-height: 450px;
        overflow: hidden;
        margin: 0;
        position: relative;
    }
    .nav-bar::after{
      content: "";
      width: 3.5vw;
      min-width: 26.25px;
      aspect-ratio: 1/1.5;
      position: absolute;
      background: url("@/assets/logo.svg");
      background-size: contain;
      background-repeat: no-repeat;
      top:1rem;
      right:0.5vw;
    }
    #nav-indicator{
      position: absolute;
      left:0;
      width:90%;
      aspect-ratio: 1/1;
      margin-left:10%;
      background-color:var(--content-color);
      border-top-left-radius: 50%;
      border-bottom-left-radius: 50%;
      transition: top 0.5s;
      transition-timing-function: ease-in-out;
    }
    #helper-rec{
      width:30%;
      height: 150%;
      background-color:var(--content-color);
      position:absolute;
      right:0;
      top: -25%;
    }
    #helper-circle-one{
      position:absolute;
      height: 60%;
      width:60%;
      background-color:var(--main-color);
      border-radius: 50%;
      right:0;
      top: -60%;
    }
     #helper-circle-two{
      position:absolute;
      height: 60%;
      width:60%;
      background-color:var(--main-color);
      border-radius: 50%;
      right:0;
      bottom: -60%;
    }
    #mode-switch{
      position:absolute;
      bottom: 2vh;
      left: 1vw;
      cursor:pointer;
      appearance: none;
      font-size: var(--font-l);
      font-family: 'Segoe UI Emoji';
      display: flex;
      justify-content: center;
      align-items: center;
    }
    #mode-switch::after{
      content: '';
      background: url("@/assets/sun.png");
      height: 2.5vw;
      min-height: 18.75px;
      aspect-ratio: 1/1;
      background-size: contain;
      background-repeat: no-repeat;
    }
    #mode-switch:checked::after{
      content: '';
      background: url("@/assets/moon.png");
      height: 2.5vw;
      min-height: 18.75px;
      aspect-ratio: 1/1;
      background-size: contain;
      background-repeat: no-repeat;
    }

</style>