<template>
  <div class="scroll-wrapper">
    <!-- <li is="todo-item" v-for="(item, index) in items" :item="item" @remove="remove"></li> -->
    <div v-for="(item, key) in items" :key="key">
        <div :data-anchor="item.anchor" class="scroll-anchor">
            {{item.name}}
        </div>
    </div>
  </div>
</template>

<script>
    import * as apis from '../api';
    export default {
        mounted () {
            apis.getTags().then(res => {
                this.items = res.items || [];
            });
        },
        data () {
            return {
                items: []
            }
        }
    }
</script>
<style>
.scroll-wrapper {
    -webkit-tap-highlight-color: transparent;
}
.scroll-anchor {
    padding: 4.666667vw 2vw;
    border-bottom: 1px solid #e8e8e8;
    background-color: #f8f8f8;
}
.scroll-anchor.active {
    color: #333;
    background-color: #fff;
}
</style>
