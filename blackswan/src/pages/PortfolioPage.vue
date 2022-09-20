<template>
  <q-page>
    <EquityComponent
      title="My Portfolio"
    ></EquityComponent>
  </q-page>
</template>

<script>

import { defineComponent, onBeforeMount } from 'vue';
import EquityComponent from '../components/EquityComponent.vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';

export default defineComponent({
  name: 'PortfolioPage',
  components: { EquityComponent },
  setup () {
    const portfolioStore = usePortfolioStore()

    onBeforeMount(async () => {
      portfolioStore.importCurrentEquities()
    });
  }
});



</script>
