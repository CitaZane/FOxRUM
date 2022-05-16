<template>
    <div class="type-box-container">
        <div class="type-box-icon"></div>
        <div class="type-box-content">
            <div class="type-message-content">
                <div class="typing-dot dot-1"></div>
                <div class="typing-dot dot-2"></div>     
                <div class="typing-dot dot-3"></div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props:["resp"]
}
</script>

<style scoped>
    .type-box-container{ 
        width: 100%;
        display:flex;
        flex-direction: row;
        flex-wrap: nowrap;
        justify-content: flex-start;
        margin-bottom: 1.5vh;
    }
    .type-box-icon{
        height: clamp(30px, 10vh, 50px);
        aspect-ratio: 1/1;
        background-image: url("@/assets/logo.svg");
        background-repeat: no-repeat;
        background-size: contain;
        background-position: center center
    }
    .type-box-content{
        max-width:50vw;
        margin-left: 1.5vw;
        background-color: var(--content-color);
        border-radius: 0.75rem;
        padding: 0.2rem 1rem 0.2rem 1rem;
        text-align: left;
        overflow-wrap: break-word;
        position: relative;
        
    }
    .type-box-content::before{
        position: absolute;
        background-color: var(--content-color);
        left: -7px;
        top: 19px;
        display: block;
        width: 12px;
        height: 10px;
        content: " ";
        transform: rotate(29deg) skew(-35deg);
    }
    .type-message-content{
        display:flex;
        flex-direction: row;
        flex-wrap: nowrap;
        gap: var(--font-xs);
        margin-top: 18px;
    }
    .typing-dot{
        width:var(--font-s);
        aspect-ratio: 1/1;
        background: var(--main-color);
        border-radius:50%;
        animation-name: typing;
        animation-duration: 1s;
        animation-iteration-count: infinite;
    }
    .dot-2 {
    animation-delay:.25s;
    }

    .dot-3 {
    animation-delay:.5s;
    }
    @keyframes typing {
  0% {
    backround:var(--main-color);
  }
  100%{
    background: var(--main-text-color);
  }
}
</style>