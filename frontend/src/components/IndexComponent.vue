<template>
  <h1 class="maintitle">{{ title }}</h1>
  <br>
  <div class="mainflexbox">
    <div class="flexitem" >
      <h2>S&P 500</h2>
      <p1>{{ indexStore.data[2] }}</p1>
      <br>
      <p1>{{ indexStore.data[3] }}</p1>
    </div>
    <div class="flexitem">
      <h2>Nasdaq</h2>
      <p1>{{ indexStore.data[6] }}</p1>
      <br>
      <p1>{{ indexStore.data[7] }}</p1>
    </div>
    <div class="flexitem">
      <h2>Dow Jones</h2>
      <p1>{{ indexStore.data[4] }}</p1>
      <br>
      <p1>{{ indexStore.data[5] }}</p1>
    </div>
  </div>
  <div class="mainflexbox">
    <div class="flexitem">
      <h2>VIX</h2>
      <p1>{{ indexStore.data[0] }}</p1>
      <br>
      <p1>{{ indexStore.data[1] }}</p1>
    </div>
    <div class="flexitem">
      <h2>SPX P/E</h2>
      <p1>0</p1>
      <br>
      <p1>(Current Ratio)</p1>
    </div>
    <div class="flexitem">
      <h2>VIX</h2>
      <p1>0%</p1>
      <br>
      <p1>(Current Index)</p1>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useIndexStore} from 'src/stores/indexes-store';
import { storeToRefs } from 'pinia';

export default defineComponent({
  name: 'IndexComponent',
  props: {
    title: {
      type: String,
      required: true
    },
  },
  setup () {
    const indexStore = useIndexStore()

    const { indexes } = storeToRefs(indexStore)

    indexStore.getData()

    return {
      indexes,
      indexStore
    }
  }
});
</script>

