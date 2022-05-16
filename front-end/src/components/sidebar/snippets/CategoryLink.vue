<template>
    <router-link :class="btnClass(category)" :to="linkURL(category)">
    <p :class="titleClass(category)">{{category.name}}</p>
    </router-link>
</template>

<script>
export default{
    props: ["category"],
    methods:{
        linkURL(category){
            return "/category/"+category.name;
        },
        btnClass(category){
            return category.isActive ? 'category-link-btn-active': 'category-link-btn';
        },
        titleClass(category){
            return category.isActive ? 'category-link-title-active': 'category-link-title';
        }
    }
}
</script>

<style scoped>
.category-link-btn{
    background-color: var(--content-color);
    width: 100%;
    display:flex;
    justify-content: space-between;
    text-align: left;
    
    border-radius: 0.5rem;
    text-decoration: none;
    padding: 0.5vh 3vh 0.5vh 3vh;
    margin-bottom: 5px;

}
.category-link-btn-active{
    background-color: var(--main-color);
    width: 100%;
    display:flex;
    justify-content: space-between;
    text-align: left;
    border-radius: 0.5rem;
    text-decoration: none;
    padding: 0.5vh 3vh 0.5vh 3vh;
    margin-bottom: 5px;
}
.category-link-title{
   color: var(--main-color);
    font-size: var(--font-m); 
}
.category-link-title-active{
   color: var(--accent-text-color);
    font-size: var(--font-m);
    font-weight: 500; 
}
.category-link-btn:hover{
    background-color: var(--content-color-active);
}

</style>