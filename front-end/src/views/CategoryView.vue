<template>
    <SideFeed v-if="isLoggedIn" :categoryName="categoryName" :resp="resp"/>
    <MainFeed v-if="isLoggedIn" :categoryName="categoryName" :resp="resp"/>
</template>

<script>
import SideFeed from '@/components/sidebar/SideFeed.vue'
import MainFeed from '@/components/main/MainFeed.vue'

import {mapState, mapActions} from 'pinia'
import {useMainStore} from '@/stores/main.js'

export default {
    components: {
        SideFeed,
        MainFeed
    },
    props:["resp"],
    computed:{
        categoryName(){
            return this.$route.params.name
        },
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
    if(typeof this.isLoggedIn == 'undefined' || !this.isLoggedIn){
       this.resetStore()
      this.$router.push({name:'login'})
      return
    }
    this.fetchCategories()
  }
}
</script>