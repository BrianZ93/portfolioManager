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
  </q-page>
</template>

<script>
import { defineComponent, onBeforeMount, ref } from 'vue';
import EquityComponent from '../components/EquityComponent.vue';
import PropertyComponent from '../components/PropertyComponent.vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';

export default defineComponent({
  name: 'PortfolioPage',
  components: { EquityComponent, PropertyComponent },
  setup() {
    const portfolioStore = usePortfolioStore();

    onBeforeMount(async () => {
      portfolioStore.importCurrentEquities();
    });

    return {
      tab: ref('Full Portfolio'),
    };
  },
});
</script>
