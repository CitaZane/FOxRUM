<template>
    <SideFeed v-if="isLoggedIn" :resp="resp"/>
    <MainPost v-if="isLoggedIn" :resp="resp"/>
</template>

<script>
import SideFeed from '@/components/sidebar/SideFeed.vue'
import MainPost from '@/components/main/MainPost.vue'

import {mapState, mapActions} from 'pinia'
import {useMainStore} from '@/stores/main.js'

export default {
    components: {
        SideFeed,
        MainPost
    },
    props:["resp"],
    computed:{
        postId(){
            return parseInt(this.$route.params.id)
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
        },
        fetchPost(){
            this.ws.send(JSON.stringify({
                    action: 'get-post',
                    message: this.postId.toString(),
            }))
        },
        fetchComments(offset){
            this.ws.send(JSON.stringify({
                    action: 'get-comments',
                    message: offset,
                    target: this.postId.toString(),
            }))
        },
    },
     created(){
         if(typeof this.isLoggedIn == 'undefined' ||!this.isLoggedIn){
            this.resetStore()
            this.$router.push({name:'login'})
            return
        }
        this.fetchCategories()
        this.fetchPost()
        this.fetchComments("0")
     }

}
</script>