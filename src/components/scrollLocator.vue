<template>
  <div>
      <div>123{{name}}</div>
    <slot></slot>
  </div>
</template>

<script>
    export default {
        data () {
            return {
                name: 'schoeu'
            }
        }
    }
</script>
