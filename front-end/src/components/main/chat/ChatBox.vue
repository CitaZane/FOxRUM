<template>
    <div v-if="name != message.sender" class="chat-box-container">
        <div class="chat-box-icon"></div>
        <div class="chat-box-content">
            <div class="chat-message-time">{{message.sender}} at: {{messageTime()}}</div>
            <div class="chat-message-content">{{message.content}}</div>
        </div>
    </div>
    <div v-else class="chat-box-container-author">  
        <div class="chat-box-content-author">
            <div class="chat-message-time">{{message.sender}} at: {{messageTime()}}</div>
            <div class="chat-message-content">{{message.content}}</div>
        </div>
        <div class="chat-box-icon-author"></div>
    </div>
</template>

<script>
import { mapState} from 'pinia' //helper for store
import {useMainStore} from '@/stores/main.js'
export default {
    props:["message"],
    computed:{
        ...mapState(useMainStore,["name"]),
        notAuthor(){
            return this.name != this.message.sender
        }
    },
    methods:{
        messageTime(){
            const date = new Date(Number(this.message.time)*1000)
            const today = new Date()
            let minutes= (Number(date.getMinutes())<10)? "0"+date.getMinutes(): date.getMinutes();
            return (today.getDate() === date.getDate())? date.getHours()+":"+minutes: date.getDate()+"/"+date.getMonth()+"/" + date.getFullYear()
        }
    }
}
</script>

<style scoped>
    .chat-box-container{ 
        width: 100%;
        display:flex;
        flex-direction: row;
        flex-wrap: nowrap;
        justify-content: flex-start;
        margin-bottom: 1.5vh;
    }
    .chat-box-icon{
        height: clamp(30px, 10vh, 50px);
        aspect-ratio: 1/1;
        background-image: url("@/assets/logo.svg");
        background-repeat: no-repeat;
        background-size: contain;
        background-position: center center
    }
    .chat-box-content{
        max-width:50vw;
        margin-left: 1.5vw;
        background-color: var(--content-color);
        border-radius: 0.75rem;
        padding: 0.2rem 1rem 0.2rem 1rem;
        text-align: left;
        overflow-wrap: break-word;
        position: relative;
        
    }
    .chat-box-content::before{
        position: absolute;
        background-color: var(--content-color);
        left: -7px;
        top: 19px;
        display: block;
        width: 12px;
        height: 10px;
        content: " ";
        transform: rotate(29deg) skew(-35deg);
    }
    .chat-message-time{
        color: var(--content-color-dark);
        font-size: var(--font-xs);
    }
    .chat-message-content{
        color: var(--main-color);
        font-size: var(--font-s);
    }

    /* Author */
        .chat-box-container-author{ 
        width: 100%;
        display:flex;
        flex-direction: row;
        flex-wrap: nowrap;
        justify-content: flex-end;
        margin-bottom: 1.5vh;
    }
    .chat-box-icon-author{
        height: clamp(30px, 10vh, 50px);
        aspect-ratio: 1/1;
        background-image: url("@/assets/logo.svg");
        background-repeat: no-repeat;
        background-size: contain;
        background-position: center center;
        filter:invert(49%) sepia(14%) saturate(5855%) hue-rotate(352deg) brightness(90%) contrast(96%);
    }
    .chat-box-content-author{
        max-width:50vw;
        margin-right: 1.5vw;
        background-color: var(--content-color);
        border-radius: 0.75rem;
        padding: 0.2rem 1rem 0.2rem 1rem;
        text-align: left;
        overflow-wrap: break-word;
        position: relative;
        
    }
    .chat-box-content-author::before{
        position: absolute;
        background-color: var(--content-color);
        right: -7px;
        top: 19px;
        display: block;
        width: 12px;
        height: 10px;
        content: " ";
        transform: rotate(29deg) skew(-35deg);
    }

</style>