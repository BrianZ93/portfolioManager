<template>
  <q-page>
    <div class="flex flex-center">
      <!-- Cards on Top -->
      <div class="row cards">
        <!-- Net Worth Card -->
        <div class="col section">
          <q-card class="card" v-ripple>
            <q-card-section>
              <div>Net Worth</div>
              <p1>Ticker or Description</p1>
            </q-card-section>
          </q-card>
        </div>

        <!-- Real Estate Card -->
        <div class="col section">
          <q-card class="card">
            <q-card-section>
              <div>Real Estate</div>
            </q-card-section>
          </q-card>
        </div>

        <!-- Equities Card -->
        <div class="col section">
          <q-card class="card">
            <q-card-section>
              <div>Equities</div>
            </q-card-section>
          </q-card>
        </div>

        <!-- Debt Card -->
        <div class="col section">
          <q-card class="card">
            <q-card-section>
              <div>Debt</div>
            </q-card-section>
          </q-card>
        </div>
      </div>

      <!-- Mid Charts -->
      <div class="row">
        <!-- Total Assets Chart -->
        <div
          class="col section summary"
          style="background-color: rgba(30, 30, 30, 1)"
        >
          <div class="title">Total Assets</div>
          <portfolioCompositionChart></portfolioCompositionChart>
        </div>

        <!-- RE Chart -->

        <div
          class="col section propertysummary"
          style="background-color: rgba(30, 30, 30, 1)"
        >
          <div class="title">Real Estate</div>
          <realEstateChart></realEstateChart>
        </div>

        <!-- Equities Chart -->

        <div
          class="col section equitysummary"
          style="background-color: rgba(30, 30, 30, 1)"
        >
          <div class="title">Equities</div>
          <equitiesChart></equitiesChart>
        </div>
      </div>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent, defineAsyncComponent, onBeforeMount } from 'vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';

const portfolioStore = usePortfolioStore();

const portfolioCompositionChart = defineAsyncComponent(
  () => import('components/charts/portfoliocomposition.vue')
);

const realEstateChart = defineAsyncComponent(
  () => import('components/charts/realesteatecomposition.vue')
);

const equitiesChart = defineAsyncComponent(
  () => import('components/charts/equitescomposition.vue')
);

export default defineComponent({
  name: 'DashboardComponent',
  components: {
    portfolioCompositionChart,
    realEstateChart,
    equitiesChart,
  },
  setup() {
    onBeforeMount(() => {
      portfolioStore.importCurrentEquities();
      portfolioStore.importCurrentProperties();
    });

    return {};
  },
});
</script>

<style lang="scss">
.title {
  text-align: center;
  font-size: 2rem;
  padding: auto;
  color: gray;
}

.section {
  transition: all 0.2s ease-in-out;
  -webkit-animation-duration: 2s;
  animation-duration: 2s;
  -webkit-animation-fill-mode: both;
  animation-fill-mode: both;
  -webkit-animation-name: fadeIn;
  animation-name: fadeIn;
  border-radius: 1rem;
}

.section:hover {
  transform: scale(1.1);
}

@-webkit-keyframes fadeIn {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

@keyframes fadeIn {
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
}

.summary {
  margin-left: 1vw;
  margin-right: 1vw;
}

.propertysummary {
  margin-left: 1vw;
  margin-right: 1vw;
}

.equitysummary {
  margin-left: 1vw;
  margin-right: 1vw;
}

/* Cards */

.cards {
  padding-bottom: 3vh;
}

.card {
  max-width: 15vw;
  min-width: 15vw;
  max-height: 10vh;
  min-height: 10vh;
  margin-left: 1vw;
  margin-right: 1vw;
  background-color: $primary;
}
</style>
