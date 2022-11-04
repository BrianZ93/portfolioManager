<template>
  <q-page>
    <div style="text-align: center">
      <q-btn
        class="glossy"
        rounded
        color="primary"
        label="New Debt"
        @click="persistentDebt = true"
      ></q-btn>
    </div>
  </q-page>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { usePortfolioStore } from 'src/stores/portfolio-store';

export default defineComponent({
  name: 'DebtComponent',
  setup() {
    let portfolioStore = usePortfolioStore();

    const persistentDebt = ref(false);

    return {
      persistentDebt,
    };
  },
});
</script>
