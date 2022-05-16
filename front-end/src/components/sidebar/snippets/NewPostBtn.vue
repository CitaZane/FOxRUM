<template>
    <router-link class="side-post-btn" to="/newpost" title="Create new Post">
    <p class="side-post-btn-title">Howl</p>
    <p class="side-post-btn-plus">+</p>
    </router-link>
</template>

<script>
export default {
}
</script>

<style scoped>
.side-post-btn{
    background-color: var(--accent-color-dark);
    width: 100%;
    height: 28px;

    display:flex;
    justify-content: space-between;

    border-radius: 0.5rem;
    text-decoration: none;
    padding: 0.5vh 3vh 0.5vh 3vh;
}
.side-post-btn:hover{
    background-color: var(--accent-color-light);
}
.side-post-btn-title{
    color: var(--accent-text-color);
    font-size: var(--font-m);
    font-weight: 500;
}
.side-post-btn-plus{
    color: var(--accent-text-color);
    font-size: var(--font-m);
    font-weight: 500;
}
</style>