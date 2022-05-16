<template>
    <router-link :class="linkClass" class="sc-link" @click="joinChat()" :to="linkURL()">
        <div  :class="iconClass" class="active-icon" ></div>
        <div :class="titleClass()" class="sc-title">{{user.name}}</div>
    </router-link>
</template>


<script>  
import { mapState} from 'pinia'
import {useMainStore} from '@/stores/main.js'

    export default{
        props:["user", "chatRoom", "resp"],
        data(){
            return{
            }
        },
        computed:{
            ...mapState(useMainStore,["ws"]), //holds ws conn
            ...mapState(useMainStore,["name"]), // active user name
            linkClass(){ 
                if (this.user.name == this.chatRoom)return 'user-link-active'
                if (this.resp.activeChats[this.findChatRoomName()]?.unread)return 'user-link-notification'
                return "user-link"
            },
            iconClass(){
                return (!this.user.online)? "offline": ""
            }
        },
        methods:{
            titleClass(){
                return this.user.name == this.chatRoom ? 'user-link-title-active': 'user-link-title';
            },
             linkURL(){
                return "/chat/"+this.user.name;
            },
            joinChat(){
                 this.ws.send(JSON.stringify({
                    action: 'join-chat',
                    target: this.user.name
                }))
            },
            findChatRoomName(){
                let connection = ""
                if (this.user.name<this.name){
                    connection+=this.user.name+":"+this.name
                }else{
                    connection+=this.name+":"+this.user.name
                }
                return connection
            }
        }
    }
</script>

<style scoped>
/* general link class */
    .sc-link{
        width: 100%;
        border-radius: 0.5rem;
        padding: 0.5vh 1vw 0.5vh 1vw ;
        margin-bottom: 5px;
        display: flex;
        justify-content: flex-start;
        text-decoration: none;
        text-align: left;
        gap: 1vw;
        align-items: center;
    }
    .user-link{
        background-color: var(--content-color);
    }
    .user-link-notification{
        background-color: var(--accent-color-notification);  
    }
    .user-link-active{
        background-color: var(--main-color);
    }
    .active-icon{
        background: url("@/assets/logo.svg");
        background-size: contain;
        background-repeat: no-repeat;
        height: clamp(1rem, 2vh, 1.25rem);
        aspect-ratio: 1/1.2;
    }
    .offline{
       filter: grayscale(100%);
    }
    /* general side chat title class */
    .sc-title{
        font-size: var(--font-m);
        word-wrap: break-word;
        width: 10vw;
        cursor: pointer; 
    }
    @media (max-width: 900px){
    .sc-title{
        width: 26vw;
    }
}
    .user-link-title{
    color: var(--main-color);
    }
    .user-link-title-active{
    color: var(--accent-text-color);
    font-weight: 500; 
    }
    .user-link:hover{
    background-color: var(--content-color-active);
    }
    .user-link-notification:hover{
    background-color: var(--accent-color-dark);
    }
    
</style>