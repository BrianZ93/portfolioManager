<template>
  <q-page>
    <div class="flex">
      <div class="row">
        <!-- Total Assets Chart -->
        <div
          class="col-grow section summary"
          style="background-color: rgba(30, 30, 30, 1)"
        >
          <div class="title">Total Assets</div>
          <portfolioCompositionChart></portfolioCompositionChart>
        </div>
        <!-- RE Chart -->
        <div
          class="col-grow section summary"
          style="background-color: rgba(30, 30, 30, 1)"
        >
          <div class="title">Real Estate</div>
          <realEstateChart></realEstateChart>
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

export default defineComponent({
  name: 'DashboardComponent',
  components: {
    portfolioCompositionChart,
    realEstateChart,
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

<style>
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
  margin-left: 2vw;
}
</style>
