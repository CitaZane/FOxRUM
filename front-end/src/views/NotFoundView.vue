<template>
    <SideFeed v-if="isLoggedIn" :resp="resp"/>
    <MainNotFound v-if="isLoggedIn" :resp="resp"/>
</template>

<script>
import SideFeed from '@/components/sidebar/SideFeed.vue'
import MainNotFound from '@/components/main/MainNotFound.vue'

import { mapState, mapActions} from 'pinia' //helper for store
import {useMainStore} from '@/stores/main.js'

export default {
    name: 'NotFoundView',
    components: {
    SideFeed,
    MainNotFound
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
      if(typeof this.isLoggedIn == 'undefined' || !this.isLoggedIn || this.ws.readyState != 1){
      this.resetStore()
      this.$router.push({name:'login'})
      return
    }
    this.fetchCategories()
  }
    
}
</script>