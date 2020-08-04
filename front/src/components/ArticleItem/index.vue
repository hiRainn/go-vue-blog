<template>
  <div>
	  asddsa
  </div>
</template>

<script>


export default {
  name: 'ArticleItem',
  props: {
    list: {
      type: Array,
      required: true
    },
  },
  computed: {
    isExternal() {
      return isExternal(this.iconClass)
    },
  }
}
</script>

<style scoped>

</style>
