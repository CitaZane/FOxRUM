<template>
    <router-link class="nav-link" :to="linkURL(link)" :title="link.name">
    <span :class="[defaultClass, linkClass(link), notification]" :alt="link.name" :id="linkId(link)"></span>
    </router-link>
</template>

<script>
export default{
    props: ["link"],
    data(){
        return{
            defaultClass: 'nav-icon'
        }
    },
    computed:{
        notification(){
            return this.link.notificationActive ? "notify": "";
        }
    },
    methods:{
        linkClass(link){
            return "icon-"+link.name;
        },
        linkURL(link){
            if(link.name == "feed") return "/"
            return "/"+link.name;
        },
        linkId(link){
            return "nav-link-"+ link.name;
        }
    }
}
</script>

<style>
    .nav-link{
    display:block;
    margin-left: 25%;
    padding: 3vh 0 3vh 0;
    width: 50%;
    }
    .nav-icon{
      display: block;
      width: 100%;
      aspect-ratio: 1/1;
    }
    .icon-profile{
        background-image: url("@/assets/icons/user.svg");
        background-repeat: no-repeat;
        filter:invert(49%) sepia(14%) saturate(5855%) hue-rotate(352deg) brightness(90%) contrast(96%);
    }
    .icon-feed{
        background-image: url("@/assets/icons/feed.svg");
        background-repeat: no-repeat;
        filter:invert(49%) sepia(14%) saturate(5855%) hue-rotate(352deg) brightness(90%) contrast(96%);
    }
    .icon-logout{
        background-image: url("@/assets/icons/logout.svg");
        background-repeat: no-repeat;
        filter:invert(49%) sepia(14%) saturate(5855%) hue-rotate(352deg) brightness(90%) contrast(96%);
    }
    .icon-chat{
        position:relative;
        background-image: url("@/assets/icons/chat.svg");
        background-repeat: no-repeat;
        filter:invert(49%) sepia(14%) saturate(5855%) hue-rotate(352deg) brightness(90%) contrast(96%);
    }
    .icon-login{
        background-image: url("@/assets/icons/login.svg");
        background-repeat: no-repeat;
        filter:invert(49%) sepia(14%) saturate(5855%) hue-rotate(352deg) brightness(90%) contrast(96%);
    }
    .notify::after{
        content: "";
        position: absolute;
        top: -0.25vw;
        right: -0.25vw;
        height: 1vw;
        aspect-ratio: 1/1;
        border-radius: 100%;
        background-color:#e33d14;
    }
</style>