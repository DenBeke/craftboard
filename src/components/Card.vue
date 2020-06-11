<template>
  <div class="card">
    <button @click="startEditCard">Edit</button>
    <div ref="cardContent" :contenteditable="edit" v-html="card.content" @blur="endEditCard"></div>
  </div><!-- .card -->
</template>

<script>

import { EventBus, AddNewCardEvent } from '../eventbus.js';


export default {
  name: 'Card',
  props: {
    card: Object
  },
  data: function(){
    return {
      edit: false
    }
  },
  mounted: function() {
    EventBus.$on(AddNewCardEvent, cardID => {
      if(this.card.id == cardID) {
        console.log(`[${AddNewCardEvent}] event received that matched this card id: ${cardID}`)
        this.startEditCard()
      }
    });
  },
  methods: {
    endEditCard: function(e) {
      this.card.content = e.target.innerHTML
      this.edit = false
    },
    startEditCard: function() {
      console.log("click")
      this.edit = true
      
      var self = this

      setTimeout(function() {
        self.$refs["cardContent"].focus()
      }, 0);
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
