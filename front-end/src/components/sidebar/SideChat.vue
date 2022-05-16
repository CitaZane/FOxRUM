<template>
    <div class="side-wrapper">
        <div class="side-container">
            <h3>Who is in the forest?</h3>
            <hr class=" divider side-divider">
            <div class="side-user-list">
                <chat-user v-for="user in filteredUserList" :key="user.name" :user="user" :chatRoom="chatRoom" :resp="resp"></chat-user>
            </div>
        </div>
    </div>
</template>


<script>
import ChatUser from '@/components/sidebar/snippets/ChatUser.vue'
export default {
    props:["resp", "ws", "chatRoom"],
    components: { ChatUser },
    computed:{
       filteredUserList(){
           /* ---------------------- sort user list alphabetically --------------------- */
           let sortedList = this.resp.userList
            sortedList = sortedList.sort((a,b)=>{
              let fa = a.name.toLowerCase(), fb = b.name.toLowerCase();
                if (fa < fb) return -1 
                if (fa > fb) return 1
                return 0  
            })
            /* --------------------- sort list again by last message time -------------------- */
            sortedList = sortedList.sort((a,b)=>{
              let timeA = new Date(Number(a.lastMsgTime)*1000).getTime(), timeB = new Date(Number(b.lastMsgTime)*1000).getTime();
              if (isNaN(timeA)) timeA = 0; if (isNaN(timeB)) timeB = 0
                return timeB-timeA 
            })
           return sortedList
       } 
    }
}
</script>

<style scoped>
h3{
    font-size: var(--font-l);
}

.side-user-list{
    height: 85%;
    overflow-y: auto;
}

</style>