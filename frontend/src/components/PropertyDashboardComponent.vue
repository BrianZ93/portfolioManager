<template>
  <q-card>
    <q-tabs v-model="tab" indicator-color="primary" class="text-primary">
      <q-tab name="tenants" icon="people" label="Tenants" />
      <q-tab name="ledger" icon="newspaper" label="Ledger" />
      <q-tab name="amortization" icon="hourglass_bottom" label="Amortization" />
    </q-tabs>

    <q-separator></q-separator>

    <q-card>
      <q-scroll-area style="height: 70vh; max-width: 80vw">
        <GeneralLedger v-if="tab == 'ledger'"></GeneralLedger>
        <TenantComponent v-if="tab == 'tenants'"></TenantComponent>
      </q-scroll-area>
    </q-card>
  </q-card>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import GeneralLedger from './GeneralLedger.vue';
import TenantComponent from './TenantComponent.vue';

export default defineComponent({
  name: 'PropertyDashboardComponent',
  components: {
    GeneralLedger,
    TenantComponent,
  },
  setup() {
    const tab = ref('tenants');

    return {
      tab,
    };
  },
});
</script>
