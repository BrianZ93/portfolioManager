<template>
  <q-page>
    <!-- Navigation for Portfolio Page -->
    <div class="q-pa-md">
      <div class="q-gutter-y-md">
        <q-tabs v-model="tab" narrow-indicator dense align="justify">
          <q-tab
            class="text-purple"
            name="Dashboard"
            icon="dashboard"
            label="Dashboard"
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
          <q-tab
            class="text-yellow"
            name="Budget"
            icon="pages"
            label="Budget"
          />
          <q-tab class="text-red" name="Debt" icon="pages" label="Debt" />
          <q-tab class="text-blue" name="Market" icon="pages" label="Market" />
        </q-tabs>
      </div>
    </div>

    <!-- Tab Components -->

    <DashboardComponent v-if="tab == 'Dashboard'"> </DashboardComponent>

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
      <q-toolbar class="ticker-tape">
        <q-toolbar-title class="col-shrink" style="color: gray"
          >Market Pulse</q-toolbar-title
        >
        <div class="ticker">
          <q-item style="text-align: center" class="ticker__item">
            <q-item-section
              :style="
                indexStore.data[3] < 0 ? { color: 'red' } : { color: 'green' }
              "
              >S&P 500 (GSPC)
              {{
                indexStore.data[2].toLocaleString('en-US', {
                  style: 'currency',
                  currency: 'USD',
                })
              }}
              {{ parseFloat(indexStore.data[3]).toFixed(2) + '%' }}
            </q-item-section>
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
          </q-item>

          <q-item style="text-align: center" class="ticker__item">
            <q-item-section
              :style="
                indexStore.data[7] < 0 ? { color: 'red' } : { color: 'green' }
              "
              >NASDAQ (IXIC)
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
          </q-item>

          <q-item style="text-align: center" class="ticker__item">
            <q-item-section
              :style="
                indexStore.data[5] < 0 ? { color: 'red' } : { color: 'green' }
              "
              >Dow Jones (DJI)
              {{
                indexStore.data[4].toLocaleString('en-US', {
                  style: 'currency',
                  currency: 'USD',
                })
              }}
              {{
                parseFloat(indexStore.data[5]).toFixed(2) + '%'
              }}</q-item-section
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
          </q-item>

          <q-item style="text-align: center" class="ticker__item">
            <q-item-section
              :style="
                indexStore.data[1] < 0 ? { color: 'red' } : { color: 'green' }
              "
              >VIX (VIX) {{ indexStore.data[0] }}
              {{ parseFloat(indexStore.data[1]).toFixed(2) + '%' }}
            </q-item-section>
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
          </q-item>
        </div>
      </q-toolbar>
    </q-footer>
  </q-page>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, ref } from 'vue';
import EquityComponent from '../components/EquityComponent.vue';
import PropertyComponent from '../components/PropertyComponent.vue';
import DashboardComponent from '../components/DashboardComponent.vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';
import { useIndexStore } from 'src/stores/indexes-store';

export default defineComponent({
  name: 'PortfolioPage',
  components: { EquityComponent, PropertyComponent, DashboardComponent },
  setup() {
    const portfolioStore = usePortfolioStore();
    const indexStore = useIndexStore();

    const indexString = indexStore.data[1] + 'Hi there';

    onBeforeMount(async () => {
      portfolioStore.importCurrentEquities();
      indexStore.getData();
    });

    return {
      tab: ref('Dashboard'),
      indexStore,
      indexString,
    };
  },
});
</script>

<style>
:root {
  --height: 7vh;
  --speed: 25s;
}

.ticker-tape {
  display: flex;

  width: 100%;
  /* overflow: hidden; */
  height: var(--height);
  padding-left: 100%;
}

.ticker-tape .ticker {
  height: var(--height);
  line-height: var(--height);
  white-space: nowrap;
  padding-right: 100%;

  -webkit-animation-iteration-count: infinite;
  -webkit-animation-timing-function: linear;
  -webkit-animation-name: ticker;
  -webkit-animation-duration: var(--speed);

  animation-iteration-count: infinite;
  animation-timing-function: linear;
  animation-name: ticker;
  animation-duration: var(--speed);
}

.ticker-tape .ticker__item {
  display: inline-block;
  padding: 0 0.5rem;
  font-size: 2rem;
  font-weight: 900;
  letter-spacing: 1rem;
  color: rgba(0, 0, 0, 0.2);
}

@-webkit-keyframes ticker {
  0% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    visibility: visible;
  }
  100% {
    -webkit-transform: translateX(-151000%);
    transform: translateX(-100%);
  }
}

@keyframes ticker {
  0% {
    -webkit-transform: translateX(0);
    transform: translateX(0);
    visibility: visible;
  }
  100% {
    -webkit-transform: translateX(-100%);
    transform: translateX(-100%);
  }
}
</style>
