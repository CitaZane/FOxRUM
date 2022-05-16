<template>
    <main class="main-wrapper">
        <div class="post-container main-container">
            <div class="main-header">
                <div class="post-stats">
                <p class="post-author">Fox: {{resp.post.author}}</p>
                <p>Howled at:  {{postTime()}}</p>
                </div>
                <h2 class="post-title">{{resp.post.title}}</h2>
                <div class="post-content">{{resp.post.content}}</div>
                <hr class="divider">
            </div>

            <div class="main-body">

                 <div class="post-commenting-container">
                    <input v-model="comment" name="comment" placeholder="Got nothing to say?" class="input-field" @keyup.enter.exact="postComment()">
                    <div @mousedown="postComment()" id="post-new-comment-btn">Bark back</div>
                </div>

                <div class="post-comment-list">
                    <div v-if="commentListEmpty" class="emptyList">Hmm... No one barked back yet. You better start barking!</div>
                <single-comment v-for="comment in resp.comments" :key="comment.id" :comment="comment"></single-comment>
                 <button @click="backToTop" id="topBtn" title="Go to top"></button>
                </div>

            </div>
            
        </div>
    </main>
</template>

<script>
import _ from 'lodash';
import SingleComment from "@/components/main/feed/SingleComment.vue";
import { mapState} from 'pinia' //helper for store
import {useMainStore} from '@/stores/main.js'

    export default{
        components:{SingleComment},
        props:["resp"],
        computed:{
            ...mapState(useMainStore,["ws"]),
            postId(){
                return parseInt(this.$route.params.id)
            },
            commentListEmpty(){
                return (this.resp.comments.length == 0)? true: false;
            }
        },
        data: function(){
            return{
                comment:""
            }  
        },
        methods:{
            postTime(){
                if (!this.resp.post.time)return
                const date = new Date(Number(this.resp.post.time) *1000)
                let minutes= (Number(date.getMinutes())<10)? "0"+date.getMinutes(): date.getMinutes();
                return date.getHours()+":"+minutes+"  "+ date.getDate()+"/"+date.getMonth()+"/" + date.getFullYear()
            },
            postComment(){
                if(this.comment){
                    this.ws.send(JSON.stringify({
                        action: 'new-comment',
                        message: this.comment,
                        target: this.resp.post.id.toString(),
                    }))
                    this.comment = "";
                }
            },
            fetchComments(offset){
            this.ws.send(JSON.stringify({
                    action: 'get-comments',
                    message: offset,
                    target: this.postId.toString(),
            }))
        },
            backToTop(){
            let lastMessage = document.querySelector(".post-comment-list").firstElementChild;
            lastMessage.scrollIntoView({behavior: "smooth"});
            },
            handleScroll(){
                let commentBox = document.querySelector(".post-comment-list");
                if (commentBox.childElementCount == 0)return;
                // // no scroll present
                if (commentBox.clientHeight == commentBox.scrollHeight) return 
                // console.log("Scroll:  ",  commentBox.scrollTop,  commentBox.clientHeight, commentBox.scrollHeight)
                let position = commentBox.scrollTop + commentBox.clientHeight
                // // give 100 px offset as fetching border
                if (position > commentBox.scrollHeight-100){ 
                    let offset = this.resp.comments?.length || 0;
                    this.fetchComments(offset.toString())
                }
                let topBtn = document.querySelector('#topBtn');
                if(position> 1000){
                    topBtn.style.display = "block"
                }else{
                    topBtn.style.display = "none"
                }
                },
        },
          mounted(){
            let commentBox = document.querySelector(".post-comment-list");
            commentBox.addEventListener("scroll", _.debounce(this.handleScroll, 150))
        }
    }
</script>

<style scoped>
.main-body{
 position: relative;
 flex:1;
}
.main-header{
    flex:0;
}
.post-container{
    text-align: left;
    display: flex;
    flex-flow: column nowrap;
    justify-content: flex-start;
}
.post-stats p{
    font-size:var(--font-xs);
    display:inline;
}
.post-author{
    margin-right:5vh;
}
.post-title{
    font-size: var(--font-l);
    font-weight: bold;
    color: var(--accent-color-dark);
}
.post-content{
    font-size: var(--font-m);
}
.post-comment-list{
    position: absolute;
    left: 0; top: calc(var(--font-s) + 2vh + 4vh); right: 0; bottom: 0;
    overflow-y: auto;
}
.emptyList{
    text-align: center;

}
.post-commenting-container{
    display:flex;
    justify-content: space-between;
    flex-wrap: nowrap;
    gap: 1vw;
    margin-bottom: 2vh;
}
.post-commenting-container input{
    background-color: var(--content-color);
    border-radius: 0.5rem;
    border-style: none;
}
#post-new-comment-btn{
    background-color: var(--accent-color-dark);
    color:var(--accent-text-color);
    font-size: var(--font-s); 
    font-weight: bold; 
    text-align: center;
    padding: 0.5vh 3vh 0.5vh 3vh;
    border-radius: 0.5rem;
    white-space: nowrap;
    cursor: pointer;
}
#topBtn{
    display: none;
    position: fixed;
    bottom: 8vh;
    right: 10vw;
    border: none;
    outline: none;
    background-color: var(--main-color);
    cursor: pointer;
    height: 50px;
    aspect-ratio: 1/1;
    border-radius:50%;
    border: 2px solid var(--main-color);
    transition: background-color 0.3s;
}
#topBtn::after{
    content: "";
    position: absolute;
    background-image: url("@/assets/icons/up.svg");
    background-repeat: no-repeat;
    filter:invert(49%) sepia(14%) saturate(5855%) hue-rotate(352deg) brightness(90%) contrast(96%);
    height:40px;
    aspect-ratio: 1/1;
    top:3px;
    left:8px;
}
#topBtn:hover{
    background-color: var(--content-color);
    border: 2px solid var(--main-color);
}

</style>