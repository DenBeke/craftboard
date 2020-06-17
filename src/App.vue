<template>
  <div id="app">
    <div class="cards">


      <div class="column" :class="cardTypeKey" v-for="(cardTypeValue,cardTypeKey) in cards" :key="cardTypeKey">
        
        <div class="column-meta">
          <h3>{{cardTypeValue.name}}</h3>
        </div><!-- .column-meta -->
        
        <draggable :list="cardTypeValue.cards" group="cards" class="card-draggable" @change="changedBoardByDragDrop">
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

import { EventBus, AddNewCardEvent, UpdatedCardEvent } from './eventbus.js';
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
  created: function() {
    var self = this

    // Listen for board updates
    EventBus.$on(UpdatedCardEvent, function(cardID){
      console.log(UpdatedCardEvent + " " + cardID)
      self.saveBoard()
    })

    // Get the initial board
    self.getBoard()

  },
  methods: {
    addNewCard: function(type) {
      console.log(`[addNewCard] with type ${type}`)

      let id = createUUID()

      this.cards[type].cards.push({
        id: id,
        content: "",
        created_at: new Date().toJSON(),
        updated_at: new Date().toJSON()
      })
      this.$nextTick(function () {
        EventBus.$emit(AddNewCardEvent, id);
      })
      
    },
    changedBoardByDragDrop: function() {
      EventBus.$emit(UpdatedCardEvent);
    },
    saveBoard: function() {

      var self = this

      self.axios.post('http://localhost:1234/api/v1/board',
        self.cards
      )
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
    },
    getBoard: function() {

      var self = this

      self.axios.get('http://localhost:1234/api/v1/board')
      .then(function (response) {
        self.cards = response.data
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
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
          cursor: pointer;
        }
      }

    }
  }

}
</style>
