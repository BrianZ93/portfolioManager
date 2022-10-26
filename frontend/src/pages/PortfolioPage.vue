<template>
  <q-page>
    <!-- Navigation for Portfolio Page -->
    <div class="q-pa-md">
      <div class="q-gutter-y-md" style="max-width: 400px">
        <q-tabs v-model="tab" narrow-indicator dense align="justify">
          <q-tab
            class="text-purple"
            name="Full Portfolio"
            icon="card_travel"
            label="Full Portflio"
          />
          <q-tab
            class="text-orange"
            name="Equities"
            icon="request_page"
            label="Equities"
          />
          <q-tab
            class="text-teal"
            name="Real Estate"
            icon="maps_home_work"
            label="Real Estate"
          />
        </q-tabs>
      </div>
    </div>

    <!-- Tab Components -->
    <EquityComponent
      v-if="tab == 'Equities'"
      title="Equities"
    ></EquityComponent>

    <PropertyComponent v-if="tab == 'Real Estate'" title="Properties">
    </PropertyComponent>

    <q-footer reveal elevated>
      <q-toolbar>
        <q-toolbar-title class="col-shrink">Market Pulse</q-toolbar-title>

        <q-section class="row" style="justify-content: center">
          <q-item>S&P 500</q-item>
          <q-item>{{ indexStore.data[2] }}</q-item>
          <q-item>{{ indexStore.data[3] }}</q-item>
          <q-item>NASDAQ</q-item>
          <q-item>{{ indexStore.data[6] }}</q-item>
          <q-item>{{ indexStore.data[7] }}</q-item>
          <q-item>DOW JONES</q-item>
          <q-item>{{ indexStore.data[4] }}</q-item>
          <q-item>{{ indexStore.data[5] }}</q-item>
          <q-item>VIX</q-item>
          <q-item>{{ indexStore.data[0] }}</q-item>
          <q-item>{{ indexStore.data[1] }}</q-item>
        </q-section>
      </q-toolbar>
    </q-footer>
  </q-page>
</template>

<script>
import { defineComponent, onBeforeMount, ref } from 'vue';
import EquityComponent from '../components/EquityComponent.vue';
import PropertyComponent from '../components/PropertyComponent.vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';
import { useIndexStore } from 'src/stores/indexes-store';

export default defineComponent({
  name: 'PortfolioPage',
  components: { EquityComponent, PropertyComponent },
  setup() {
    const portfolioStore = usePortfolioStore();
    const indexStore = useIndexStore();

    onBeforeMount(async () => {
      portfolioStore.importCurrentEquities();
      indexStore.getData();
    });

    return {
      tab: ref('Full Portfolio'),
      indexStore,
    };
  },
});
</script>
