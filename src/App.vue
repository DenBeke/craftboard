<template>
  <div id="app">
    <div class="cards">
      <Card :card="card" v-for="card in cards" :key="card.id"/>
    </div><!-- .cards -->


    <div>
      <pre><code>{{ cards }}</code></pre>
    </div>

  </div>
</template>

<script>
import Card from './components/Card.vue'

export default {
  name: 'App',
  components: {
    Card
  },
  data: function(){
    return {
      cards: [
        {
          "id": 1,
          "content": "Lorem Ipsum"
        },
        {
          "id": 2,
          "content": "Foo bar"
        }
      ]
    }
  }
}
</script>

<style lang="scss">

body,
html {
  margin: 0;
  padding: 0;
}

#app {

  .cards {
    width: 33%;
  }

}
</style>
