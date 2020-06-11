<template>
  <div class="card" v-html="card.content" contenteditable @blur="somefunc">
  </div><!-- .card -->
</template>

<script>


export default {
  name: 'Card',
  props: {
    card: Object
  },
  methods: {
    somefunc: function(e) {
      this.card.content = e.target.innerHTML
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

.card {
  background-color: teal;
  color: white;
  border-radius: 2px;
  
  margin: 1em;
  padding: 1em;
}

</style>
