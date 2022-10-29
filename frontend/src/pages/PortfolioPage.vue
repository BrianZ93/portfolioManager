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

    <q-footer
      style="background-color: rgb(16, 16, 16)"
      reveal
      elevated
      bordered
    >
      <q-toolbar>
        <q-toolbar-title class="col-shrink" style="color: gray"
          >Market Pulse</q-toolbar-title
        >

        <q-item style="text-align: center">
          <q-item-section
            :style="
              indexStore.data[3] < 0 ? { color: 'red' } : { color: 'green' }
            "
            >S&P 500</q-item-section
          >
          <q-item-section
            :style="
              indexStore.data[3] < 0 ? { color: 'red' } : { color: 'green' }
            "
          >
            {{
              indexStore.data[2].toLocaleString('en-US', {
                style: 'currency',
                currency: 'USD',
              })
            }}
            {{
              parseFloat(indexStore.data[3]).toFixed(2) + '%'
            }}</q-item-section
          >
          <q-icon
            v-if="indexStore.data[3] < 0"
            style="color: red"
            name="arrow_drop_down"
            size="lg"
          ></q-icon>
          <q-icon
            v-if="indexStore.data[3] > 0"
            style="color: green"
            name="arrow_drop_up"
            size="lg"
          ></q-icon>
          <q-item-section></q-item-section>
        </q-item>

        <q-item style="text-align: center">
          <q-item-section
            :style="
              indexStore.data[7] < 0 ? { color: 'red' } : { color: 'green' }
            "
            >NASDAQ</q-item-section
          >
          <q-item-section
            :style="
              indexStore.data[7] < 0 ? { color: 'red' } : { color: 'green' }
            "
          >
            {{
              indexStore.data[6].toLocaleString('en-US', {
                style: 'currency',
                currency: 'USD',
              })
            }}
            {{
              parseFloat(indexStore.data[7]).toFixed(2) + '%'
            }}</q-item-section
          >
          <q-icon
            v-if="indexStore.data[7] < 0"
            style="color: red"
            name="arrow_drop_down"
            size="lg"
          ></q-icon>
          <q-icon
            v-if="indexStore.data[7] > 0"
            style="color: green"
            name="arrow_drop_up"
            size="lg"
          ></q-icon>
          <q-item-section></q-item-section>
        </q-item>

        <q-item style="text-align: center">
          <q-item-section
            :style="
              indexStore.data[5] < 0 ? { color: 'red' } : { color: 'green' }
            "
            >Dow Jones</q-item-section
          >
          <q-item-section
            :style="
              indexStore.data[5] < 0 ? { color: 'red' } : { color: 'green' }
            "
            >{{
              indexStore.data[4].toLocaleString('en-US', {
                style: 'currency',
                currency: 'USD',
              })
            }}
            {{
              parseFloat(indexStore.data[5]).toFixed(2) + '%'
            }}</q-item-section
          >
          <q-icon
            v-if="indexStore.data[5] < 0"
            style="color: red"
            name="arrow_drop_down"
            size="lg"
          ></q-icon>
          <q-icon
            v-if="indexStore.data[5] > 0"
            style="color: green"
            name="arrow_drop_up"
            size="lg"
          ></q-icon>
          <q-item-section></q-item-section>
        </q-item>

        <q-item style="text-align: center">
          <q-item-section
            :style="
              indexStore.data[1] < 0 ? { color: 'red' } : { color: 'green' }
            "
            >VIX</q-item-section
          >
          <q-item-section
            :style="
              indexStore.data[1] < 0 ? { color: 'red' } : { color: 'green' }
            "
            >{{ indexStore.data[0] }}
            {{ parseFloat(indexStore.data[1]).toFixed(2) + '%' }}
          </q-item-section>
          <q-icon
            v-if="indexStore.data[1] < 0"
            style="color: red"
            name="arrow_drop_down"
            size="lg"
          ></q-icon>
          <q-icon
            v-if="indexStore.data[1] > 0"
            style="color: green"
            name="arrow_drop_up"
            size="lg"
          ></q-icon>
          <q-item-section></q-item-section>
        </q-item>
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
import { storeToRefs } from 'pinia';

export default defineComponent({
  name: 'PortfolioPage',
  components: { EquityComponent, PropertyComponent },
  setup() {
    const portfolioStore = usePortfolioStore();
    const indexStore = useIndexStore();

    const { equityRows } = storeToRefs(portfolioStore);

    onBeforeMount(async () => {
      portfolioStore.importCurrentEquities();
      indexStore.getData();
    });

    return {
      tab: ref('Full Portfolio'),
      indexStore,
      equityRows,
    };
  },
});
</script>
