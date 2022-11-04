<template>
  <q-page>
    <div style="text-align: center">
      <q-btn
        class="glossy"
        rounded
        color="primary"
        label="New Income Source"
        @click="persistentProperty = true"
      ></q-btn>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';

export default defineComponent({
  name: 'BudgetComponent',
  setup() {
    let portfolioStore = usePortfolioStore();
  },
});
</script>
