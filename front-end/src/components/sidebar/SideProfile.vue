<template>
    <div class="side-wrapper">
        <div class="side-container">
            <h3 class="side-profile-username">{{ name }}</h3>
            <hr class="side-horizontal-divider">
            <div class="side-profile-subcontainer">

                <div class="side-profile-stats">
                    <p class="profile-stats-info">Howling activity level</p>
                    <p class="profile-stats-number">{{resp.userStats.posts}}</p>
                    <p class="profile-stats-info">Barking activity level</p>
                    <p class="profile-stats-number">{{resp.userStats.comments}}</p>
                    <p class="profile-stats-info">Joined the forest in</p>
                    <p class="profile-stats-number">{{resp.userStats.joined}}</p>
                </div>

                <!-- <div class="side-profile-leave-btn">Leave the forest forever</div> -->
            </div>
        </div>
    </div>
</template>

<script>
import {mapState} from 'pinia'
import {useMainStore} from '@/stores/main.js'

export default {
    props:["resp"],
    computed:{
        ...mapState(useMainStore,["name"]),
        ...mapState(useMainStore,["ws"])
    },
    methods:{
       getUserStats(){
           this.ws.send(JSON.stringify({
            action: 'user-stats'
      }))
       } 
    },
    mounted(){
        this.getUserStats()
    }
}
</script>


<style scoped>
.side-container{
    height: 90vh;
    margin-top: 5vh;
    padding: 0vh 5vh 0vh 5vh;
    overflow: hidden;
}
.side-profile-username{
    color:var(--accent-color-dark);
    font-size: var(--font-l)
}
.side-profile-subcontainer{
    height:76vh;
    position: relative;
}
.profile-stats-info{
    font-size:var(--font-m);
    display:inline;
}
.profile-stats-number{
    font-size:var(--font-m);
    font-weight: bold;
    color: var(--accent-color-dark);
    margin-bottom: 2vh;
}
/* .side-profile-leave-btn{
    position: absolute;
    bottom: 0;
    background-color: var(--main-color);
    color:var(--accent-text-color);
    font-size: var(--font-m);
    border-radius: 0.5rem;
    padding: 0.5vh 3vh 0.5vh 3vh;
    cursor: pointer;

}
.side-profile-leave-btn:hover{
    background-color: var(--main-color-active);
} */

.side-horizontal-divider{
    width:100%;
    height: 0.5px;
    border: 0 none;
    background-color: var(--content-color-dark);
    margin-top:4vh;
    margin-bottom: 5vh;
    margin-left: auto;
    margin-right: auto;
}
</style>