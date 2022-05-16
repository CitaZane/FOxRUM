<template>
    <main class="main-wrapper">
        <div class="main-container">
        <h2 class="main-title">Your Howls</h2>
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
  name: 'MainProfile',
  components:{PostList},
  props:["resp"],
  computed:{
    ...mapState(useMainStore,["ws"]),
  },
  methods:{
    fetchUserPosts(offset){
    this.ws.send(JSON.stringify({
          action: 'get-user-posts',
          message: offset, 
      }))
    },
    handleScroll(){
      let postBox = document.querySelector(".post-list");
      if (postBox.childElementCount == 0)return;
      if (postBox.clientHeight == postBox.scrollHeight) return 
      let position = postBox.scrollTop + postBox.clientHeight
      if (position > postBox.scrollHeight-100){ 
        let offset = this.resp.posts?.length || 0;
          this.fetchUserPosts(offset.toString())
      }
    }
  },
  mounted(){
    let postBox = document.querySelector(".post-list");
    postBox.addEventListener("scroll", _.debounce(this.handleScroll, 150))
  },
}
</script>


<style scoped>

</style>