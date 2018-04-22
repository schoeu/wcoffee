<template>
  <div class="scroll-wrapper goods-list">
    <!-- <li is="todo-item" v-for="(item, index) in items" :item="item" @remove="remove"></li> -->
    <div v-for="(item, key) in list" :key="key">
        <div class="goods-title" :id="item.anchor">{{item.typename}}</div>
        <div v-for="(i, k) in item.list" :key="k">
            <article class="goods-article" :id="i.id">
                <div class="flexbox" on="">
                    <div class="goods-img">
                        <img class="list-img" :src="i.img"/>
                    </div>
                    <section class="goods-info">
                    <div class="flexbox">
                        <div class="goods-titles">
                            <div class="good-stt">{{i.name}}</div>
                            <div class="goods-alias">{{i.alias}}</div>
                        </div>
                        <div class="goods-price">
                        <span class="goods-yentag">¥</span>
                        {{i.price / 100}}
                        </div>
                    </div>
                    <div class="goods-discribe clamp2">{{i.desc}}</div>
                    </section>
                </div>
            </article>
        </div>
    </div>
  </div>
</template>

<script>
    import * as apis from '../api';
    import * as scroll from '../../static/common/js/scroll';

    console.log(scroll, 11);
    export default {
        mounted () {
            // get list data from server.
            apis.getList().then(res => {
                this.list = res || [];
            });

            this.$el.parentElement.addEventListener('scroll', this.throttle(this.scrollHandle, 200))
        },
        data () {
            return {
                list: []
            }
        },
        methods: {
            scrollHandle: () => {
                console.log(1);
            },
            /**
             * Throttle a function.
             * @param {Function} fn
             * @param {number} delay The run time interval
             * @return {Function}
             */
            throttle: (fn, delay) => {
                var context, args, timerId;
                var execTime = 0;
                !delay && (delay = 10);
                function exec() {
                    timerId = 0;
                    execTime = Date.now();
                    fn.apply(context, args);
                };
                return function () {
                    var delta = Date.now() - execTime;
                    context = this;
                    args = arguments;
                    clearTimeout(timerId);
                    if (delta >= delay) {
                        exec();
                    } else {
                        timerId = setTimeout(exec, delay - delta);
                    }
                }
            }
        }
    }
</script>
<style>
.goods-list{
    padding-left: 2.666667vw;
}
.goods-title{
    border-bottom: 1px solid #eee;
    font-size: 13px;
    color: #666;
    padding: 2vw 8vw 2vw 0;
    font-weight: 700;
}
.goods-article{
    padding: 2.666667vw 2.666667vw 2.666667vw 0;
}
.goods-info{
    flex: 1;
}
.goods-img{
    width: 20.266667vw;
    height: 20.266667vw;
    vertical-align: top;
    -webkit-box-flex: 0;
    -webkit-flex: none;
    flex: none;
    margin-right: 2.666667vw;
}
.list-img {
    width: 100%;
    height: 100%;
}
.goods-price{
    flex: 1;
    color: #f60;
    font-size: 17px;
    text-align: right;
}
.goods-titles{
    flex: 3;
}
.good-stt{
    font-size: 16px;
}
.goods-alias {
    font-size: 10px;
    margin: 1vw 0;
    color: #999;
}
.goods-yentag{
    color: #f60;
    font-size: 13px;
    margin-right: .533333vw;
}
.goods-discribe{
    font-size: 13px;
    color: #666;
}
</style>
