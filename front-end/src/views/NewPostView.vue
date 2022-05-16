<template>
    <SideFeed v-if="isLoggedIn" :resp="resp"/>
    <MainNewPost v-if="isLoggedIn" :resp="resp"/>
</template>

<script>
import SideFeed from '@/components/sidebar/SideFeed.vue'
import MainNewPost from '@/components/main/MainNewPost.vue'

import {mapState, mapActions} from 'pinia'
import {useMainStore} from '@/stores/main.js'

export default {
  name: 'NewPostView',
  components: {
    SideFeed,
    MainNewPost
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
    }
  },
  created(){
    if(typeof this.isLoggedIn == 'undefined' ||!this.isLoggedIn){
      this.resetStore()
      this.$router.push({name:'login'})
      return
    }
    this.fetchCategories()
  }
}
</script>