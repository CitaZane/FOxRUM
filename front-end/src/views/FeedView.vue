<template>
    <SideFeed v-if="isLoggedIn" :resp="resp"/>
    <MainFeed v-if="isLoggedIn" :resp="resp"/>
</template>

<script>
// @ is an alias to /src
import SideFeed from '@/components/sidebar/SideFeed.vue'
import MainFeed from '@/components/main/MainFeed.vue'

import {mapState, mapActions} from 'pinia'
import {useMainStore} from '@/stores/main.js'

export default {
  name: 'FeedView',
  components: {
    SideFeed,
    MainFeed
  },
  props:["resp"],
  computed:{
    ...mapState(useMainStore,["isLoggedIn"]),
    ...mapState(useMainStore,["ws"])
  },
  methods:{
    ...mapActions(useMainStore, {resetStore : 'reset'}),
    fetchCategories(){
      this.ws.send(JSON.stringify({
              action: 'get-categories'
      }))
    },
    fetchPosts(){
    this.ws.send(JSON.stringify({
          action: 'get-posts',
          message: "0"
      }))
    },
     fetch(){
       this.fetchCategories()
        this.fetchPosts()
     }
  },
  created(){
    if(typeof this.isLoggedIn == 'undefined' || !this.isLoggedIn){
      this.resetStore()
      this.$router.push({name:'login'})
      return
    }
     if (this.ws.readyState != 1){
        this.ws.onopen = this.fetch
      }else{
        this.fetch()
      }
  }
}
</script>