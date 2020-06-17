<template>
  <div class="card">
    <div class="card-meta">
      <font-awesome-icon class="icon" icon="pencil-alt" @click="startEditCard" />
      <font-awesome-icon class="icon" icon="trash" @click="deleteCard" />  
    </div><!-- .card-meta -->
    <div class="card-content" ref="cardContent" :contenteditable="edit" v-html="card.content" @blur="endEditCard"></div>
  </div><!-- .card -->
</template>

<script>

import { EventBus, AddNewCardEvent, UpdatedCardEvent, DeleteCardEvent } from '../eventbus.js';


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
      this.card.updated_at = new Date().toJSON()
      this.edit = false
      EventBus.$emit(UpdatedCardEvent, this.card.id)
    },
    startEditCard: function() {
      console.log("click")
      this.edit = true
      
      var self = this

      setTimeout(function() {
        self.$refs["cardContent"].focus()
      }, 0);
    },
    deleteCard: function() {
      if (confirm("Do you really want to delete this card?") === true) { 
          EventBus.$emit(DeleteCardEvent, this.card.id)
      } else { 
          // ignore
      } 
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
  

  .card-meta {
    text-align: right;
    height: 1em;
    line-height: 1em;
    padding: 0;
    padding: 0.1em;
    padding-right: 0.2em;
    box-sizing: border-box;

    .icon {
      font-size: 0.8em;
      padding-left: 0.2em;
      cursor: pointer;
    }
  }

  .card-content {
    padding: 1em;
    padding-top: 0em;

    &::after {
      content: "\200B";
    }

    &:focus {
      outline: none;
    }
  }

  &:focus-within {
    box-shadow: 0 0 5px red;
  }
}

</style>
