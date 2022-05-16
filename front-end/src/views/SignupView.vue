<template>
    <SideLogin :path="path"/>
    <MainSignup :resp="resp"/>
</template>

<script>
// @ is an alias to /src
import SideLogin from '@/components/sidebar/SideLogin.vue'
import MainSignup from '@/components/main/MainSignup.vue'

import {mapState} from 'pinia'
import {useMainStore} from '@/stores/main.js'

export default {
  name: 'SignupView',
  components: {
    SideLogin, 
    MainSignup
  },
  props:["resp"],
  computed:{
    ...mapState(useMainStore,["isLoggedIn"]),
  },
  data: function(){
    return {
      path: "signup"
    }
  },
    created(){
     if(this.isLoggedIn){
      this.$router.push({name:'feed'})
    }
  }
}
</script>