<template>
    <div class="side-wrapper">
        <div class="side-container">
            <new-post-btn></new-post-btn>
            <hr class="side-horizontal-divider">
            <div class="side-category-list">
            <category-link v-for="category in categories" :key="category.id" :category="category"></category-link>
            </div>
        </div>
    </div>
</template>

<script>
import NewPostBtn from '@/components/sidebar/snippets/NewPostBtn.vue'
import CategoryLink from '@/components/sidebar/snippets/CategoryLink.vue'



export default {
    name: 'SideBar',
    components: { NewPostBtn, CategoryLink },
    props:["categoryName", "resp"],
    computed:{
        categories(){
            return this.filterCategories()
        }
    },
    methods:{
        filterCategories(){
            let categories = [{name: "All howls", id: 1, isActive:false}]
            let id = 2
            let linkActivated = false
            this.resp.categories.forEach(cat => {
                let category = {name: cat, id: id, isActive : false}
                if (this.$route.name == 'category' && this.categoryName == cat || this.$route.name == 'post' && this.resp.post.category == cat) {
                    category.isActive = true;
                    linkActivated = true
                }
                categories.push(category)
                id++
            });
            if (!linkActivated) categories[0].isActive=true
            return categories
        }
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
.side-content{
    background-color: rgba(137, 43, 226, 0.233);
    width: 100%;
    height: 100%;
}
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
.side-category-list{
    height: 85%;
    overflow-y: auto;
}

</style>