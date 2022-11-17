<template>
  <q-card>
    <div class="row">
      <div class="col-2">
        <q-item>
          <q-item-section avatar>
            <q-icon color="primary" name="people"></q-icon>
          </q-item-section>

          <q-item-section class="text-h6 text-primary">
            Tenants
          </q-item-section>
        </q-item>
        <q-list>
          <q-item
            clickable
            v-for="tenant in selectedProperty.tenants"
            :key="tenant"
            :value="tenant"
          >
            <q-item-section avatar>
              <q-icon color="primary" name="person"> </q-icon>
            </q-item-section>

            <q-item-section>
              <q-item-label> {{ tenant.name }} {{ tenant.unit }} </q-item-label>
            </q-item-section>
          </q-item>
        </q-list>
      </div>

      <!-- Tenant Dashboard Component -->
      <div class="col-10"></div>
    </div>
  </q-card>
</template>

<script lang="ts">
import { usePortfolioStore } from 'src/stores/portfolio-store';
import { defineComponent, ref, computed } from 'vue';

export default defineComponent({
  name: 'TenantComponent',
  setup() {
    let portfolioStore = usePortfolioStore();

    const persistentTenants = ref(false);
    const selectedProperty = computed(
      () => portfolioStore.form.selectedProperty
    );

    return {
      persistentTenants,
      selectedProperty,
    };
  },
});
</script>
