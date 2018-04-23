<template>
  <div class="scroll-wrapper">
    <!-- <li is="todo-item" v-for="(item, index) in items" :item="item" @remove="remove"></li> -->
    <div v-for="(item, key) in items" :key="key">
        <div :data-anchor="item.anchor" :class="['scroll-anchor', {'scroll-active': item.anchor == anchor}]" @click="emitHandle">
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

            bus.$on('locator-active', e => {
                if (e) {
                    this.anchor = e;
                }
            });
        },
        data () {
            return {
                items: [],
                anchor: ''
            }
        },
        methods: {
            emitHandle: function (e){
                let archor = e.target.dataset.anchor || '';
                bus.$emit('locatorChange', archor);
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
.scroll-active{
    color: #333;
    background-color: #fff;
}
</style>
