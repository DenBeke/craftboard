<template>
  <div id="app">
    <div class="cards">


      <div class="column" :class="cardTypeKey" v-for="(cardTypeValue,cardTypeKey) in cards" :key="cardTypeKey">
        
        <div class="column-meta">
          <h3>{{cardTypeValue.name}}</h3>
        </div><!-- .column-meta -->
        
        <draggable :list="cardTypeValue.cards" group="cards" class="card-draggable">
          <Card :card="card" v-for="card in cardTypeValue.cards" :key="card.id" :ref="'card-ref-'+card.id" />
        </draggable>

        <div class="column-footer">
          <font-awesome-icon class="icon" icon="plus-square" @click="addNewCard(cardTypeKey)" />
        </div><!-- .column-footer -->

      </div><!-- .column -->

      
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

@import url('./assets/nunito-sans/nunito-sans.scss');

body,
html {
  margin: 0;
  padding: 0;

  font-family: 'Nunito Sans';
}

#app {

  .cards {

    display: grid;
    grid-template-columns: 33.33% 33.33% 33.33%;

    .column {

      .column-meta {
        text-align: center;
      }

      .card-draggable {
        min-height: 50px;
        //height: 100%;
      }

      .column-footer {
        text-align: center;
        

        .icon {
          color: teal;
          font-size: 1.5em;
        }
      }

    }
  }

}
</style>
