<template>
  <div class="dialog-wrapper">
      <transition name="fade">
        <div class="dialog-mask"></div>
      </transition>
      <transition>
        <div class="dialog-content">
            <slot></slot>
        </div>
      </transition>
  </div>
</template>

<script>
    export default {

    };
</script>

<style>
.dialog-wrapper {
    
}
.fade-enter-active, .fade-leave-active {
    transition: opacity .5s;
}
.fade-enter, .fade-leave-to {
    opacity: 0;
}
</style>

