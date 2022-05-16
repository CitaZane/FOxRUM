<template>
    <main class="main-wrapper">
        <div class="main-container">
        <h2 class="main-title">{{ feedTitle }}</h2>
        <hr class="divider">
        <post-list :posts="resp.posts"></post-list>
        </div>
    </main>
</template>

<script>
import _ from 'lodash';
import {mapState} from 'pinia'
import {useMainStore} from '@/stores/main.js'
import PostList from "@/components/main/feed/PostList.vue"

export default {
  name: 'MainContent',
  components:{PostList},
  props:["categoryName", "resp"],
  computed:{
    ...mapState(useMainStore,["ws"]),
    feedTitle(){
      return (this.categoryName)? this.categoryName: "What does the fox say?"
    }
  },
  methods:{
    fetchPosts(offset="0"){
    this.ws.send(JSON.stringify({
          action: 'get-posts',
          message: offset, 
      }))
    },
    fetchCategoryPosts(offset="0"){
    this.ws.send(JSON.stringify({
          action: 'get-category-posts',
          message: offset, 
          target: this.categoryName
      }))
      return
    },
    handleScroll(){
      let postBox = document.querySelector(".post-list");
      if (postBox.childElementCount == 0)return;
      // no scroll present
      if (postBox.clientHeight == postBox.scrollHeight) return 
      // console.log("Scroll:  ", postBox.scrollTop, postBox.clientHeight, postBox.scrollHeight)
      let position = postBox.scrollTop + postBox.clientHeight
      // give 100 px offset as fetching border
      if (position > postBox.scrollHeight-100){ 
        let offset = this.resp.posts?.length || 0;
        if (!this.categoryName || this.categoryName == "All howls"){
          this.fetchPosts(offset.toString())
        }else{
          this.fetchCategoryPosts(offset.toString())
        }
      }
      let topBtn = document.querySelector('#topBtn');
      if(position> 1000){
        topBtn.style.display = "block"
      }else{
        topBtn.style.display = "none"
      }
    },
    scrollIntoView(){
      let lastMessage = document.querySelector(".post-list").firstElementChild;
      if(!lastMessage)return
      lastMessage.scrollIntoView({behavior: "smooth"});
    }
  },
  beforeMount(){
    if (!this.categoryName || this.categoryName == "All howls"){
      if (this.ws.readyState == 1){
        this.fetchPosts()
      }
    }else{
      if (this.ws.readyState == 1){
        this.fetchCategoryPosts()
      }
    }
    
  },
  mounted(){
    let postBox = document.querySelector(".post-list");
    postBox.addEventListener("scroll", _.debounce(this.handleScroll, 150))
  },
  updated(){
    this.scrollIntoView()
  },
   watch:{
    categoryName(category){
      if (category == "All howls"){
      this.fetchPosts("0")
    }else{
      this.fetchCategoryPosts("0")
    }
    }
  }
}
</script>


<style scoped>
.main-title{
    font-size: var(--font-l);
    font-weight: bold;
    color: var(--main-color);
}

</style>