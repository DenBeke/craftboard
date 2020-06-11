<template>
  <div id="app">
    <div class="cards">


      <div class="col" :class="cardTypeKey" v-for="(cardTypeValue,cardTypeKey) in cards" :key="cardTypeKey">
        <h3>{{cardTypeValue.name}}</h3>
        <button @click="addNewCard(cardTypeKey)">Add new</button>
        <draggable :list="cardTypeValue" group="cards" class="card-draggable">
          <Card :card="card" v-for="card in cardTypeValue.cards" :key="card.id" :ref="'card-ref-'+card.id" />
        </draggable>
      </div><!-- .col -->

      
    </div><!-- .cards -->


    <div>
      <pre><code>{{ cards }}</code></pre>
    </div>

  </div>
</template>

<script>
import Card from './components/Card.vue'
import draggable from 'vuedraggable'

import { EventBus, AddNewCardEvent } from './eventbus.js';
import { createUUID } from './createUUID.js';

export default {
  name: 'App',
  components: {
    Card,
    draggable
  },
  data: function(){
    return {
      cards: {
        "todo": {
          name: "Todo",
          cards: [
            {
              id: "1",
              content: "I should do this"
            },
            {
              id: "2",
              content: "Foo bar"
            }
          ],
        },
        "in-progress": {
          name: "In Progress",
          cards: [
            {
              id: "3",
              content: "Working on something"
            },
          ],
        },
        "done": {
          name: "Done",
          cards: [
            {
              id: "4",
              content: "Finished this shit"
            },
          ]
        }
      }
    }
  },
  methods: {
    addNewCard: function(type) {
      console.log(`[addNewCard] with type ${type}`)

      let id = createUUID()

      this.cards[type].cards.push({
        "id": id,
        "content": ""
      })
      this.$nextTick(function () {
        EventBus.$emit(AddNewCardEvent, id);
      })
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

    display: grid;
    grid-template-columns: 33.33% 33.33% 33.33%;

    .col {

      .card-draggable {
        min-height: 50px;
        height: 100%;
      }

    }
  }

}
</style>
