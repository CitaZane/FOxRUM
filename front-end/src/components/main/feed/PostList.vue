<template>
    <div class="post-list">
        <div v-if="postListEmpty">Hmm... The list is not missing, simply there is nothing there yet.</div>
        <single-post v-for="post in posts" :key="post.id" :post="post"></single-post>
        <button @click="backToTop" id="topBtn" title="Go to top"></button>
    </div>
</template>

<script>
import SinglePost from "@/components/main/feed/SinglePost.vue";

export default {
 components: { SinglePost },
 props: ["posts"],
 computed:{
     postListEmpty(){
         return (!this.posts)? true: false;
     }
 },
 methods:{
     backToTop(){
         let lastMessage = document.querySelector(".post-list").firstElementChild;
         lastMessage.scrollIntoView({behavior: "smooth"});
     }
 }
};
</script> 

<style scoped>
.post-list{
    height: 85%;
    overflow-y: auto;
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